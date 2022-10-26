package service

import (
	"strings"

	"github.com/reonbs/pokemoeapi/entity"
)

type PokemonCSVService interface {
	ProcessPokemonCSVData() ([]entity.Pokemon, error)
}

type csvservice struct{}

var (
	csvUtility CSVUtilityService
)

func NewPokemonCSVService(csvUtil CSVUtilityService) PokemonCSVService {
	csvUtility = csvUtil
	return &csvservice{}
}

func (*csvservice) ProcessPokemonCSVData() ([]entity.Pokemon, error) {
	pokemonsInCSV, err := csvUtility.GetPokemonFromCSV()
	if err != nil {
		return nil, err
	}

	var pokemons []entity.Pokemon

	for _, pokemonInCsv := range pokemonsInCSV {
		if pokemonInCsv.Lengendary {
			continue
		}

		if strings.ToLower(pokemonInCsv.Type1) == "ghost" || strings.ToLower(pokemonInCsv.Type2) == "ghost" {
			continue
		}

		if strings.ToLower(pokemonInCsv.Type1) == "steel" || strings.ToLower(pokemonInCsv.Type2) == "steel" {
			pokemonInCsv.HP = pokemonInCsv.HP * 2
		}

		if strings.ToLower(pokemonInCsv.Type1) == "fire" || strings.ToLower(pokemonInCsv.Type2) == "fire" {
			pokemonInCsv.Attack = pokemonInCsv.Attack * 0.9
		}
		//check to use slice for this
		if strings.ToLower(pokemonInCsv.Type1) == "bug" || strings.ToLower(pokemonInCsv.Type2) == "bug" || strings.ToLower(pokemonInCsv.Type1) == "flying" || strings.ToLower(pokemonInCsv.Type2) == "flying" {
			pokemonInCsv.SPAttack = pokemonInCsv.SPAttack * 1.1
		}

		if strings.HasPrefix(pokemonInCsv.Name, "G") {
			pokemonName := strings.Join(strings.Fields(pokemonInCsv.Name), "")
			pokemmonNamelength := len(pokemonName) - strings.Count(pokemonName, "G")
			pokemonInCsv.Defence = pokemonInCsv.Defence + float32(pokemmonNamelength*5)
		}

		pokemon := mapCSVTODBModel(pokemonInCsv)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func mapCSVTODBModel(pokemonCSV entity.PokemonCSV) entity.Pokemon {
	var pokemon entity.Pokemon

	pokemon.Name = pokemonCSV.Name
	pokemon.Type1 = pokemonCSV.Type1
	pokemon.Type2 = pokemonCSV.Type2
	pokemon.Total = pokemonCSV.Total
	pokemon.HP = pokemonCSV.HP
	pokemon.Attack = pokemonCSV.Attack
	pokemon.Defence = pokemonCSV.Defence
	pokemon.SPAttack = pokemonCSV.SPAttack
	pokemon.SPDefence = pokemonCSV.SPDefence
	pokemon.Speed = pokemonCSV.Speed
	pokemon.Generation = pokemonCSV.Generation
	pokemon.Lengendary = pokemonCSV.Lengendary

	return pokemon
}
