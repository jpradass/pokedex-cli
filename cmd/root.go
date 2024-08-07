package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pokedex",
	Short: "Pokedex-CLI is the pokedex in your terminal",
	Long:  `Transform your terminal in your pokedex where you can find useful information about your favorite pokemons`,
	Run: func(cmd *cobra.Command, args []string) {
		in := `# Pokedex-CLI
Transform your terminal in your pokedex where you can find useful information about your favorite pokemons.

Powered by [PokeApi](https://pokeapi.co)
    `
		out, _ := glamour.Render(in, "dark")
		fmt.Print(out)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
