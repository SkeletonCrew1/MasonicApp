CREATE TABLE IF NOT EXISTS sightings
(
    id SERIAL PRIMARY KEY,

    name TEXT NOT NULL,

    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,

    discovery_date DATE NOT NULL,

    picture TEXT NOT NULL,

    description TEXT NOT NULL
);