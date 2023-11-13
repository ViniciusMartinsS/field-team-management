.PHONY: dependencies run unit-tests tests coverage

dependencies:
	go mod vendor

run:
	go run .

unit-tests:
	GIN_MODE=release go test -v ./... --cover -tags="unit" ./...

tests:
	GIN_MODE=release go test -v ./... -coverprofile=coverage.out -tags="unit integration" ./...

coverage:
	go tool cover -html=coverage.out -o coverage.html

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down