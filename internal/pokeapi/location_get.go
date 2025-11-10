package pokeapi

import (
	"encoding/json"

	"io"
	"net/http"
)


func (c *Client) GetLocation(locationName string) (Location,error){
	url := baseURL + "/location-area/" + locationName

	//check if searchkey already exist in the cache
	 if val, ok := c.cache.Get(locationName); ok {
		 locationResp := Location{}
			err := json.Unmarshal(val,&locationResp)
		 if err!= nil {
			return Location{},err
		 }	
			return locationResp, nil
	 }
	
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		return Location{},err
	}

	res, err :=  c.httpClient.Do(req)
	if err != nil {
		return Location{},err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{},err
	}
	var resp Location
	if err = json.Unmarshal(data,&resp); err != nil{
		return Location{}, err
	}
	
	c.cache.Add(locationName,data)
	return resp,nil 
}
