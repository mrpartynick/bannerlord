CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE Feature (
    id SERIAL PRIMARY KEY

);

CREATE TABLE Tag (
    id SERIAL PRIMARY KEY

);

CREATE TABLE Banner (
                        id SERIAL PRIMARY KEY,
                        feature INTEGER references Feature(id),
                        contents TEXT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP,
                        is_active BOOLEAN
);

CREATE TABLE Banner_Tag (
                            banner_id int,
                            tag_id int
);


INSERT INTO Feature DEFAULT VALUES;
INSERT INTO Feature DEFAULT VALUES;
INSERT INTO Feature DEFAULT VALUES;

INSERT INTO Tag DEFAULT VALUES;
INSERT INTO Tag DEFAULT VALUES;
INSERT INTO Tag DEFAULT VALUES;

CREATE TABLE Users (
                       id SERIAL PRIMARY KEY,
                       login VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL
);

CREATE TABLE Admins (
                        id SERIAL PRIMARY KEY,
                        login VARCHAR(255) UNIQUE NOT NULL,
                        password VARCHAR(255) NOT NULL

);

INSERT INTO Users (login, password)
VALUES ('user', crypt('123', 'placeholder')) ;

INSERT INTO Admins (login, password)
VALUES ('admin', crypt('123', 'placeholder')) ;

INSERT INTO banner (
    id, feature, contents, created_at, updated_at, is_active) VALUES (

                                                                '1', '1', '{"kek":"lol"}', '2024-04-07 11:00:00', '2024-04-07 11:00:00', true),
                                                                  ('2', '2', '{"kek":"lol"}', '2024-04-07 11:00:00', '2024-04-07 11:00:00', true);
insert into banner_tag values
                           (1,1),
                           (1,2),
                           (2,1);