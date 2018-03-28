drop type if exists user_row cascade;
drop trigger if exists take_log on users;

create type user_row as (
    user_id int,
    email citext,
    is_admin boolean
);

create or replace function all_users()
returns setof user_row as $$
declare
    usr user_row%rowtype;
begin
    for usr in
        select user_id, email, is_admin
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
    select user_id, email, is_admin
        into usr
        from users
        where email = _email and password = crypt(_password, password);
    return usr;
end
$$ language plpgsql;

create or replace function create_user(_email citext, _password varchar(255))
returns void as $$
    INSERT INTO users (email, password)
        VALUES(_email, crypt(_password, gen_salt('bf', 8)));
$$ language sql;

create trigger take_log after insert or update or delete on users
for each row execute procedure create_log(' on users');
