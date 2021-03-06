build:
	go build

run:
	go build
	./trigram

clean:
	rm trigram

test:
	go test ./...

cover:
	go test ./... -timeout 30s -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out