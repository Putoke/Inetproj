all:
	
rebuilddb:
	mysql --host=mysql-vt2015.csc.kth.se --user=jedluadmin --password=mel4Q8mc < database/createtables.sql
	mysql --host=mysql-vt2015.csc.kth.se --user=jedluadmin --password=mel4Q8mc < database/defaultvalues.sql
	mysql --host=mysql-vt2015.csc.kth.se --user=jedluadmin --password=mel4Q8mc < database/procedures.sql
