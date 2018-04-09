drop type if exists payments_row cascade;
drop trigger if exists take_log on payments;
drop trigger if exists donate on payments;

create type payments_row as (
    id int,
    user_id int,
    project_id int,
    moment timestamp,
    amount numeric
);

create or replace function all_payments(_num_per_page int, _idx_page int)
returns setof payments_row as $$
declare
    pay payments_row%rowtype;
    pay_row_cursor refcursor;
    i int;
begin
    insert into logs(content, log_level)
        values ('Select all payments', 1);
    open pay_row_cursor for
        select *
        from payments;
    move absolute (_idx_page - 1) * _num_per_page from pay_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch pay_row_cursor into pay;
        return next pay;
    end loop;
    close pay_row_cursor;
    return;
end
$$ language plpgsql;

create or replace function all_project_payments(_project_id int, _num_per_page int, _idx_page int)
returns setof payments_row as $$
declare
    pay payments_row%rowtype;
    pay_row_cursor refcursor;
    i int;
begin
    insert into logs(project_id, content, log_level)
        values (_project_id, 'Select project_s payments', 1);
    open pay_row_cursor for
        select *
        from payments
        where project_id = _project_id;
    move absolute (_idx_page - 1) * _num_per_page from pay_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch pay_row_cursor into pay;
        return next pay;
    end loop;
    close pay_row_cursor;
    return;
end
$$ language plpgsql;

create or replace function all_user_payments(_user_id int, _num_per_page int, _idx_page int)
returns setof payments_row as $$
declare
    pay payments_row%rowtype;
    pay_row_cursor refcursor;
    i int;
begin
    insert into logs(user_id, content, log_level)
        values (_user_id, 'Select all payments', 1);
    open pay_row_cursor for
        select *
        from payments
        where user_id = _user_id;
    move absolute (_idx_page - 1) * _num_per_page from pay_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch pay_row_cursor into pay;
        return next pay;
    end loop;
    close pay_row_cursor;
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