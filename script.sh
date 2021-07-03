if [ "$1" == "testc" ]
then
  go test ./app/... -count=1 -cover -tags test -coverprofile=t.out
  go tool cover -html=t.out
elif [ "$1" == "test" ]
then
  air -c .air_test.toml
elif [ "$1" == "linux" ]
then
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jxc cmd/jxc/main.go
elif [ "$1" == "dev" ]
then
  upMigrate=999 air
elif [ "$1" == "seeds" ]
then
  go run cmd/seeds/main.go
else
  echo "please input: test |testc | build | dev"
fi