drop type if exists cate_row cascade;
drop trigger if exists take_log on categories;

create type cate_row as (
    name citext
);

create or replace function all_categories()
returns setof cate_row as $$
declare
    cate cate_row%rowtype;
begin
    insert into logs(content, log_level)
        values ('Select all categories', 1);
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

create trigger take_log after insert or update or delete on categories
for each row execute procedure create_log(' on categories');
