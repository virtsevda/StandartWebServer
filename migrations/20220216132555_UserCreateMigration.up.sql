CREATE TABLE users(
    id bigserial not null primary key,
    login varchar not null,
    password varchar not null
) 