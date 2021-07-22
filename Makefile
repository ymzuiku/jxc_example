test:
	cor go test ./...
test-one:
	cor go test ./... -count=1 -test.run TestLoadCompanys
cover:
	cor go test ./... -count=1 -cover -tags test -coverprofile=t.out
	go tool cover -html=t.out

dev:
	up_migrate=all cor air
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jxc cmd/jxc/main.go
up:
	up_migrate=all only_migrate=1 cor go run cmd/jxc/main.go
down:
	down_migrate=all cor go run cmd/jxc/main.go