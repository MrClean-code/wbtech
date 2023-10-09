package repository

import (
	"context"
	"fmt"
	"github.com/MrClean-code/wbtech"
	"github.com/MrClean-code/wbtech/pkg/nats"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type OrderPostgres struct {
	db *pgx.Conn
}

func NewOrderPostgres(db *pgx.Conn) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

func (r *OrderPostgres) CreateOrder(order wbtech.Order) (int, error) {
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var ord int
	createItemQuery := fmt.Sprintf(`
		INSERT INTO %s (tracknumber, entry, locale, internal_signature,
		                customer_id, delivery_service, shard_key,
		                sm_id, date_created, oof_shard, delivery_id,
		                payment_id, item_id)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`, OrdersTable)

	row := tx.QueryRow(ctx, createItemQuery, order.TrackNumber, order.Entry,
		order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService,
		order.ShardKey, order.SmId, order.DateCreated, order.OofShard,
		order.Delivery, order.Payment, order.Item)

	if err := row.Scan(&ord); err != nil {
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return ord, nil
}

func (r *OrderPostgres) GetOrders(c *gin.Context) ([]wbtech.Order, error) {

	sc, err := nats.ConnectNATSStreaming()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DeliverAllAvailable())

	if err != nil {
		fmt.Printf("Error subscribing: %v\n", err)
	}

	var orders []wbtech.Order
	query := "SELECT orders.id, orders.tracknumber, orders.entry, " +
		"orders.locale, orders.internal_signature, " +
		"orders.customer_id, orders.delivery_service, orders.shard_key, orders.sm_id, " +
		"orders.date_created, orders.oof_shard, " +
		"(SELECT row_to_json(delivery) FROM delivery WHERE delivery.id = orders.delivery_id) AS delivery, " +
		"(SELECT row_to_json(payment) FROM payment WHERE payment.id = orders.payment_id) AS payment " +
		"FROM orders " +
		"INNER JOIN delivery ON delivery.id = orders.delivery_id " +
		"INNER JOIN payment ON payment.id = orders.payment_id "

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order wbtech.Order
		err := rows.Scan(
			&order.ID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard,
			&order.Delivery,
			&order.Payment,
		)
		if err != nil {
			return nil, err
		}

		rows.Close()

		itemsQuery := "SELECT row_to_json(item) FROM item WHERE item.id = $1"
		itemRows, err := r.db.Query(context.Background(), itemsQuery, order.ID)
		if err != nil {
			fmt.Println("Ошибка выполнения запроса для элементов:", err)
			return nil, err
		}

		for itemRows.Next() {
			var item wbtech.Item
			err := itemRows.Scan(&item)
			if err != nil {
				return nil, err
			}
			order.Item = append(order.Item, item)
		}
		itemRows.Close()

		orders = append(orders, order)

	}

	sub.Unsubscribe()

	defer sc.Close()

	return orders, nil
}

func (r *OrderPostgres) GetOrderByID(id int) (wbtech.Order, error) {
	ctx := context.Background()
	logrus.Printf("%d", id)
	query := `
		SELECT
            orders.id,
            orders.tracknumber,
            orders.entry,
            orders.locale,
            orders.internal_signature,
            orders.customer_id,
            orders.delivery_service,
            orders.shard_key,
            orders.sm_id,
            orders.date_created,
            orders.oof_shard,
            (SELECT row_to_json(delivery) FROM delivery WHERE delivery.id = orders.delivery_id) AS delivery,
            (SELECT row_to_json(payment) FROM payment WHERE payment.id = orders.payment_id) AS payment
        FROM orders
        INNER JOIN delivery ON delivery.id = orders.delivery_id
        INNER JOIN payment ON payment.id = orders.payment_id
        WHERE orders.id = $1`

	var order wbtech.Order
	err := r.db.QueryRow(ctx, query, id).Scan(
		&order.ID,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard,
		&order.Delivery,
		&order.Payment,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return wbtech.Order{}, fmt.Errorf("Order with ID %d not found", id)
		}
		return wbtech.Order{}, err
	}

	itemsQuery := "SELECT row_to_json(item) FROM item WHERE item.id = $1"
	itemRows, err := r.db.Query(ctx, itemsQuery, id)
	if err != nil {
		return wbtech.Order{}, err
	}
	defer itemRows.Close()

	for itemRows.Next() {
		var item wbtech.Item
		err := itemRows.Scan(&item)
		if err != nil {
			return wbtech.Order{}, err
		}
		order.Item = append(order.Item, item)
	}

	return order, nil
}
