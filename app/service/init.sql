create database xreport;

create table service (
     id serial unique,
     title varchar not null unique,
     smtp_ip varchar not null,
     smtp_port int not null,
     smtp_email varchar not null,
     smtp_password varchar not null,
     created_at timestamp default now()
);

create table driver (
    id serial unique,
    title varchar not null unique,
    code varchar not null,
    created_at timestamp default now()
);

create table source (
    id serial unique,
    driver_id int references driver(id),
    title varchar not null unique,
    host varchar not null,
    port int not null,
    username varchar not null,
    password varchar not null,
    database varchar not null,
    created_at timestamp default now()
);

create table report (
    id serial unique,
    service_id int references service(id),
    enabled bool default true,
    title varchar not null unique,
    send_time varchar not null,
    recipients varchar[],
    source_id integer references source(id),
    data_query text,
    template text,
    xlsx bool default false,
    created_at timestamp default now()
);

