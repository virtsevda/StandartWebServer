CREATE TABLE articles(
    id bigserial not null primary key,
    title VARCHAR not null unique,
    author VARCHAR not null,
    content VARCHAR not null
)