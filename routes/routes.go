package routes

import (
	"farmacare/pokemon/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	//pokemon owned
	e.GET("/pokemons", controllers.GetAllCapturedPokemon)  //list pokemon dimiliki
	e.POST("/filter", controllers.FilterPokemon)           //filter pokemon
	e.DELETE("/ability", controllers.DeletePokemonAbility) //delete ability
	e.POST("/ability", controllers.AddPokemonAbility)      //add ability
	e.GET("/rank", controllers.GetPokemonRanking)          //pokemon rangking
	//wild pokemon
	e.GET("/wild", controllers.CaptureWildPokemon) //tangkap pokemon liar (random)

	return e
}
