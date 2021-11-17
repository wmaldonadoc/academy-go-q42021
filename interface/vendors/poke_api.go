package vendors

type Response struct {
	Name      string             `json:"name"`
	Pokemon   []Pokemon          `json:"pokemon_entries"`
	Abilities []PokemonAbilities `json:"abilities"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

type PokemonAbilities struct {
	Ability  PokemonAbility `json:"ability"`
	IsHidden bool           `json:"is_hidden"`
	Slot     int            `json:"slot"`
}

type PokemonAbility struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
