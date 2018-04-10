-- SOURCE: https://stackoverflow.com/a/3698777

-- categories
\copy categories FROM '/docker-entrypoint-initdb.d/categories.csv' CSV HEADER;

-- users
\copy users FROM '/docker-entrypoint-initdb.d/users.csv' CSV HEADER;
update users set password = crypt(password, gen_salt('bf', 8));
select setval(pg_get_serial_sequence('users', 'user_id'), coalesce(max(user_id),0) + 1, false) FROM users;

-- projects
\copy projects FROM '/docker-entrypoint-initdb.d/projects.csv' CSV HEADER;
select setval(pg_get_serial_sequence('projects', 'project_id'), coalesce(max(project_id),0) + 1, false) FROM projects;

-- tags
\copy tags FROM '/docker-entrypoint-initdb.d/tags.csv' CSV HEADER;

-- payments
\copy payments FROM '/docker-entrypoint-initdb.d/payments.csv' CSV HEADER;
select setval(pg_get_serial_sequence('payments', 'id'), coalesce(max(id),0) + 1, false) FROM payments;

-- comments
\copy comments FROM '/docker-entrypoint-initdb.d/comments.csv' CSV HEADER;
select setval(pg_get_serial_sequence('comments', 'id'), coalesce(max(id),0) + 1, false) FROM comments;
