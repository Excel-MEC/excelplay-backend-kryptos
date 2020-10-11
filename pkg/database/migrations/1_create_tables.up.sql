create table if not exists levels (
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

create table if not exists kuser (
id int primary key,
name varchar(100) not null,
curr_level int not null,
last_anstime timestamp
);

create table if not exists answer_logs (
id int references kuser(id),
name varchar(100),
attempt varchar(100),
time timestamp		
)