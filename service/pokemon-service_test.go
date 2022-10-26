package service

import (
	"errors"
	"testing"

	"github.com/reonbs/pokemoeapi/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPokemonRepository struct {
	mock.Mock
}

type MockPokemonPokeMonCSVService struct {
	mock.Mock
}

func (mock *MockPokemonRepository) SaveAll(pokemons []entity.Pokemon) ([]entity.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Pokemon), args.Error(1)
}

func (mock *MockPokemonRepository) FindAll(queryParameter map[string]string) ([]entity.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Pokemon), args.Error(1)
}

func (mock *MockPokemonPokeMonCSVService) ProcessPokemonCSVData() ([]entity.Pokemon, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Pokemon), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockPokemonRepository)
	mockCsvService := new(MockPokemonPokeMonCSVService)

	var identifier int64 = 1

	pokemon := entity.Pokemon{
		ID:         1,
		Name:       "Bulbasaur",
		Type1:      "Grass",
		Type2:      "Poison",
		Total:      19,
		HP:         10,
		Attack:     10,
		Defence:    10,
		SPAttack:   10,
		SPDefence:  10,
		Speed:      10,
		Generation: 10,
		Lengendary: false,
	}

	//setup expectations
	mockRepo.On("FindAll").Return([]entity.Pokemon{pokemon}, nil)

	pokemonService := NewPokemonService(mockRepo, mockCsvService)

	result, _ := pokemonService.FindAll(map[string]string{"hp[eq]": "100"})

	//Mock Assertion : Behavioural
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Bulbasaur", result[0].Name)
	assert.Equal(t, "Grass", result[0].Type1)
}

func TestMigratePokemonData(t *testing.T) {
	mockRepo := new(MockPokemonRepository)
	mockCsvService := new(MockPokemonPokeMonCSVService)

	pokemon := entity.Pokemon{
		Name:       "Bulbasaur",
		Type1:      "Grass",
		Type2:      "Poison",
		Total:      19,
		HP:         10,
		Attack:     10,
		Defence:    10,
		SPAttack:   10,
		SPDefence:  10,
		Speed:      10,
		Generation: 10,
		Lengendary: false,
	}

	//setup expectations
	mockRepo.On("SaveAll").Return([]entity.Pokemon{pokemon}, nil)
	mockCsvService.On("ProcessPokemonCSVData").Return([]entity.Pokemon{pokemon}, nil)

	pokemonService := NewPokemonService(mockRepo, mockCsvService)

	err := pokemonService.MigratePokemonData()

	//Mock Assertion : Behavioural
	mockRepo.AssertExpectations(t)
	mockCsvService.AssertExpectations(t)

	//Data Assertion
	assert.Nil(t, err)
}

func TestMigratePokemonDataNoDataToMigrate(t *testing.T) {
	mockRepo := new(MockPokemonRepository)
	mockCsvService := new(MockPokemonPokeMonCSVService)

	//setup expectations
	mockCsvService.On("ProcessPokemonCSVData").Return([]entity.Pokemon{}, nil)

	pokemonService := NewPokemonService(mockRepo, mockCsvService)

	err := pokemonService.MigratePokemonData()

	//Mock Assertion : Behavioural
	mockRepo.AssertNotCalled(t, "SaveAll")
	mockCsvService.AssertExpectations(t)

	//Data Assertion
	assert.NotNil(t, err)
}

func TestMigratePokemonDataFailedToSaveAll(t *testing.T) {
	mockRepo := new(MockPokemonRepository)
	mockCsvService := new(MockPokemonPokeMonCSVService)

	pokemon := entity.Pokemon{
		ID:         1,
		Name:       "Bulbasaur",
		Type1:      "Grass",
		Type2:      "Poison",
		Total:      19,
		HP:         10,
		Attack:     10,
		Defence:    10,
		SPAttack:   10,
		SPDefence:  10,
		Speed:      10,
		Generation: 10,
		Lengendary: false,
	}

	//setup expectations
	mockRepo.On("SaveAll").Return([]entity.Pokemon{pokemon}, errors.New("failed"))
	mockCsvService.On("ProcessPokemonCSVData").Return([]entity.Pokemon{pokemon}, nil)

	pokemonService := NewPokemonService(mockRepo, mockCsvService)

	err := pokemonService.MigratePokemonData()

	//Mock Assertion : Behavioural
	mockRepo.AssertExpectations(t)
	mockCsvService.AssertExpectations(t)

	//Data Assertion
	assert.NotNil(t, err)
}
