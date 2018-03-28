CREATE EXTENSION citext;
CREATE EXTENSION pgcrypto;
drop table if exists users CASCADE;
drop table if exists categories CASCADE;
drop table if exists projects CASCADE;
drop table if exists payments CASCADE;
drop table if exists comments CASCADE;
drop table if exists logs CASCADE;
drop view if exists whole_project_info CASCADE;

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    email CITEXT UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS categories (
    name CITEXT UNIQUE PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS projects (
    project_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    start_time TIMESTAMP NOT NULL DEFAULT now(),
    end_time TIMESTAMP NOT NULL CHECK (start_time <= end_time),
    amount_required NUMERIC(10, 2) NOT NULL CHECK (amount_required > 0), --10 sf, 2dp--
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    category CITEXT NOT NULL REFERENCES categories (name) ON UPDATE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users (user_id) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (user_id) ON UPDATE CASCADE,
    project_id INTEGER NOT NULL REFERENCES projects (project_id) ON UPDATE CASCADE,
    moment TIMESTAMP NOT NULL DEFAULT now(),
    amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0) --10 sf, 2dp--
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users (user_id) ON UPDATE CASCADE,
    project_id INTEGER NOT NULL REFERENCES projects (project_id) ON UPDATE CASCADE,
    moment TIMESTAMP NOT NULL DEFAULT now(),
    content TEXT NOT NULL
);

create table if not exists logs (
    id serial primary key,
    user_id integer default null, 
    project_id integer default null,
    moment timestamp not null default now(),
    content text not null,
    log_level integer not null
);

create view whole_project_info as 
    SELECT p.project_id, p.title, p.description,
        p.start_time, p.end_time,
        p.amount_required, COALESCE(SUM(f.amount), 0) as amount_raised,
        p.verified, p.category, p.user_id
    FROM projects p LEFT OUTER JOIN payments f
    ON p.project_id = f.project_id
    GROUP BY p.project_id, p.user_id;

-- -- SOURCE: https://www.meetspaceapp.com/2016/04/12/passwords-postgresql-pgcrypto.html
-- PREPARE create_user (TEXT, VARCHAR) AS
--     INSERT INTO users (email, password)
--     VALUES($1, crypt($2, gen_salt('bf', 8)));

-- PREPARE select_user (TEXT, VARCHAR) AS
--     SELECT id, email, is_admin FROM users
--     WHERE email = $1
--     AND password = crypt($2, password);

-- \copy categories FROM '/docker-entrypoint-initdb.d/categories.csv' CSV HEADER;
-- \copy users FROM '/docker-entrypoint-initdb.d/users.csv' CSV HEADER;
-- -- SOURCE: https://stackoverflow.com/a/3698777
-- SELECT setval(pg_get_serial_sequence('users', 'id'), coalesce(max(id),0) + 1, false) FROM users;
