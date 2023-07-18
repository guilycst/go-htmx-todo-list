package loadenv

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	l := len(os.Args)
	if l > 1 {
		var env = os.Args[l-1]
		godotenv.Load(env)
	} else {
		godotenv.Load()
	}
}
