drop function all_users;
drop function get_user;
drop function create_user;
drop type user_row;

create type user_row as (
    id int,
    email citext,
    is_admin boolean
);

create or replace function all_users()
returns setof user_row as $$
declare
    usr user_row%rowtype;
begin
    for usr in
        select *
        from users
    loop
        return next usr;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function get_user(_email citext, _password varchar(255))
returns user_row as $$
declare
    usr user_row%rowtype;
begin
    select *
        into usr
        from users
        where email = _email and password = crypt(_password);
    return usr;
end
$$ language plpgsql;

create or replace function create_user(_email citext, _password varchar(255))
returns text as $$
    INSERT INTO users (email, password)
        VALUES(_email, crypt(_password, gen_salt('bf', 8)));
    select 'create OK';
$$ language sql;