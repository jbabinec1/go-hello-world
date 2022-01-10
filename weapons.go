package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var result map[string]string

func readAndUpdateJson() {
	//original json file
	loadAssets, _ := ioutil.ReadFile("test-modify-assets.json")

	json.Unmarshal([]byte(loadAssets), &result)

	result["case 2"] = "case 3"

	fmt.Println(result)

	loadJson, _ := json.Marshal(result)

	fmt.Println(string(loadJson))

	ioutil.WriteFile("test-modify-assets.json", []byte(string(loadJson)), 0666)

}
