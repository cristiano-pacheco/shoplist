# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Install dependencies

install-libs:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest


# ==============================================================================
# Administration

ping:
	curl -il http://localhost:4000/api/v1/ping

run:
	go run .

migrate:
	go run ./main.go db:migrate

# ==============================================================================
# Running tests within the local computer

lint:
	golangci-lint run ./...

vuln-check:
	govulncheck -show verbose ./... 

test: test-only lint vuln-check

test-race: test-r lint vuln-check

tests:
	CGO_ENABLED=0 go test -count=1 ./...

tests-coverage:
	CGO_ENABLED=0 go test -v -coverprofile cover.out ./...
	CGO_ENABLED=0 go tool cover -html cover.out -o cover.html
	open cover.html
