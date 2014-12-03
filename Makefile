

all:
	cd src; make

webserver:
	gccgo -g -o webserver src/webserver.go

clean:
	rm -f webserver
