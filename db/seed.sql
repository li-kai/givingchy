CREATE EXTENSION citext;
CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email CITEXT UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS categories (
    name CITEXT UNIQUE PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    start_time TIMESTAMP NOT NULL DEFAULT now(),
    end_time TIMESTAMP NOT NULL CHECK (start_time <= end_time),
    amount_required NUMERIC(10, 2) NOT NULL CHECK (amount_required > 0), --10 sf, 2dp--
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    category CITEXT NOT NULL REFERENCES categories (name) ON UPDATE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users (id) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS funds (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON UPDATE CASCADE,
    moment TIMESTAMP NOT NULL DEFAULT now(),
    amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0) --10 sf, 2dp--
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON UPDATE CASCADE,
    moment TIMESTAMP NOT NULL DEFAULT now(),
    content TEXT NOT NULL
);

-- SOURCE: https://www.meetspaceapp.com/2016/04/12/passwords-postgresql-pgcrypto.html
PREPARE create_user (TEXT, VARCHAR) AS
    INSERT INTO users (email, password)
    VALUES($1, crypt($2, gen_salt('bf', 8)));

PREPARE select_user (TEXT, VARCHAR) AS
    SELECT id, email, is_admin FROM users
    WHERE email = $1
    AND password = crypt($2, password);

\copy categories FROM '/docker-entrypoint-initdb.d/categories.csv' CSV HEADER;
\copy users FROM '/docker-entrypoint-initdb.d/users.csv' CSV HEADER;
-- SOURCE: https://stackoverflow.com/a/3698777
SELECT setval(pg_get_serial_sequence('users', 'id'), coalesce(max(id),0) + 1, false) FROM users;
