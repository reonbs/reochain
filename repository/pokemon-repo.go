package repository

import "github.com/reonbs/pokemoeapi/entity"

type PokemonRepository interface {
	SaveAll(pokemons []entity.Pokemon) ([]entity.Pokemon, error)
	FindAll(queryParameter map[string]string) ([]entity.Pokemon, error)
}
