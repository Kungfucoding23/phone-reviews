package main

import (
	"github.com/Kungfucoding23/API-Go-mysql-chi/gadgets/smartphones/web"
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/database"
	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/logs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/golang-migrate/migrate/v4"
	migration "github.com/golang-migrate/migrate/v4/database/mysql"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()

	client := database.NewSQLClient("root:root@tcp(localhost:3306)/phones_review")
	doMigrate(client, "phones_review")

	handler := web.NewCreateSmartphoneHandler(client)
	mux := Routes(handler)
	server := NewServer(mux)
	server.Run()
}

func doMigrate(client *database.MySQLClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}
