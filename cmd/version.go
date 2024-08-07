package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Pokedex-CLI version",
	Long:  "Every app has a version. This is Pokedex-CLI's",
	Run: func(cmd *cobra.Command, args []string) {
		version := `       ⣀⣤⣴⣶⣿⣿⣿⣿⣿⣿⣶⣦⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣠⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣄⠀⠀⠀⠀
⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀
⠀⠀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠀⠀
⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠁⠀⠀⠀⠉⠻⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀
⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠀⢠⣶⣶⣦⠀⠀⠸⣿⣿⣿⣿⣿⣿⣿⣿⡆
⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⢻⣿⠀⠀⠸⣿⣿⡿⠁⠀⢀⣿⠈⠉⠙⠛⠿⢿⣿⣷
⣿⣿⣿⣿⣿⠿⠋⠁⠀⠀⠈⢿⣦⡀⠀⠈⠉⠀⠀⢀⣾⠏⠀⠀⠀⠀⠀⠀⢸⣿
⢿⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠙⠿⣶⣤⣤⣴⡾⠟⠃⠀⠀⠀⠀⠀⠀⠀⢸⣿
⠸⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⠇
⠀⢻⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⡟⠀
⠀⠀⠻⣧⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⠟⠀⠀
⠀⠀⠀⠙⢿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⡿⠋⠀⠀⠀
⠀⠀⠀⠀⠀⠙⠻⣶⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⠟⠋⠀⠀⠀⠀⠀Pokedex-CLI v0.0.1 -- HEAD
⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠻⠷⣶⣶⣶⣶⣶⣶⠾⠟⠛⠉`
		fmt.Print(version)
	},
}
