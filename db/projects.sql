drop type if exists project_row cascade;
drop trigger if exists take_log on projects;
drop trigger if exists update_cate on projects;

create type project_row as (
    project_id int,
    title varchar(100),
    user_id int,
    category citext,
    description text,
    verified boolean,
    tags citext[],
    image citext,
    amount_raised numeric,
    amount_required numeric,
    start_time timestamp,
    end_time timestamp
);

create or replace function all_projects(_num_per_page int, _idx_page int)
returns setof project_row as $$
declare
    proj project_row%rowtype;
    proj_row_cursor refcursor;
    i int;
begin
    insert into logs(content, log_level)
        values ('Select all projects', 1);
    open proj_row_cursor for 
        select 
            p.project_id,
            p.title,
            p.user_id,
            p.category,
            p.description,
            p.verified,
            array_remove(array_agg(t.tag_name), NULL) as tags,
            p.image,
            p.amount_raised,
            p.amount_required,
            p.start_time,
            p.end_time
        from projects p
        left join tags t on p.project_id = t.project_id
        group by p.project_id;
    move absolute (_idx_page - 1) * _num_per_page from proj_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch proj_row_cursor into proj;
        exit when not found;
        return next proj;
    end loop;
    close proj_row_cursor;
    return;
end
$$ language plpgsql;

create or replace function search_projects(_keyword citext, _num_per_page int, _idx_page int)
returns setof project_row as $$
declare
    proj project_row%rowtype;
    proj_row_cursor refcursor;
    i int;
begin
    insert into logs(content, log_level)
        values ('Search projects', 1);
    open proj_row_cursor for 
        select *
        from projects
        where title like '%' || _keyword || '%';
    move absolute (_idx_page - 1) * _num_per_page from proj_row_cursor;
    i := 0;
    loop
        if i >= _num_per_page then
            exit;
        end if;
        i := i + 1;
        fetch proj_row_cursor into proj;
        exit when not found;
        return next proj;
    end loop;
    close proj_row_cursor;
    return;
end
$$ language plpgsql;


create or replace function create_project(
    _title varchar(100),
    _user_id int,
    _category citext,
    _description text,
    _image citext,
    _amount_required numeric,
    _end_time timestamp)
returns integer as $$
    insert into projects (title, user_id, category, description, image, amount_required, end_time)
        values (_title, _user_id, _category, _description, _image, _amount_required, _end_time);
    select max(project_id)
        from projects
$$ language sql;

create or replace function get_project(_project_id int)
returns project_row as $$
declare
    proj_row project_row%rowtype;
begin
    insert into logs(project_id, content, log_level)
        values (_project_id, 'Select project', 1);
    select 
            p.project_id,
            p.title,
            p.user_id,
            p.category,
            p.description,
            p.verified,
            array_remove(array_agg(t.tag_name), NULL) as tags,
            p.image,
            p.amount_raised,
            p.amount_required,
            p.start_time,
            p.end_time
        from projects p
        into proj_row
        left join tags t on p.project_id = t.project_id
        group by p.project_id
        having p.project_id = _project_id;
    return proj_row;
end
$$ language plpgsql;

create or replace function update_cate_num()
returns trigger as $$
begin
    if (tg_op = 'DELETE') then
        update categories
            set proj_num = proj_num - 1
            where name = old.category;
        return old;
    elsif (tg_op = 'UPDATE') then
        update categories
            set proj_num = proj_num - 1
            where name = old.category;
        update categories
            set proj_num = proj_num + 1
            where name = new.category;
        return new;
    elsif (tg_op = 'INSERT') then
        update categories
            set proj_num = proj_num + 1
            where name = new.category;
        return new;
    end if;
    return null;
end
$$ language plpgsql;

create trigger update_cate after insert or delete on projects
for each row execute procedure update_cate_num();

create trigger take_log after insert or update or delete on projects
for each row execute procedure create_log_user_proj(' on projects');