create table if not exists customer (
    id               varchar not null primary key,
    name             varchar not null,
    surname          varchar not null,
    created_at       timestamptz,
    updated_at       timestamptz
);
