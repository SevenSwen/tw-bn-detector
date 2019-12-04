package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func InterpretResponse(response io.ReadCloser) map[string]map[string]interface{} {
	res := make(map[string]map[string]interface{})
	b, _ := ioutil.ReadAll(response)
	json.Unmarshal([]byte(b), &res)
	return res
}

//func InterpretMarketInfoResponse(response io.ReadCloser) models.MarketInfo {
//	res := models.MarketInfo{}
//	b, _ := ioutil.ReadAll(response)
//	json.Unmarshal([]byte(b), &res)
//	return res
//}
//
//func InterpretBarChartsResponse(response io.ReadCloser) models.HystoResponse {
//	res := models.HystoResponse{}
//	b, _ := ioutil.ReadAll(response)
//	json.Unmarshal([]byte(b), &res)
//	return res
//}
//
//func GetBars(url string, symbol string) []models.Bar {
//	response, err := http.Get(url)
//	Check(fmt.Sprintf("Cannot receive response %s", url), err)
//	res := InterpretBarChartsResponse(response.Body)
//	if res.Response != "Success" {
//		panic(fmt.Sprintf("Error get data: %s! Response is not succeed", symbol))
//	}
//	return res.Data
//}
