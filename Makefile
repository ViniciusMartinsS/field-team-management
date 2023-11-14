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

kubernetes:
	k3d cluster create field-cluster -p "8080:30080@agent:0" --agents 2 && kubectl apply -f k8s/manifest.yml

kubernetes-pod-status:
	kubectl get pods

kubernetes-stop:
	kubectl delete pods --all && k3d cluster stop field-cluster && k3d cluster delete field-cluster