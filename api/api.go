package api

import (
	"Les-Crackhead_groupie_tracker/structure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetTokenList() []structure.Token {

	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=bitcoin&names=Bitcoin&symbols=btc&category=layer-1&price_change_percentage=1h"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, url, nil)

	if errReq != nil {
		fmt.Println("une erreur est survenue :", errReq.Error())
	}

	res, errResp := httpClient.Do(req)

	if errResp != nil {
		fmt.Println("une erreur est survenue : ", errResp.Error())
		return []structure.Token{}
	}

	defer res.Body.Close()

	body, errBody := io.ReadAll(res.Body)

	if errBody != nil {
		fmt.Println("une erreur est survenue,", errBody.Error())
		return []structure.Token{}
	}

	//fmt.Println(string(body))

	var decodeData []structure.Token

	json.Unmarshal(body, &decodeData)
	//Exemple
	// for i := range decodeData.Albums {
	// 	fmt.Println("Nom :", decodeData.Albums[i].Name)
	// 	fmt.Println("Release date :", decodeData.Albums[i].Releasedate)
	// 	fmt.Println("Total tracks :", decodeData.Albums[i].Totaltrack)
	// }

	// for _, token := range decodeData {
	// 	fmt.Printf("%s : %.2f $\n", token.Symbol, token.CurrentPrice)
	// }

	return decodeData

}
