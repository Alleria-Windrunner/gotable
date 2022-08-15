package main

import (
	"fmt"

	"github.com/Alleria-Windrunner/gotable"
)

func main() {
	fmt.Println("sss")
	table, err := gotable.Create("China", "US", "French")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	// Use map
	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington, D.C."
	row["French"] = "Paris"
	err = table.AddRow(row)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	// Use Slice
	row2 := []string{"Yinchuan", "Los Angeles", "Orleans"}
	err = table.AddRow(row2)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	row4 := []string{"=", "~", "+"}
	err = table.AddRow(row4)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	table.SetPNColumnsTag(0, 0)

	// Add new part with new columns
	table.AddPart("name", "salary")

	rows := make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["name"] = fmt.Sprintf("employee-%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}
	table.AddRows(rows)

	// Add row to previous part
	row3 := []string{"WuHan", "April", "Blank"}
	err = table.AddPNRow(0, row3)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	table.AddPart("name", "salary")

	rows = make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["name"] = fmt.Sprintf("employee-%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}
	table.AddRows(rows)
	table.DelRows()
	table.DelPNRows(0)

	table.SetBorder(1) //0" " 1"-" 2"=" 3"~" 4"+"

	//get columns and maps in specific part
	fmt.Println(table.GetPNColumns(1))
	fmt.Println(table.GetPNValues(0))

	//change any column length in any part
	// table.SetColumnMaxLength(1, "name", 15)
	// table.SetColumnMaxLength(1, "salary", 10)
	// table.SetColumnMaxLength(2, "name", 30)
	// table.SetColumnMaxLength(2, "salary", 15)
	table.AdaptColLen(0, 1, "salary")
	table.AdaptColLen(1, 2, "name")
	table.SetBorder(1)
	table.SetPNColumnsTag(2, 3)

	fmt.Println(table)

	//table.CloseBorder()

	//fmt.Println(table)
	// outputs:
	// +----------+------------------+---------+
	// |  China   |        US        | French  |
	// +----------+------------------+---------+
	// | Beijing  | Washington, D.C. |  Paris  |
	// | Yinchuan |   Los Angeles    | Orleans |
	// +----------+------------------+---------+
}
