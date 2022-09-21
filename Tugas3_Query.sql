create schema if not exists efishery_academy;

create table if not exists postgres.efishery_academy.customers(
	id int primary key unique not null,
	customer_name char(50) not null
);

insert into efishery_academy.customers (id,customer_name) values (1,'Zulfi'),(2,'Asep'),(3,'Ujang'),(4,'Mamat');

create table if not exists postgres.efishery_academy.product(
	id int primary key unique not null,
	product_name char(50) not null
	price float not null
);

insert into efishery_academy.product (id,product_name,price) values (1,'feeder ikan',2500000),(2,'feeder udang',2000000);

create table if not exists postgres.efishery_academy.orders(
	id int primary key unique not null,
	customer_id int not null REFERENCES efishery_academy.customers (id),
	product_id int  not null REFERENCES efishery_academy.product (id),
	order_date date not null
	quantity int not null
);

insert into efishery_academy.orders (id,customer_id,product_id,order_date,quantity) values (1,1,2,now(),1),(2,2,1,now(),1),(3,4,1,now(),2);

select o.id, o.order_date, c.customer_name , p.product_name , p.price, o.quantity, price*quantity as total
from efishery_academy.orders o
inner join efishery_academy.customers c on o.customer_id = c.id 
right join efishery_academy.product p on o.product_id = p.id;


