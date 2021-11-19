package presenter_test

import (
	"net/http"

	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/infrastructure/api"
	"github.com/wmaldonadoc/academy-go-q42021/interface/presenter"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Presenter", func() {
	Context("Mapping pokemons", func() {
		It("Should return the Pokemon given without any mutation", func() {
			// Arrange
			var pokemon model.Pokemon
			faker.FakeData(&pokemon)
			// Act
			pr := presenter.NewPokemonPresenter()
			result := pr.ResponsePokemon(&pokemon)
			// Assert
			Expect(result.Ability).Should(Equal(pokemon.Ability))
			Expect(result.ID).Should(Equal(pokemon.ID))
			Expect(result.Name).Should(Equal(pokemon.Name))
		})
		It("Should return a Pokemon given a API response", func() {
			// Arrange
			client := api.NewApiClient()
			resp, _ := client.Get("https://pokeapi.co/api/v2/pokemon/ditto")
			// Act
			pr := presenter.NewPokemonPresenter()
			result := pr.ResponseMappedPokemonFromAPI(resp)
			// Assert
			Expect(resp.HTTPStatus).Should(Equal(http.StatusOK))
			Expect(result).ShouldNot(BeNil())
			Expect(result.ID).Should(BeNumerically(">", 0))
			Expect(result.Name).ShouldNot(BeNil())
			Expect(result.Ability).ShouldNot(BeNil())
		})
	})
})
