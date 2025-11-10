package pokeapi

import (
	"encoding/json"
	
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	Url string `json:"url"`
} 


type RespShallowLocationAreas struct{
	Id float64 `json:"id"`
	Name string `json:"name"`
	PokemonsFound []struct{
		P Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}


func (c *Client) FindPokemons(searchKey string) (RespShallowLocationAreas,error){
	url := baseURL + "/location-area/" + searchKey 

	//check if searchkey already exist in the cache
	 if val, ok := c.cache.Get(searchKey); ok {
		 locationAreas := RespShallowLocationAreas{}
			err := json.Unmarshal(val,&locationAreas)
		 if err!= nil {
			return RespShallowLocationAreas{},err
		 }	
			return locationAreas, nil
	 }
	
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		return RespShallowLocationAreas{},err
	}

	res, err :=  c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocationAreas{},err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocationAreas{},err
	}
	var resp RespShallowLocationAreas
	if err = json.Unmarshal(data,&resp); err != nil{
		return RespShallowLocationAreas{}, err
	}
	
	c.cache.Add(searchKey,data)
	return resp,nil 
}
