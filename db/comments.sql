drop type if exists comment_row;
create type comment_row as (
    id int,
    user_id int,
    project_id int,
    moment timestamp,
    content text
);

create or replace function all_comments()
returns setof comment_row as $$
declare
    comm comment_row%rowtype;
begin
    for comm in
        select *
        from comments
    loop
        return next comm;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function all_project_comments(_project_id int)
returns setof comment_row as $$
declare
    comm comment_row%rowtype;
begin
    for comm in
        select *
        from comments
        where project_id = _project_id
    loop
        return next comm;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function all_user_comments(_user_id int)
returns setof comment_row as $$
declare
    comm comment_row%rowtype;
begin
    for comm in
        select *
        from comments
        where user_id = _user_id
    loop
        return next comm;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function create_comment(_user_id int, _project_id int, _content text)
returns void as $$
    insert into comments(user_id, project_id, content)
        values(_user_id, _project_id, _content);
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
