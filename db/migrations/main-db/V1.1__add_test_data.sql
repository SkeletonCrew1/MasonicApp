INSERT INTO users (UserDisplayName, UserPassword, UserStatus, UserEmail, UserIsInquisitor)
VALUES ('roman1', 'pass', 'golden', 'roro1@gmail.com', '0');

INSERT INTO users (UserDisplayName, UserPassword, UserStatus, UserEmail, UserIsInquisitor)
VALUES ('roman2', 'pass', 'golden', 'roro2@gmail.com', '0');


INSERT INTO bannedips ( BannedIp)
VALUES ('10.10.10.10.');


INSERT INTO messages ( MessageContext, MessageRecieverStatus, MessageStatus, SenderId,CreatedAt)
VALUES ('bob died',  'gold', 'sent', '1','2013-11-03 00:00:00-07');