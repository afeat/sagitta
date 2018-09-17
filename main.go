package main

import (
	"context"
	"encoding/json"
	"github.com/afeat/sagitta/model"
	"github.com/arangodb/go-driver"
	arango "github.com/arangodb/go-driver/http"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://palmtale:ec52266f9b460717440686116f5b166779c67f7d@flightxml.flightaware.com/json/FlightXML3/AirportInfo?airport_code=ZSSS"
	var out []byte
	if res, err := http.Get(url); err != nil {
		return
	} else {
		if out, err = ioutil.ReadAll(res.Body); err != nil {
			return
		}
	}

	airport := &struct {
		AirportInfoResult *model.Airport `json:"AirportInfoResult"`
	}{}
	json.Unmarshal(out, airport)

	conn, err := arango.NewConnection(arango.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		// Handle error
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("aquila", "aquila"),
	})
	if err != nil {
		// Handle error
	}
	ctx := context.Background()
	db, err := client.Database(ctx, "river_drum")
	col, err := db.Collection(ctx, "airport")
	meta, err := col.CreateDocument(ctx, airport.AirportInfoResult.ToAirportDoc())
	if err != nil {
		println(err)
	}
	println(meta.Key)
}
