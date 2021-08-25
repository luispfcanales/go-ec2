create table cloud(
user_id serial PRIMARY KEY,
username VARCHAR(50) UNIQUE NOT NULL,
email VARCHAR(50) UNIQUE NOT NULL,
created_on TIMESTAMP NOT NULL,
);

insert into luis.cloud values(0,'luispf','@pf',now());
insert into luis.cloud values(0,'carlos','@carlos',now());
insert into luis.cloud values(0,'maria','@maria',now());
insert into luis.cloud values(0,'juanperez','@juanperez',now());
