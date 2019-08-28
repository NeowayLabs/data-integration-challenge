--CREATE USER dci;
--CREATE DATABASE dci;
--GRANT ALL PRIVILEGES ON DATABASE dci TO dci;

create table company (
    id bigserial ,
    name varchar(200),
    zipCode varchar(5)
);

alter table company add constraint PK_COMPANY primary key (id);