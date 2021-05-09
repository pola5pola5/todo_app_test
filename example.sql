drop table sessions;
drop table users;
drop table todos;

create table users (
    id         serial primary key,
    uuid       varcher(64) not null unique,
    name       varcher(255),
    email      varcher(255) not null unique,
    password   varcher(255) not null,
    created_at timestamp not null
);

create table sessions (
    id         serial primary key,
    uuid       varcher(64) not null unique,
    name       varcher(255),
    user_id    integer references users(id),
    created_at timestamp not null
);

create table todos (
    id         serial primary key,
    uuid       varcher(64) not null unique,
    content    text,
    user_id    integer references users(id),
    created_at timestamp not null
);