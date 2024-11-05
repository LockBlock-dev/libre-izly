package main

import (
	"fmt"
	"os"

	"github.com/LockBlock-dev/libre-izly/core"
	"github.com/LockBlock-dev/libre-izly/lib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var User string
var Password string
var ActivationCode string
var Language string

func main() {
	rootCmd := &cobra.Command{
		Use:   "libre-izly",
		Short: "Libre and open source Izly client implementation",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	login := &cobra.Command{
		Use:   "login",
		Short: "Authenticate on Izly",
		Long:  `Either provide a password (--password/-p) to send the 2FA or provide the 2FA code (--code/-c) to fully log in.`,
		Run: func(cmd *cobra.Command, args []string) {
			user, err := cmd.Flags().GetString("user")
			if err != nil {
				cmd.Help()
				os.Exit(1)
			}

			password, err := cmd.Flags().GetString("password")
			if err != nil {
				cmd.Help()
				os.Exit(1)
			}

			code, err := cmd.Flags().GetString("code")
			if err != nil {
				cmd.Help()
				os.Exit(1)
			}

			language, err := cmd.Flags().GetString("language")
			if err != nil {
				cmd.Help()
				os.Exit(1)
			}

			if password == "" && code == "" {
				cmd.Help()
				os.Exit(1)
			}

			c := lib.NewSoapClient()

			if code != "" {
				actCode, err := lib.FetchActivationCode(user, code)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				res, err := c.LogonSecondStep(lib.NewLogonSecondStepParams(user, actCode, language))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				fmt.Printf("%+v\n", res)
			}

			if password != "" {
				_, err := c.Logon(lib.NewLogonParams(user, password, language))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		},
	}

	login.Flags().StringVarP(&User, "user", "u", "", "Your Izly user id (required)")
	login.Flags().StringVarP(&Password, "password", "p", "", "Your Izly password")
	login.Flags().StringVarP(&ActivationCode, "code", "c", "", "Your Izly activation code (used for 2FA)")
	login.Flags().StringVarP(&Language, "language", "l", "fr", "Your desired language")
	login.MarkFlagRequired("user")

	// ToDo: qr validation command?
	qr := &cobra.Command{
		Use:   "qr",
		Short: "Generate a payment QR code",
		Run: func(cmd *cobra.Command, args []string) {
			if !lib.IsAuthDataPersisted() {
				fmt.Println("You must be logged in to Izly to generate a payment QR code!")
				os.Exit(1)
			}

			qr, err := lib.GenerateQRCodeDataWithPersistedAuthentification(
				core.QR_CODE_MODE_IZLY,
				core.QR_CODE_VERSION_THREE,
			)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(qr)
		},
	}

	clean := &cobra.Command{
		Use:   "clean",
		Short: "Clean libre-izly files",
		Long:  "Remove persisted data on the disk.",
		Run: func(cmd *cobra.Command, args []string) {
			if !lib.IsAuthDataPersisted() {
				return
			}

			err := lib.DeletePersistedAuthData()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(login)
	rootCmd.AddCommand(qr)
	rootCmd.AddCommand(clean)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
