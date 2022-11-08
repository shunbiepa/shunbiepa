package main

import (
	"fmt"
	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
	"math/rand"
	"strings"
	"time"
)

const (
	TABLE_IG = "IPAexG"
	TABLE_MD = "MPBOLD"
	TABLE_MY = "微软雅黑"
)

// CreateSignTable 创建签到表
func CreateSignTable() {
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
	r.Cell(200, 55, "申请日期：")
	r.Cell(300, 55, "审批状态：")
	// 注册个表格生成器
	r.RegisterExecutor(core.Executor(SimpleTableExecutor), core.Detail)
	//r.Execute("sign.pdf")
	//
	//
	//
	//r.AddNewPage(false)
	r.Image("img.png", 80, 500, 200, 700)
	//r.pag
	r.Execute("sign.pdf")
	fmt.Println(r.GetCurrentPageNo())
}

func SimpleTableExecutor(report *core.Report) {
	/*lineSpace := 0.0
	// 每行文字的间距
	lineHeight := 8.0

	// 创建一个表格，4列，30行，宽度600
	table := gopdf.NewTable(7, 20, 600, lineHeight, report)
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
	//c20 := table.NewCellByRange(1, 1)
	//c21 := table.NewCellByRange(6, 1)

	//c30 := table.NewCellByRange(1, 1)
	//c31 := table.NewCellByRange(2, 2)
	//c32 := table.NewCellByRange(3, 3)
	//c33 := table.NewCellByRange(4, 4)
	//c34 := table.NewCellByRange(5, 5)
	//c35 := table.NewCellByRange(6, 6)
	//c36 := table.NewCellByRange(7, 7)
	f1 := core.Font{Family: TABLE_MY, Size: 5, Style: ""}
	// 这里设置的是字内容和边框的距离
	border := core.NewScope(1.0, 1.0, 0, 0)
	// table.GetColWidth中, (0,0)第1行，第1列，(1,0)，第2行，第1列，(1,2),第2行，第3列，虽然第2行，合并成只有两个单元格，但获取长度还是按原来默认表格计数
	// 这里可以设置背景色，字体对齐样式等。
	c00.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("申请编码"))
	c01.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("007123"))
	c10.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("申请人"))
	c11.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三张三"))

	//c20.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("申请人部门"))
	//c21.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("总经理办公室－财务部"))

	//c30.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("日期"))
	//c31.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("部门"))
	//c32.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("费用纳入项目"))
	//c33.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("报销金额"))
	//c34.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("报销类别"))
	//c35.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("费用明细"))
	//c36.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).VerticalCentered().SetContent("费用明细"))
	//

	f1 = core.Font{Family: TABLE_MY, Size: 10}
	//border = core.NewScope(4.0, 4.0, 0, 0)

	for i := 0; i < 18; i++ {
		cells := make([]*gopdf. gopdf.TableCell, 7)
		for j := 0; j < 7; j++ {
			cells[j] = table.NewCell()
		}
	}

	table.GenerateAtomicCell()*/

	lineSpace := 1.0
	lineHeight := 18.0

	rows, cols := 100, 5
	table := gopdf.NewTable(cols, rows, 415, lineHeight, report)
	table.SetMargin(core.Scope{})

	for i := 0; i < rows; i += 5 {
		key := rand.Intn(3)
		//key := (i+1)%2 + 1
		f1 := core.Font{Family: TABLE_MY, Size: 10}
		border := core.NewScope(4.0, 4.0, 0, 0)

		switch key {
		case 0:
			for row := 0; row < 5; row++ {
				for col := 0; col < cols; col++ {
					conent := fmt.Sprintf("%v-(%v,%v)", 0, i+row, col)
					cell := table.NewCell()
					txt := gopdf.NewTextCell(table.GetColWidth(i+row, col), lineHeight, lineSpace, report).SetBackColor(GetRandColor())
					txt.SetFont(f1).SetBorder(border).SetContent(conent + GetRandStr(1))
					cell.SetElement(txt)
				}
			}

		case 1:
			c00 := table.NewCellByRange(1, 5)
			c01 := table.NewCellByRange(2, 2)
			c03 := table.NewCellByRange(2, 3)
			c21 := table.NewCellByRange(2, 1)
			c31 := table.NewCellByRange(4, 1)
			c41 := table.NewCellByRange(4, 1)

			t00 := gopdf.NewTextCell(table.GetColWidth(i+0, 0), lineHeight, lineSpace, report).SetBackColor(GetRandColor()).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t00"))
			t01 := gopdf.NewTextCell(table.GetColWidth(i+0, 1), lineHeight, lineSpace, report).SetBackColor(GetRandColor()).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t01"))
			t03 := gopdf.NewTextCell(table.GetColWidth(i+0, 3), lineHeight, lineSpace, report).SetBackColor(GetRandColor()).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t03"))
			t21 := gopdf.NewTextCell(table.GetColWidth(i+2, 1), lineHeight, lineSpace, report).SetBackColor(GetRandColor()).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t21"))
			t31 := gopdf.NewTextCell(table.GetColWidth(i+3, 1), lineHeight, lineSpace, report).SetBackColor(GetRandColor()).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t31"))
			t41 := gopdf.NewTextCell(table.GetColWidth(i+4, 1), lineHeight, lineSpace, report).SetBackColor(GetRandColor()).SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t41"))

			//t00.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t00"))
			//t01.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t01"))
			//t03.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t03"))
			//t21.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t21"))
			//t31.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t31"))
			//t41.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t41"))

			c00.SetElement(t00)
			c01.SetElement(t01)
			c03.SetElement(t03)
			c21.SetElement(t21)
			c31.SetElement(t31)
			c41.SetElement(t41)

		case 2:
			c00 := table.NewCellByRange(3, 2)
			c03 := table.NewCellByRange(2, 3)
			c20 := table.NewCellByRange(1, 2)
			c21 := table.NewCellByRange(2, 3)
			c33 := table.NewCellByRange(2, 2)
			c40 := table.NewCellByRange(1, 1)

			t00 := gopdf.NewTextCell(table.GetColWidth(i+0, 0), lineHeight, lineSpace, report).SetBackColor(GetRandColor())
			t03 := gopdf.NewTextCell(table.GetColWidth(i+0, 3), lineHeight, lineSpace, report).SetBackColor(GetRandColor())
			t20 := gopdf.NewTextCell(table.GetColWidth(i+2, 0), lineHeight, lineSpace, report).SetBackColor(GetRandColor())
			t21 := gopdf.NewTextCell(table.GetColWidth(i+2, 1), lineHeight, lineSpace, report).SetBackColor(GetRandColor())
			t33 := gopdf.NewTextCell(table.GetColWidth(i+3, 3), lineHeight, lineSpace, report).SetBackColor(GetRandColor())
			t40 := gopdf.NewTextCell(table.GetColWidth(i+4, 0), lineHeight, lineSpace, report).SetBackColor(GetRandColor())

			t00.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t00_1"))
			t03.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t03_2"))
			t20.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t20_3"))
			t21.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t21_4"))
			t33.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t33_5"))
			t40.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("t40_6"))

			c00.SetElement(t00)
			c03.SetElement(t03)
			c20.SetElement(t20)
			c21.SetElement(t21)
			c33.SetElement(t33)
			c40.SetElement(t40)
		}

	}

	table.GenerateAtomicCell()

}

var (
	seed = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GetRandColor() (color string) {
	r, g, b := seed.Intn(256), seed.Intn(256), seed.Intn(256)
	if float64(r)*0.299+float64(g)*0.578+float64(b)*0.114 >= 192 {
		color = fmt.Sprintf("%v,%v,%v", r, g, b)
		return color
	}

	return GetRandColor()
}

func GetRandStr(l ...int) string {
	str := "0123456789ABCDEFGHIGKLMNOPQRSTUVWXYZ"
	l = append(l, 8)
	r := seed.Intn(l[0]*11) + 8
	data := strings.Repeat(str, r/36+1)
	return data[:r]
}

func main() {
	CreateSignTable()
}
