\copy categories FROM '/docker-entrypoint-initdb.d/categories.csv' CSV HEADER;
\copy users FROM '/docker-entrypoint-initdb.d/users.csv' CSV HEADER;
update users
set password = crypt(password, gen_salt('bf', 8));
-- SOURCE: https://stackoverflow.com/a/3698777
select setval(pg_get_serial_sequence('users', 'user_id'), coalesce(max(user_id),0) + 1, false) FROM users;
\copy projects FROM '/docker-entrypoint-initdb.d/projects.csv' CSV HEADER;
select setval(pg_get_serial_sequence('projects', 'project_id'), coalesce(max(project_id),0) + 1, false) FROM projects;
\copy payments FROM '/docker-entrypoint-initdb.d/payments.csv' CSV HEADER;
select setval(pg_get_serial_sequence('payments', 'id'), coalesce(max(id),0) + 1, false) FROM payments;
\copy comments FROM '/docker-entrypoint-initdb.d/comments.csv' CSV HEADER;
select setval(pg_get_serial_sequence('comments', 'id'), coalesce(max(id),0) + 1, false) FROM comments;