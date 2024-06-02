package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wildanfaz/go-market/configs"
	"github.com/wildanfaz/go-market/internal/routers"
	"github.com/wildanfaz/go-market/migrations"
)

var email string

func InitCmd(ctx context.Context) {
	var rootCmd = &cobra.Command{
		Short: "Go Market",
	}

	rootCmd.PersistentFlags().StringVar(&email, "email", "", "Email address of user")

	rootCmd.AddCommand(startFiberApp, addBalance)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}

var startFiberApp = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		routers.InitRouter()
	},
}

var addBalance = &cobra.Command{
	Use:   "add-balance",
	Short: "Add user balance",
	Run: func(cmd *cobra.Command, args []string) {
		config := configs.InitConfig()

		db := configs.InitMySql(config.DatabaseDSN)

		err := migrations.AddBalance(db, email)
		if err != nil {
			panic(err)
		}

		fmt.Println("Add balance success")
	},
}