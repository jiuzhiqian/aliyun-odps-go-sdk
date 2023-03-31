package main

import (
	"log"
	"os"

	"github.com/jiuzhiqian/aliyun-odps-go-sdk/odps"
	"github.com/jiuzhiqian/aliyun-odps-go-sdk/odps/account"
)

func main() {
	conf, err := odps.NewConfigFromIni(os.Args[1])
	if err != nil {
		log.Fatalf("%+v", err)
	}

	aliAccount := account.NewAliyunAccount(conf.AccessId, conf.AccessKey)
	odpsIns := odps.NewOdps(aliAccount, conf.Endpoint)
	odpsIns.SetDefaultProjectName(conf.ProjectName)

	table := odpsIns.Table("data_type_demo")
	err = table.AddPartition(true, "p1=20,p2='hangzhou'")
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
