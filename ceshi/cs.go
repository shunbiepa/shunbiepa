package main

import (
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"net/http"
)

func main() {
	//err := GeneratePdf("hello.pdf")
	//if err != nil {
	//	log.Print(err)
	//}

	imgUrl := "https://www.twle.cn/static/i/img1.jpg"

	// Get the data
	resp, err := http.Get(imgUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("img1.jpg", data, 0644)
}

// GeneratePdf generates our pdf by adding text and images to the page
// then saving it to a file (name specified in params).
func GeneratePdf(filename string) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	//lineSpace := 0.0
	//// 每行文字的间距
	//lineHeight := 18.0
	//report :=
	//table := gopdf.NewTable(4, 30, 600, lineHeight, report)
	//table.SetMargin(core.Scope{})
	pdf.SetXY(0, 10)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 7, "Welcome to topgoer.com", "0", 0, "CM", false, 0, "")

	pdf.Link(1, 2, 3, 4, 0)

	//图片
	pdf.AddPage()
	//循环放图片
	pdf.Image("img.png", 0, 0, 40, 80, false, "", 0, "")

	return pdf.OutputFileAndClose(filename)
}
