drop type if exists user_row cascade;
drop trigger if exists take_log on users;

create type user_row as (
    user_id int,
    email citext,
    username citext,
    total_donation numeric(10, 2),
    mobile_number citext,
    address citext,
    occupation citext,
    image citext,
    motto citext,
    is_admin boolean,
    bank_account citext,
    birthday timestamp,
    sex varchar(10)
);

create or replace function all_users()
returns setof user_row as $$
declare
    usr user_row%rowtype;
begin
    insert into logs(content, log_level)
        values ('Select all users', 1);
    for usr in
        select user_id, email, username, total_donation, mobile_number, address, 
            occupation, image, motto, is_admin, bank_account, birthday, sex
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
    select user_id, email, username, total_donation, mobile_number, address, 
            occupation, image, motto, is_admin, bank_account, birthday, sex
        into usr
        from users
        where email = _email and password = crypt(_password, password);
    insert into logs(user_id, content, log_level)
        values (usr.user_id, 'Select user', 1);
    return usr;
end
$$ language plpgsql;

create or replace function create_user(
    _email citext, 
    _password varchar(255),
    _username citext,
    _mobile_number citext,
    _address citext,
    _occupation citext,
    _image citext,
    _motto citext,
    _bank_account citext,
    _birthday timestamp,
    _sex varchar(10))
returns integer as $$
    INSERT INTO users (email, password, username, mobile_number, address, occupation, 
            image, motto, bank_account, birthday, sex)
        VALUES(_email, crypt(_password, gen_salt('bf', 8)), _username, _mobile_number,
            _address, _occupation, _image, _motto, _bank_account, _birthday, _sex);
    select max(user_id)
        from users
$$ language sql;

create trigger take_log after insert or update or delete on users
for each row execute procedure create_log(' on users');
