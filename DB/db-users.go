package DB

import (
	"time"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := oracle.BuildUrl("127.0.0.1", 1521, "XE", "sys", "senha123", nil)
	db, err := gorm.Open(oracle.New(oracle.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar ao banco de dados")
	}
	return db

	sqlDB, err := db.DB()
	if err != nil {
		panic("falha ao obter a conexão do banco de dados")
	}
	// pool de conexões
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	// É necessário importar "time" para usar time.Hour
	// Adicione: import "time" no início do arquivo
	sqlDB.SetConnMaxLifetime(time.Hour)

}
