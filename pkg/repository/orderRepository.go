package repository

import (
	"fmt"
	"github.com/MrClean-code/wbtech"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateOrder(order wbtech.Order) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	//createItemQuery := fmt.Sprintf(`
	//			INSERT INTO %s (track_number, entry, locale, internal_signature, customer_id,
	//			                customer_id, delivery_service, shard_key, sm_id, date_created,
	//			                oof_shard)
	//			values ($1, $2) RETURNING id`, todoItemsTable)
	//
	//row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	//err = row.Scan(&itemId)
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}
	//
	//createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	//_, err = tx.Exec(createListItemsQuery, orderId, itemId)
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}

	return itemId, tx.Commit()
}

func (r *OrderPostgres) GetOrders() ([]wbtech.Order, error) {
	logrus.Print("GetOrders orderRepository is ok")
	var items []wbtech.Order
	query := fmt.Sprintf(`SELECT o.id, o.tracknumber, o.entry, o.locale, o.internal_signature,
							     o.customer_id, o.delivery_service, o.shard_key, o.sm_id,
							     o.date_created, o.oof_shard, d.id, d.name, d.phone, d.city,
							     d.address, d.region, d.email, i.id, i.track_number, i.price,
							     i.rid, i.name, i.sale, i.size, i.total_price, i.nmid, i.brand,
							     i.status, p.id, p.request_id, p.currency, p.provider, p.amount,
							     p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee
								 FROM orders o
									 INNER JOIN delivery d on d.id = o.deliveryid
									 INNER JOIN item i 	   on i.id = o.paymentid
									 INNER JOIN payment p  on p.id = o.itemid`)
	err := r.db.Select(&items, query)
	fmt.Println("items")
	fmt.Println(items)

	if err != nil {
		return nil, err
	}

	return items, nil
}
