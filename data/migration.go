package data

import (
	"database/sql"
	"log"
	"github.com/rubenv/sql-migrate"
)

// migration represents the mgrations that can be done
type migration struct {
	db         *sql.DB
	driver     string
	migrations migrate.MigrationSource
}

var migrations = []*migrate.Migration{
	&migrate.Migration{
		Id: "1",
		Up: []string{
			`CREATE TABLE chat_members (
				chat VARCHAR(64),
				name VARCHAR(64),
				PRIMARY KEY (chat,name)
			)`,
		},
		Down: []string{
			`DROP TABLE IF EXISTS chat_members`,
		},
	},
	&migrate.Migration{
		Id: "2",
		Up: []string{
			`CREATE TABLE nostot (
				chat VARCHAR(64),
				name VARCHAR(64),
				created TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
				PRIMARY KEY (chat, name, created)
			)`,
		},
		Down: []string{
			`DROP TABLE IF EXISTS nostot`,
		},
	},
}

func newMigration(db *sql.DB, driver string) *migration {
	return &migration{
		db,
		driver,
		&migrate.MemoryMigrationSource{
			Migrations: migrations,
		},
	}
}

// Do runs all migrations
func (m *migration) Do(reapply bool) error {
	if reapply {
		if _, err := migrate.ExecMax(m.db, m.driver, m.migrations, migrate.Down, 1); err != nil {
			return err
		}
	}
	n, err := migrate.Exec(db, m.driver, m.migrations, migrate.Up)
	log.Infof("executed %d migrations", n)
	return err
}

// Reset resets the database
func (m *migration) Reset() error {
	n, err := migrate.Exec(m.db, m.driver, m.migrations, migrate.Down)
	log.Infof("database reset, rolled back %d steps", n)
	return err
}
