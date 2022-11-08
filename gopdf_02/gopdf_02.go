package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	TABLE_IG = "IPAexG"
	TABLE_MD = "MPBOLD"
	TABLE_MY = "微软雅黑"
)

// CreateSignTable 创建签到表
func CreateSignTable(person Person) {

	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: TABLE_IG,
		// 要将字体文件，放到程序可以访问的地方
		FileName: "ttf/SIMYOU.TTF",
	}
	font2 := core.FontMap{
		FontName: TABLE_MD,
		FileName: "ttf/SIMYOU.TTF",
	}
	font3 := core.FontMap{
		FontName: TABLE_MY,
		FileName: "ttf/SIMYOU.TTF",
	}
	r.SetFonts([]*core.FontMap{&font1, &font2, &font3})
	r.SetPage("A4", "P")

	// 一定要设置字体，否则执行执行r.Cell会报错
	r.SetFont(TABLE_IG, 15)
	// 这里坐标是不断调整后的结果
	r.Cell(270, 30, "报销")
	r.SetFont(TABLE_IG, 10)
	r.Cell(90, 55, "湖南心镜科技有限公司")
	r.Cell(220, 55, fmt.Sprintf("申请日期:%s", person.ApprovalDate))
	r.Cell(340, 55, fmt.Sprintf("审批状态:%s", person.ApprovalState))

	//r.Cell(90, 248, "打印时间：")
	//r.Cell(200, 248, "打印人：")
	// 注册个表格生成器
	r.RegisterExecutor(func(report *core.Report) {
		SimpleTableExecutor(report, person)
	}, core.Detail)

	for i, url := range person.Urls {
		if i >= 3 {
			k := 0
			j := float64(k * 150)
			urls := strings.Split(url, "/")
			r.Image(urls[len(urls)-1], 80+j, 625, 200+j, 825)
			k++
		} else {
			j := float64(i * 150)
			urls := strings.Split(url, "/")
			r.Image(urls[len(urls)-1], 80+j, 400, 200+j, 600)
		}

	}

	//r.pag
	content := fmt.Sprintf("%s_报销.pdf", person.ApprovalID)
	r.Execute(content)
	//fmt.Println(r.GetCurrentPageNo())
	//SimpleTableExecutor(ret,"string")
}

func SimpleTableExecutor(report *core.Report, per Person) {

	lineSpace := 0.0
	// 每行文字的间距
	lineHeight := 8.0

	// 创建一个表格，4列，30行，宽度600
	table := gopdf.NewTable(7, 21, 600, lineHeight, report)
	table.SetMargin(core.Scope{
		Left:   0.0,
		Top:    0.0,
		Right:  0.0,
		Bottom: 0.0,
	})

	// 合并单元格
	c00 := table.NewCellByRange(1, 1) // 第1行，1个单元格合并
	c01 := table.NewCellByRange(6, 1) // 第1行，6个单元格合并
	c10 := table.NewCellByRange(1, 1)
	c11 := table.NewCellByRange(6, 1)
	c20 := table.NewCellByRange(1, 1)
	c21 := table.NewCellByRange(6, 1)

	c30 := table.NewCellByRange(1, 1)
	c31 := table.NewCellByRange(1, 1)
	c32 := table.NewCellByRange(1, 1)
	c33 := table.NewCellByRange(1, 1)
	c34 := table.NewCellByRange(1, 1)
	c35 := table.NewCellByRange(1, 1)
	c36 := table.NewCellByRange(1, 1)

	c40 := table.NewCellByRange(1, 3)
	c41 := table.NewCellByRange(1, 3)
	c42 := table.NewCellByRange(1, 3)
	c43 := table.NewCellByRange(1, 3)
	c44 := table.NewCellByRange(1, 3)
	c45 := table.NewCellByRange(1, 3)
	c46 := table.NewCellByRange(1, 3)

	c50 := table.NewCellByRange(1, 1)
	c51 := table.NewCellByRange(6, 1) //7

	c60 := table.NewCellByRange(1, 1) //8
	c61 := table.NewCellByRange(6, 1)

	c70 := table.NewCellByRange(1, 2)
	c71 := table.NewCellByRange(6, 2) //10

	c80 := table.NewCellByRange(1, 10) //11
	c81 := table.NewCellByRange(6, 10)
	//c82 := table.NewCellByRange(6, 2)
	//c83 := table.NewCellByRange(6, 1)
	//c84 := table.NewCellByRange(6, 1)
	//c85 := table.NewCellByRange(6, 1)
	//c86 := table.NewCellByRange(6, 1)
	//c87 := table.NewCellByRange(6, 1)
	//c88 := table.NewCellByRange(6, 1)
	//c89 := table.NewCellByRange(6, 1)

	f1 := core.Font{Family: TABLE_MY, Size: 5, Style: ""}
	// 这里设置的是字内容和边框的距离
	border := core.NewScope(1.0, 1.0, 0, 0)
	// table.GetColWidth中, (0,0)第1行，第1列，(1,0)，第2行，第1列，(1,2),第2行，第3列，虽然第2行，合并成只有两个单元格，但获取长度还是按原来默认表格计数
	// 这里可以设置背景色，字体对齐样式等。
	c00.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("申请编码"))
	c01.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.ApprovalID))
	c10.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("申请人"))
	c11.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.Name))
	c20.SetElement(gopdf.NewTextCell(table.GetColWidth(2, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("申请人部门"))
	c21.SetElement(gopdf.NewTextCell(table.GetColWidth(2, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.ApprovalDepartment))

	c30.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent(""))
	c31.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("日期"))
	c32.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("部门"))
	c33.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 3), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("费用纳入项目"))
	c34.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 4), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("报销金额"))
	c35.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("报销类别"))
	c36.SetElement(gopdf.NewTextCell(table.GetColWidth(3, 6), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("费用明细"))

	c40.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("报销明细"))
	c41.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.ApprovalDate))
	c42.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 2), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.Department))
	c43.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 3), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.CostProject))
	c44.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 4), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.Money))
	c45.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 5), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.CostProject))
	c46.SetElement(gopdf.NewTextCell(table.GetColWidth(4, 6), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.ExpenseDetails))

	c50.SetElement(gopdf.NewTextCell(table.GetColWidth(7, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent(""))
	c51.SetElement(gopdf.NewTextCell(table.GetColWidth(7, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(fmt.Sprintf("报销金额(元)%s", per.Money)))

	c60.SetElement(gopdf.NewTextCell(table.GetColWidth(8, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("发票情况"))
	c61.SetElement(gopdf.NewTextCell(table.GetColWidth(8, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.ExpensesClass))

	c70.SetElement(gopdf.NewTextCell(table.GetColWidth(9, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("图片(请上传凭证)"))
	c71.SetElement(gopdf.NewTextCell(table.GetColWidth(9, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("1个"))

	c80.SetElement(gopdf.NewTextCell(table.GetColWidth(10, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("审批流程"))
	c81.SetElement(gopdf.NewTextCell(table.GetColWidth(11, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent(per.Record))
	//c82.SetElement(gopdf.NewTextCell(table.GetColWidth(12, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("9"))
	//c83.SetElement(gopdf.NewTextCell(table.GetColWidth(14, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("10"))
	//c84.SetElement(gopdf.NewTextCell(table.GetColWidth(15, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("11"))
	//c85.SetElement(gopdf.NewTextCell(table.GetColWidth(16, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("12"))
	//c86.SetElement(gopdf.NewTextCell(table.GetColWidth(17, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("13"))
	//c87.SetElement(gopdf.NewTextCell(table.GetColWidth(18, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("14"))
	//c88.SetElement(gopdf.NewTextCell(table.GetColWidth(19, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("15"))
	//c89.SetElement(gopdf.NewTextCell(table.GetColWidth(20, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("16"))

	f1 = core.Font{Family: TABLE_MY, Size: 10}
	//border = core.NewScope(4.0, 4.0, 0, 0)

	cells := make([]*gopdf.TableCell, 7)
	for j := 0; j < 7; j++ {
		cells[j] = table.NewCell()
	}

	table.GenerateAtomicCell()

}

type Person struct {
	ApprovalID         string   // 审批编号
	ApprovalState      string   // 审批状态
	Name               string   // 申请人姓名
	ApprovalDepartment string   // 申请部门
	Record             string   // 审批记录
	ApprovalDate       string   // 日期
	Department         string   // 部门
	CostProject        string   // 费用纳入项目
	Money              string   //费用
	ExpensesClass      string   //报销类型
	ExpenseDetails     string   //费用明细
	UrlSun             string   //图片数量
	Urls               []string //图片路径
}

func ReadExcel() (per []Person) {

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
			if i != 0 {
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
				var listStr []string
				urlList := str[23]
				listStr = strings.Split(urlList, "\n")
				for _, url := range listStr {
					if url != "" {
						temp.Urls = append(temp.Urls, url)
					}
				}
				temp.UrlSun = strconv.Itoa(len(temp.Urls))

				//下载好图片
				DemUrlJpg(temp.Urls)

				old := "|"
				new1 := ":"
				temp.Record = strings.Replace(temp.Record, old, new1, -1)
				per = append(per, temp)

			}

		}

	}

	return per
}

func DemUrlJpg(imgUrls []string) {
	for _, url := range imgUrls {
		urls := strings.Split(url, "/")

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(urls[len(urls)-1], data, 0644)
	}

}

func main() {

	per := ReadExcel()
	for _, v := range per {
		CreateSignTable(v)
	}

}
