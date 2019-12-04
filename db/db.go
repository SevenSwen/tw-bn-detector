package db

import (
	"github.com/sevenswen/tw-bn-detector/db/models"
	"github.com/sevenswen/tw-bn-detector/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() *gorm.DB {
	dirPath := utils.GetOutDirectory()
	filePath := dirPath + "/gorm.db"

	db, err := gorm.Open("sqlite3", filePath)

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Currency{})

	return db
}

func Save(db *gorm.DB) {
	//// сохраняем новых победителей в базу данных
	//for i := range bestCurrencies {
	//	db.Create(&bestCurrencies[i])
	//}
	//
	//// сохраняем информацию о переводе средств в базу данных
	//for _, currency := range newInvestedCurrencies {
	//	db.Create(&currency)
	//
	//	var newLogInvestedCurrencies models.LogInvestedCurrency
	//	db.Where(models.LogInvestedCurrency{Name: currency.Name}).
	//		Attrs(models.LogInvestedCurrency{Time: currency.Time}).
	//		FirstOrCreate(&newLogInvestedCurrencies)
	//
	//	newLogInvestedCurrencies.TotalCost = currency.TotalCost
	//	newLogInvestedCurrencies.Amount = currency.Amount
	//	newLogInvestedCurrencies.Price = currency.Price
	//	newLogInvestedCurrencies.Number = currency.Number
	//	newLogInvestedCurrencies.Type = currency.Type
	//
	//	db.Save(&newLogInvestedCurrencies)
	//}
}

func Load(db *gorm.DB, initInvestedCurrencies func()) {
	//
	//// загружаем информацию о том, в какой валюте у нас сейчас деньги
	//db.Find(&investedCurrencies)
	//// если этой информации нет - создает новую запись (считаем что деньги в BTC)
	//if len(investedCurrencies) == 0 {
	//	investedCurrencies = initInvestedCurrencies()
	//}
	//
	//// загружаем сохраненных крипто-победителей с прошлого анализа
	//db.Find(&prevBestCurrencies)
	//
	//// удаляем старые записи
	//db.Delete(models.Currency{})
	//db.Delete(models.InvestedCurrency{})
	//
	//return investedCurrencies, prevBestCurrencies
}
