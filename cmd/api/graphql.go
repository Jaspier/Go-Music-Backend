package main

import (
	"backend/models"
	"io"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

var songs []*models.Song

// graphql schema definition
var fields = graphql.Fields{
	"song": &graphql.Field{
		Type:        songType,
		Description: "Get song by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				for _, song := range songs {
					if song.ID == id {
						return song, nil
					}
				}
			}
			return nil, nil
		},
	},
	"list": &graphql.Field{
		Type:        graphql.NewList(songType),
		Description: "Get all songs",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return songs, nil
		},
	},
}

var songType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"release_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"duration": &graphql.Field{
				Type: graphql.Int,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
			"riaa_rating": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

func (app *application) songsGraphQL(w http.ResponseWriter, r *http.Request) {
	songs, _ = app.models.DB.All()

	q, err := io.ReadAll(r.Body)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	query := string(q)

	log.Println(query)
}