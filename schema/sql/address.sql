create table if not exists address (
    id               varchar not null primary key,
    street           varchar not null,
    zip              varchar not null,
    city             varchar not null,
    country          varchar not null,
    created_at       timestamptz,
    updated_at       timestamptz
);
