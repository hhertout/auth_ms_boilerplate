package migrations

import (
	"auth_ms/src/database"
	"database/sql"
	"errors"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Migration struct {
	dbPool   *sql.DB
	basePath string
}

func NewMigration(basePath string) *Migration {
	return &Migration{
		nil,
		basePath,
	}
}

func (m *Migration) Migrate(filename string) error {
	db, err := database.Connect()
	if err != nil {
		return errors.New("failed to connect to db")
	}
	m.dbPool = db.DbPool

	if err := m.migrateFromFile(filename); err != nil {
		return err
	}

	if err = m.dbPool.Close(); err != nil {
		return errors.New("failed to close db connection after executing migrations")
	}

	return nil
}

func (m *Migration) MigrateAll() error {
	db, err := database.Connect()
	if err != nil {
		log.Println("Failed to connect to db")
	}
	m.dbPool = db.DbPool

	migrationFiles, err := m.GetMigrationFiles("/")
	log.Println(migrationFiles)
	if err != nil {
		return errors.New("failed to retrieve migration files")
	}
	if len(migrationFiles) == 0 {
		log.Println("No migration file found ! To add one, run 'make migration-generate'.")
	} else {
		for _, f := range migrationFiles {
			err := m.migrateFromFile(f)
			if err != nil {
				return err
			}
		}
		log.Println("Migration complete, all migration file are executed")
	}

	if err = m.dbPool.Close(); err != nil {
		log.Println("Failed to close db connection after migration")
	}

	return nil
}

func (m *Migration) migrateFromFile(filename string) error {
	workingDir, _ := os.Getwd()
	fileOpen, err := os.Open(workingDir + m.basePath + "/migrations/" + filename)
	if err != nil {
		return err
	}
	defer fileOpen.Close()

	content, err := io.ReadAll(fileOpen)
	if err != nil {
		return err
	}

	queries := string(content)
	queriesSplit := strings.Split(queries, "--")

	for _, query := range queriesSplit {
		if strings.TrimSpace(query) == "" {
			continue
		}

		_, err = m.dbPool.Exec(query + ";")
		if err != nil {
			if err.Error() == "trigger set_viewed_param already exists" {
				continue
			}
			return err
		}
	}
	return nil
}

func (m *Migration) GetMigrationFiles(basePath string) ([]string, error) {
	var res []string
	baseDir := "migrations"
	workingDir, _ := os.Getwd()

	dir, err := os.ReadDir(workingDir + basePath + baseDir)
	if err != nil {
		return nil, err
	}

	for _, dirEntry := range dir {
		if !dirEntry.IsDir() && strings.HasSuffix(dirEntry.Name(), ".sql") {
			res = append(res, dirEntry.Name())
		}
	}

	sort.Strings(res)

	return res, nil
}
