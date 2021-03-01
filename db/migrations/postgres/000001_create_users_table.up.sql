CREATE SEQUENCE Users_seq;

CREATE TABLE IF NOT EXISTS Users
(
    ID INT NOT NULL UNIQUE DEFAULT NEXTVAL
(
    'Users_seq'
),
    Username VARCHAR
(
    127
) NOT NULL UNIQUE,
    Phone VARCHAR
(
    127
) NOT NULL UNIQUE,
    Address VARCHAR
(
    255
),
    City VARCHAR
(
    255
),
    Password VARCHAR
(
    127
) NOT NULL,
    PRIMARY KEY
(
    ID
)
    )