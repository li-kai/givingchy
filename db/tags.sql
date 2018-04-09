drop type if exists tag_row cascade;
drop trigger if exists take_log on tags;

create type tag_row as (
    project_id int,
    tag_name citext
);

create or replace function all_tags()
returns setof citext as $$
declare 
    _tag_name citext;
begin
    insert into logs(content, log_level)
        values ('Select all tags', 1);
    for _tag_name in
        select distinct tag_name
        from tags
    loop
        return next _tag_name;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function create_tag(
    _project_id int,
    _tag_name citext)
returns citext as $$
    insert into tags (project_id, tag_name)
        values (_project_id, _tag_name);
    select _tag_name;
$$ language sql;

create or replace function get_project_s_tags(
    _project_id int)
returns setof citext as $$
declare 
    _tag_name citext;
begin
    insert into logs(project_id, content, log_level)
        values (_project_id, 'Select project_s tags', 1);
    for _tag_name in
        select tag_name
        from tags
        where project_id = _project_id
    loop
        return next _tag_name;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function get_tag_s_projects(
    _tag_name citext)
returns setof int as $$
declare
    _project_id int;
begin
    insert into logs(content, log_level)
        values ('Select tag_s projects', 1);
    for _project_id in
        select project_id
        from tags
        where tag_name = _tag_name
    loop
        return next _project_id;
    end loop;
    return;
end
$$ language plpgsql;

create trigger take_log after insert or update or delete on tags
for each row execute procedure create_log_proj(' on tags');