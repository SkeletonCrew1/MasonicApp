INSERT INTO users (UserDisplayName, UserPassword, UserStatus, UserEmail, UserIsInquisitor)
VALUES ('bob', 'pass', 'bronze', 'bob@gmail.com', '1');

INSERT INTO bannedips ( BannedIp)
VALUES ('10.10.10.10.');


INSERT INTO messages ( MessageContext, MessageRecieverStatus, MessageStatus, SenderId,CreatedAt)
VALUES ('bob died',  'gold', 'sent', '1','2013-11-03 00:00:00-07');