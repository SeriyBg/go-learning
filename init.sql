CREATE USER goprojects WITH PASSWORD 'goprojects';

CREATE DATABASE goprojects;

GRANT ALL PRIVILEGES ON DATABASE goprojects to goprojects;

CREATE TABLE IF NOT EXISTS PAGES(
    id        SERIAL PRIMARY KEY,
    title     TEXT   NOT NULL,
    content   TEXT   NOT NULL
);

CREATE TABLE IF NOT EXISTS POSTS(
    id            SERIAL PRIMARY KEY,
    title         TEXT   NOT NULL,
    content       TEXT   NOT NULL,
    date_created  DATE   NOT NULL
);

CREATE TABLE IF NOT EXISTS COMMENTS(
    id            SERIAL PRIMARY KEY,
    title         TEXT   NOT NULL,
    content       TEXT   NOT NULL,
    date_created  DATE   NOT NULL,
    post_id       INT    REFERENCES POSTS(id)
);
