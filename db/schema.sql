-- Creation Order: Users, Media

create table Users (
    user_id int unsigned auto_increment primary key,
    username varchar(50) not null unique,
    password_hash varchar(255) not null,
    created_at timestamp not null default current_timestamp
);

create table Media (
    media_id int unsigned auto_increment primary key,
    user_id int unsigned not null,
    object_key varchar(255) not null,
    media_type enum('photo', 'video') not null,
    created_at timestamp not null default current_timestamp,

    constraint fk_media_user
        foreign key (user_id) references Users(user_id)
        on delete cascade
        on update cascade,

    index idx_media_user_created(user_id, created_at)
);
