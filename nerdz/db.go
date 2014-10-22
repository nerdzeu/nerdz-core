package nerdz

import (
	"flag"
	"fmt"
	"github.com/galeone/gorm"
	_ "github.com/lib/pq"
	"os"
)

var db gorm.DB

func init() {
	flag.Parse()
	args := flag.Args()
	envVar := os.Getenv("CONF_FILE")

	var file string
	if len(args) == 1 {
		file = args[0]
	} else if envVar != "" {
		file = envVar
	} else {
		panic(fmt.Sprintln("Configuration file is required.\nUse: CONF_FILE environment variable or cli args"))
	}

	var err error

	if err = InitConfiguration(file); err != nil {
		panic(fmt.Sprintf("[!] %v\n", err))
	}

	var connectionString string
	if connectionString, err = Configuration.ConnectionString(); err != nil {
		panic(err.Error())
	}

	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database: '%v'\n", err))
	}

	enableLog := os.Getenv("ENABLE_LOG")
	if enableLog != "" {
		db.LogMode(true)
	}
}
