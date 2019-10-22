package html2pdf

import (
	//"bytes"
	//"fmt"
	"log"
	"strings"

	"../md2html/md2html"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func Html2pdf(fi md2html.Fileinfo) {

	html := fi.Html

	pdfg := wkhtmltopdf.NewPDFPreparer()
	var err error
	// Create new PDF generator
	/*
		pdfg, err := wkhtmltopdf.NewPDFGenerator()
		if err != nil {
			log.Print("PDF generator 作成失敗")
			log.Print(fi.Pdfpath)
			log.Fatal(err)
			return
		}
		log.Print("Create new PDF generator")
	*/

	// Set global options
	pdfg.Dpi.Set(300)
	//pdfg.Orientation.Set(OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))

	// Set options for this page
	//page.FooterRight.Set("[page]")
	//page.FooterFontSize.Set(10)
	//page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()

	if err != nil {
		log.Print("PDF doc internal buffer 作成失敗")
		log.Fatal(err)
		return
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(fi.Pdfpath)
	if err != nil {
		log.Print("buffer 書き込み失敗")
		log.Print(fi.Pdfpath)
		log.Fatal(err)
		return
	}
}
