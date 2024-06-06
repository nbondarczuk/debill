create table if not exists account (
    id               varchar not null primary key,
    name             varchar not null,
    created_at       timestamptz,
    updated_at       timestamptz
);
