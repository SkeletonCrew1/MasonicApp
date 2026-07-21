
\connect auth_service;

create table users (
  UserId serial primary key,
  UserFakename varchar(30) not null,
  UserPassword text not null,
  UserStatus varchar(30),
  UserEmail text not null
);
