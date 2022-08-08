package database

import (
	"encoding/json"
	"farmacare/pokemon/config"
	"farmacare/pokemon/models"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// func GetCapturedPokemon() ([]models.Pokemons, error) {
// 	// strQuery := "SELECT p.pokemon_name, p.height, p.weight, a.ability_name, s.species_name
// 	// FROM pokemon AS p
// 	// LEFT JOIN detail_abilities AS da
// 	// ON p.Id = da.pokemon_id
// 	// LEFT JOIN abilities AS a
// 	// ON da.abilities_id = a.id
// 	// LEFT JOIN detail_species as ds
// 	// ON p.Id = ds.pokemon_id
// 	// LEFT JOIN species as s
// 	// ON ds.species_id = s.id
// 	// "
// 	// strQuery := "SELECT p.pokemon_name, p.height, p.weight, a.ability_name, s.species_name FROM pokemon AS p LEFT JOIN detail_abilities AS da ON p.Id = da.pokemon_id LEFT JOIN abilities AS a ON da.abilities_id = a.id LEFT JOIN detail_species as ds ON p.Id = ds.pokemon_id LEFT JOIN species as s ON ds.species_id = s.id"

// 	var pokemons []models.Pokemons
// 	rows, err := config.DB.Query("SELECT * FROM pokemon")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var pok models.Pokemons
// 		if err := rows.Scan(&pok.Id, &pok.Pokemon_Name, &pok.Height, &pok.Weight,
// 			&pok.CreateAt, &pok.UpdateAt); err != nil {
// 			return pokemons, err
// 		}
// 		pokemons = append(pokemons, pok)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return pokemons, err
// 	}

// 	return pokemons, nil
// }
func GetCapturedPokemon() ([]models.ListPokemons, error) {

	strQuery := "SELECT p.id, p.pokemon_name, p.height, p.weight, a.ability_name, s.species_name FROM pokemon AS p LEFT JOIN detail_abilities AS da ON p.Id = da.pokemon_id LEFT JOIN abilities AS a ON da.abilities_id = a.id LEFT JOIN detail_species as ds ON p.Id = ds.pokemon_id LEFT JOIN species as s ON ds.species_id = s.id WHERE species_name IS NOT NULL"

	var pokemons []models.ListPokemons
	rows, err := config.DB.Query(strQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pok models.ListPokemons
		if err := rows.Scan(&pok.Id, &pok.Pokemon_Name, &pok.Height, &pok.Weight, &pok.Ability_Name, &pok.Species_Name); err != nil {
			return pokemons, err
		}
		pokemons = append(pokemons, pok)
	}
	if err = rows.Err(); err != nil {
		return pokemons, err
	}

	return pokemons, nil
}

func GetWildPokemon() (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	keyRand := fmt.Sprint(rand.Intn(1154))
	fmt.Println("https://pokeapi.co/api/v2/pokemon/" + keyRand + "/")

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + keyRand + "/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resObj models.ResDetailPokemon
	json.Unmarshal(responseData, &resObj)

	dt := time.Now()

	queryStmt := `INSERT INTO pokemon (pokemon_name, height, weight, "createAt", "updateAt") VALUES ('` + resObj.Name + `', ` + strconv.Itoa(resObj.Height) + `, ` + strconv.Itoa(resObj.Weight) + `, '` + string(dt.Format("01-02-2006")) + `', '` + string(dt.Format("01-02-2006")) + `') RETURNING Id`
	fmt.Println(queryStmt)
	rowsPokemon, er := config.DB.Query(queryStmt)
	if er != nil {
		panic(er)
	}
	defer rowsPokemon.Close()

	var pokemonsId int

	for rowsPokemon.Next() {
		var pok models.ResId
		if err := rowsPokemon.Scan(&pok.Id); err != nil {
			return pokemonsId, err
		}
		pokemonsId = pok.Id
	}

	//insert table detail abilities
	fmt.Println(resObj.Abilities)
	var abilityList []models.ResId
	for _, element := range resObj.Abilities {
		queryAbility := `SELECT Id FROM Abilities where ability_name = '` + element.Ability.Name + `'`
		rowAbility, errAb := config.DB.Query(queryAbility)
		if errAb != nil {
			panic(errAb)
		}
		defer rowAbility.Close()

		for rowAbility.Next() {
			var resId models.ResId
			if err := rowAbility.Scan(&resId.Id); err != nil {
				return resId, err
			}
			abilityList = append(abilityList, resId)
		}
	}

	//insert table detail species
	querySpecies := `SELECT Id FROM species where species_name = '` + resObj.Species.Name + `'`
	rowSpecies, errSp := config.DB.Query(querySpecies)
	if errSp != nil {
		panic(errSp)
	}

	defer rowSpecies.Close()
	var speciesId int
	for rowSpecies.Next() {
		var resId models.ResId
		if err := rowSpecies.Scan(&resId.Id); err != nil {
			return resId, err
		}
		speciesId = resId.Id
	}
	//insert detail abilities
	for _, element := range abilityList {
		queryDetailAbility := `INSERT INTO detail_abilities(abilities_id, pokemon_id, "createAt", "updateAt") VALUES (` + strconv.Itoa(element.Id) + `, ` + strconv.Itoa(pokemonsId) + `, '` + string(dt.Format("01-02-2006")) + `', '` + string(dt.Format("01-02-2006")) + `')`
		fmt.Println(queryDetailAbility)
		_, ers := config.DB.Query(queryDetailAbility)
		if ers != nil {
			panic(ers)
		}
	}

	//insert detail species
	queryDetailSpecies := `INSERT INTO detail_species(species_id, pokemon_id, "createAt", "updateAt") VALUES (` + strconv.Itoa(speciesId) + `, ` + strconv.Itoa(pokemonsId) + `, '` + string(dt.Format("01-02-2006")) + `', '` + string(dt.Format("01-02-2006")) + `')`
	fmt.Println(queryDetailSpecies)
	_, ers := config.DB.Query(queryDetailSpecies)
	if ers != nil {
		panic(ers)
	}

	return resObj, nil
}

func GetFilteredPokemon(cat models.PokemonFilter) ([]models.ListPokemons, error) {

	queryStmt := `SELECT p.id, p.pokemon_name, p.height, p.weight, a.ability_name, s.species_name FROM pokemon AS p LEFT JOIN detail_abilities AS da ON p.Id = da.pokemon_id LEFT JOIN abilities AS a ON da.abilities_id = a.id LEFT JOIN detail_species as ds ON p.Id = ds.pokemon_id LEFT JOIN species as s ON ds.species_id = s.id WHERE species_name IS NOT NULL`
	fmt.Println(len(cat.Name))
	fmt.Println(len(cat.Ability))
	fmt.Println(len(cat.Species))

	if len(cat.Name) > 0 {
		tempName := "('"
		for _, valName := range cat.Name {
			tempName = tempName + valName + "', '"
		}
		tempName = tempName[:len(tempName)-3] + ")"
		queryStmt = queryStmt + ` AND p.pokemon_name in ` + tempName
	}

	if len(cat.Height) > 0 {
		queryStmt = queryStmt + " AND p.height " + cat.Height
	}
	if len(cat.Weight) > 0 {
		queryStmt = queryStmt + " AND p.weight " + cat.Weight
	}
	if len(cat.Ability) > 0 {
		tempAbility := "('"
		for _, valName := range cat.Ability {
			tempAbility = tempAbility + valName + "', '"
		}
		tempAbility = tempAbility[:len(tempAbility)-3] + ")"
		queryStmt = queryStmt + ` AND a.ability_name in ` + tempAbility
	}
	if len(cat.Species) > 0 {
		tempSpecies := "('"
		for _, valName := range cat.Species {
			tempSpecies = tempSpecies + valName + "', '"
		}
		tempSpecies = tempSpecies[:len(tempSpecies)-3] + ")"
		queryStmt = queryStmt + ` AND s.species_name in ` + tempSpecies
	}

	var pokemons []models.ListPokemons
	rows, err := config.DB.Query(queryStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pok models.ListPokemons
		if err := rows.Scan(&pok.Id, &pok.Pokemon_Name, &pok.Height, &pok.Weight, &pok.Ability_Name, &pok.Species_Name); err != nil {
			return pokemons, err
		}
		pokemons = append(pokemons, pok)
	}
	if err = rows.Err(); err != nil {
		return pokemons, err
	}

	return pokemons, nil

}

func DeleteAbility(abilityId int, pokemonId int) (int, error) {
	var resCode int

	queryDelete := `DELETE FROM detail_abilities WHERE abilities_id = ` + strconv.Itoa(abilityId) + ` and pokemon_id = ` + strconv.Itoa(pokemonId) + `;	`
	_, erd := config.DB.Query(queryDelete)
	if erd != nil {
		panic(erd)
	}
	resCode = 200
	return resCode, nil
}

func AddAbility(abilityId int, pokemonId int) (int, error) {
	var resCode int
	dt := time.Now()

	queryAdd := `INSERT INTO public.detail_abilities(abilities_id, pokemon_id, "createAt", "updateAt") VALUES (` + strconv.Itoa(abilityId) + `, ` + strconv.Itoa(abilityId) + `, '` + string(dt.Format("01-02-2006")) + `', '` + string(dt.Format("01-02-2006")) + `');`
	_, erd := config.DB.Query(queryAdd)
	if erd != nil {
		panic(erd)
	}
	resCode = 200
	return resCode, nil
}

func PokemonRank() ([]models.PokemonRank, error) {
	queryRank := `SELECT p.pokemon_name as name, COUNT(*) as total FROM pokemon AS p LEFT JOIN detail_abilities AS da ON p.Id = da.pokemon_id LEFT JOIN abilities AS a ON da.abilities_id = a.id  LEFT JOIN detail_species as ds ON p.Id = ds.pokemon_id LEFT JOIN species as s ON ds.species_id = s.id WHERE species_name IS NOT NULL GROUP BY p.pokemon_name ORDER BY total DESC	`

	var pokemons []models.PokemonRank
	rows, err := config.DB.Query(queryRank)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pok models.PokemonRank
		if err := rows.Scan(&pok.Name, &pok.AbilityTotal); err != nil {
			return pokemons, err
		}
		pokemons = append(pokemons, pok)
	}
	if err = rows.Err(); err != nil {
		return pokemons, err
	}

	return pokemons, nil
}
