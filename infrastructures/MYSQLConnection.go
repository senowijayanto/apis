package infrastructures

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type IMYSQLConnection interface {
	MarketDataRead() *sql.DB
	CloseConnection()
}

type MYSQLConnection struct{}

var (
	dbMarketDataRead *sql.DB
	err              error
)

func createDbConnection(descriptor string, maxIdle, maxOpen int) *sql.DB {
	conn, err := sql.Open("mysql", descriptor)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "connection for mysql",
			"event":  "mysql_error_connection",
		}).Error(err)
		os.Exit(0)
	}

	conn.SetMaxIdleConns(maxIdle)
	conn.SetMaxOpenConns(maxOpen)
	return conn
}

//GetReadDb used for connect to read database
func (s *MYSQLConnection) MarketDataRead() *sql.DB {
	if dbMarketDataRead == nil {
		dbMarketDataRead = createDbConnection(
			viper.GetString("database.marketdata_live.read"),
			viper.GetInt("database.marketdata_live.max_idle"),
			viper.GetInt("database.marketdata_live.max_cons"))
	}
	if dbMarketDataRead.Ping() != nil {
		dbMarketDataRead = createDbConnection(
			viper.GetString("database.marketdata_live.read"),
			viper.GetInt("database.marketdata_live.max_idle"),
			viper.GetInt("database.marketdata_live.max_cons"))
	}

	return dbMarketDataRead
}

// CloseConnection used for close database connection
func (s *MYSQLConnection) CloseConnection() {

	if dbMarketDataRead != nil {
		err = dbMarketDataRead.Close()
	}

	if err != nil {
		log.Errorf("db Close Connection Error: %s", err)
	}
}
