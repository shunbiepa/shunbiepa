package main

import (
	"fmt"
	"github.com/tealeg/xlsx" // v2版本的，V3版本太复杂，没研究明白
)

type Person struct {
	ApprovalID         string // 审批编号
	ApprovalState      string // 审批状态
	Name               string // 申请人姓名
	ApprovalDepartment string // 申请部门
	Record             string // 审批记录
	ApprovalDate       string // 日期
	Department         string // 部门
	CostProject        string // 费用纳入项目
	Money              string //费用
	ExpensesClass      string //报销类型
	ExpenseDetails     string //费用明细
}

func ReadExcel() {
	var per []Person

	// 打开 xlsx文件
	xlFile, err := xlsx.OpenFile("报销.xlsx")
	if err != nil {
		fmt.Println("打开文件失败", err.Error())
		return
	}

	// 遍历 sheet 页
	for _, sheet := range xlFile.Sheets {

		// 行
		for i, row := range sheet.Rows {
			if i != 1 {
				// 列
				var temp Person

				// 将excel每一列文件读取放在字符串切片中
				var str []string
				for _, cell := range row.Cells {
					str = append(str, cell.String())
				}
				temp.ApprovalID = str[0]
				temp.ApprovalState = str[2]
				temp.Name = str[9]
				temp.ApprovalDepartment = str[10]
				temp.Record = str[12]
				temp.ApprovalDate = str[15]
				temp.Department = str[16]
				temp.CostProject = str[17]
				temp.Money = str[18]
				temp.ExpensesClass = str[19]
				temp.ExpenseDetails = str[20]

				per = append(per, temp)
			}

		}

	}

	fmt.Println(per)
}

func main() {
	ReadExcel()

}
