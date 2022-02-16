CREATE TABLE users(
    id integer not null primary key,
    name varchar not null,
    email varchar not null,
    age integer not null,
    password_hash varchar not null
) 