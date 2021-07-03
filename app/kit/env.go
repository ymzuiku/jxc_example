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
}

var Env = TheEnv{}
var loaded = false
var EnvDir = ""

func loadDotEnvFile(twd string) string {
	str := path.Join(twd, ".env")
	if !PathExists(str) {
		fmt.Println(path.Join(twd, ".."))
		return loadDotEnvFile(path.Join(twd, ".."))
	}
	EnvDir = path.Dir(str)
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
	Env = TheEnv{
		IsDev: os.Getenv("DEV") != "",
	}
}
