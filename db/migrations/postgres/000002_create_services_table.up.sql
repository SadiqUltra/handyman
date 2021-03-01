CREATE SEQUENCE Services_seq;

CREATE TABLE IF NOT EXISTS Services
(
    ID INT NOT NULL UNIQUE DEFAULT NEXTVAL
(
    'Services_seq'
),
    Title VARCHAR
(
    255
),
    Description TEXT,
    Price DOUBLE PRECISION,
    UserID INT,
    FOREIGN KEY
(
    UserID
) REFERENCES Users
(
    ID
),
    PRIMARY KEY
(
    ID
)
    )