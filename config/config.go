package config

type DatabaseConfig struct {
	Type     string // e.g., "postgres", "inmemory"
	Host     string
	User     string
	Password string
	SSLMode  string
	DBName   string
}

var InMemory DatabaseConfig = DatabaseConfig{Type: "inmemory", Host: "N/A", User: "", Password: ""}

var Postgres DatabaseConfig = DatabaseConfig{Type: "postgres", DBName: "acme", Host: "localhost", User: "postgres", Password: "pandusia1", SSLMode: "disable"}
