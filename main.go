package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	"os"
	"time"
)
// This stuff all initializes the Beego Orm database subsystem.
func init() {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// registerModels()

	if _, err := orm.GetDB(); err != nil {
		// Connect default DB
		driverName := os.Getenv("DB_DRIVER")
		dataSource := os.Getenv("DB_SOURCE")
		if driverName == "" {
			driverName = beego.AppConfig.String("dbDriver")
		}
		if dataSource == "" {
			dataSource = beego.AppConfig.String("dbSource")
		}
		maxIdle, _ := beego.AppConfig.Int("dbMaxIdle")
		maxConn, _ := beego.AppConfig.Int("dbMacConn")
		orm.RegisterDataBase("default", driverName, dataSource, maxIdle, maxConn)
		orm.DefaultTimeLoc = time.UTC
	}

	orm.Debug, _ = beego.AppConfig.Bool("ormDebug")
}


//var (
//	fileExcludedMap = make(map[string]int) // keeping track of excluded files
//	// it's useful for making log less noisy & for (Prometheus) monitoring
//)


func main() {
	var Lists []orm.ParamsList
	o := orm.NewOrm()
	query := "select classification_id from PalSpaces;"
	fmt.Println("Trying: " + query)
	num, err := o.Raw(query).ValuesList(&Lists)
	if ((err == nil) && (num > 0)) {
		fmt.Println(Lists)
	}
}
