CREATE TABLE Feature (
                         id SERIAL PRIMARY KEY

);

CREATE TABLE Tag (
                     id SERIAL PRIMARY KEY

);

CREATE TABLE Banner (
                        id SERIAL PRIMARY KEY,
                        feature INTEGER references Feature(id),
                        tag INTEGER references Tag(id),
                        title VARCHAR(255),
                        text TEXT,
                        url TEXT,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP,
                        is_active BOOLEAN
);


INSERT INTO Feature DEFAULT VALUES;
INSERT INTO Feature DEFAULT VALUES;
INSERT INTO Feature DEFAULT VALUES;

INSERT INTO Tag DEFAULT VALUES;
INSERT INTO Tag DEFAULT VALUES;
INSERT INTO Tag DEFAULT VALUES;


INSERT INTO Banner (feature, tag, title, text, url, created_at, updated_at, is_active) VALUES
                                                                                        (1, 1, 'Summer Sale', 'Huge discounts on summer collections!', 'http://example.com/summer-sale', '2024-04-07 10:00:00', '2024-04-07 10:00:00', TRUE),
                                                                                        (2, 2, 'New Arrivals', 'Discover the latest arrivals in fashion!', 'http://example.com/new-arrivals', '2024-04-07 11:00:00', '2024-04-07 11:00:00', TRUE),
                                                                                        (3, 3, 'Clearance', 'Limited-time clearance sale!', 'http://example.com/clearance', '2024-04-07 12:00:00', '2024-04-07 12:00:00', TRUE);

