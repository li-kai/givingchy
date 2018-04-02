drop type if exists project_row cascade;
drop trigger if exists take_log on projects;

create type project_row as (
    project_id int,
    title varchar(100),
    user_id int,
    category citext,
    description text,
    likes int,
    verified boolean,
    image citext,
    reward citext,
    URL citext,
    people_viewed_number int,
    people_attend_number int,
    bank_info citext,
    completed boolean,
    amount_raised numeric,
    amount_required numeric,
    start_time timestamp,
    end_time timestamp
);

create or replace function all_projects()
returns setof project_row as $$
declare
    proj project_row%rowtype;
begin
    insert into logs(content, log_level)
        values ('Select all projects', 1);
    for proj in
        select *
        from projects
    loop
        return next proj;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function create_project(
    _title varchar(100),
    _user_id int,
    _category citext,
    _description text,
    _image citext,
    _reward citext,
    _URL citext,
    _bank_info citext,
    _amount_required numeric,
    _end_time timestamp)
returns integer as $$
    insert into projects (title, user_id, category, description, image, reward, URL, bank_info,
            amount_required, end_time)
        values (_title, _user_id, _category, _description, _image, _reward, _URL, _bank_info,
            _amount_required, _end_time);
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
    select *
        from projects
        into proj_row
        where project_id = $1;
    return proj_row;
end
$$ language plpgsql;

create trigger take_log after insert or update or delete on projects
for each row execute procedure create_log(' on projects');

create or replace function complete_trigger()
returns trigger as $$
begin
    if (old.amount_raised != new.amount_raised) then
        if (new.amount_raised >= new.amount_required) then
            update projects
                set completed = true
                where project_id = new.project_id;
        else 
            update projects
                set completed = false
                where project_id = new.project_id;
        end if;
    end if;
    return new;
end
$$ language plpgsql;

create trigger complete_proj after update on projects
for each row execute procedure complete_trigger();