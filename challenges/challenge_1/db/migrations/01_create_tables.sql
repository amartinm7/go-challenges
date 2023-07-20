CREATE TABLE IF NOT EXISTS ads
(
    id          VARCHAR(50) NOT NULL CHECK(id != ''),
    title       VARCHAR(50) NOT NULL CHECK(id != ''),
    description VARCHAR(50) NOT NULL CHECK(id != ''),
    price       INT         NOT NULL,
    time_stamp  VARCHAR(10) NOT NULL
);

ALTER TABLE ads
    ADD CONSTRAINT pk_ad UNIQUE (id);

CREATE INDEX IF NOT EXISTS idx_a_id ON ads (id);