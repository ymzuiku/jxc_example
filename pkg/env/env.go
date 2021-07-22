package env

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/ymzuiku/gewu_jxc/pkg/pathex"

	"github.com/joho/godotenv"
	"github.com/ymzuiku/env_migrate"
)

var IsDev bool
var IgnoreSQLLog bool
var Jwt []byte
var JwtIss string
var Session string
var Sha256Slat string
var RootDir string
var FileDir string
var LoggerLevel int
var onceEnvInit sync.Once

func loadFileDir() {
	if FileDir == "" {
		file, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		FileDir = file
	}
}

func loadRootDir(twd string) {
	str := path.Join(twd, "go.mod")
	if !pathex.Exists(str) {
		fmt.Println(path.Join(twd, ".."))
		loadRootDir(path.Join(twd, ".."))
		return
	}
	RootDir = path.Dir(str)
}

func envInit() {
	loadFileDir()
	loadRootDir(FileDir)

	envLocal := path.Join(RootDir, ".env")
	if pathex.Exists(envLocal) {
		if err := godotenv.Load(envLocal); err != nil {
			log.Fatalln(err)
		}
	}

	Jwt = []byte(os.Getenv("JWT"))
	JwtIss = os.Getenv("JWTISS")
	Session = os.Getenv("SESSION")
	Sha256Slat = os.Getenv("SHA256_SALT")

	if !IsDev {
		IsDev = os.Getenv("DEV") != ""
	}

	if !IgnoreSQLLog {
		IgnoreSQLLog = os.Getenv("sql_log") == ""
	}

	env_migrate.BaseRootDir = RootDir
}

func Init() {
	onceEnvInit.Do(envInit)
}
