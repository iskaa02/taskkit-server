CREATE TYPE repeat_enum AS ENUM ('daily','weekly','monthly');
CREATE TABLE theme (
    id SERIAL PRIMARY KEY,
    "primary"  VARCHAR(6) NOT NULL,
    secondary VARCHAR(6) NOT NULL
);
CREATE TABLE list (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    theme_id INT NOT NULL REFERENCES theme(id),
    last_modified TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE task (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    subtasks JSONB, 
    list_id VARCHAR NOT NULL REFERENCES list(id), 
    description  TEXT,
    reminder DATE,
    repeat repeat_enum,
    is_completed BOOLEAN NOT NULL DEFAULT false,
    last_modified TIMESTAMP DEFAULT current_timestamp 
);
