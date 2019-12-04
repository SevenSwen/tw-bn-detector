package logger

import (
	"github.com/sevenswen/tw-bn-detector/utils"
	"log"
	"os"
)

var (
	Log                         *log.Logger
)

func Init() {
	dirPath := utils.GetOutDirectory()
	filePath := dirPath + "/logs.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		file, err = os.Create(filePath)
		Log = log.New(file, "", log.LstdFlags)
		Log.Println("\"BOTNETDETECTOR\"'s LogFile: " + filePath)
		if err != nil {
			panic(err)
		}
	} else {
		Log = log.New(file, "[BOTNETDETECTOR]", log.LstdFlags)
	}
}

func OpenOrCreateOutputFile(filename string) *os.File {
	dirPath := utils.GetOutDirectory()
	fullname := dirPath + filename
	res, err := os.OpenFile(fullname, os.O_RDWR|os.O_APPEND, 0660)

	if err != nil {
		res, err = os.Create(fullname)
		if err != nil {
			panic(err)
		}
	}
	return res
}

//func LogCurrencies(file *os.File, currencies []models.Currency, startTime time.Time) {
//	Log.Println("**************************")
//	Log.Println(fmt.Sprintf("Most relevant Currencies(%d): ", len(currencies)))
//	for i := range currencies {
//		Log.Println(fmt.Sprintf("- type: %v \"%v\".With weight: %v price: %v", currencies[i].Type, currencies[i].Name, currencies[i].Weight, currencies[i].Price))
//		row := strings.Join(
//			[]string{
//				fmt.Sprintf("%v", currencies[i].Type),
//				fmt.Sprintf("\"%v\"", startTime.Unix()),
//				fmt.Sprintf("\"%s\"", currencies[i].Name),
//				fmt.Sprintf("%v", currencies[i].Price),
//				fmt.Sprintf("%v", currencies[i].Weight),
//				"\n"},
//			";")
//		writeLine(file, row)
//	}
//	writeLine(file, "\n")
//	Log.Println("**************************")
//}


func writeLine(file *os.File, text string) {
	bytes := []byte(text)
	bytesWritten, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
