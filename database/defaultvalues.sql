use jedlu;

#Testusers
insert into users (id, name, lastname, email, password) values (1, "Mr", "Workout", "foo@bar.com", md5("password"));

#Default excersises
insert into exercises (userid, name, musclegroup) values (1, "Bench press", "Chest");
insert into exercises (userid, name, musclegroup) values (1, "Legp ress", "Legs");
insert into exercises (userid, name, musclegroup) values (1, "Deadlift", "Back");
insert into exercises (userid, name, musclegroup) values (1, "Pushups", "Chest");
insert into exercises (userid, name, musclegroup) values (1, "Situps", "Abdominals");
insert into exercises (userid, name, musclegroup) values (1, "Leg extension", "Legs");
insert into exercises (userid, name, musclegroup) values (1, "Bicep curl", "Arms");
insert into exercises (userid, name, musclegroup) values (1, "Shoulder press", "Arms");

#Default workouts
insert into workouts (userid, name, exercise, sets, reps) values (1, "Upper Body", "Bench press", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Upper Body", "Pushups", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Upper Body", "Deadlift", 3, 10);

insert into workouts (userid, name, exercise, sets, reps) values (1, "Lower Body", "Leg press", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Lower Body", "Leg extension", 3, 15);


insert into workouts (userid, name, exercise, sets, reps) values (1, "Arms", "Bicep curl", 3, 15);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Arms", "Shoulder press", 3, 15);

#Default schedule
insert into schedules (userid, name, workout, day) values (1, "Schema 1", "Upper Body", "Monday");
insert into schedules (userid, name, workout, day) values (1, "Schema 1", "Lower Body", "Tuesday");
insert into schedules (userid, name, workout, day) values (1, "Schema 2", "Arms", "Wednesday");
insert into schedules (userid, name, workout, day) values (1, "Schema 2", "Upper Body", "Friday");

