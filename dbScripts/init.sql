-- Table: public.orders

-- DROP TABLE IF EXISTS public.orders;

CREATE TABLE IF NOT EXISTS public.orders
(
    order_uid text COLLATE pg_catalog."default" NOT NULL,
    track_number text COLLATE pg_catalog."default" NOT NULL,
    entry text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    phone text COLLATE pg_catalog."default" NOT NULL,
    zip text COLLATE pg_catalog."default" NOT NULL,
    city text COLLATE pg_catalog."default" NOT NULL,
    address text COLLATE pg_catalog."default" NOT NULL,
    region text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    transaction text COLLATE pg_catalog."default" NOT NULL,
    request_id text COLLATE pg_catalog."default" NOT NULL,
    currency text COLLATE pg_catalog."default" NOT NULL,
    provider text COLLATE pg_catalog."default" NOT NULL,
    amount bigint NOT NULL,
    payment_dt bigint NOT NULL,
    bank text COLLATE pg_catalog."default" NOT NULL,
    delivery_cost bigint NOT NULL,
    goods_total bigint NOT NULL,
    custom_fee bigint NOT NULL,
    locale text COLLATE pg_catalog."default" NOT NULL,
    internal_signature text COLLATE pg_catalog."default" NOT NULL,
    customer_id text COLLATE pg_catalog."default" NOT NULL,
    delivery_service text COLLATE pg_catalog."default" NOT NULL,
    shardkey text COLLATE pg_catalog."default" NOT NULL,
    sm_id bigint NOT NULL,
    date_created timestamp with time zone NOT NULL,
    oof_shard text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT orders_pkey PRIMARY KEY (order_uid)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.orders
    OWNER to userwb;

-- Table: public.items

-- DROP TABLE IF EXISTS public.items;

CREATE TABLE IF NOT EXISTS public.items
(
    chrt_id bigint NOT NULL,
    track_number text COLLATE pg_catalog."default" NOT NULL,
    price bigint NOT NULL,
    rid text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    sale bigint NOT NULL,
    size text COLLATE pg_catalog."default" NOT NULL,
    total_price bigint NOT NULL,
    nm_id bigint NOT NULL,
    brand text COLLATE pg_catalog."default" NOT NULL,
    status bigint NOT NULL,
    CONSTRAINT items_pkey PRIMARY KEY (chrt_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.items
    OWNER to userwb;