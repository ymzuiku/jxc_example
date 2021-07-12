package kit

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/joho/godotenv"
	"github.com/ymzuiku/env_migrate"
)

type TheEnv struct {
	IsDev      bool
	Jwt        []byte
	JwtIss     string
	Session    string
	Sha256Slat string
	RootDir    string
	FileDir    string
}

var onceEnvInit sync.Once
var Env = &TheEnv{}

func loadFileDir() {
	if Env.FileDir == "" {
		file, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		Env.FileDir = file
	}
}

func loadRootDir(twd string) {
	str := path.Join(twd, "go.mod")
	if !PathExists(str) {
		fmt.Println(path.Join(twd, ".."))
		loadRootDir(path.Join(twd, ".."))
		return
	}
	Env.RootDir = path.Dir(str)
}

func envInit() {
	loadFileDir()
	loadRootDir(Env.FileDir)

	envLocal := path.Join(Env.RootDir, ".env")
	if PathExists(envLocal) {
		if err := godotenv.Load(envLocal); err != nil {
			log.Fatalln(err)
		}
	}

	if err := godotenv.Load(path.Join(Env.RootDir, ".base.env")); err != nil {
		log.Fatalln(err)
	}

	Env.IsDev = os.Getenv("DEV") != ""
	Env.Jwt = []byte(os.Getenv("JWT"))
	Env.JwtIss = os.Getenv("JWTISS")
	Env.Session = os.Getenv("SESSION")
	Env.Sha256Slat = os.Getenv("SHA256_SALT")

	env_migrate.BaseRootDir = Env.RootDir
}

func EnvInit() {
	onceEnvInit.Do(envInit)
}
