build: 
	go build

run: 
	go run main.go

test:
	go test -v ./...

clean: 
	rm -fr asset-analysis
