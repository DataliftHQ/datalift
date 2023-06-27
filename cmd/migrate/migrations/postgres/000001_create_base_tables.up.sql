CREATE TABLE IF NOT EXISTS application
(
    id         serial primary key,
    name       varchar(255) not null,
    created_at timestamp with time zone default now(),
    created_by varchar(255) not null
);

CREATE TABLE IF NOT EXISTS authn_tokens(
(
    user_id       text not null,
    provider      text not null,
    access_token  bytea,
    refresh_token bytea,
    id_token      bytea,
    expiry        timestamp with time zone,
    primary key (user_id, provider)
);


