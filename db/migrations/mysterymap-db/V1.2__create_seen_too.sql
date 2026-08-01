CREATE TABLE IF NOT EXISTS SeenToo
(
    SeenId SERIAL PRIMARY KEY,
    post TEXT NOT NULL,
    user_identifier TEXT NOT NULL,
    CONSTRAINT unique_user_post_approval UNIQUE (post, user_identifier)
);