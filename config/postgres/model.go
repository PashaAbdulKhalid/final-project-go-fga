package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"database_name"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

type PostgresClient interface {
	GetClient() *gorm.DB
}

type PostgresClientImpl struct {
	cln    *gorm.DB
	config Config
}

func NewPostgresConnection(config Config) PostgresClient {
	connectionStrinf := fmt.Sprintf(`
		host=%s
		port=%s
		user=%s
		password=%s
		dbname=%s`,
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DatabaseName)
	// open gorm connection to postgres
	db, err := gorm.Open(postgres.Open(connectionStrinf), &gorm.Config{})
	// check if error connecting to postgres
	if err != nil {
		panic(err)
	}
	return &PostgresClientImpl{cln: db, config: config}
}

func (p *PostgresClientImpl) GetClient() *gorm.DB {
	return p.cln
}
