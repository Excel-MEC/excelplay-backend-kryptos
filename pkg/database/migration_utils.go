package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// DBParams is used to construct the connection string
type DBParams struct {
	host     string
	port     string
	user     string
	password string
	name     string
}

type migration struct {
	id      int
	fname   string
	content string
}

// Migrate is the master function used to run migrations
// path should be a string giving the path to the directory containing the SQL files.
// config is a pointer to an instance of DBParams, and is used to establish connection to the database.
// driverName is the name of the driver to be passed to sql.Open()
func Migrate(path string, config *DBParams, driverName string) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.host,
		config.port,
		config.user,
		config.password,
		config.name,
	)

	db, err := connectDB(driverName, connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	lastMigrationID, err := getLastMigrationID(db)
	if err != nil {
		return err
	}

	migrationsToApply, err := getMigrationsToApply(lastMigrationID, path)
	if err != nil {
		return err
	}
	err = applyMigrations(db, migrationsToApply)
	return nil
}

func applyMigrations(db *sql.DB, m []migration) error {
	lastMigration := 0
	for _, migration := range m {
		fmt.Println("Applying " + migration.fname)
		_, err := db.Exec(migration.content)
		lastMigration = max(lastMigration, migration.id)
		if err != nil {
			return err
		}
	}
	db.Exec("UPDATE meta_migration_data SET id = $1", lastMigration)
	return nil
}

// Migrations must follow the format number_filename.extension
// This function extracts and returns the migration number
func getMigrationIDFromFileName(name string) (int, error) {
	value, err := strconv.Atoi(strings.Split(name, "_")[0])
	if err != nil {
		return -1, errors.Wrap(err, "Invalid name for migration file '"+name+"'")
	}
	return value, nil
}

// This function checks every file in the migrations directory against the migration ID
// and returns an array of migration objects that has to be applied to the database.
func getMigrationsToApply(id int, path string) ([]migration, error) {
	var files []string
	err := filepath.Walk(path, func(fpath string, info os.FileInfo, err error) error {
		// Only attempt to read files
		if !info.IsDir() {
			mid, err := getMigrationIDFromFileName(filepath.Base(fpath))
			if err != nil {
				return err
			}
			if mid > id {
				files = append(files, fpath)
			}
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "Could not read directory with migration files")
	}
	var migrations []migration
	// Read content of each file in it's entirety as a string
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, errors.Wrap(err, "Could not read migration files")
		}
		// Can ignore error here because every file in files[] is validated to have a correct ID above
		mid, _ := getMigrationIDFromFileName(filepath.Base(file))
		migrations = append(migrations, migration{mid, filepath.Base(file), string(content)})
	}
	return migrations, nil
}

// This function creates the migrations meta table if it doesn't already exist,
// then fetches the latest migration ID as stored in the database and returns this as an int
func getLastMigrationID(db *sql.DB) (int, error) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS meta_migration_data (id INTEGER DEFAULT 0)")
	_, err = db.Exec("INSERT INTO meta_migration_data (id) SELECT 0 WHERE NOT EXISTS (SELECT * FROM meta_migration_data)")
	if err != nil {
		return -1, errors.Wrap(err, "Could not create table in database to track applied migrations")
	}
	rows, err := db.Query("SELECT * FROM meta_migration_data")
	if err != nil {
		return -1, errors.Wrap(err, "Could not read last applied migration number")
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return -1, errors.Wrap(err, "Could not read last applied migration ID")
		}
	}
	err = rows.Err()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func connectDB(driverName, connectionString string) (*sql.DB, error) {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Could not establish connection to database")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Could not ping the db to establish conection")
	}
	return db, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
