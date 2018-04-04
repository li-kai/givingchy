drop type if exists payments_row cascade;
drop trigger if exists take_log on payments;
drop trigger if exists donate_trigger on payments;

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
    insert into logs(content, log_level)
        values ('Select all payments', 1);
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
    insert into logs(project_id, content, log_level)
        values (_project_id, 'Select project_s payments', 1);
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
    insert into logs(project_id, content, log_level)
        values (_user_id, 'Select uesr_s payments', 1);
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
returns integer as $$
    insert into payments(user_id, project_id, amount)
        values(_user_id, _project_id, _amount);
    select max(id) 
        from payments;
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

create trigger take_log after insert or update or delete on payments
for each row execute procedure create_log_user_proj(' on payments');

create or replace function donate_trigger()
returns trigger as $$
begin
    if (tg_op = 'DELETE') then
        update users
            set total_donation = total_donation - old.amount
            where user_id = old.user_id;
        update projects
            set amount_raised = amount_raised - old.amount
            where project_id = old.project_id;
        return old;
    elsif (tg_op = 'UPDATE') then
        update users
            set total_donation = total_donation - old.amount
            where user_id = old.user_id;
        update projects
            set amount_raised = amount_raised - old.amount
            where project_id = old.project_id;
        update users
            set total_donation = total_donation + new.amount
            where user_id = new.user_id;
        update projects
            set amount_raised = amount_raised + new.amount
            where project_id = new.project_id;
        return new;
    elsif (tg_op = 'INSERT') then
        update users
            set total_donation = total_donation + new.amount
            where user_id = new.user_id;
        update projects
            set amount_raised = amount_raised + new.amount
            where project_id = new.project_id;
        return new;
    end if;
end;
$$ language plpgsql;

create trigger donate after insert or update or delete on payments
for each row execute procedure donate_trigger();