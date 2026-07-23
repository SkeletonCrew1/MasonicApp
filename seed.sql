
\connect auth-service;

create table Users (
    UserId serial primary key,
    UserDisplayName varchar(40) not null,
    UserPassword text not null,
    UserStatus varchar(40) not null,
    UserEmail text not null,
    UserIsInquisitor bool not null
);

create table BannedIps (
    UserID serial primary key references Users(UserID) not null,
    BannedIp varchar(40) not null
);

create table Messages (
    MessageId serial primary key not null,
    MessageContext text not null,
    MessageRecieverEmail varchar(40) not null,
    MessageStatus varchar(20) not null,
    SenderId int not null references Users(UserId) not null,
    CreatedAt timestamp with time zone not null
);