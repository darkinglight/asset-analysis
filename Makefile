build: 
	go build

run: 
	go run main.go

test:
	go test -v ./...

clean: 
	rm -fr asset-analysis

dbuild:
	docker build -t my-golang-app .

drun:
	docker run -it --rm --name my-running-app my-golang-app

dgobuild:
	docker run --rm -v `$PWD`:/usr/src/myapp -w /usr/src/myapp golang:1.13 go build -v
