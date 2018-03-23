drop function all_projects;
drop function create_project;
drop function get_project;
drop type proj_row;

create type project_row as (
    id int,
    title varchar(100),
    description text,
    start_time timestamp,
    end_time timestamp,
    amount_required numeric,
    amount_raised numeric,
    verified boolean,
    category citext,
    user_id int
);

create or replace function all_projects()
returns setof project_row as $$
declare
    proj project_row%rowtype;
begin
    for proj in
        WITH total_funds AS (
            SELECT project_id, sum(amount) AS raised
            FROM payments
            GROUP BY project_id
        )
        SELECT p.id, p.title, p.description,
               p.start_time, p.end_time,
               p.amount_required, COALESCE(f.raised, 0) as amount_raised,
               p.verified, p.category, p.user_id
        FROM projects p LEFT OUTER JOIN total_funds f
        ON p.id = f.project_id
    loop
        return next proj;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function create_project(
    _title varchar(100),
    _description text,
    _amount_required numeric,
    _end_time timestamp,
    _category citext,
    _user_id int) 
returns text as $$
    insert into projects (title, description, amount_required, end_time, category, user_id)
        values (_title, _description, _amount_required, _end_time, _category, _user_id);
    select 'Insert OK';
$$ language sql;

create or replace function get_project(_project_id int)
returns project_row as $$
declare
    proj_row project_row%rowtype;
begin
    SELECT p.id, p.title, p.description,
        p.start_time, p.end_time,
        p.amount_required, COALESCE(SUM(f.amount), 0) as amount_raised,
        p.verified, p.category, p.user_id
        into proj_row
        FROM projects p LEFT OUTER JOIN payments f
        ON p.id = f.project_id
        WHERE p.id = $1
        GROUP BY p.id, p.user_id;
    return proj_row;
end
$$ language plpgsql;