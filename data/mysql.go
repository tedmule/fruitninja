package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type DB struct {
	Cli *sql.DB
}
type Fruit struct {
	ID     int64
	Name   string
	Amount int
}

func NewMysqlClient(address, username, password, dbname string) (*DB, error) {
	// dsn := fmt.Sprintf("%s:%s@%s/%s", username, password, address, dbname)
	cfg := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 address,
		DBName:               dbname,
		AllowNativePasswords: true,
	}
	log.Debug(cfg.FormatDSN())
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	pingErr := conn.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	return &DB{
		Cli: conn,
	}, nil
}

func (db *DB) GetFruits() ([]Fruit, error) {
	var fruits []Fruit

	rows, err := db.Cli.Query("SELECT * FROM fruit")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fruit Fruit
		if err := rows.Scan(&fruit.ID, &fruit.Name, &fruit.Amount); err != nil {
			return nil, err
		}
		fruits = append(fruits, fruit)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return fruits, nil
}
func (db *DB) AddAmount(fruit string) error {
	sql := fmt.Sprintf("SELECT * FROM fruit where name=%s", fruit)
	rows, err := db.Cli.Query(sql)
	if err != nil {
		log.Debug(err)
		return err
	}

	defer rows.Close()

}

func (db *DB) GetSingleFruit(name string) (Fruit, error) {
	var fruit Fruit

	sql := fmt.Sprintf("SELECT * FROM fruit where name=%s", name)
	row := db.Cli.QueryRow(sql)
	if err := row.Scan(&fruit.ID, &fruit.Name, &fruit.Amount); err != nil {
		if err == sql.ErrNoRows {
			return fruit, fmt.Errorf("Fruit %s: not found", fruit)
		}
		return fruit, err
	}

	return fruit, nil

}
