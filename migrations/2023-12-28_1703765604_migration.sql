/* Caution : each query must be separated with empty comment */

CREATE TABLE IF NOT EXISTS "user" (
    id text PRIMARY KEY not null,
    created_at date,
    deleted_at date,
    email varchar(255) not null unique,
    password varchar(255) not null
)