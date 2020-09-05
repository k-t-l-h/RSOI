package database

import (
	"RSOI/internal/pkg/config"
	"gopkg.in/jackc/pgx.v3"
	"time"
)

type PostgresDB struct {
	dbPool *pgx.ConnPool
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (db *PostgresDB) Open(con config.DatabaseConfig)  {
	connConfig := pgx.ConnConfig{
		Host:     con.Host,
		Port:     uint16(con.Port),
		Database: con.Database,
		User:     con.User,
		Password: con.Password,
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 50,
		AcquireTimeout: 10 * time.Second,
		AfterConnect:   nil,
	}



	db.dbPool, _ = pgx.NewConnPool(poolConfig)
}
