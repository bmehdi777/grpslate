.PHONY: build clean

build:
	go build -o bin/grpslate cmd/grpslate/main.go

clean:
	rm bin/*
