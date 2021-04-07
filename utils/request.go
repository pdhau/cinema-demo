package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func GetJSONBody(c echo.Context) map[string]interface{} {
	requestBody, err := c.Request().GetBody()
	if err != nil {
		panic(err)
	}
	requestBodyBytes, err := ioutil.ReadAll(requestBody)
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(requestBodyBytes, &data)
	if err != nil {
		panic(err)
	}
	return data
}
