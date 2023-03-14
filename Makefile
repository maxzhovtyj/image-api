.PHONY: run
run:
		go run ./cmd/main.go

.PHONY: golint
golint:
		golangci-lint run --no-config --disable-all --enable=revive


.PHONY: appUp
appUp:
		docker-compose up -d --build api
