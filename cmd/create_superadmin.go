package cmd

import (
	"fmt"
	"regexp"

	"github.com/gogjango/gjango/config"
	"github.com/gogjango/gjango/manager"
	"github.com/gogjango/gjango/repository"
	"github.com/gogjango/gjango/secret"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var email string
var password string
var createSuperAdminCmd = &cobra.Command{
	Use:   "create_superadmin",
	Short: "create_superadmin creates a superadmin user that has access to manage all other users in the system",
	Long:  `create_superadmin creates a superadmin user that has access to manage all other users in the system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create_superadmin called")

		email, _ = cmd.Flags().GetString("email")
		fmt.Println(email)
		if !validateEmail(email) {
			fmt.Println("Invalid email provided; superadmin user not created")
			return
		}

		password, _ = cmd.Flags().GetString("password")
		fmt.Println(password)
		if password == "" {
			password, _ = secret.GenerateRandomString(16)
			fmt.Printf("No password provided, so we have generated one for you: %s\n", password)
		}

		db := config.GetConnection()
		log, _ := zap.NewDevelopment()
		defer log.Sync()
		accountRepo := repository.NewAccountRepo(db, log, secret.New())
		roleRepo := repository.NewRoleRepo(db, log)

		m := manager.NewManager(accountRepo, roleRepo, db)
		m.CreateSuperAdmin(email, password)
	},
}

func init() {
	localFlags := createSuperAdminCmd.Flags()
	localFlags.StringVarP(&email, "email", "e", "", "SuperAdmin user's email")
	localFlags.StringVarP(&password, "password", "p", "", "SuperAdmin user's password")
	createSuperAdminCmd.MarkFlagRequired("email")
	rootCmd.AddCommand(createSuperAdminCmd)
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
