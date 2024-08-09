package api

import (
	"encoding/json"

	"github.com/jpradass/pokedex-cli/api/models"
)

var baseURL = "https://pokeapi.co/api/v2/"

func ListAllPokemons() (*models.PokemonListResponse, error) {
	res, err := getRequest(baseURL + "pokemon?limit=1500")
	if err != nil {
		return nil, err
	}

	var result *models.PokemonListResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func GetPokemonInfo(name string) (*models.PokemonInfo, error) {
	res, err := getRequest(baseURL + "pokemon/" + name)
	if err != nil {
		return nil, err
	}

	var result *models.PokemonInfo
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
