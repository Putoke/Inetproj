all:
	go run server.go
	
rebuilddb:
	mysql --host=mysql-vt2015.csc.kth.se --user=jedluadmin --password=mel4Q8mc < database/createtables.sql
	mysql --host=mysql-vt2015.csc.kth.se --user=jedluadmin --password=mel4Q8mc < database/defaultvalues.sql
	mysql --host=mysql-vt2015.csc.kth.se --user=jedluadmin --password=mel4Q8mc < database/procedures.sql

setupgo:
	go get github.com/go-sql-driver/mysql
	go get github.com/gorilla/mux