# wbtech

http://localhost:8000/api/orders -- данные о заказе
http://localhost:8000/api/orders?id -- получить данные о заказе по id

модель данных:
[
    {
        "order_uuid": 1,
        "track_number": "123",
        "entry": "vds",
        "locale": "ds",
        "internal_signature": "das",
        "customer_id": "das",
        "delivery_service": "dsa",
        "shard_key": "123",
        "sm_id": 1,
        "date_created": "321",
        "oof_shard": "vcx",
        "delivery": {
            "id": 1,
            "name": "Ivan",
            "phone": "86754234",
            "city": "Volga",
            "address": "v2",
            "region": "vs",
            "email": "32@es"
        },
        "payment": {
            "transaction": 0,
            "request_id": "1",
            "currency": "231",
            "provider": "sad",
            "amount": 433,
            "payment_dt": 213,
            "bank": "BB",
            "delivery_cost": 123,
            "goods_total": 43,
            "custom_fee": 312
        },
        "items": [
              {
                "chrt_id": 0,
                "track_number": "3cx21z",
                "price": 123,
                "rid": 1,
                "name": "Celli",
                "sale": 123,
                "size": "32c12",
                "total_price": 231,
                "nm_id": 0,
                "brand": "vyews",
                "status": 123
            }
        ]
    }
]
