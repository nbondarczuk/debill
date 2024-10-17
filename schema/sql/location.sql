create table if not exists location (
    id               varchar not null primary key,
    name             varchar not null,
    latitude         float,
    longitude        float,
    created_at       timestamptz,
    updated_at       timestamptz
);
