all: gen build
gen:
	thrift -r --gen go example.thrift
build:
	go build -o bin/client client/main.go
	go build -o bin/server server/main.go
clean:
	rm -rf bin/*
clean-all: clean
	rm -rf gen-go
