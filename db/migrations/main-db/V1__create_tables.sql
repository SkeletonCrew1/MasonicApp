create table if not exists Users (
    UserId serial primary key,
    UserDisplayName varchar(40) not null,
    UserPassword text not null,
    UserStatus varchar(40) not null,
    UserEmail text not null,
    UserIsInquisitor bool not null
);

create table if not exists BannedIps (
    IpId serial primary key ,
    BannedIp varchar(40) not null
);

create table if not exists Messages (
    MessageId serial primary key not null,
    MessageContext text not null,
    MessageRecieverStatus varchar(40) not null,
    MessageStatus varchar(20) not null,
    SenderId int not null references Users(UserId) not null,
    CreatedAt timestamp with time zone not null
);
