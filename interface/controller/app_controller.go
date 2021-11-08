package controller

type AppController struct {
	Pokemon interface{ PokemonController }
	Health  interface{ HealthController }
}
