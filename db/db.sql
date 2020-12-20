-- Create tables.
	DROP TABLE IF EXISTS usersCatalog;
	DROP TABLE IF EXISTS sites;
	DROP TABLE IF EXISTS interestCatalog;
	DROP TABLE IF EXISTS users;
	CREATE TABLE users
	(
		id   SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL UNIQUE
	);

	CREATE TABLE interestCatalog
	(
		id   SERIAL PRIMARY KEY,
		interest VARCHAR(50) NULL,
		userID INT NOT NULL,
		FOREIGN KEY (userID)
		REFERENCES users(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	CREATE TABLE sites
	(
		id   SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL UNIQUE,
		theme VARCHAR(50) NOT NULL UNIQUE
	);

	CREATE TABLE usersCatalog
	(
		id   SERIAL PRIMARY KEY,
		sitesID INT NOT NULL,
		FOREIGN KEY (sitesID)
		REFERENCES sites(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE,
		userID INT NULL,
		FOREIGN KEY (userID)
		REFERENCES users(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	-- Insert demo data.
	INSERT INTO users (name) VALUES ('Franko');
	INSERT INTO users (name) VALUES ('Dimon');
	INSERT INTO users (name) VALUES ('Tom');
	INSERT INTO interestCatalog (interest, userID) VALUES ('Films', 2);
	INSERT INTO interestCatalog (interest, userID) VALUES ('Films', 1);
	INSERT INTO interestCatalog (interest, userID) VALUES ('Magazines', 3);
	INSERT INTO interestCatalog (interest, userID) VALUES ('Magazines', 3);
	INSERT INTO sites (name, theme) VALUES ('Anime', 'Naruto');
	INSERT INTO sites (name, theme) VALUES ('TV-show', 'Supernatural');
	INSERT INTO sites (name, theme) VALUES ('Books', 'Adventure books');
	INSERT INTO sites (name, theme) VALUES ('Films', 'Marvel');
	INSERT INTO usersCatalog (sitesID, userID) VALUES (2, 1);
	INSERT INTO usersCatalog (sitesID, userID) VALUES (2, 2);
	INSERT INTO usersCatalog (sitesID, userID) VALUES (4, 4);
	INSERT INTO usersCatalog (sitesID, userID) VALUES (3, 3);
