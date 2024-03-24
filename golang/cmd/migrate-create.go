package main

import (
	"fmt"
	"github.com/oooiik/test_09.03.2024/internal/database/migration"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		logger.Fatal("please indicate the driver and file name (cmd driver name).")
	}
	driver := os.Args[1]
	name := os.Args[2]

	mig := migration.New(driver)
	create, err := mig.Create(name)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(fmt.Sprintf("Success created file %s", create))
}
