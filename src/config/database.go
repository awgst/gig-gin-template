package config

type DatabaseConfig interface {
	DbDsn() string
}
