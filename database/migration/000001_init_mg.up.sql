CREATE TABLE IF NOT EXISTS userlist (
    username VARCHAR(255) NOT NULL PRIMARY KEY UNIQUE,
    hashedpassword VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    isadmin BOOLEAN NOT NULL
)