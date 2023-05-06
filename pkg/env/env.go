package env

import "os"

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
