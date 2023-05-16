package database

import "fmt"

type PostgreSql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	TimeZone string
}

func (p *PostgreSql) DbDsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		p.Host,
		p.Username,
		p.Password,
		p.Database,
		p.Port,
		p.TimeZone,
	)
}
