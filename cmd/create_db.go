package cmd

import (
	"fmt"
	"maker/config"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/cobra"
)

// createCmd represents the migrate command
var createdbCmd = &cobra.Command{
	Use:   "create_db",
	Short: "create_db creates a database user and database from database parameters declared in config",
	Long:  `create_db creates a database user and database from database parameters declared in config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create_db called")
		p := config.GetPostgresConfig()

		// connection to db as postgres superuser
		dbSuper := config.GetPostgresSuperUserConnection()
		defer dbSuper.Close()

		CreateDatabaseUserIfNotExist(dbSuper, p)
		CreateDatabaseIfNotExist(dbSuper, p)
	},
}

func init() {
	rootCmd.AddCommand(createdbCmd)
}

func CreateDatabaseIfNotExist(db *pg.DB, p *config.PostgresConfig) {
	statement := fmt.Sprintf(`SELECT 1 AS result FROM pg_database WHERE datname = '%s';`, p.Database)
	res, _ := db.Exec(statement)
	if res.RowsReturned() == 0 {
		fmt.Println("creating database")
		statement = fmt.Sprintf(`CREATE DATABASE %s WITH OWNER %s;`, p.Database, p.User)
		_, err := db.Exec(statement)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf(`Created database %s`, p.Database)
		}
	} else {
		fmt.Printf("Database named %s already exists\n", p.Database)
	}
}

func CreateDatabaseUserIfNotExist(db *pg.DB, p *config.PostgresConfig) {
	statement := fmt.Sprintf(`SELECT * FROM pg_roles WHERE rolname = '%s';`, p.User)
	res, _ := db.Exec(statement)
	if res.RowsReturned() == 0 {
		statement = fmt.Sprintf(`CREATE USER %s WITH PASSWORD '%s';`, p.User, p.Password)
		_, err := db.Exec(statement)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf(`Created user %s`, p.User)
		}
	} else {
		fmt.Printf("Database user %s already exists\n", p.User)
	}

}
