drop type if exists payments_row;
create type payments_row as (
    id int,
    user_id int,
    project_id int,
    moment timestamp,
    amount numeric
);

create or replace function all_payments()
returns setof payments_row as $$
declare
    pay payments_row%rowtype;
begin
    for pay in
        select *
        from payments
    loop
        return next pay;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function all_project_payments(_project_id int)
returns setof payments_row as $$
declare
    pay payments_row%rowtype;
begin
    for pay in
        select *
        from payments
        where project_id = _project_id
    loop
        return next pay;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function all_user_payments(_user_id int)
returns setof payments_row as $$
declare
    pay payments_row%rowtype;
begin
    for pay in
        select *
        from payments
        where user_id = _user_id
    loop
        return next pay;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function create_payment(_user_id int, _project_id int, _amount numeric)
returns void as $$
    insert into payments(user_id, project_id, amount)
        values(_user_id, _project_id, _amount);
$$ language sql;

create or replace function update_payment(_payment_id int, _amount numeric)
returns void as $$
    update payments
        set amount = _amount
        where id = _payment_id;
$$ language sql;

create or replace function delete_payment(_payment_id int)
returns void as $$
    delete from payments
        where id = _payment_id;
$$ language sql;
