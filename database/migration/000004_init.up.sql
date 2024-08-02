CREATE TABLE convertq (
    C_Id int(10) unsigned NOT NULL AUTO_INCREMENT UNIQUE,
    username varchar(255) NOT NULL UNIQUE,
    FOREIGN KEY (username) REFERENCES userlist(username),
    primary key(C_Id)
);