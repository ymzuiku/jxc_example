package kit

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/joho/godotenv"
)

type TheEnv struct {
	IsDev      bool
	Dir        string
	Jwt        []byte
	JwtIss     string
	Session    string
	Sha256Slat string
}

var onceEnvInit sync.Once
var Env = &TheEnv{}

func loadDotEnvFile(twd string) string {
	str := path.Join(twd, ".env")
	if !PathExists(str) {
		fmt.Println(path.Join(twd, ".."))
		return loadDotEnvFile(path.Join(twd, ".."))
	}
	Env.Dir = path.Dir(str)
	return str
}

func envInit() {
	file, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	file = loadDotEnvFile(file)

	if err := godotenv.Load(file); err != nil {
		log.Fatalln(err)
	}

	Env.IsDev = os.Getenv("DEV") != ""
	Env.Jwt = []byte(os.Getenv("JWT"))
	Env.JwtIss = os.Getenv("JWTISS")
	Env.Session = os.Getenv("SESSION")
	Env.Sha256Slat = os.Getenv("SHA256_SALT")
}

func EnvInit() {
	onceEnvInit.Do(envInit)
}
