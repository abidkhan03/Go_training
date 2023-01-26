CREATE TABLE objects IF NOT EXISTS (
    id serial PRIMARY KEY NOT NULL,
    name varchar(50) NOT NULL,
    description varchar(250) NOT NULL
);
