package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)



func (c *Client) GetPokemon(name string) (Pokemon,error){
	url := baseURL + "/pokemon/" + name


	if val, ok := c.cache.Get(url); ok {
		var pokemonResp Pokemon

		err := json.Unmarshal(val,&pokemonResp)
		if err != nil {
			return Pokemon{},err
		}

		return pokemonResp, nil
	}

	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	var pokemonResp Pokemon
	if err = json.Unmarshal(data, &pokemonResp); err != nil{
		return Pokemon{}, err
	}


	c.cache.Add(url,data)
	return pokemonResp, nil
} 