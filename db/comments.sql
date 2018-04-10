drop type if exists comment_row cascade;
drop trigger if exists take_log on comments;

create type comment_row as (
    id int,
    user_id int,
    project_id int,
    moment timestamp,
    content text
);

create or replace function all_comments(_num_per_page int, _idx_page int)
returns setof comment_row as $$
declare
    comm comment_row%rowtype;
    comm_row_cursor refcursor;
    i int;
    num_items int;
begin
    insert into logs(content, log_level)
        values ('Select all comments', 1);
    open comm_row_cursor for
        select *
        from comments;
    move absolute (_idx_page - 1) * _num_per_page from comm_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch comm_row_cursor into comm;
        exit when not found;
        return next comm;
    end loop;
    close comm_row_cursor;
    return;
end
$$ language plpgsql;

create or replace function all_project_comments(_project_id int, _num_per_page int, _idx_page int)
returns setof comment_row as $$
declare
    comm comment_row%rowtype;
    comm_row_cursor refcursor;
    i int;
begin
    insert into logs(project_id, content, log_level)
        values (_project_id, 'Select project_s comments', 1);
    open comm_row_cursor for
        select *
        from comments
        where project_id = _project_id;
    move absolute (_idx_page - 1) * _num_per_page from comm_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch comm_row_cursor into comm;
        exit when not found;
        return next comm;
    end loop;
    close comm_row_cursor;
    return;
end
$$ language plpgsql;

create or replace function all_user_comments(_user_id int, _num_per_page int, _idx_page int)
returns setof comment_row as $$
declare
    comm comment_row%rowtype;
    comm_row_cursor refcursor;
    i int;
begin
    insert into logs(user_id, content, log_level)
        values (_user_id, 'Select user_s comments', 1);
    open comm_row_cursor for
        select *
        from comments
        where user_id = _user_id;
    move absolute (_idx_page - 1) * _num_per_page from comm_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch comm_row_cursor into comm;
        exit when not found;
        return next comm;
    end loop;
    close comm_row_cursor;
    return;
end
$$ language plpgsql;

create or replace function create_comment(_user_id int, _project_id int, _content text)
returns integer as $$
    insert into comments(user_id, project_id, content)
        values(_user_id, _project_id, _content);
    select max(id)
        from comments;
$$ language sql;

create or replace function update_comment(_comment_id int, _content text)
returns void as $$
    update comments
        set content = _content
        where id = _comment_id;
$$ language sql;

create or replace function delete_comment(_comment_id int)
returns void as $$
    delete from comments
        where id = _comment_id;
$$ language sql;

create trigger take_log after insert or update or delete on comments
for each row execute procedure create_log_user_proj(' on comments');
