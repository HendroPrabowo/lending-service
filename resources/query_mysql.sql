CREATE TABLE account (
	id int PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	name VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	fcm_token VARCHAR (255),
	created_at TIMESTAMP NOT null default NOW(),
	updated_at TIMESTAMP NOT null default NOW()
);

CREATE TABLE loan (
	id serial PRIMARY KEY,
	lender int not null,
	borrower int not null,
	amount varchar(50) not null,
	status varchar(50) not null,
	description varchar(255) not null,
	created_at TIMESTAMP NOT null default NOW(),
	updated_at TIMESTAMP NOT null default NOW()
);