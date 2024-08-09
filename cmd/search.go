package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	fzf "github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"

	"github.com/jpradass/pokedex-cli/api"
	"github.com/jpradass/pokedex-cli/api/models"
)

func init() {
	rootCmd.AddCommand(searchCommand)
}

var searchCommand = &cobra.Command{
	Use:   "search",
	Short: "Pokedex-CLI search engine",
	Long:  "with search you can search any pokemon you pass as argument",
	Run: func(cmd *cobra.Command, args []string) {
		pokemons, err := api.ListAllPokemons()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		var names []string
		for _, pokemon := range pokemons.Results {
			names = append(names, pokemon.Name)
		}

		f, err := fzf.New()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		idxs, err := f.Find(names, func(i int) string { return names[i] })
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		result := pokemons.Results[idxs[0]]
		pinfo, err := api.GetPokemonInfo(result.Name)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// fmt.Println(pokemons.Results[idxs[0]])
		displayPokemonInfo(strings.ToUpper(string(result.Name[0]))+result.Name[1:], pinfo)
		// displayPokemon(struct {
		// 	Name string
		// 	Url  string
		// }(pokemons.Results[idxs[0]]))
	},
}

func displayPokemonInfo(pname string, pinfo *models.PokemonInfo) {
	in := fmt.Sprintf(`# %s`, pname)
	out, _ := glamour.Render(in, "light")
	fmt.Print(out)
}
