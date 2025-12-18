-- Creation Order: Users, Albums, Media, Vault, Album Items

create table Users (
    user_id int unsigned auto_increment primary key,
    username varchar(50) not null unique,
    password_hash varchar(255) not null,
    vault_pin_hash varchar(255) null, -- optional
    created_at timestamp not null default current_timestamp
);

create table Albums (
    album_id int unsigned auto_increment primary key,
    user_id int unsigned not null, -- defines the column in the table Albums only
    title varchar(100) not null,
    created_at timestamp not null default current_timestamp,

    constraint fk_albums_user
        foreign key (user_id) references Users(user_id) -- the user_id in Albums must match a user_id in Users
        on delete cascade -- if a user is deleted, their albums are deleted too
        on update cascade, -- if a user_id changes, update it in Albums too (which probably won't happen)

    index idx_albums_user_created (user_id, created_at) -- pre-organize to make searching faster
);

create table Media (
    media_id int unsigned auto_increment primary key,
    user_id int unsigned not null,
    storage_path varchar(255) not null,
    media_type enum('photo', 'video') not null,
    created_at timestamp not null default current_timestamp,

    constraint fk_media_user
        foreign key (user_id) references Users(user_id)
        on delete cascade
        on update cascade,

    index idx_media_user_created(user_id, created_at)
);

create table Vault (
    vault_id int unsigned auto_increment primary key,
    user_id int unsigned not null,
    media_id int unsigned not null,
    created_at timestamp not null default current_timestamp,

    constraint fk_vault_user
        foreign key (user_id) references Users(user_id)
        on delete cascade
        on update cascade,
    constraint fk_vault_media
        foreign key (media_id) references Media(media_id)
        on delete cascade
        on update cascade,
    constraint uq_vault_media unique (media_id),

    index idx_vault_user_created (user_id, created_at)
);

create table AlbumItems (
    album_id int unsigned not null,
    media_id int unsigned not null,
    created_at timestamp not null default current_timestamp,

    primary key (album_id, media_id), -- this combination must be unique (no duplicate media in the same album)
    constraint fk_albumitems_album
        foreign key (album_id) references Albums(album_id)
        on delete cascade
        on update cascade,
    constraint fk_albumitems_media
        foreign key (media_id) references Media(media_id)
        on delete cascade
        on update cascade,
    
    index idx_albumitems_media (media_id)
);



