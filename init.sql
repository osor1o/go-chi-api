CREATE TABLE todos (
    id              SERIAL              PRIMARY KEY,
    title           VARCHAR,
    description     TEXT,
    done            BOOL                DEFAULT false
);
