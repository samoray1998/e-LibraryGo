package schemas

import (
	"errors"
	"testGoGraph/models"

	"github.com/graphql-go/graphql"
)

var genreType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Genre",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

const (
	ScienceFictionID         = 1
	MysteryID                = 2
	ThrillerID               = 3
	RomanceID                = 4
	FantasyID                = 5
	HorrorID                 = 6
	HistoricalFictionID      = 7
	ComedyID                 = 8
	DramaID                  = 9
	ActionAdventureID        = 10
	BiographyAutobiographyID = 11
)

// Define a slice of genres to represent all possible book genres
var genres = []models.Genre{
	{ID: ScienceFictionID, Name: "Science Fiction"},
	{ID: MysteryID, Name: "Mystery"},
	{ID: ThrillerID, Name: "Thriller"},
	{ID: RomanceID, Name: "Romance"},
	{ID: FantasyID, Name: "Fantasy"},
	{ID: HorrorID, Name: "Horror"},
	{ID: HistoricalFictionID, Name: "Historical Fiction"},
	{ID: ComedyID, Name: "Comedy"},
	{ID: DramaID, Name: "Drama"},
	{ID: ActionAdventureID, Name: "Action/Adventure"},
	{ID: BiographyAutobiographyID, Name: "Biography/Autobiography"},
}

func addGenre(params graphql.ResolveParams) (interface{}, error) {
	name, ok := params.Args["name"].(string)
	if ok {
		newGenre := models.Genre{ID: len(genres) + 1, Name: name}
		genres = append(genres, newGenre)
		return newGenre, nil
	}
	return nil, errors.New("the name parameter is required")
}

func getGenres(params graphql.ResolveParams) (interface{}, error) {
	return genres, nil
}
func getGenreById(params graphql.ResolveParams) (interface{}, error) {
	genreId, ok := params.Args["id"].(int)
	if ok {
		for _, genre := range genres {
			if genre.ID == genreId {
				return genre, nil
			}
		}
		return nil, errors.New("genre not found")
	}

	return nil, errors.New("Id not valid")
}
