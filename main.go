package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	ID     string `json:"id" gorm:"type:varchar(36);primaryKey;"`
	UserID string `json:"userId" gorm:"type:varchar(36)"`
}

func main() {

	dbDSN := "user=postgres password=root host=localhost dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	dbSql, err := sql.Open("postgres", dbDSN)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSql,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	fmt.Println(err)

	db.AutoMigrate(
		&Student{},
	)
}
