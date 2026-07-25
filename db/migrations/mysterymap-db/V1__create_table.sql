CREATE TABLE IF NOT EXISTS sightings
(
    SightingId SERIAL PRIMARY KEY,

    SightingName TEXT NOT NULL,

    SightingLatitude DOUBLE PRECISION NOT NULL,
    SightingLongitude DOUBLE PRECISION NOT NULL,

    SightingDiscoveryDate DATE NOT NULL,

    SightingPicture TEXT NOT NULL,

    SightingDescription TEXT NOT NULL
);