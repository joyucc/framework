package config

import (
	"fmt"
	"time"
)

const (
	// PostgreSQLDatabaseAdapter is the adapter name for PostgreSQL.
	PostgreSQLDatabaseAdapter DatabaseAdapter = "postgres"
	// MySQLDatabaseAdapter is the adapter name for MySQL.
	MySQLDatabaseAdapter DatabaseAdapter = "mysql"
	// SQLiteDatabaseAdapter is the adapter name for SQLite.
	SQLiteDatabaseAdapter DatabaseAdapter = "sqlite3"
)

// DatabaseAdapter represents a database adapter constant.
type DatabaseAdapter string

// Database holds database connection parameters.
type Database struct {
	Adapter     string `default:"postgres"`
	Hostname    string
	Port        int
	Username    string
	Password    string
	Database    string
	Singularize bool   `default:"false"`
	UTC         bool   `default:"true"`
	Description string `default:""`
	Params      map[string]interface{}

	Active      int           `json:"active"`
	Idle        int           `json:"idle"`
	IdleTimeout time.Duration `json:"idle_timeout"`
}

// URL returns a connection string for the database.
func (d *Database) URL() string {
	return parseConnConfig(d.Adapter, d.Hostname, d.Database, d.Username, d.Password, d.Description, d.Port)
}

func realDSN(driver, dbname, username, password, addr, charset string) string {
	connStr := ""
	switch driver {
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True&loc=Local", username, password, addr, dbname, charset)
	case "postgres":
		connStr = fmt.Sprintf("%s dbname=%s user=%s password=%s", addr, dbname, username, password)
	case "mssql":
		connStr = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", username, password, addr, dbname)
	}
	return connStr
}

func parseConnConfig(dialect, host, dbName, username, password, charset string, port int) string {
	connHost := ""
	switch dialect {
	case "mysql":
		connHost = fmt.Sprintf("tcp(%s:%d)", host, port)
	case "postgres":
		connHost = fmt.Sprintf("host=%s port=%d", host, port)
	case "mssql":
		connHost = fmt.Sprintf("%s:%d", host, port)
	}
	s := realDSN(dialect, dbName, username, password, connHost, charset)
	return s
}
