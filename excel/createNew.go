package excel

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func create() {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellValue("Sheet1", "A1", 50)
	if err := f.SaveAs("simple.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func goodNew(){
	fmt.Println("Good new!")
}
