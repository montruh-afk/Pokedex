package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := BaseURL + DefaultLocation
	if pageURL != nil {
		
		url = *pageURL
	}
	// Cache check for stored available resource
	if val, ok := c.cash.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
		return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}


	// Network Call
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	c.cash.Add(url, dat)
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) Catch(id *string) (RespShallowLocations, error) {
	if id == nil {
		return RespShallowLocations{}, errors.New("Invalid Pokemon... see 'help' for usage")
	}
	
	url := BaseURL + "/pokemon/" + *id
	// Cache check for stored available resource
	if val, ok := c.cash.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
		return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	// Network Call
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	c.cash.Add(url, dat)
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		fmt.Println("Error processing returned json")
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}