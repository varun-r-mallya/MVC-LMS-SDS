CREATE DATABASE IF NOT EXISTS LMS;
USE LMS;

CREATE TABLE IF NOT EXISTS userlist (
    username VARCHAR(255) NOT NULL PRIMARY KEY UNIQUE,
    hashedpassword VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    isadmin BOOLEAN NOT NULL
);

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

CREATE TABLE transactions (
    T_Id int(10) unsigned NOT NULL AUTO_INCREMENT UNIQUE,
    username varchar(255) NOT NULL,
    B_Id int(10) unsigned NOT NULL, -- book Identification number
    CheckOutAccepted boolean DEFAULT NULL,
    CheckInAccepted boolean DEFAULT NULL,
    DateBorrowed date DEFAULT NULL,
    DateReturned date DEFAULT NULL,
    OverDueFine int(10) unsigned DEFAULT 0,
    primary key(T_Id),
    FOREIGN KEY (username) REFERENCES userlist(username),
    FOREIGN KEY (B_Id) REFERENCES booklist(B_Id)
);

CREATE TABLE convertq (
    C_Id int(10) unsigned NOT NULL AUTO_INCREMENT UNIQUE,
    username varchar(255) NOT NULL UNIQUE,
    FOREIGN KEY (username) REFERENCES userlist(username),
    primary key(C_Id)
);
