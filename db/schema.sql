-- Creation Order: Users, Albums, Media, Vault, Album Items

create table Users (
    user_id int unsigned auto_increment primary key,
    username varchar(10) not null unique,
    password_hash varchar(255) not null,
    created_at timestamp not null default current_timestamp
);


