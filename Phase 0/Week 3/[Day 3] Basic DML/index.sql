--SHOW DATABASE
SHOW DATABASES;

--create database tourism
CREATE DATABASE tourism;

--use database tourism
USE tourism;

--show tables
SHOW TABLES;

--create table destination
CREATE TABLE destinations (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	country VARCHAR(50) NOT NULL,
	description VARCHAR(255) NOT NULL,
	rating INT NOT NULL
);

--describe table destinations
DESCRIBE destinations;

--create table hotels
CREATE TABLE hotels (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	rating INT NOT NULL,
	address VARCHAR(255) NOT NULL,
	destination_id INT,
	FOREIGN KEY (destination_id) REFERENCES destinations(id)
);

--describe entity hotels
DESCRIBE hotels;

--create entity bookings
CREATE TABLE bookings (
	id INT PRIMARY KEY AUTO_INCREMENT,
	guest_name VARCHAR(50) NOT NULL,
	guest_number INT NOT NULL,
	check_in_date DATE NOT NULL,
	check_out_date DATE NOT NULL,
	hotel_id INT,
	FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

--describe entity bookings
DESCRIBE bookings;

--insert data entity destinations
INSERT INTO destinations (name, country, description, rating)
VALUES ('Paris',	'France',	'Known for the Eiffel Tower, world-class museums, and romantic ambiance.',	5),
('Kyoto',	'Japan',	'Famous for its classical Buddhist temples, beautiful gardens, and traditional wooden houses.',	4),
('Cape Town',	'South Africa',	'Offers stunning beaches, scenic mountains, and vibrant local culture.',	5),
('Santorini',	'Greece',	'Beautiful island known for its white-washed buildings, stunning sunsets, and clear blue waters.',	4),
('Sydney',	'Australia',	'Features iconic landmarks like the Sydney Opera House and beautiful harbor views.',	5);

--insert data entity hotels
INSERT INTO hotels (name, rating, address, destination_id) VALUES
('Hotel Le Parisien',	5,	'15 Rue de Paris, Paris, France',	1),
('Kyoto Garden Inn',	4,	'1-1-1 Higashiyama, Kyoto, Japan',	2),
('Cape Sun Hotel',	5,	'123 Ocean Drive, Cape Town, S. Africa',	3),
('Sunset Villas',	4,	'Fira, Santorini, Greece',	4),
('Harbor View Suites',	5,	'45 Opera Rd, Sydney, Australia',	5);

--insert data entity bookings
INSERT INTO bookings (guest_name, guest_number, check_in_date, check_out_date, hotel_id) VALUES
('Alice Johnson',	2,	'2024-11-20',	'2024-11-25',	1),
('Michael Roberts',	1,	'2024-12-05',	'2024-12-10',	2),
('Emma Green',	3,	'2024-12-15',	'2024-12-20',	3),
('Liam Thompson',	4,	'2024-12-22',	'2024-12-30',	4),
('Sophia Martinez',	2,	'2025-01-10',	'2025-01-15',	5);

--retrieve all data in destinations
SELECT * FROM destinations;

--retrieve all data in hotels
SELECT * FROM hotels;

--retrieve all data in bookings
SELECT * FROM bookings;

--update rating detinations
UPDATE destinations
SET rating = 2
WHERE id = 5;

--update address hotels
UPDATE hotels
SET address = '789 Island Blvd, Thira, Greece'
WHERE id = 5;

--delete data from bookings
DELETE FROM bookings WHERE id = 2;

--create table reviews
CREATE TABLE reviews (
	id INT PRIMARY KEY AUTO_INCREMENT,
	hotel_id INT,
	user_id INT NOT NULL,
	review_date DATE NOT NULL,
	rating INT NOT NULL,
	review_text VARCHAR(100)
);

--describe entity reviews
DESCRIBE reviews;

--add foreign key hotel_id
ALTER TABLE reviews
ADD FOREIGN KEY (hotel_id) REFERENCES hotels(id);

--insert data table reviews
INSERT INTO reviews (hotel_id, user_id, review_date, rating, review_text) VALUES
(1, 101, '2023-01-15', 4, 'Great service and comfortable rooms.'),
(2, 102, '2023-02-10', 5, 'Excellent experience, will visit again!'),
(3, 103, '2023-03-05', 3, 'Average stay, nothing special.'),
(1, 104, '2023-04-20', 5, 'Absolutely loved it!'),
(2, 105, '2023-05-30', 2, 'Not worth the price.');

--retrieve list of hotels that have a rating greater than 4
SELECT * FROM hotels
WHERE rating > 4;

--count the number of reviews given to each hotel
SELECT
h.name AS hotel_name,
COUNT(r.id) AS review_number
FROM hotels h
LEFT JOIN reviews r ON h.id = r.hotel_id
GROUP BY h.id;

--list of hotel names, destination countries, number of reviews
SELECT
h.name AS hotel_name,
d.country AS destination_country,
COUNT(r.id) AS review_number
FROM hotels h
JOIN destinations d ON h.destination_id = d.id
LEFT JOIN reviews r ON h.id = r.hotel_id
GROUP BY h.id, d.country;

--retrieve hotels in a specific destination based on user input
SELECT *
FROM hotels h
LEFT JOIN destinations d ON h.destination_id = d.id
WHERE d.name = 'Kyoto'
GROUP BY h.id, d.id;

--calculate and display the average rating of hotels in a particular destination
SELECT
h.name AS hotel_name,
d.name AS destination_country,
AVG(r.rating) AS review_average
FROM hotels h
JOIN destinations d ON h.destination_id = d.id
LEFT JOIN reviews r ON h.id = r.hotel_id
WHERE d.name = 'Kyoto'
GROUP BY h.id, d.country;