-- Creation Order: Users, Albums, Media, Vault, Album Items

create table Users (
    user_id int unsigned auto_increment primary key,
    username varchar(10) not null unique,
    password_hash varchar(255) not null,
    created_at timestamp not null default current_timestamp
);

create table Albums (
    album_id int unsigned auto_increment primary key,
    user_id int unsigned not null, -- defines the column in the table Albums only
    title varchar(10) not null,
    created_at timestamp not null default current_timestamp,
    constraint fk_albums_user
        foreign key (user_id) references User(user_id) -- the user_id in Albums must match a user_id in Users
        on delete cascade -- if a user is deleted, their albums are deleted too
        on update cascade -- if a user_id changes, update it in Albums too (which probably won't happen)
);




