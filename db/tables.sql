CREATE TABLE IF NOT EXISTS serials (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    season INTEGER NOT NULL,
    episode INTEGER NOT NULL,
    resolution TEXT NOT NULL
);
