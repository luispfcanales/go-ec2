create table cloud(
user_id serial PRIMARY KEY,
username VARCHAR(50) UNIQUE NOT NULL,
email VARCHAR(50) UNIQUE NOT NULL,
created_on TIMESTAMP NOT NULL
);

insert into cloud values(1,'luispf','@pf',now());
insert into cloud values(2,'carlos','@carlos',now());
insert into cloud values(3,'maria','@maria',now());
insert into cloud values(4,'juanperez','@juanperez',now());
