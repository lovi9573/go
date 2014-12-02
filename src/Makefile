PATH=/usr/bin:bin

all: create

create:
	gccgo -g -o files files.go directory.go parallelmergesort.go

clean:
	rm files
