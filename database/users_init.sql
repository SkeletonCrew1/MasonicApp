CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL,
    email TEXT NOT NULL,
    is_inquisitor BOOL NOT NULL DEFAULT FALSE
);

INSERT INTO users (login, status, email, is_inquisitor) VALUES
('roman1', 'golden', 'romanmissuryn@gmail.com', 'false'),
('roman2', 'silver', 'romanmissuryn20@gmail.com', 'false'),
('roman3', 'bronze', 'ntwrkghostany1@gmail.com', 'false')
ON CONFLICT (login) DO NOTHING;