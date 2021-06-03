package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type WarpKey struct {
	MqttURL        string `json:"mqttURL"`
	MqttMyId       string `json:"mqttMyId"`
	MqttPairId     string `json:"mqttPairId"`
	SerialBaudRate string `json:"serialBaudRate"`
}

func main() {

	jsonFile, err := os.Open("warpKey.json")
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var warpKey WarpKey
	json.Unmarshal(byteValue, &warpKey)

	fmt.Println(warpKey.MqttURL)
	fmt.Println(warpKey.MqttMyId)

	// get unstructed data
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result["serialBaudRate"])

}
