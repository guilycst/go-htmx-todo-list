package loadenv

import (
	"flag"

	"github.com/joho/godotenv"
)

func LoadEnv(env *string) {
	flag.Parse()
	if env == nil || *env == "" {
		return
	}
	godotenv.Load(*env)
}
