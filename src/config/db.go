package config

import "gig-gin-template/pkg/env"

type conf map[string]map[string]string

// Return type conf of database configurations from each driver
func Database() conf {
	return conf{
		"mysql": {
			"host":     env.Get("DB_HOST", "localhost"),
			"port":     env.Get("DB_PORT", "3306"),
			"username": env.Get("DB_USERNAME", "root"),
			"password": env.Get("DB_PASSWORD", ""),
			"database": env.Get("DB_DATABASE", ""),
		},
		// Your database configuration here
	}
}
