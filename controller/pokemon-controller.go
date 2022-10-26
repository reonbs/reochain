package controller

import (
	"net/http"

	"github.com/reonbs/pokemoeapi/service"
	"github.com/reonbs/pokemoeapi/utilities"
)

type controller struct{}

var (
	pokemonService service.PokemonService
	util           utilities.Utilities
)

type PokemonController interface {
	GetPokemon(response http.ResponseWriter, request *http.Request)
}

func NewPokemonController(service service.PokemonService, utility utilities.Utilities) PokemonController {
	pokemonService = service
	util = utility
	return &controller{}
}

func (*controller) GetPokemon(response http.ResponseWriter, request *http.Request) {
	queryParameters := util.GetUrlParams(request)
	pokemons, err := pokemonService.FindAll(queryParameters)
	if err != nil {
		util.ErrorJson(response, err)
		return
	}

	util.WriteJSON(response, http.StatusOK, pokemons, "pokemons")
}
