package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MrClean-code/wbtech"
	"github.com/MrClean-code/wbtech/pkg/nats"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
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
	//tx, err := r.db.Begin()
	//if err != nil {
	//	return 0, err
	//}

	var ord int
	//createItemQuery := fmt.Sprintf(`
	//			INSERT INTO %s (tracknumber, entry, locale, internal_signature,
	//			                customer_id, delivery_service, shard_key,
	//			                sm_id, date_created, oof_shard, deliveryid,
	//			                paymentid, itemid)
	//			values
	//			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	//			RETURNING id`, OrdersTable)
	//row := tx.QueryRow(createItemQuery, order.TrackNumber, order.Entry,
	//	order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService,
	//	order.ShardKey, order.SmId, order.DateCreated, order.OofShard, order.Delivery,
	//	order.Payment, order.Items)
	//err = row.Scan(&ord)
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}

	//return ord, tx.Commit()
	return ord, nil
}

func (r *OrderPostgres) GetOrders(c *gin.Context) ([]wbtech.Order, error) {
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

	message, err := json.Marshal(orders)
	if err != nil {
		log.Fatalf("Failed marshal orders")
	}

	sc, err := nats.ConnectNATSStreaming()
	err = sc.Publish("order-service", message)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	defer sc.Close()

	return orders, nil

}
