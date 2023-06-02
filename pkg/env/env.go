package env

import (
	"os"

	"github.com/joho/godotenv"
)

// Load env file
func Load(filenames ...string) error {
	if len(filenames) > 0 {
		return godotenv.Load(filenames...)
	}

	return godotenv.Load(".env")
}

// Get env value using its key set the alternative value if the key is undefined
func Get(key string, alt ...string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}

	if len(alt) > 0 {
		return alt[0]
	}

	return ""
}
