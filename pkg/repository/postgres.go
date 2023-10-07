package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

const (
	deliveriesTable = "deliveries"
	itemsTable      = "items"
	OrdersTable     = "orders"
	paymentsTable   = "payments"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	//db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = db.Ping()
	//if err != nil {
	//	return nil, err
	//}
	//
	//logrus.Print("db connected")
	//return db, nil
	return conn, nil
}
