package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func generateCertificate(recipient, course, date string) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddUTF8Font("DejaVu", "", "assets/fonts/DejaVuSans.ttf")
	pdf.AddPage()

	bgWidth, bgHeight := 297.0, 210.0
	pdf.ImageOptions("assets/img/background.jpg", 0, 0, bgWidth, bgHeight, false, gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true}, 0, "")

	pdf.SetFont("DejaVu", "", 36)
	pdf.CellFormat(0, 20, "Сертификат", "", 1, "C", false, 0, "")

	pdf.SetFont("DejaVu", "", 20)
	pdf.Ln(10)
	pdf.CellFormat(0, 10, "Табысталады", "", 1, "C", false, 0, "")
	pdf.SetFont("DejaVu", "", 28)
	pdf.CellFormat(0, 15, recipient, "", 1, "C", false, 0, "")

	pdf.SetFont("DejaVu", "", 20)
	pdf.Ln(10)
	pdf.CellFormat(0, 10, "курсты сәтті аяқтады:", "", 1, "C", false, 0, "")
	pdf.SetFont("DejaVu", "", 24)
	pdf.CellFormat(0, 12, course, "", 1, "C", false, 0, "")

	pdf.SetFont("DejaVu", "", 18)
	pdf.Ln(15)
	pdf.CellFormat(0, 10, fmt.Sprintf("Күні %s", date), "", 1, "C", false, 0, "")

	pdf.Ln(30)
	y := pdf.GetY()
	lineY := y + 5

	pdf.Line(50, lineY, 100, lineY)
	pdf.Line(200, lineY, 250, lineY)
	pdf.SetXY(50, lineY+5)
	pdf.CellFormat(50, 10, "Нұсқаушы", "", 0, "C", false, 0, "")
	pdf.SetXY(200, lineY+5)
	pdf.CellFormat(50, 10, "Директор", "", 0, "C", false, 0, "")

	err := pdf.OutputFileAndClose("certificate.pdf")
	if err != nil {
		fmt.Println("Қате:", err)
		return
	}
	fmt.Println("Сертификат сәтті жасалды!")
}

func main() {
	recipient := "Бекзаттегі Бекзат Бекзатұлы"
	course := "Өмір сүргені үшін"
	date := "25 мамыр, 2025 жыл"

	generateCertificate(recipient, course, date)
}
