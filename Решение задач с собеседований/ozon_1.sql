--
-- ПОДГОТОВКА ТАБЛИЦ
--
create table "user"
(
    id        integer,
    firstname varchar,
    lastname  varchar,
    birth     date
);

alter table "user"
    owner to postgres;

create table purchase
(
    sku     integer,
    price   integer,
    user_id integer,
    date    date
);

alter table purchase
    owner to postgres;

create table ban_list
(
    user_id   integer,
    date_from date
);

alter table ban_list
    owner to postgres;

insert into ban_list values (1, 2021-03-08);
insert into ban_list values (1, 2021-03-08);
insert into purchase values (1,5500,1,2021-02-15);
insert into purchase values (1,5700,1,2021-01-15);
insert into purchase values (3,8000,2,2021-03-15);
insert into purchase values (4,400,2,2021-03-02);
insert into purchase values (2,4000,3,2021-02-15);
insert into "user" values (1,'ivan','petrov',1996-05-01);
insert into "user" values (2,'anna','petrova',1999-06-01);
insert into "user" values (4,'abba','petrova','1990-10-02');

--
--
--

-- 1. Вывести уникальные комбинации пользователя
-- и id товара для всех покупок,
-- совершенных пользователями до того, как их забанили.
-- Отсортировать сначала по имени пользователя,
-- потом по SKU
-- 2. Найти пользователей,
-- которые совершили покупок на сумму больше 5000р.
-- Вывести их имена в формате
-- id пользователя | имя | фамилия | сумма покупок

select u.id, u.firstname, p.sku
from "user" u
         inner join purchase p on u.id = p.user_id
where u.id not in (
    select bl.user_id from ban_list bl
                               inner join purchase p on p.user_id = bl.user_id
    where p.date < bl.date_from)
order by u.firstname asc,
         p.sku asc;


select u.id as id, u.firstname as name, u.lastname as lastname, sum(p.price) as sum_price
from "user" u
         inner join purchase p on u.id = p.user_id
GROUP BY
    u.id, u.firstname, u.lastname
HAVING
    SUM(p.price) > 5000;