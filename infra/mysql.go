package infra

import (
	"fmt"
	"log"
	"sync"
	"time"

	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/cenkalti/backoff"
	"github.com/cj/scraper/config"
	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMySQL(cfg *config.MySQLConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get DB instance failed: %v", err)
		return nil, err
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTimeMiliseconds) * time.Millisecond)

	return db, nil
}

func CreateDBAndMigrate(cfg *config.MySQLConfig, migrationFile string) *gorm.DB {
	var db *gorm.DB
	// Wait to create store DB first
	boff := backoff.NewExponentialBackOff()

	// Wait to create operator DB
	err := backoff.Retry(func() error {
		var errNested error
		db, errNested = InitMySQL(cfg)
		if errNested != nil {
			fmt.Printf("Connect mysql error %s \n", errNested.Error())
		} else {
			fmt.Println("Connect mysql successful.")
		}
		return errNested
	}, boff)
	if err != nil {
		panic(err)
	}

	Migrate(migrationFile, cfg.MigrationConnURL)
	return db
}

func Migrate(source string, connStr string) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Migrating....")
	fmt.Printf("Source=%+v Connection=%+v\n", source, connStr)

	db, _ := sql.Open("mysql", connStr)
	driver, _ := migrateMysql.WithInstance(db, &migrateMysql.Config{})
	mg, _ := migrate.NewWithDatabaseInstance(
		source,
		"mysql",
		driver,
	)
	defer mg.Close()

	version, dirty, err := mg.Version()
	if err != nil && err.Error() != migrate.ErrNilVersion.Error() {
		panic(err)
	}

	if dirty {
		mg.Force(int(version) - 1) // nolint
	}

	err = mg.Up()

	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		panic(err)
	}

	fmt.Println("Migration done...")
}
