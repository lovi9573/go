

all: server

gccserver:
	gccgo -g webserver.go -o ../webserver

server:
	go build webserver.go
	mv webserver ../
