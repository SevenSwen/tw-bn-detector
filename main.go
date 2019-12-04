package main

import (
	"github.com/sevenswen/tw-bn-detector/db"
	"github.com/sevenswen/tw-bn-detector/utils/logger"
	"github.com/jinzhu/gorm"
)

var (
	// база данных
	currenciesDB *gorm.DB
)

func init() {
	// инициализируем все файлы выходных данных и log файл
	logger.Init()
	// инициализируем базу данных
	currenciesDB = db.Init()
}

func main() {
	defer currenciesDB.Close()

}
