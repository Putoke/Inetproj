use jedlu;

#Testusers
insert into users (id, name, lastname, email, password) values (1, "Dr", "Mugg", "Mugg@EvilDR.se", md5("dasseborg"));

#Default excersises
insert into exercises (userid, name, musclegroup) values (1, "Bänkpress", "Bröst");
insert into exercises (userid, name, musclegroup) values (1, "Benpress", "Ben");
insert into exercises (userid, name, musclegroup) values (1, "Marklyft", "Rygg");
insert into exercises (userid, name, musclegroup) values (1, "Armhävningar", "Rygg");
insert into exercises (userid, name, musclegroup) values (1, "Situps", "Rygg");

#Default workouts
insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass", "Bänkpress", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass", "Benpress", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass", "Marklyft", 3, 10);

insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass 2", "Marklyft", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass 2", "Armhävningar", 3, 15);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass 2", "Situps", 3, 15);

insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass 3", "Situps", 3, 15);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass 3", "Marklyft", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (1, "Exempelpass 3", "Benpress", 3, 15);

#Default schedule
insert into schedules (userid, name, workout, day) values (1, "Exempelschema", "Exempelpass", "Måndag");
insert into schedules (userid, name, workout, day) values (1, "Exempelschema", "Exempelpass 2", "Tisdag");
insert into schedules (userid, name, workout, day) values (1, "Exempelschema 2", "Exempelpass 3", "Onsdag");
insert into schedules (userid, name, workout, day) values (1, "Exempelschema 2", "Exempelpass", "Fredag");

