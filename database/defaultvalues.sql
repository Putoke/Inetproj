use jedlu;

#Testusers
insert into users (id, name, lastname, email, password) values (0, "Dr", "Mugg", "Mugg@EvilDR.se", md5("dasseborg"));

#Default excersises
insert into exercises (userid, name, musclegroup) values (0, "Bänkpress", "Bröst");
insert into exercises (userid, name, musclegroup) values (0, "Benpress", "Ben");
insert into exercises (userid, name, musclegroup) values (0, "Marklyft", "Rygg");
insert into exercises (userid, name, musclegroup) values (0, "Armhävningar", "Rygg");
insert into exercises (userid, name, musclegroup) values (0, "Situps", "Rygg");

#Default workouts
insert into workouts (userid, name, exercise, sets, reps) values (0, "Exempelpass", "Bänkpress", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (0, "Exempelpass", "Benpress", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (0, "Exempelpass", "Marklyft", 3, 10);

insert into workouts (userid, name, exercise, sets, reps) values (0, "Exempelpass 2", "Marklyft", 3, 10);
insert into workouts (userid, name, exercise, sets, reps) values (0, "Exempelpass 2", "Armhävningar", 3, 15);
insert into workouts (userid, name, exercise, sets, reps) values (0, "Exempelpass 2", "Situps", 3, 15);


#Default schedule
insert into schedules (userid, name, workout, day) values (0, "Exempelschema", "Exempelpass", "Måndag");
insert into schedules (userid, name, workout, day) values (0, "Exempelschema", "Exempelpass 2", "Tisdag");
insert into schedules (userid, name, workout, day) values (0, "Exempelschema 2", "Exempelpass", "Fredag");