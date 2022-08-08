package controllers

import (
	"farmacare/pokemon/lib/database"
	"farmacare/pokemon/models"
	"fmt"
	"os"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllCapturedPokemon(c echo.Context) error {
	pokemons, e := database.GetCapturedPokemon()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   pokemons,
	})
}

// func CaptureWildPokemon(c echo.Context) error {
// 	reponse, er := http.Get("https://pokeapi.co/api/v2/pokemon/1/")

// 	if er != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, er.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "success",
// 		"data":   reponse,
// 	})
// }

func CaptureWildPokemon(c echo.Context) error {
	wild, e := database.GetWildPokemon()
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   wild,
	})

	// response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=100000&offset=0")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var resObj models.Response
	// json.Unmarshal(responseData, &resObj)

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"status": "success",
	// 	"data":   resObj,
	// })
}

func FilterPokemon(c echo.Context) error {
	filterModel := models.PokemonFilter{}
	c.Bind(&filterModel)
	filtered, e := database.GetFilteredPokemon(filterModel)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   filtered,
	})
}

func DeletePokemonAbility(c echo.Context) error {
	IdDel := models.ReqId{}
	c.Bind(&IdDel)
	deleted, e := database.DeleteAbility(IdDel.AbilityId, IdDel.PokemonId)

	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   deleted,
	})
}

func AddPokemonAbility(c echo.Context) error {
	IdIns := models.ReqId{}
	c.Bind(&IdIns)
	inserted, e := database.AddAbility(IdIns.AbilityId, IdIns.PokemonId)

	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   inserted,
	})
}

func GetPokemonRanking(c echo.Context) error {
	pokemons, e := database.PokemonRank()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   pokemons,
	})
}
