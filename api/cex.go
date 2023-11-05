package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"nexflare.com/crypto/model"
)

const url = "https://cex.io/api/ticker/%s/USD"


// reference to model.Rate is returned so that nil can be returned if there is a error
func GetRate(currency string) (*model.Rate, error){
	res, err := http.Get(fmt.Sprintf(url, strings.ToUpper(currency)))
	var cryptoResponse model.CexResponseObject
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		// We use ReadAll here because we get stream of data from UDP or TCP protocol
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		// The value needs to be passed by reference for data to be stored in obj after unmashalling
		err = json.Unmarshal(bodyBytes, &cryptoResponse)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Printing right value %v \n", cryptoResponse)
	} else {
		return nil, fmt.Errorf("Something went wrong %v", res.StatusCode)
	}

	rate := model.Rate{Currency: currency, Price: cryptoResponse.Ask}

	return &rate, nil
}