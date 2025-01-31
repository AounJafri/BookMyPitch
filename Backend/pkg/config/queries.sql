-- DATABSE CREATION QUERY

Create database BookMyPitch;

-- TABLES CREATION QUERIES

Create Table Users(
	id integer Primary key auto_increment,
	username VARCHAR(50) UNIQUE,
    password VARCHAR(50),
    email varchar(55),
    age integer,
	role VARCHAR(15),
    location VARCHAR(255)
)

Create Table Grounds(
	id integer Primary key auto_increment,
	name VARCHAR(100) ,
    location VARCHAR(255),
	priceperday integer,
    image varchar(255)
) 
 
CREATE TABLE Bookings(
	id integer Primary key auto_increment,
    groundid integer not null,
    userid integer not null,
    bookingdate Date,
    timeslot VARCHAR(30),
    status varchar(30),
    foreign key(groundid) references Grounds(id),
    foreign key(userid) references Users(id)
)
