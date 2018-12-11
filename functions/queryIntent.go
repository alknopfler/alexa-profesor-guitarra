package function

import (
	"fmt"
	"googlemaps.github.io/maps"
	geo "github.com/martinlindhe/google-geolocate"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	cfg "github.com/alknopfler/alexa-cenas/config"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"context"
)

func QueryIntent(context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response){
	log.Println("Tengo apikey: " + cfg.ApiKey)

	resp := doRequest(http.MethodGet, cfg.URLGetCountryCode1 + aContext.System.Device.DeviceID + cfg.URLGetCountryCode2, aContext.System.APIAccessToken, nil )

	log.Println(resp.StatusCode)
	log.Println(resp.Body)

	var outData CountryResponse
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body,&outData)


	client := geo.NewGoogleGeo(cfg.ApiKey)
	res, _ := client.Geocode(outData.PostalCode+","+outData.CountryCode)
	fmt.Println(res)

	c, _ := maps.NewClient(maps.WithAPIKey(cfg.ApiKey))
	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: res.Lat,
			Lng: res.Lng,
		},
		Radius: 3000,
		Name: request.Intent.Slots["query"].Value,

	}

	respOut , _:= c.NearbySearch(context, r)
	for _,val := range respOut.Results{
		fmt.Println(val.Name+" "+val.Vicinity)
		fmt.Println(val.Rating)
	}
}
