create table users(
    user_id serial,
    user_name varchar(25),
    user_password varchar(255),
    user_phone varchar(15),
    user_role varchar(15),
    user_token varchar(255),
    constraint user_id_pk primary key(user_id),
    constraint user_name_uq unique(user_name),
    constraint user_phone_uq unique(user_phone)
);

create table roles(
    role_id serial,
    role_name varchar(15),
    constraint role_id_pk primary key(role_id),
    constraint role_name_uq unique(role_name)
);

create table user_roles(
	usro_user_id int,
	usro_role_id int,
	constraint usro_user_role_pk primary key(usro_user_id,usro_role_id),
	constraint usro_user_id_fk foreign key (usro_user_id) references users(user_id),
	constraint usro_role_id_fk foreign key (usro_role_id) references roles(role_id)
);