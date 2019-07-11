package main

import (
	"database/sql"
	"fmt"
	"goLangTrainingAssignments/Day2/go/connect"

	"github.com/tealeg/xlsx"
)

func dbToExcel(db *sql.DB) {
	excelPath := "../documents/db_to_excel.xlsx"

	employee, err := db.Query("SELECT * FROM employee_details;")
	if err != nil {
		panic(err)
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	var column [2]string
	for employee.Next() {
		row := sheet.AddRow()
		employee.Scan(&column[0], &column[1])
		fmt.Println(column[0], column[1])
		for _, value := range column {
			cell := row.AddCell()
			cell.Value = value
		}
	}

	file.Save(excelPath)
}

func excelToDb(db *sql.DB) {
	excelPath := "../documents/excel_to_db.xlsx"

	xlFile, err := xlsx.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var value = make([]string, 4)
			for key, cell := range row.Cells {
				value[key] = cell.String()
				fmt.Printf("%s ", value[key])
			}

			_, err = db.Exec("INSERT INTO employee_details (employee_id, name) VALUES ($1, $2)", value[0], value[1])
			if err != nil {
				println(err)
				panic(err)
			}
			fmt.Printf("\n")
		}
	}
}

func main() {
	db := connect.Dbconnect()

	excelToDb(db)
	dbToExcel(db)

	defer db.Close()
}
