use jedlu;

#Testusers
insert into users (id, name, lastname, email, password) values (1, "Dr", "Mugg", "Mugg@EvilDR.se", md5("dasseborg"));

#Default excersises
insert into exercises (userid, name, musclegroup) values (1, "Benchpress", "Chest");
insert into exercises (userid, name, musclegroup) values (1, "Legpress", "Legs");
insert into exercises (userid, name, musclegroup) values (1, "Deadlift", "Back");
insert into exercises (userid, name, musclegroup) values (1, "Pushups", "Chest");
insert into exercises (userid, name, musclegroup) values (1, "Situps", "Abdominals");

#Default workouts
insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 1", "Benchpress", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 1", "Legpress", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 1", "Deadlift", 3, 10);

insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 2", "Deadlift", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 2", "Pushups", 3, 15);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 2", "Situps", 3, 15);

insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 3", "Situps", 3, 15);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 3", "Deadlift", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Workout 3", "Legpress", 3, 15);

#Default schedule
insert into schedules (userid, name, workout, day) values (1, "Schema 1", "Workout 1", "Monday");
insert into schedules (userid, name, workout, day) values (1, "Schema 1", "Workout 2", "Tuesday");
insert into schedules (userid, name, workout, day) values (1, "Schema 2", "Workout 3", "Wednesday");
insert into schedules (userid, name, workout, day) values (1, "Schema 2", "Workout 1", "Friday");

