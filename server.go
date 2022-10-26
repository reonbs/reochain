package main

import (
	"fmt"
	"net/http"

	"github.com/reonbs/pokemoeapi/controller"
	router "github.com/reonbs/pokemoeapi/http"
	"github.com/reonbs/pokemoeapi/repository"
	"github.com/reonbs/pokemoeapi/service"
	"github.com/reonbs/pokemoeapi/utilities"
)

var (
	csvUtil           service.CSVUtilityService    = service.NewCsvUtilityService()
	csvservice        service.PokemonCSVService    = service.NewPokemonCSVService(csvUtil)
	util              utilities.Utilities          = utilities.NewUtilities()
	pokemonRepository repository.PokemonRepository = repository.NewPostgresRepository()
	pokemonService    service.PokemonService       = service.NewPokemonService(pokemonRepository, csvservice)
	pokemonController controller.PokemonController = controller.NewPokemonController(pokemonService, util)
	httpRouter        router.Router                = router.NewMuxRouter()
)

func main() {
	//run on start up to migrate pokemon data from to db
	pokemonService.MigratePokemonData()
	const port string = ":8000"
	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and Running...")
	})

	httpRouter.GET("/pokemon", pokemonController.GetPokemon)
	httpRouter.SERVE(port)
}
