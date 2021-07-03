package kit

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

type TheEnv struct {
	IsDev bool
	Dir   string
}

var Env = &TheEnv{}
var loaded = false

func loadDotEnvFile(twd string) string {
	str := path.Join(twd, ".env")
	if !PathExists(str) {
		fmt.Println(path.Join(twd, ".."))
		return loadDotEnvFile(path.Join(twd, ".."))
	}
	Env.Dir = path.Dir(str)
	return str
}

func EnvInit() {
	if loaded {
		return
	}
	loaded = true

	file, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}
	file = loadDotEnvFile(file)
	err = godotenv.Load(file)
	if err != nil {
		log.Fatalln(err)
	}
	Env.IsDev = os.Getenv("DEV") != ""
}
