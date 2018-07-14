package controllers

import (
	"TestWork/model"
	"fmt"
	"encoding/json"
)

type GatewayController struct {
	BaseController
}

func (c *GatewayController)MobileGateway()string{

	mobileNumber := c.Ctx.Input.Param(":mobileNumber")
	prefix := mobileNumber[:4]
	result :="true"

	dbStatus,Result :=model.InsertIntoDb(prefix)
	switch dbStatus {
	case true:
		w :=c.Ctx.ResponseWriter
		json.NewEncoder(w).Encode(Result)
		jsonResp, _ := json.Marshal(Result)
		fmt.Println(string(jsonResp))
		result :=string(jsonResp)
		return result
	}
	return result

}
