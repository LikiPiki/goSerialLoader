CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    season INTEGER NOT NULL,
    episode INTEGER NOT NULL,
    resolution TEXT NOT NULL
);
