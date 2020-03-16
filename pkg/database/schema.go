package database

var schema = `create table if not exists levels (
number int not null primary key,
question varchar(1000),
image_level boolean not null,
level_file varchar(1000),
answer varchar(100)
);

create table if not exists hints (
id serial primary key,
number int references levels(number) not null,
content varchar(2000)
);

create extension if not exists "uuid-ossp";

create table if not exists kuser (
id uuid primary key default uuid_generate_v1(),
name varchar(100) not null,
curr_level int references levels(number) not null,
last_anstime timestamp
);

create table if not exists answer_logs (
id uuid references kuser(id),
name varchar(100),
attempt varchar(100),
time timestamp		
)
`
