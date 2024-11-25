create table carts(
	cart_id serial,
	customer_id varchar(5),
	product_id integer,
	unit_price real,
	qty integer,
	cart_created_on date,
	constraint cart_cust_id_pk primary key (cart_id,customer_id,product_id),
	constraint cart_customer_fk foreign key (customer_id) references customers(customer_id),
	constraint cart_product_fk foreign key (product_id) references products (product_id)
)