CREATE TABLE "customer.customers"(
    customer_id bigserial primary key,
    phone varchar not null,
    email varchar not null,
    hashed_password varchar not null,
    create_at timestamptz not null default 'now()',
    updated_at timestamptz not null default 'now()'
);

CREATE TABLE "producer.producers"(
    producer_id bigserial primary key,
    phone varchar not null,
    email varchar not null,
    hashed_password varchar not null,
    create_at timestamptz not null default 'now()',
    updated_at timestamptz not null default 'now()'
);

CREATE TABLE "courier.couriers" (
    courier_id bigserial primary key,
    producer_id bigint not null,
    phone varchar not null,
    email varchar not null,
    hashed_password varchar not null,
    create_at timestamptz not null default 'now()',
    updated_at timestamptz not null default 'now()'
);

ALTER TABLE "courier.couriers" ADD FOREIGN KEY ("producer_id") REFERENCES "producer.producers" ("producer_id");