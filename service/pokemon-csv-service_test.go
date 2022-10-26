package service

import (
	"testing"

	"github.com/reonbs/pokemoeapi/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCSVUtility struct {
	mock.Mock
}

func (mock *MockCSVUtility) GetPokemonFromCSV() ([]entity.PokemonCSV, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.PokemonCSV), args.Error(1)
}

func TestProcessPokemonCSVData_NameStartWithG(t *testing.T) {
	mockCSVUtility := new(MockCSVUtility)

	pokemonsCsv := []entity.PokemonCSV{
		{
			Name:       "Gastly",
			Type1:      "Water",
			Type2:      "Poison",
			Total:      10,
			HP:         10,
			Attack:     10,
			Defence:    10,
			SPAttack:   10,
			SPDefence:  10,
			Speed:      10,
			Generation: 10,
			Lengendary: false,
		},
	}

	var expectedDefense float32 = 35

	//setup expectations
	mockCSVUtility.On("GetPokemonFromCSV").Return(pokemonsCsv, nil)

	pokemonCSVService := NewPokemonCSVService(mockCSVUtility)

	pokemons, _ := pokemonCSVService.ProcessPokemonCSVData()

	//Mock Assertion : Behavioural
	mockCSVUtility.AssertExpectations(t)

	assert.Equal(t, 1, len(pokemons))
	assert.Equal(t, expectedDefense, pokemons[0].Defence)
}

func TestProcessPokemonCSVData_ExcludeLengendary(t *testing.T) {
	mockCSVUtility := new(MockCSVUtility)

	pokemonsCsv := []entity.PokemonCSV{
		{
			Name:       "Gastly",
			Type1:      "Water",
			Type2:      "Poison",
			Total:      10,
			HP:         10,
			Attack:     10,
			Defence:    10,
			SPAttack:   10,
			SPDefence:  10,
			Speed:      10,
			Generation: 10,
			Lengendary: true,
		},
	}

	//setup expectations
	mockCSVUtility.On("GetPokemonFromCSV").Return(pokemonsCsv, nil)

	pokemonCSVService := NewPokemonCSVService(mockCSVUtility)

	pokemons, _ := pokemonCSVService.ProcessPokemonCSVData()

	//Mock Assertion : Behavioural
	mockCSVUtility.AssertExpectations(t)

	assert.Equal(t, 0, len(pokemons))
}

func TestProcessPokemonCSVData_ExcludeTypeGhost(t *testing.T) {
	mockCSVUtility := new(MockCSVUtility)

	pokemonsCsv := []entity.PokemonCSV{
		{
			Name:       "Gastly",
			Type1:      "Ghost",
			Type2:      "Poison",
			Total:      10,
			HP:         10,
			Attack:     10,
			Defence:    10,
			SPAttack:   10,
			SPDefence:  10,
			Speed:      10,
			Generation: 10,
			Lengendary: true,
		},
	}

	//setup expectations
	mockCSVUtility.On("GetPokemonFromCSV").Return(pokemonsCsv, nil)

	pokemonCSVService := NewPokemonCSVService(mockCSVUtility)

	pokemons, _ := pokemonCSVService.ProcessPokemonCSVData()

	//Mock Assertion : Behavioural
	mockCSVUtility.AssertExpectations(t)

	assert.Equal(t, 0, len(pokemons))
}

func TestProcessPokemonCSVData_SteelType_DoubleHP(t *testing.T) {
	mockCSVUtility := new(MockCSVUtility)

	pokemonsCsv := []entity.PokemonCSV{
		{
			Name:       "Gastly",
			Type1:      "Steel",
			Type2:      "Steel",
			Total:      10,
			HP:         10,
			Attack:     10,
			Defence:    10,
			SPAttack:   10,
			SPDefence:  10,
			Speed:      10,
			Generation: 10,
			Lengendary: false,
		},
	}

	var expectedHp float32 = 20

	//setup expectations
	mockCSVUtility.On("GetPokemonFromCSV").Return(pokemonsCsv, nil)

	pokemonCSVService := NewPokemonCSVService(mockCSVUtility)

	pokemons, _ := pokemonCSVService.ProcessPokemonCSVData()

	//Mock Assertion : Behavioural
	mockCSVUtility.AssertExpectations(t)

	assert.Equal(t, expectedHp, pokemons[0].HP)
}

func TestProcessPokemonCSVData_FireType_LowerAttack(t *testing.T) {
	mockCSVUtility := new(MockCSVUtility)

	pokemonsCsv := []entity.PokemonCSV{
		{
			Name:       "Gastly",
			Type1:      "Fire",
			Type2:      "Fire",
			Total:      10,
			HP:         10,
			Attack:     10,
			Defence:    10,
			SPAttack:   10,
			SPDefence:  10,
			Speed:      10,
			Generation: 10,
			Lengendary: false,
		},
	}

	//lower attack by 10%
	var expectedAttack float32 = pokemonsCsv[0].Attack * 0.9

	//setup expectations
	mockCSVUtility.On("GetPokemonFromCSV").Return(pokemonsCsv, nil)

	pokemonCSVService := NewPokemonCSVService(mockCSVUtility)

	pokemons, _ := pokemonCSVService.ProcessPokemonCSVData()

	//Mock Assertion : Behavioural
	mockCSVUtility.AssertExpectations(t)

	assert.Equal(t, expectedAttack, pokemons[0].Attack)
}

func TestProcessPokemonCSVData_BugType_IncreaseAttackSpeed(t *testing.T) {
	mockCSVUtility := new(MockCSVUtility)

	pokemonsCsv := []entity.PokemonCSV{
		{
			Name:       "Gastly",
			Type1:      "Bug",
			Type2:      "Flyin",
			Total:      10,
			HP:         10,
			Attack:     10,
			Defence:    10,
			SPAttack:   10,
			SPDefence:  10,
			Speed:      10,
			Generation: 10,
			Lengendary: false,
		},
	}

	//lower attack by 10%
	var expectedAttack float32 = pokemonsCsv[0].SPAttack * 1.1

	//setup expectations
	mockCSVUtility.On("GetPokemonFromCSV").Return(pokemonsCsv, nil)

	pokemonCSVService := NewPokemonCSVService(mockCSVUtility)

	pokemons, _ := pokemonCSVService.ProcessPokemonCSVData()

	//Mock Assertion : Behavioural
	mockCSVUtility.AssertExpectations(t)

	assert.Equal(t, expectedAttack, pokemons[0].SPAttack)
}
