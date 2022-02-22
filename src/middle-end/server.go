package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Detail struct {
	Key        string `json:"key"`
	Title      string `json:"title"`
	Time       string `json:"time"`
	Instructor string `json:"instructor"`
}

type Day map[string][]Detail

type Room map[string]Day

type Buidling struct {
	Rooms map[string]Room
}

// Info -> Building: Room: Day: Details
type Info map[string]map[string]map[string][]Detail

// var userType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "details",
// 		Fields: graphql.Fields{
// 			"Room": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"Key": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"Title": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"Days": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"Times": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"Instructor": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 		},
// 	},
// )

// var queryType = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "Query",
// 		Fields: graphql.Fields{
// 			"key": &graphql.Field{
// 				Type: userType,
// 				Args: graphql.FieldConfigArgument{
// 					"id": &graphql.ArgumentConfig{
// 						Type: graphql.String,
// 					},
// 				},
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					idQuery, isOK := p.Args["id"].(string)
// 					if isOK {
// 						return data[idQuery], nil
// 					}
// 					return nil, nil
// 				},
// 			},
// 		},
// 	})

// var schema, _ = graphql.NewSchema(
// 	graphql.SchemaConfig{
// 		Query: queryType,
// 	},
// )

func main() {

	// Working
	// var dat Info

	// content, err := ioutil.ReadFile("out.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = json.Unmarshal(content, &dat)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// CTR := dat["KUPF"]

	// for rooms := range CTR {
	// 	fmt.Println(rooms)
	// }

	// for _, val := range CTR {
	// 	// fmt.Println("i: [", i, "]")
	// 	fmt.Println("val: [", val["M"][0].Time, "]")
	// }
	// fmt.Println(CTR)

	// Not working
	// var data Buidling
	// content, err := ioutil.ReadFile("out.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = json.Unmarshal(content, &data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(data.Rooms)

	var dat Info

	content, err := ioutil.ReadFile("out.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &dat)
	if err != nil {
		log.Fatal(err)
	}

	CTR := dat["CTR"]

	// for rooms := range CTR {
	// 	fmt.Println(rooms)
	// }

	// for _, val := range CTR {
	// 	// fmt.Println("i: [", i, "]")
	// 	fmt.Println("val: [", val["M"][0].Time, "]")
	// }
	fmt.Println(CTR)

	// _ = importJSONDataFromFile("out.json", &data)
	// enc := json.NewEncoder(os.Stdout)
	// enc.Encode(dat)

	// http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	// 	result := graphql.Do(graphql.Params{
	// 		Schema: schema,
	// 		RequestString: r.URL.Query().Get("query"),
	// 	})

	// 	if len(result.Errors) > 0 {
	// 		fmt.Printf("unexpected errors: %v", result.Errors)
	// 	}

	// 	json.NewEncoder(w).Encode(result)
	// })

	// fmt.Println("Now server is running on port 8080")
	// fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	// http.ListenAndServe(":8080", nil)
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
