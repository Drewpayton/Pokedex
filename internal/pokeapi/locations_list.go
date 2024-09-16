package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// This will be our api that we will call
	url := baseURL + "/location-area"
	
	if pageURL != nil {
		url = *pageURL
	}

	// Get request that we will call later to get our data 
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// calling the Do method from httpClient which sends our request and returns our data.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// defer runs after function and it will close the data
	defer resp.Body.Close()

	// Grabbing the data from the responses body.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Setting up our object that is going to hold the locations.
	locationResp := RespShallowLocations{}

	// Unmarshal is used to decode the data into our struct, in this case the locationResp.
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationResp, nil
} 