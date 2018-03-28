drop type if exists cate_row;
create type cate_row as (
    name citext
);

create or replace function all_categories()
returns setof cate_row as $$
declare
    cate cate_row%rowtype;
begin
    for cate in
        select name
        from categories
    loop
        return next cate;
    end loop;
    return;
end
$$ language plpgsql;

create or replace function create_categories(name citext)
returns void as $$
    insert into categories
        values(name);
$$ language sql;
