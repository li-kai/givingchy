delete from payments;
delete from comments;
delete from projects;
delete from categories;
delete from users;

-- test for users.sql

select create_user('1@gmail.com', '123');
select create_user('2@gmail.com', '1234');
select * from all_users();
select * from get_user('1@gmail.com', '123');

-- test for categories.sql

select create_categories('Art');
select create_categories('Science');
select * from all_categories();

-- test for projects.sql

select create_project('hello', 'hello', 100.0, '2018-08-20 14:52:49'::timestamp, 'Art', (
    select max(id) - 1
    from users
));
select create_project('bye', 'bye', 100.0, '2018-08-20 14:52:49'::timestamp, 'Art', (
    select max(id)
    from users
));
select * from all_projects();
select * from get_project((
    select max(id)
    from projects
));

-- test for payments.sql

select create_payment((
    select max(id)
    from users
), (
    select max(id)
    from projects
), 20.0);
select create_payment((
    select max(id)
    from users
), (
    select max(id)
    from projects
), 20.0);
select update_payment((
    select max(id)
    from payments
), 30.0);
select * from all_project_payments((
    select max(id)
    from projects
));
select * from all_user_payments((
    select max(id)
    from users
));
select delete_payment((
    select max(id) - 1
    from payments
));
select * from all_payments();

-- test for comments.sql

select create_comment((
    select max(id)
    from users
), (
    select max(id)
    from projects
), 'hello');
select create_comment((
    select max(id)
    from users
), (
    select max(id)
    from projects
), 'hello');
select update_comment((
    select max(id)
    from comments
), 'bye');
select * from all_project_comments((
    select max(id)
    from projects
));
select * from all_user_comments((
    select max(id)
    from users
));
select delete_comment((
    select max(id) - 1
    from comments
));
select * from all_comments();