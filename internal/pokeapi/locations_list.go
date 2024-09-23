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

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationResp, nil
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

	c.cache.Add(url, data)
	return locationResp, nil
} 


func (c *Client) GetLocation(areaName string) (Location, error) {
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}
	
	c.cache.Add(url, data)
	return locationResp, nil
}