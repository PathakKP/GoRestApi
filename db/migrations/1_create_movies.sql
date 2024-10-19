-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS movies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Unique ID for each movie
    title VARCHAR(255) NOT NULL,                    -- Movie title
    category VARCHAR(100) NOT NULL,                 -- Movie category (genre)
    year VARCHAR(10) NOT NULL,                      -- Year of release
    imdb_rating VARCHAR(10) NOT NULL                -- IMDb rating
);

-- +migrate Down
-- DROP TABLE IF EXISTS movies;