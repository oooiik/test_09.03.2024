package migration

import (
	"fmt"
	"github.com/oooiik/test_09.03.2024/internal/config"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const BaseDir = "/app/migrations"
const SlotPath = "./slot.sql"

type Interface interface {
	Create(name string) (string, error)
	//Status() []string // TODO
	//Up(count *uint) error // TODO
	//Down(count *uint) error // TODO
}

type migration struct {
	driver string
	addr   string
}

func New(driver string) Interface {
	var addr string
	switch driver {
	case "postgres":
	case "pg":
		driver, addr = config.Load().Postgres.Driver()
		break
	case "clickhouse":
	case "ch":
		driver, addr = config.Load().Clickhouse.Driver()
		break
	default:
		logger.Fatal(fmt.Sprintf("not found driver: %s", driver))
		return nil
	}

	return &migration{
		driver: driver,
		addr:   addr,
	}
}

func (m *migration) dir() string {
	return path.Join(BaseDir, m.driver)
}

func (m *migration) files() ([]string, error) {
	var res []string

	files, err := os.ReadDir(m.dir())
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			res = append(res, file.Name())
		}
	}

	return res, nil
}

func (m *migration) Create(name string) (string, error) {
	nowTime := time.Now()
	slotTimeFormat := nowTime.String()
	fileTimeFormat := nowTime.Format("20060102150405") // YYYYMMDDHHMMSS
	fileName := fmt.Sprintf("%s_%s.sql", fileTimeFormat, name)

	if _, err := os.Stat(m.dir()); os.IsNotExist(err) {
		err := os.MkdirAll(m.dir(), 0755)
		if err != nil {
			logger.Error(err)
			return "", err
		}
	}

	filePath := filepath.Join(m.dir(), fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	slotByte, err := os.ReadFile(path.Join(BaseDir, SlotPath))
	if err != nil {
		logger.Error(err)
		return "", err
	}

	slot := string(slotByte)

	slot = strings.Replace(slot, "{{NAME}}", name, -1)
	slot = strings.Replace(slot, "{{TIME}}", slotTimeFormat, -1)

	_, err = file.WriteString(slot)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
