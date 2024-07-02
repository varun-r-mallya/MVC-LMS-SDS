CREATE TABLE IF NOT EXISTS booklist (
    B_Id int(10) unsigned NOT NULL AUTO_INCREMENT,
    Title varchar(255) NOT NULL UNIQUE,
    Author varchar(255) DEFAULT NULL,
    Genre varchar(255) DEFAULT NULL,
    NumberofCopies int(10) unsigned DEFAULT 0,
    NumberofCopiesAvailable int(10) unsigned DEFAULT 0,
    NumberofCopiesBorrowed int(10) unsigned DEFAULT 0,
    DueTime int(10) unsigned DEFAULT 0,
    CONSTRAINT check_copies CHECK (NumberofCopies = NumberofCopiesAvailable + NumberofCopiesBorrowed),
    PRIMARY KEY (B_Id)
);