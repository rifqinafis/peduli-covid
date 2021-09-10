SET TIMEZONE = 'Etc/GMT-7';

CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
END;
$$ language 'plpgsql';

create table users(
	id SERIAL not null primary key,
	role_id int not null references roles(id) default 0,
	email text not null default '',
	password text not null default '',
	phone text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp 
);
CREATE TRIGGER users BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into users(id, role_id, email, password, phone, deleted_at) values(0, 0, '', '', '', current_timestamp);

create table roles(
	id SERIAL not null primary key,
	name text not null default '',
	code text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER roles BEFORE UPDATE ON roles FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into roles(id, name, code, deleted_at) values(0, '', '', current_timestamp);

create table payment_methods(
	id SERIAL not null primary key,
	code text not null default '',
	name text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER payment_methods BEFORE UPDATE ON payment_methods FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into payment_methods(id, name, code, deleted_at) values(0, '', '', current_timestamp);

create table reservations(
	id SERIAL not null primary key,
	user_id int not null references users(id) default 0,
	hospital_id int not null default 0,
	bed_type text not null default '',
	status text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER reservations BEFORE UPDATE ON reservations FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into reservations(id, user_id, hospital_id, bed_type, deleted_at) values(0, 0, 0, '', current_timestamp);

create table admins(
	id SERIAL not null primary key,
	role_id int not null references roles(id) default 0,
	hospital_id int not null default 0,
	email text not null default '',
	password text not null default '',
	phone text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER admins BEFORE UPDATE ON admins FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into admins(id, role_id, hospital_id, email, password, phone, deleted_at) values(0, 0, 0, '', '', '', current_timestamp);

create table notifications(
	id SERIAL not null primary key,
	user_id int not null references users(id) default 0,
	admin_id int not null references admins(id) default 0,
	code text not null default '',
	details text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER notifications BEFORE UPDATE ON notifications FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into notifications(id, user_id, admin_id, code, details, deleted_at) values(0, 0, 0, '', '', current_timestamp);

create table payments(
	id SERIAL not null primary key,
	payment_method_id int not null references payment_methods(id) default 0,
	reservation_id int not null references reservations(id) default 0,
	amount numeric(20,3) not null default 0,
	date date not null default now(),
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER payments BEFORE UPDATE ON payments FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into payments(id, payment_method_id, reservation_id, amount, date, deleted_at) values(0, 0, 0, 0, now(), current_timestamp);

create table invoices(
	id SERIAL not null primary key,
	reservation_id int not null references reservations(id) default 0,
	code text not null default '',
	total numeric(20,3) not null default 0,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER invoices BEFORE UPDATE ON invoices FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into invoices(id, reservation_id, code, total, deleted_at) values(0, 0, '', 0, current_timestamp);
