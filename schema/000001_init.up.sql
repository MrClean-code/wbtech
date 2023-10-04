-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE delivery
(
    id      serial primary key,
    name    varchar(255),
    phone   varchar(255),
    city    varchar(255),
    address varchar(255),
    region  varchar(255),
    email   varchar(255)
);

CREATE TABLE payment
(
    id            serial primary key,
    request_id    varchar(255),
    currency      varchar(255),
    provider      varchar(255),
    amount        int,
    payment_dt    int,
    bank          varchar(255),
    delivery_cost int,
    goods_total   int,
    custom_fee    int
);

CREATE TABLE item
(
    id           serial primary key,
    track_number varchar(255),
    price        int,
    rid          serial,
    name         varchar(255),
    sale         int,
    size         varchar(255),
    total_price  int,
    NmId         int,
    brand        varchar(255),
    status       int
);

CREATE TABLE orders
(
    id                 serial primary key,
    trackNumber        varchar(255),
    entry              varchar(255),
    locale             varchar(255),
    internal_signature varchar(255),
    customer_id        varchar(255),
    delivery_service   varchar(255),
    shard_key          varchar(255),
    sm_id              int,
    date_created       varchar(255),
    oof_shard          varchar(255),
    deliveryId         int,
    paymentId          int,
    itemId             int,
    foreign key (deliveryId) references delivery (id) on delete cascade,
    foreign key (paymentId) references payment (id) on delete cascade,
    foreign key (itemId) references item (id) on delete cascade
);