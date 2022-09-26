package services

import (
	"Assigment3/helpers"
	"Assigment3/models"
	"Assigment3/params"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type WebService struct {
}

func NewWebService() *WebService {
	return &WebService{}
}

func (w *WebService) GetStatus() (params.Data, error) {

	//read data
	jsonFile, err := os.OpenFile("data.json", os.O_RDWR, 0644)
	if err != nil {
		return params.Data{}, err
	}
	defer jsonFile.Close()

	var dataFromFile models.Data

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &dataFromFile)
	if err != nil {
		return params.Data{}, err
	}

	water := helpers.RandInt(1, 100)
	wind := helpers.RandInt(1, 100)

	dataFromFile.Status.Water = water
	dataFromFile.Status.Wind = wind

	//out
	var itmData = params.ItemData{
		Water:       water,
		StatusWater: helpers.CheckStatusWater(water),
		Wind:        wind,
		StatusWind:  helpers.CheckStatusWind(wind),
	}

	// out
	var statusData = params.Data{
		Status: itmData,
	}

	//update data.json
	output, err := json.Marshal(dataFromFile)
	if err != nil {
		return params.Data{}, err
	}
	log.Default().Println("output : ", string(output))

	// make data.json empty before new one
	if err := os.Truncate("./data.json", 0); err != nil {
		log.Println("Failed to truncate : %v", err)
	}
	Write(output)

	return statusData, err
}
func Write(output []byte) (params.Data, error) {
	jsonFile, err := os.OpenFile("data.json", os.O_RDWR, 0644)
	if err != nil {
		return params.Data{}, err
	}
	defer jsonFile.Close()

	_, err = jsonFile.WriteAt(output, 0)
	if err != nil {
		return params.Data{}, err
	}
	return params.Data{}, err
}