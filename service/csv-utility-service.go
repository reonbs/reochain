package service

import (
	"io/ioutil"
	"log"

	"github.com/gocarina/gocsv"
	"github.com/reonbs/pokemoeapi/entity"
)

type csvUtil struct{}

type CSVUtilityService interface {
	GetPokemonFromCSV() ([]entity.PokemonCSV, error)
}

func NewCsvUtilityService() CSVUtilityService {
	return &csvUtil{}
}

func (*csvUtil) GetPokemonFromCSV() ([]entity.PokemonCSV, error) {
	bytes, err := ioutil.ReadFile("Data/pokemon.csv")
	if err != nil {
		log.Fatalf("Failed to read pokenmon csv %v", err)
		return nil, err
	}

	var pokemons []entity.PokemonCSV
	gocsv.UnmarshalBytes(bytes, &pokemons)

	return pokemons, nil
}
