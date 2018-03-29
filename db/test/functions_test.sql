delete from logs;
delete from payments;
delete from comments;
delete from projects;
delete from categories;
delete from users;

-- test for users.sql

select create_user('1@gmail.com', '123', '1', '1', null, null, null, null, null, null, null);
select create_user('2@gmail.com', '1234', '2', '2', null, null, null, null, null, null, null);
select create_user('3@gmail.com', '12345', '3', '3', null, null, null, null, null, null, null);
select * from all_users();
select * from get_user('1@gmail.com', '123');
delete from users where email = '3@gmail.com';

-- test for categories.sql

select create_categories('Art');
select create_categories('Science');
select * from all_categories();

-- test for projects.sql

select create_project('hello', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', null, null, null, '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('bye', (
    select max(user_id) - 1
    from users
), 'Art', 'bye', null, null, null, '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select * from all_projects();
select * from get_project((
    select max(project_id)
    from projects
));

-- test for payments.sql

select create_payment((
    select max(user_id)
    from users
), (
    select max(project_id)
    from projects
), 20.0);
select create_payment((
    select max(user_id)
    from users
), (
    select max(project_id)
    from projects
), 20.0);
select update_payment((
    select max(id)
    from payments
), 300.0);
select * from all_project_payments((
    select max(project_id)
    from projects
));
select * from all_user_payments((
    select max(user_id)
    from users
));
select delete_payment((
    select max(id) - 1
    from payments
));
select * from all_payments();

-- test for comments.sql

select create_comment((
    select max(user_id)
    from users
), (
    select max(project_id)
    from projects
), 'hello');
select create_comment((
    select max(user_id)
    from users
), (
    select max(project_id)
    from projects
), 'hello');
select update_comment((
    select max(id)
    from comments
), 'bye');
select * from all_project_comments((
    select max(project_id)
    from projects
));
select * from all_user_comments((
    select max(user_id)
    from users
));
select delete_comment((
    select max(id) - 1
    from comments
));
select * from all_comments();