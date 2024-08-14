package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/resrrdttrt/VOU/admin"
	thhttpapi "github.com/resrrdttrt/VOU/admin/api/http"
	"github.com/resrrdttrt/VOU/admin/postgres"
	"github.com/resrrdttrt/VOU/pkg/common"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/logger"
)

const (
	DefHTTPPort       = "3000"
	DefLogLevel       = "info"
	ConnectionTimeout = 10

	DefDBHost      = "171.251.89.96"
	DefDBPort      = "8907"
	DefDBPortRead  = "8907"
	DefDBPortWrite = "8907"
	DefDBUser      = "postgres"
	DefDBPass      = "newpassword"
	DefDbName      = "admin"
	DefSSLMode     = "disable"
	DefSSLCert     = ""
	DefSSLKey      = ""
	DefSSLRootCert = ""

	MongoHost    = "localhost"
	MongoUser    = "root"
	MongoPass    = "1"
	MongoReplica = "rs0"
	MongoDbName  = "admin"
)

type config struct {
	logLevel string
	dbConfig postgres.Config
	httpPort string
}

func loadConfig() config {
	dbConfig := postgres.Config{
		Host:        common.Env("DB_HOST", DefDBHost),
		Port:        common.Env("DB_PORT", DefDBPort),
		PortRead:    common.Env("DB_PORTREAD", DefDBPortRead),
		PortWrite:   common.Env("DB_PORTWRITE", DefDBPortWrite),
		User:        common.Env("DB_USER", DefDBUser),
		Pass:        common.Env("DB_PASS", DefDBPass),
		Name:        common.Env("DB_NAME", DefDbName),
		SSLMode:     common.Env("DB_SSLMODE", DefSSLMode),
		SSLCert:     common.Env("DB_SSLCERT", DefSSLCert),
		SSLKey:      common.Env("DB_SSLKEY", DefSSLKey),
		SSLRootCert: common.Env("DB_ROOTCERT", DefSSLRootCert),
	}

	return config{
		logLevel: common.Env("LOG_LEVEL", DefLogLevel),
		dbConfig: dbConfig,
		httpPort: common.Env("HTTP_PORT", DefHTTPPort),
	}
}

func main() {
	cfg := loadConfig()

	// fake auth service
	admin.ConnectToPostgres()

	// logger
	logging, err := logger.New(os.Stdout, cfg.logLevel)
	if err != nil {
		log.Fatal(err.Error())
	}

	// postgres
	rdb := connectToDBRead(cfg.dbConfig, logging)
	wdb := connectToDBWrite(cfg.dbConfig, logging)
	defer rdb.Close()
	defer wdb.Close()

	// mongoDriver := db.ConnectToMongoDB(db.GetMongoConfig(), logging)
	// commonMongo := db.NewMongoTransactions(mongoDriver)
	// svc := newService(logging, rdb, wdb, commonMongo)

	svc := newService(logging, rdb, wdb)
	errs := make(chan error)
	go startHTTPServer(thhttpapi.MakeHandler(svc), cfg, logging, make(chan error))
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	err = <-errs
	logging.Log(fmt.Sprintf("Admin service terminated: %s", err))

}

func newService(logger logger.Logger, rdb *sqlx.DB, wdb *sqlx.DB) admin.Service {
	database := db.NewReadWrite(rdb, wdb)
	userRepo := postgres.NewUserRepository(database, logger)
	statisticRepo := postgres.NewStatisticRepository(database, logger)
	authRepo := postgres.NewAuthRepository(database, logger)
	svc := admin.NewAdminService(logger, userRepo, statisticRepo, authRepo)
	return svc
}

func startHTTPServer(handler http.Handler, cfg config, logger logger.Logger, errs chan error) {
	p := fmt.Sprintf(":%s", cfg.httpPort)
	logger.Info(fmt.Sprintf("HTTP service start using http on %s", p))
	errs <- http.ListenAndServe(p, handler)
}

func connectToDBRead(dbConfig postgres.Config, logger logger.Logger) *sqlx.DB {
	logger.Info(fmt.Sprintf("Connecting to read database: %s", dbConfig))
	dbConfig.Port = dbConfig.PortRead
	db, err := postgres.ConnectRead(dbConfig)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to read database: %s", err))
		os.Exit(1)
	}
	db.SetConnMaxIdleTime(50)
	db.SetConnMaxLifetime(200)
	return db
}

func connectToDBWrite(dbConfig postgres.Config, logger logger.Logger) *sqlx.DB {
	logger.Info(fmt.Sprintf("Connecting to write database: %s", dbConfig))
	dbConfig.Port = dbConfig.PortWrite
	db, err := postgres.ConnectWrite(dbConfig)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to write database: %s", err))
		os.Exit(1)
	}
	db.SetConnMaxIdleTime(50)
	db.SetConnMaxLifetime(200)
	return db
}
