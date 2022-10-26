package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/reonbs/pokemoeapi/entity"
)

type posgresrepo struct{}

func NewPostgresRepository() PokemonRepository {
	return &posgresrepo{}
}

func GetOperator(operator string) string {
	operatorStrs := make(map[string]string)
	operatorStrs["eq"] = "="
	operatorStrs["ne"] = "<>"
	operatorStrs["gt"] = ">"
	operatorStrs["lt"] = "<"
	operatorStrs["gte"] = ">="
	operatorStrs["lte"] = "<="

	return operatorStrs[operator]
}

func (*posgresrepo) FindAll(queryParameter map[string]string) ([]entity.Pokemon, error) {
	db, err := openDb()

	if err != nil {
		return nil, err
	}

	var queryFromParams []string

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select id, Name, Type1, Type2, Total, HP, Attack, Defence, SPAttack, SPDefence, Speed, Generation, Lengendary from pokemon"
	page := 1
	pageSize := 10
	count := 0
	condition := ""

	for k, v := range queryParameter {
		if strings.Contains(k, "|") {
			column := strings.Split(k, "|")[0]
			operator := strings.Split(k, "|")
			operatorValue := GetOperator(operator[1])

			if operatorValue == "" && strings.Contains(k, "|") {
				return nil, errors.New("invalid operator")
			}

			str := fmt.Sprintf(" %s %s %s ", column, operatorValue, v)
			if count != 0 {
				condition = "AND"
			}

			queryFromParams = append(queryFromParams, condition+str)
			count++
		}

		if strings.ToLower(k) == "page" {
			pageNo, _ := strconv.Atoi(v)
			page = pageNo
		}

		if strings.ToLower(k) == "pagesize" {
			perPage, _ := strconv.Atoi(v)
			pageSize = perPage
		}

		if strings.ToLower(k) == "name" && strings.Trim(v, " ") != "" {
			if count != 0 {
				condition = "AND"
			}

			str := fmt.Sprintf(" %s lower(name) LIKE '%%%s%%' ", condition, strings.ToLower(v))

			queryFromParams = append(queryFromParams, str)
			count++
		}
	}

	if len(queryFromParams) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(queryFromParams, " "))
	}

	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, pageSize, (page-1)*pageSize)

	log.Println(query)

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var pokemons []entity.Pokemon

	for rows.Next() {
		var pokemon entity.Pokemon

		err = rows.Scan(
			&pokemon.ID,
			&pokemon.Name,
			&pokemon.Type1,
			&pokemon.Type2,
			&pokemon.Total,
			&pokemon.HP,
			&pokemon.Attack,
			&pokemon.Defence,
			&pokemon.SPAttack,
			&pokemon.SPDefence,
			&pokemon.Speed,
			&pokemon.Generation,
			&pokemon.Lengendary,
		)

		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func (*posgresrepo) SaveAll(pokemons []entity.Pokemon) ([]entity.Pokemon, error) {

	db, err := openDb()

	if err != nil {
		log.Fatalf("error while instantiating db %d", err)
		return pokemons, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for _, pokemon := range pokemons {

		stmt := `insert into pokemon (Name, Type1, Type2, Total, HP, Attack, Defence, SPAttack, SPDefence, Speed, Generation, Lengendary)
				 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

		_, err1 := db.ExecContext(ctx, stmt,
			pokemon.Name,
			pokemon.Type1,
			pokemon.Type2,
			pokemon.Total,
			pokemon.HP,
			pokemon.Attack,
			pokemon.Defence,
			pokemon.SPAttack,
			pokemon.SPDefence,
			pokemon.Speed,
			pokemon.Generation,
			pokemon.Lengendary,
		)

		if err1 != nil {
			log.Println(err1)
			log.Fatalf("error while saviing record to db %d ", err1)
			return pokemons, err1
		}
	}

	return pokemons, err
}

func openDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:Password@123@localhost/pokemondb?sslmode=disable")

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
