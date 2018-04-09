delete from logs cascade;
delete from payments cascade;
delete from comments cascade;
delete from tags cascade;
delete from projects cascade;
delete from categories cascade;
delete from users cascade;

-- test for users.sql

select create_user('1@gmail.com', '123', '1', '1');
select create_user('2@gmail.com', '1234', '2', '2');
select create_user('3@gmail.com', '12345', '3', '3');
select create_user('4@gmail.com', '123', '1', '1');
select create_user('5@gmail.com', '1234', '2', '2');
select create_user('6@gmail.com', '12345', '3', '3');
select create_user('7@gmail.com', '123', '1', '1');
select create_user('8@gmail.com', '1234', '2', '2');
select create_user('9@gmail.com', '12345', '3', '3');
select * from all_users(8, 2);
select * from get_user('1@gmail.com', '123');
delete from users where email = '3@gmail.com';

-- test for categories.sql

select create_categories('Art');
select create_categories('Science');
select * from all_categories();

-- test for projects.sql

select create_project('hello1', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('bye2', (
    select max(user_id) - 1
    from users
), 'Art', 'bye', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('hello3', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('hello4', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('hello5', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('hello6', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('hello7', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select create_project('hello8', (
    select max(user_id) - 1
    from users
), 'Art', 'hello', '1', 100.0, '2018-08-20 14:52:49'::timestamp);
select * from all_projects(2, 3);
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
), 2, 1);
select * from all_user_payments((
    select max(user_id)
    from users
), 2, 1);
select delete_payment((
    select max(id) - 1
    from payments
));
select * from all_payments(1, 2);

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
), 2, 1);
select * from all_user_comments((
    select max(user_id)
    from users
), 2, 1);
select delete_comment((
    select max(id) - 1
    from comments
));
select * from all_comments(1, 2);

-- test for tags.sql
select create_tag((
    select max(project_id)
    from projects
), 'A');
select create_tag((
    select max(project_id)
    from projects
), 'B');
select create_tag((
    select max(project_id) - 1
    from projects
), 'A');
select * from all_tags();
select * from get_project_s_tags((
    select max(project_id)
    from projects
));
select * from get_tag_s_projects('A');