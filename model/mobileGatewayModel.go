package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)
type MobileDataResult struct {
	Prefix 		string  ` json:”Prefix” `
	GatewayName string ` json:”GatewayName” `
	IpAddress 	string ` json:”IpAddress” `
}

func InsertIntoDb(prefixString string)(bool,MobileDataResult){
	db := dbConn()

	//create prefix table
	_, err := db.Query("CREATE  TABLE IF NOT EXISTS `PrefixTable`(`prefix` VARCHAR(8),`gateway` VARCHAR (30))")
	if err != nil {
		panic(err.Error())
	}
	prefixArray :=[]int{123,1234,9194}
	gateWay :=[]string{"airtel","vodafone","tata"}
	for i:=0;i<len(prefixArray);i++{
		prefix :=prefixArray[i]
		gateway :=gateWay[i]
		insForm, err := db.Prepare("INSERT INTO PrefixTable(prefix,gateway) VALUES(?,?)")
			if err != nil {
				panic(err.Error())
		}
		insForm.Exec(prefix,gateway)
		}
	//defer db.Close()
	//create ip table
	_, err = db.Query("CREATE  TABLE IF NOT EXISTS `IpAddress`(`gateway` VARCHAR(20),`ip` VARCHAR (50))")
	if err != nil {
		panic(err.Error())
	}
	ipArray :=[]string{"12.12.12.12, 13.13.13.13","14.14.14.14 ","15.15.15.15, 16.16.16.16"}
		gateWayData :=[]string{"airtel","vodafone","tata"}
		for i:=0;i<len(gateWayData);i++{
			ip :=ipArray[i]
			gateway :=gateWayData[i]
			insForm, err := db.Prepare("INSERT INTO IpAddress(gateway,ip) VALUES(?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(gateway,ip)
		}
		//find gateway

		preDB, err := db.Query("SELECT prefix,gateway FROM PrefixTable")
		var prefixNo string
		var gatewayNo string
		var ResultGateway string
		if err != nil {
			panic(err.Error())
		}
		for preDB.Next() {
			err := preDB.Scan(&prefixNo, &gatewayNo)
			if err != nil {
				log.Fatal(err)
			}
			if prefixNo==prefixString{

				ResultGateway=gatewayNo
				}
			}
			if len(ResultGateway)==0{
				ResultGateway="airtel"
				prefixString="123"
			}
			//select ipaddress
		ipDB, err := db.Query("SELECT * FROM IpAddress WHERE gateway=?", ResultGateway)
		if err != nil {
			panic(err.Error())
		}
		Result :=MobileDataResult{}
		var resultedIp string

		for ipDB.Next() {
			var gateWayIp string
			var resultIp string
			err = ipDB.Scan(&gateWayIp, &resultIp)
			if err != nil {
				panic(err.Error())
			}
			resultedIp=resultIp
		}
		Result.GatewayName =ResultGateway
		Result.IpAddress=resultedIp
		Result.Prefix=prefixString
		fmt.Println("result",Result)

	return true,Result
}
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "MobileData"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}