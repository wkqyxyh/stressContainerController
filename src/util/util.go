package util

import (
	"fmt"
	"os"
	"encoding/json"
)

type configuration struct {
    RampUpPeriod int
    RampUpFinalContainerNum int
    DurationPeriod int
    RampDownPeriod int
    ImageName string
}


func ReadConfigFile(filePath string) (configuration) {
	file, _ := os.Open(filePath)
	defer file.Close()
    decoder := json.NewDecoder(file)
    conf := configuration{}
    err := decoder.Decode(&conf)
    if err != nil {
        fmt.Println("Error:", err)
    }
    return conf
}