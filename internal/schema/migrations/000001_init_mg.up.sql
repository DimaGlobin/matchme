CREATE TYPE user_sex AS ENUM ('male', 'female', 'not_defined');
CREATE TYPE user_role AS ENUM ('basic', 'premium', 'admin');

CREATE TABLE users (
    id serial primary key,
    email varchar(255) unique, --TODO:not null
    phone_number varchar(255) not null, --TODO:unique
    name varchar(255) not null,
    password_hash text, --TODO: not null
    sex user_sex default 'not_defined',
    age integer not null,
    birth_date date not null,
    city varchar(255) not null,
    description text,
    role user_role default 'basic',
    max_age integer not null
    -- radius integer not null default 100,
    -- last_ip varchar(255),
    -- latitude float,
    -- longitude float
);