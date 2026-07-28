create table if not exists WhiteList (
    InvitationId serial primary key ,
    InvitedEmail varchar(40) not null
);
