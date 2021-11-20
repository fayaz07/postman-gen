dev:
	go run src/main.go

all: install

install:
	echo "Installing go modules..." && \
	go mod download && \
	echo "Completed" 

build:
	go build src/main.go
