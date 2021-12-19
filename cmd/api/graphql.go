package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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
	"search": &graphql.Field{
		Type:        graphql.NewList(songType),
		Description: "Search songs by title",
		Args: graphql.FieldConfigArgument{
			"titleContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var theList []*models.Song
			search, ok := params.Args["titleContains"].(string)
			if ok {
				for _, currentSong := range songs {
					if strings.Contains(strings.ToLower(currentSong.Title), strings.ToLower(search)) {
						theList = append(theList, currentSong)
					}
				}
			}
			return theList, nil
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
			"artist": &graphql.Field{
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

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		app.errorJSON(w, errors.New("failed to create schema"))
		log.Println(err)
		return
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		app.errorJSON(w, errors.New(fmt.Sprintf("failed: %+v", res.Errors)))
	}

	j, _ := json.MarshalIndent(res, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
