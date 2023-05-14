package database

import "fmt"

type MySql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (m *MySql) DbDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
}
