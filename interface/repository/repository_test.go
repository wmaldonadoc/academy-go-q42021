package repository_test

import (
	"fmt"
	"os"

	"github.com/wmaldonadoc/academy-go-q42021/constants"
	"github.com/wmaldonadoc/academy-go-q42021/domain/model"
	"github.com/wmaldonadoc/academy-go-q42021/interface/repository"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	var pokemons []*model.Pokemon

	BeforeSuite(func() {
		for i := 0; i < 10; i++ {
			pokemon := model.Pokemon{
				ID:      i,
				Name:    faker.Name(),
				Ability: faker.Word(),
			}
			pokemons = append(pokemons, &pokemon)
		}
	})

	Describe("Pokemon repository", func() {
		Context("Finding a pokemon by ID", func() {
			It("Should return the pokemon record with the matching ID", func() {
				//Arrange
				id, _ := faker.RandomInt(1, 10)
				//Act
				repo := repository.NewPokemonRepository(pokemons)
				result, err := repo.FindByID(id[0])
				//Assert
				fmt.Printf("ID generated %v", id[0])
				Expect(result).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
				Expect(result.ID).Should(Equal(id[0]))
			})
			It("Should return a NotFoundError when the ID passed doesn't exists", func() {
				//Arrange
				id, _ := faker.RandomInt(11, 20)
				//Act
				repo := repository.NewPokemonRepository(pokemons)
				result, err := repo.FindByID(id[0])
				//Assert
				Expect(result).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
				Expect(err.Code).Should(Equal(constants.NotFoundExceptionCode))
			})
		})
		Context("Storing a new pokemon", func() {
			It("Should return a WritingCSVError whit a bad file location", func() {
				//Arrange
				id, _ := faker.RandomInt(0, 10)
				poke := model.Pokemon{
					ID:      id[0],
					Name:    faker.Name(),
					Ability: faker.Word(),
				}
				//Act
				repo := repository.NewPokemonRepository(pokemons)
				result, err := repo.CreateOne(&poke)
				//Assert
				Expect(result).Should(BeNil())
				Expect(err).ShouldNot(BeNil())
				Expect(err.Code).Should(Equal(constants.WritingCSVFileExceptionCode))
			})
			It("Should return a WritingCSVError whit a bad file location", func() {
				//Arrange
				os.Setenv("FILE_LOCATION", "../../infrastructure/datastore/data-test.csv")
				id, _ := faker.RandomInt(0, 10)
				poke := model.Pokemon{
					ID:      id[0],
					Name:    faker.Name(),
					Ability: faker.Word(),
				}
				//Act
				repo := repository.NewPokemonRepository(pokemons)
				result, err := repo.CreateOne(&poke)
				//Assert
				Expect(result).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
				Expect(result.ID).Should(Equal(id[0]))
			})
		})
	})
})
