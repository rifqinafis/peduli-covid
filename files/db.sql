SET TIMEZONE = 'Etc/GMT-7';

CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
END;
$$ language 'plpgsql';

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


create table paymentmethods(
	id SERIAL not null primary key,
	code text not null default '',
	name text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER paymentmethods BEFORE UPDATE ON paymentmethods FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into paymentmethods(id, name, code, deleted_at) values(0, '', '', current_timestamp);


create table provinces(
	id SERIAL not null primary key,
	code text not null default '',
	name text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER provinces BEFORE UPDATE ON provinces FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into provinces(id, code, name, deleted_at) values(0, '', '', current_timestamp);

create table cities(
	id SERIAL not null primary key,
	province_id int not null references provinces(id) default 0,
	code text not null default '',
	name text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER cities BEFORE UPDATE ON cities FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into cities(id, province_id, code, name, deleted_at) values(0, 0, '', '', current_timestamp);

create table hospitals(
	id SERIAL not null primary key,
	city_id int not null references cities(id) default 0,
	name text not null default '',
	address text not null default '',
	phone text not null default '',
	total_bed_available int not null default 0,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER hospitals BEFORE UPDATE ON hospitals FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into hospitals(id, city_id, name, address, phone, total_bed_available, deleted_at) values(0, 0, '', '', '', 0, current_timestamp);

create table bedtypes(
	id SERIAL not null primary key,
	hospital_id int not null references hospitals(id) default 0,
	name text not null default '',
	bed_available int not null default 0,
	bed_empty int not null default 0,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER bedtypes BEFORE UPDATE ON bedtypes FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into bedtypes(id, name, bed_available, bed_empty, deleted_at) values(0, '', 0, 0, current_timestamp);

create table reservations(
	id SERIAL not null primary key,
	user_id int not null references users(id) default 0,
	hospital_id int not null references hospitals(id) default 0,
	bedtype_id int not null references bedtypes(id) default 0,
	status text not null default '',
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER reservations BEFORE UPDATE ON reservations FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into reservations(id, user_id, hospital_id, bed_type_id, deleted_at) values(0, 0, 0, 0, current_timestamp);

create table payments(
	id SERIAL not null primary key,
	paymentmethod_id int not null references paymentmethods(id) default 0,
	reservation_id int not null references reservations(id) default 0,
	amount numeric(20,3) not null default 0,
	date date not null default now(),
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp,
	deleted_at timestamp
);
CREATE TRIGGER payments BEFORE UPDATE ON payments FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
insert into payments(id, paymentmethod_id, reservation_id, amount, date, deleted_at) values(0, 0, 0, 0, now(), current_timestamp);

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

create table admins(
	id SERIAL not null primary key,
	role_id int not null references roles(id) default 0,
	hospital_id int not null references hospitals(id) default 0,
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

ALTER TABLE provinces ADD UNIQUE (code);
ALTER TABLE cities ADD COLUMN province_code VARCHAR(8) NOT NULL DEFAULT '';
ALTER TABLE cities ADD CONSTRAINT fk_province_code FOREIGN KEY (province_code) REFERENCES provinces(code);
alter table bed_types add column hospital_id int not null references hospitals(id) default 0;
ALTER TABLE users ADD COLUMN hospital_id int NOT NULL DEFAULT 0;
ALTER TABLE users ADD CONSTRAINT fk_hospital_id FOREIGN KEY (hospital_id) REFERENCES hospitals(id);