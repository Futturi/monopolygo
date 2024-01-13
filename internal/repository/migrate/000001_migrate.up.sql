CREATE TABLE IF NOT EXISTS users
(
    user_id bigserial PRIMARY KEY,
    name varchar,
    username varchar,
    email varchar,
    password_hash varchar,
    refresh_token text,
    refresh_token_expiry bigint not null default 0,
    token text,
    is_email_verified boolean not null default false
);

CREATE TABLE IF NOT EXISTS rooms
(
    room_id bigserial PRIMARY KEY,
    first_player_id int references users (user_id) on delete cascade not null,
    second_player_id int references users (user_id) on delete cascade default 1,
    third_player_id int references users (user_id) on delete cascade default 1,
    fourth_player_id int references users (user_id) on delete cascade default 1
);

CREATE TABLE IF NOT EXISTS users_rooms
(
    id bigserial not null unique,
    user_id int references users (user_id) on delete cascade not null,
    room_id int references rooms (room_id) on delete cascade not null
);