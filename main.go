package main

import (
	"backend/config"
	"backend/routers"
	"backend/utils"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)
func init(){ 
	utils.InitializeLogger()

}


func main(){
	viper_config, err := utils.LoadConfig("./utils");

	if err !=nil {
		utils.SugarLogger.Error("error in loading config file with viper");
	}
	fmt.Println(viper_config.Db.DbDriver);
	config.ConnectToDB(); 
	portNumber := ":" + os.Getenv("PORT");
	fmt.Println(portNumber);
	fmt.Println(portNumber);

	r := routers.SetupRouter();

	r.Run(portNumber);
	
}