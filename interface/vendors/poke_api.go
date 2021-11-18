package vendors

// Response - Holds part of the root level of PokeAPI response.
// @Reference - https://pokeapi.co/
type Response struct {
	Name      string             `json:"name"`
	Pokemon   []Pokemon          `json:"pokemon_entries"`
	Abilities []PokemonAbilities `json:"abilities"`
}

// Pokemon - Holds the pokemon part of PokeAPI response.
// @Reference - https://pokeapi.co/
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// PokemonSpecies - Holds the pokemon species part of PokeAPI response.
// @Reference - https://pokeapi.co/
type PokemonSpecies struct {
	Name string `json:"name"`
}

// PokemonAbilities - Holds the pokemon abilities of PokeAPI response.
// @Reference - https://pokeapi.co/
type PokemonAbilities struct {
	Ability  PokemonAbility `json:"ability"`
	IsHidden bool           `json:"is_hidden"`
	Slot     int            `json:"slot"`
}

// PokemonAbility - Holds the pokemon ability part of PokeAPI response.
// @Reference - https://pokeapi.co/
type PokemonAbility struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
