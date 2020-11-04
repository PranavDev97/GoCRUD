package conf

import (
 	"github.com/joho/godotenv"
 	"github.com/kelseyhightower/envconfig"
	 "fmt"
	 //"os"
)

type Vars struct {
	Username 	string 		`envconfig:"MYSQLUSER" required:"true"`
	Password	string 		`envconfig:"MYSQLPASSWORD" required:"true"`
	DB			string 		`envconfig:"MYSQLDB" required:true`
}

type ServerConfiguraion struct {
	Port	string		`envconfig:"PORT" required:"true"` 
}

var MysqlData Vars
var Conf ServerConfiguraion

func init(){
	err := godotenv.Load("conf/conf.env")
  	if err != nil {
    	fmt.Println("ERROR: ",err)
	}
	//fmt.Println(os.Getenv("GOCRUD_MYSQLUSER"))
	appname := "GOCRUD"  
	erro := envconfig.Process(appname, &MysqlData)
	if erro != nil {
		fmt.Println("ERROR: ",erro)
	}
	er := envconfig.Process(appname, &Conf)
	if er != nil {
		fmt.Println("ERROR: ",er)
	}
	//fmt.Println(Conf.Username)
}

// func Test(){
// 	fmt.Println("conf")
// }
