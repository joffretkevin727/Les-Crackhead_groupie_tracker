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
	urls := []string{
		"https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=bitcoin&names=Bitcoin&symbols=btc&category=layer-1&price_change_percentage=1h",
		"https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=ethereum&names=Ethereum&symbols=eth&category=layer-2&price_change_percentage=1h",
	}

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	allTokens := []structure.Token{}

	for _, url := range urls {
		req, errReq := http.NewRequest(http.MethodGet, url, nil)
		if errReq != nil {
			fmt.Println("une erreur est survenue :", errReq.Error())
			continue
		}

		res, errResp := httpClient.Do(req)
		if errResp != nil {
			fmt.Println("une erreur est survenue :", errResp.Error())
			continue
		}

		defer res.Body.Close()

		body, errBody := io.ReadAll(res.Body)
		if errBody != nil {
			fmt.Println("une erreur est survenue :", errBody.Error())
			continue
		}

		var decodeData []structure.Token
		if err := json.Unmarshal(body, &decodeData); err != nil {
			fmt.Println("erreur JSON :", err.Error())
			continue
		}

		// Ajoute les tokens de cette URL au slice global
		allTokens = append(allTokens, decodeData...)
	}

	return allTokens
}

func GetTokenInfo(tokenName string) structure.TokenInfo {

	url := "https://api.coingecko.com/api/v3/coins/" + tokenName

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
		return structure.TokenInfo{}
	}

	defer res.Body.Close()

	body, errBody := io.ReadAll(res.Body)

	if errBody != nil {
		fmt.Println("une erreur est survenue,", errBody.Error())
		return structure.TokenInfo{}
	}

	//fmt.Println(string(body))

	var decodeData structure.TokenInfo

	json.Unmarshal(body, &decodeData)

	decodeData.DescriptionFinal = decodeData.Description.En
	decodeData.Image = decodeData.ImgUrl.Large
	decodeData.WebUrl = decodeData.Links.Homepage[0]

	return decodeData

}
