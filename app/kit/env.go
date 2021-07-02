package kit

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type TheEnv struct {
	IsDev bool
}

var Env = TheEnv{}
var loaded = false

func EnvInit(file string) {
	if loaded {
		return
	}
	loaded = true
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalln(err)
	}
	Env = TheEnv{
		IsDev: os.Getenv("DEV") != "",
	}
}
