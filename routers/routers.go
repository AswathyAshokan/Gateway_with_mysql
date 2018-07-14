package routers

import (
	"github.com/astaxie/beego"
	"TestWork/controllers"
)

func init(){

	beego.Router("/numberTrace/:mobileNumber",&controllers.GatewayController{},"*:MobileGateway")

}
