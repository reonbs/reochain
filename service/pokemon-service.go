package service

import (
	"errors"
	"log"

	"github.com/reonbs/pokemoeapi/entity"
	"github.com/reonbs/pokemoeapi/repository"
	
)

type PokemonService interface {
	FindAll(queryParameter map[string]string) ([]entity.Pokemon, error)
	MigratePokemonData() error
}

type service struct{}

var (
	repo       repository.PokemonRepository
	csvService PokemonCSVService
)

func NewPokemonService(repository repository.PokemonRepository, csv PokemonCSVService) PokemonService {
	repo = repository
	csvService = csv
	return &service{}
}

func (*service) MigratePokemonData() error {
	pokemons, err := csvService.ProcessPokemonCSVData()
	if err != nil {
		log.Println("unable to process csv data ")
		return err
	}

	if len(pokemons) == 0 {
		return errors.New("no records to migrate")
	}

	_, err = repo.SaveAll(pokemons)

	if err != nil {
		log.Printf("error saving pokemon to db %s", err)
		return err
	}

	return nil
}

func (*service) FindAll(queryParameter map[string]string) ([]entity.Pokemon, error) {
	return repo.FindAll(queryParameter)
}
