package main

import (
	"fmt"
	"io"
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
	sql := "select * from data_type_demo where p1=20 and p2='hangzhou';"

	sqlTask := odps.NewSqlTask("select", sql, "", nil)
	ins, err := sqlTask.RunInOdps(odpsIns, odpsIns.DefaultProjectName())
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = ins.WaitForSuccess()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	csvReader, err := sqlTask.GetSelectResultAsCsv(ins, true)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%+v", err)
		}
		fmt.Printf("%v\n", record)
	}
}
