CREATE TABLE IF NOT EXISTS users
(
    user_id bigserial PRIMARY KEY,
    name varchar,
    username varchar,
    email varchar,
    password_hash varchar
);

CREATE TABLE IF NOT EXISTS rooms
(
    room_id bigserial PRIMARY KEY,
    max_users int
);

CREATE TABLE IF NOT EXISTS users_rooms
(
    id bigserial not null unique,
    user_id int references users (user_id) on delete cascade not null,
    room_id int references rooms (room_id) on delete cascade not null
);