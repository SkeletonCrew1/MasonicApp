create table if not exists Votings (
    VotingId serial primary key,
    VotingSubject varchar(40) not null,
    VotingCategory varchar(40) not null
);

create table if not exists Votes (
    VotingId int REFERENCES Votings(VotingId) ON DELETE CASCADE,
    VoterId int not null,
    PRIMARY KEY (VotingId, VoterId)
);