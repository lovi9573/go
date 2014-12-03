PATH=/usr/bin:bin

all: create

create:
	gccgo -g -o driver driver.go files.go directory.go parallelmergesort.go

go:
	go build -o driver driver.go files.go directory.go parallelmergesort.go

clean:
	rm driver
