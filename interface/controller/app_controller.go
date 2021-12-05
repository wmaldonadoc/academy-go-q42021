package controller

// AppController - Holds the abstractions of the project's controllers.
type AppController struct {
	Pokemon interface{ PokemonController }
	Health  interface{ HealthController }
}
