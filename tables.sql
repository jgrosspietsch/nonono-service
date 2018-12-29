CREATE TABLE IF NOT EXISTS puzzle (
    id integer GENERATED ALWAYS AS IDENTITY,
    height smallint NOT NULL CHECK (height > 0 AND height % 5 = 0), -- height of the puzzle, constrained to always be a multiple of 5
    width smallint NOT NULL CHECK (width > 0 AND width % 5 = 0), -- width of the puzzle, constrained to always be a multiple of 5
    grid smallint[][] NOT NULL,
    hash bytea,
    created timestamp without time zone default (now() at time zone 'utc'),
    PRIMARY KEY (id),
    UNIQUE (hash)
);

-- CREATE TABLE IF NOT EXISTS progress_profile (
    
-- );

-- CREATE TABLE IF NOT EXISTS saved_progress (

-- );
