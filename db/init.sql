CREATE DATABASE users WITH TEMPLATE = template0
    ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';

\connect users

CREATE TABLE users (
    id serial,
    name varchar(128) not null,
    email varchar(128) not null,
    password varchar(128) not null,

    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
