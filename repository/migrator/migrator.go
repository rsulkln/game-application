package migrator

import (
	"database/sql"
	"fmt"
	"game/repository/mysql"
	"github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dbConfig   mysql.Config
	dialect    string
	migrations *migrate.FileMigrationSource
}

func New(dbConfig mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	}
	return Migrator{dbConfig: dbConfig, dialect: "mysql", migrations: migrations}
}

func (m *Migrator) connection() (*sql.DB, error) {
	return sql.Open(m.dialect, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username,
		m.dbConfig.Password,
		m.dbConfig.Host,
		m.dbConfig.Port,
		m.dbConfig.Database,
	))
}

func (m *Migrator) Up() {
	db, err := m.connection()
	if err != nil {
		panic(fmt.Errorf("connect to mysql failed: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't appliend to mysql: %+v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)

}

func (m *Migrator) Down() {
	db, err := m.connection()

	if err != nil {
		panic(fmt.Errorf("connect to mysql failed: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback to mysql: %v", err))
	}
	fmt.Printf("rollback %d migrations!\n", n)
}

func (m *Migrator) Status() {
	//TODO - you should add status code for migrations ****

}
