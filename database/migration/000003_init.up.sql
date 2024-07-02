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
