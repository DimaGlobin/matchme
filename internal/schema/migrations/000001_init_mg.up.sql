CREATE TYPE user_sex AS ENUM ('male', 'female', 'not_defined');
CREATE TYPE user_role AS ENUM ('basic', 'premium', 'admin');

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password_hash TEXT NOT NULL,
    sex user_sex DEFAULT 'not_defined',
    age INTEGER NOT NULL,
    birth_date DATE NOT NULL,
    city VARCHAR(255) NOT NULL,
    description TEXT,
    role user_role DEFAULT 'basic',
    max_age INTEGER NOT NULL,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- radius integer not null default 100,
    -- last_ip varchar(255),
    -- latitude float,
    -- longitude float
);

CREATE TABLE likes (
    like_id SERIAL PRIMARY KEY, 
    liking_id INTEGER REFERENCES users (user_id),
    liked_id INTEGER REFERENCES users (user_id),
    creation_date TIMESTAMP default CURRENT_TIMESTAMP
);

CREATE TABLE dislikes (
    dislike_id SERIAL PRIMARY KEY,
    disliking_id INTEGER REFERENCES users (user_id),
    disliked_id INTEGER REFERENCES users (user_id),
    creation_date TIMESTAMP default CURRENT_TIMESTAMP
);

CREATE TABLE matches (
    match_id SERIAL PRIMARY KEY,
    user1 INTEGER REFERENCES users (user_id),
    user2 INTEGER REFERENCES users (user_id),
    creation_date TIMESTAMP default CURRENT_TIMESTAMP
);

CREATE TABLE files (
    file_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (user_id),
    file_name VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL,
    upload_date TIMESTAMP default CURRENT_TIMESTAMP
);