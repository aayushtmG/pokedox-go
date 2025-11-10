package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURl != nil {
		url = *pageURl
	}	
	data, exist := c.Cache.Get(url)	
	if exist {
		locationsResp := RespShallowLocations{}
        err := json.Unmarshal(data,&locationsResp)
		if err != nil {
			return RespShallowLocations{},err
		}
		return locationsResp, nil
    }

	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return RespShallowLocations{},err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{},err
	}
	defer resp.Body.Close()
	data,err = io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{},err
	}

	c.Cache.Add(url,data)
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data,&locationsResp)
	if err != nil {
		return RespShallowLocations{},err
	}

	return locationsResp, nil
}
