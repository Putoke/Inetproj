ADMIN := $(shell cat database/adminuser)
ADMINPW := $(shell cat database/adminpassword)

all:
	go run server.go
	
rebuilddb:
	cat database/*.sql | mysql --host=mysql-vt2015.csc.kth.se --user=$(ADMIN) --password=$(ADMINPW)

setupgo:
	go get github.com/go-sql-driver/mysql
	go get github.com/gorilla/mux