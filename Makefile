hoc: hoc.go
	go build

hoc.go:
	goyacc -o hoc.go hoc.y

clean:
	rm hoc.go
	rm hoc

