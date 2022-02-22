package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

type key struct {
	Days		string `json:"days"`
	Instructor	string `json:"instructor"`
	Location	string `json:"location"`
	Times		string `json:"times"`
	Title		string `json:"title"`
	Course 		string `json:"course"`
	Section		string `json:"section"`
}

var data map[string]key

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Key",
		Fields: graphql.Fields{
			"Days": &graphql.Field{
				Type: graphql.String,
			},
			"Instructor": &graphql.Field{
				Type: graphql.String,
			},
			"Location": &graphql.Field{
				Type: graphql.String,
			},
			"Times": &graphql.Field{
				Type: graphql.String,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Course": &graphql.Field{
				Type: graphql.String,
			},
			"Section": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)


var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"key": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})


var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig {
		Query: queryType,
	},
)	

func main() {
	_ = importJSONDataFromFile("out.json", &data)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema: schema,
			RequestString: r.URL.Query().Get("query"), 
		})

		if len(result.Errors) > 0 {
			fmt.Printf("unexpected errors: %v", result.Errors)
		}

		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(":8080", nil)
}


//Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}