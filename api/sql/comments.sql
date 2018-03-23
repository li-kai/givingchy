-- Sql functions and triggers for comments

---- all_comments()
-- usage : select * from all_comments();

---- all_project_comments(_project_id int)
-- usage : select * from all_project_comments(1);

---- all_user_comments(_user_id int);
-- usage : select * from all_user_comments(1);

---- create_comment(_user_id int, _project_id int, _content text)
-- usage : select create_comment(1, 1, 'hello world');

---- update_comment(_comment_id int, _content text)
-- usage : select update_comment(1, 'hello world');

---- delete_comment(_comment_id int)
-- usage : select delete_comment(1);

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
returns text as $$
    insert into comments
        values((
            select count(*) + 1
            from comments
        ), _user_id, _project_id, CURRENT_TIMESTAMP, _content);
    select 'Insert OK';
$$ language sql;

create or replace function update_comment(_comment_id int, _content text)
returns text as $$
    update comments
        set content = _content
        where id = _comment_id;
    select 'Update OK';
$$ language sql;

create or replace function delete_comment(_comment_id int)
returns text as $$
    delete from comments
        where id = _comment_id;
    select 'Delete OK';
$$ language sql;