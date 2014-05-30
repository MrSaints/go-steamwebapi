package godoto

import (
	"encoding/json"
)

type Heroes struct {
	Heroes []Hero `json:"heroes"`
	Count int `json:"count"`
}

type Hero struct {
	Name string `json:"name"`
	Id int `json:"id"`
	LocalizedName string `json:"localized_name,omitempty"`
}

func GetHeroes() (heroes Heroes) {
	api := WebAPI{Endpoint: "/IEconDOTA2_570/GetHeroes"}
	result := api.GetResult()

	heroes = Heroes{}
	err := json.Unmarshal(result.Data, &heroes)
	pError(err)

	return
}