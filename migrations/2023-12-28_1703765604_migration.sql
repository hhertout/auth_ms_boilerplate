/* Caution : each query must be separated with empty comment */

CREATE TABLE IF NOT EXISTS "user"
(
    id         text PRIMARY KEY not null DEFAULT gen_random_uuid(),
    created_at date             not null,
    deleted_at date,
    email      varchar(255)     not null unique,
    password   varchar(255)     not null
);

--

CREATE OR REPLACE FUNCTION set_creation_date()
    RETURNS TRIGGER AS
$$
BEGIN
    New.created_at = NOW();
    RETURN New;
END;
$$ language 'plpgsql';

--

CREATE OR REPLACE TRIGGER auto_creation_date
    BEFORE INSERT
    ON "user"
    FOR EACH ROW
EXECUTE PROCEDURE set_creation_date();