package wbtech

type Order struct {
	ID                int32  `json:"order_uuid" db:"id"`
	TrackNumber       string `json:"track_number" db:"tracknumber"`
	Entry             string `json:"entry" db:"entry"`
	Locale            string `json:"locale" db:"locale"`
	InternalSignature string `json:"internal_signature" db:"internal_signature"`
	CustomerId        string `json:"customer_id" db:"customer_id"`
	DeliveryService   string `json:"delivery_service" db:"delivery_service"`
	ShardKey          string `json:"shard_key" db:"shard_key"`
	SmId              int32  `json:"sm_id" db:"sm_id"`
	DateCreated       string `json:"date_created" db:"date_created"`
	OofShard          string `json:"oof_shard" db:"oof_shard"`
	//DeliveryId        int32    `json:"deliveryId" db:"deliveryid"`
	//PaymentId         int32    `json:"paymentId" db:"paymentid"`
	//ItemId            int32    `json:"itemId" db:"itemid"`
	Delivery Delivery `json:"delivery"`
	Payment  Payment  `json:"payment"`
	Item     []Item   `json:"items"`
}

type Delivery struct {
	ID      int32  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Payment struct {
	ID           int32  `json:"transaction" db:"id"`
	RequestId    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int32  `json:"amount" db:"amount"`
	PaymentDt    int32  `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int32  `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int32  `json:"goods_total" db:"goods_total"`
	CustomFee    int32  `json:"custom_fee" db:"custom_fee"`
}

type Item struct {
	ID          int32  `json:"chrt_id" db:"id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int32  `json:"price" db:"price"`
	Rid         int32  `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int32  `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int32  `json:"total_price" db:"total_price"`
	NmId        int32  `json:"nm_id" db:"nmid"`
	Brand       string `json:"brand" db:"brand"`
	Status      int32  `json:"status" db:"status"`
}
