use jedlu;

drop table if exists users;
drop table if exists exercises;
drop table if exists schedules;
drop table if exists workouts;

create table users(
	id int NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	lastname varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	UNIQUE(email),
	PRIMARY KEY (id)
);

create table exercises(
	userid int NOT NULL,
	name varchar(255) NOT NULL,
	musclegroup varchar(255),
	FOREIGN KEY (userid) REFERENCES users(id),
	UNIQUE KEY name (userid, name)
);

create table schedules(
	userid int NOT NULL,
	name varchar(255),
	workout varchar(255) NOT NULL,
	day varchar(255) NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(id),
	FOREIGN KEY (workout) REFERENCES workouts(name)
);

create table workouts(
	userid int NOT NULL,
	name varchar(255) NOT NULL,
	exercise varchar(255) NOT NULL,
	sets int,
	reps int,
	FOREIGN KEY (userid) REFERENCES users(id),
	FOREIGN KEY (exercise) REFERENCES exercises(name),
	UNIQUE KEY name (userid, name, exercise)
);