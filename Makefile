

all: webserver

webserver:
	gccgo -g -o webserver src/webserver.go
	cd src; make
	cp src/driver ./driver
	cp src/driver.sh ./driver.sh

go:
	go build -o webserver src/webserver.go
	cd src; make go
	cp src/driver ./driver
	cp src/driver.sh ./driver.sh

clean:
	cd src; make clean
	rm -f webserver
	rm -f driver
	rm -f driver.sh
	rm -f html/out.html

.PHONY: webserver all clean go
