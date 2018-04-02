drop type if exists log_row cascade;
create type log_row as (
    id int,
    user_id int,
    project_id int,
    moment timestamp,
    content text,
    log_level int
);

create or replace function all_logs() 
returns setof log_row as $$
declare 
    lrow log_row%rowtype;
begin
    for lrow in
        select *
        from logs
    loop
        return next lrow;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function all_user_logs(_user_id int) 
returns setof log_row as $$
declare
    lrow log_row%rowtype;
begin
    for lrow in
        select *
        from logs
        where user_id = _user_id
    loop
        return next lrow;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function all_project_logs(_project_id int) 
returns setof log_row as $$
declare
    lrow log_row%rowtype;
begin
    for lrow in
        select *
        from logs
        where project_id = _project_id
    loop
        return next lrow;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function create_log()
returns trigger as $$
begin
    if (tg_table_name = 'users') then
        if (tg_op = 'DELETE') then
            insert into logs(user_id, content, log_level)
                values (old.user_id, 'Delete op' || tg_argv[0], 0);
            return old;
        elsif (tg_op = 'UPDATE') then
            insert into logs(user_id, content, log_level)
                values (new.user_id, 'Update op' || tg_argv[0], 0);
            return new;
        elsif (tg_op = 'INSERT') then
            insert into logs(user_id, content, log_level)
                values (new.user_id, 'Insert op' || tg_argv[0], 0);
            return new;
        end if;
    elsif (tg_table_name = 'categories') then
        if (tg_op = 'DELETE') then
            insert into logs(content, log_level)
                values ('Delete op' || tg_argv[0], 0);
            return old;
        elsif (tg_op = 'UPDATE') then
            insert into logs(content, log_level)
                values ('Update op' || tg_argv[0], 0);
            return new;
        elsif (tg_op = 'INSERT') then
            insert into logs(content, log_level)
                values ('Insert op' || tg_argv[0], 0);
            return new;
        end if;
    else 
        if (tg_op = 'DELETE') then
            insert into logs(user_id, project_id, content, log_level)
                values (old.user_id, old.project_id, 'Delete op' || tg_argv[0], 0);
            return old;
        elsif (tg_op = 'UPDATE') then
            insert into logs(user_id, project_id, content, log_level)
                values (new.user_id, new.project_id, 'Update op' || tg_argv[0], 0);
            return new;
        elsif (tg_op = 'INSERT') then
            insert into logs(user_id, project_id, content, log_level)
                values (new.user_id, new.project_id, 'Insert op' || tg_argv[0], 0);
            return new;
        end if;
    end if;
    return null;
end
$$ language plpgsql;