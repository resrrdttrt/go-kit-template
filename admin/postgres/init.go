package postgres

import (
	"database/sql"
	"fmt"

	"go-kit-template/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // required for SQL access
	migrate "github.com/rubenv/sql-migrate"
)

// Config defines the options that are used when connecting to a PostgreSQL instance
type Config struct {
	Host        string
	Port        string
	PortRead    string
	PortWrite   string
	User        string
	Pass        string
	Name        string
	SSLMode     string
	SSLCert     string
	SSLKey      string
	SSLRootCert string
}

var (
	ErrInsertDb      = errors.New("Error inserting to database")
	ErrUpdateDb      = errors.New("Error updating database")
	ErrDeleteDb      = errors.New("Error deleting from database")
	ErrSelectDb      = errors.New("Error selecting from database")
	ErrMarshalJSON   = errors.New("Error marshaling to JSON")
	ErrUnmarshalJSON = errors.New("Error unmarshaling")
	ErrGenerateToken = errors.New("Failed to generate token")
	ErrNoData        = errors.New("No data found")
)

func ConnectRead(cfg Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s \n", cfg.Host, cfg.PortRead, cfg.User, cfg.Name, cfg.Pass, cfg.SSLMode, cfg.SSLCert, cfg.SSLKey, cfg.SSLRootCert)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectWrite(cfg Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s \n", cfg.Host, cfg.PortWrite, cfg.User, cfg.Name, cfg.Pass, cfg.SSLMode, cfg.SSLCert, cfg.SSLKey, cfg.SSLRootCert)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := migrateDB(db.DB); err != nil {
		return nil, err
	}
	return db, nil
}

func migrateDB(db *sql.DB) error {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "user_table",
				Up: []string{
					`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
					`CREATE TABLE IF NOT EXISTS "users" (
						id             	UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at     	TIMESTAMP       DEFAULT NOW(),
						updated_at     	TIMESTAMP       DEFAULT NOW(),
						name       		VARCHAR(254)    NOT NULL,
						username       	VARCHAR(254)    NOT NULL,
						password       	VARCHAR(254)    NOT NULL,
						email       	VARCHAR(254)    NOT NULL,
						phone       	VARCHAR(20)    	NOT NULL,
						role       		VARCHAR(20)    	NOT NULL,
						status       	VARCHAR(20)    	NOT NULL
					)`,
				},
				Down: []string{
					`DROP TABLE "users"`,
				},
			},
			{
				Id: "game_table",
				Up: []string{
					`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
					`CREATE TABLE IF NOT EXISTS "games" (
						id             	UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at     	TIMESTAMP       DEFAULT NOW(),
						updated_at     	TIMESTAMP       DEFAULT NOW(),
						name       		VARCHAR(254)    NOT NULL,
						images       	TEXT            NOT NULL,
						type       		VARCHAR(20)    	NOT NULL,
						exchange_allow  BOOLEAN         NOT NULL,
						tutorial       	TEXT         	NOT NULL
					)`,
				},
				Down: []string{
					`DROP TABLE "games"`,
				},
			},
			{
				Id: "access_token_table",
				Up: []string{
					`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
					`CREATE TABLE IF NOT EXISTS "access_tokens" (
						id             	UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at     	TIMESTAMP       DEFAULT NOW(),
						updated_at     	TIMESTAMP       DEFAULT NOW(),
						token          	VARCHAR(254)    NOT NULL,
						user_id        	UUID            NOT NULL
					)`,
				},
				Down: []string{
					`DROP TABLE "access_tokens"`,
				},
			},
			{
				Id: "enterprise_table",
				Up: []string{
					`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
					`CREATE TABLE IF NOT EXISTS "enterprises" (
						id              UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at      TIMESTAMP       DEFAULT NOW(),
						updated_at      TIMESTAMP       DEFAULT NOW(),
						name            VARCHAR(254)    NOT NULL,
						field           VARCHAR(254)    NOT NULL,
						location        VARCHAR(254)    NOT NULL,
						gps             VARCHAR(254)    NOT NULL,
						status          VARCHAR(20)     NOT NULL
					)`,
				},
				Down: []string{
					`DROP TABLE "enterprises"`,
				},
			},
			{
				Id: "event_table",
				Up: []string{
					`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
					`CREATE TABLE IF NOT EXISTS "events" (
						id              UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at      TIMESTAMP       DEFAULT NOW(),
						updated_at      TIMESTAMP       DEFAULT NOW(),
						user_id         UUID            NOT NULL,
						name            VARCHAR(254)    NOT NULL,
						images          TEXT            NOT NULL,
						voucher_num     INTEGER         NOT NULL,
						start_time      TIMESTAMP       NOT NULL,
						end_time        TIMESTAMP       NOT NULL,
						game_id         UUID            NOT NULL
					)`,
				},
				Down: []string{
					`DROP TABLE "events"`,
				},
			},
			{
				Id: "voucher_table",
				Up: []string{
					`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
					`CREATE TABLE IF NOT EXISTS "vouchers" (
						id              UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at      TIMESTAMP       DEFAULT NOW(),
						updated_at      TIMESTAMP       DEFAULT NOW(),
						code            VARCHAR(254)    NOT NULL,
						qrcode          TEXT            NOT NULL,
						images          TEXT            NOT NULL,
						value           INTEGER         NOT NULL,
						description     TEXT            NOT NULL,
						expired_time    TIMESTAMP       NOT NULL,
						status          VARCHAR(20)     NOT NULL,
						event_id        UUID            NOT NULL,
					)`,
				},
				Down: []string{
					`DROP TABLE "vouchers"`,
				},
			},
		},
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	return err
}
