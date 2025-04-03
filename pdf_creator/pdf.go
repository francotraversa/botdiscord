package pdf_creator

import (
	"fmt"
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func PDF(ticker string, cierre float64, info string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	// Fuente y colores base
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(32, 32, 128)

	// Título
	pdf.Cell(0, 10, "SAPISTOCK")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 10, "Financial Reports")
	pdf.Ln(10)

	// Fecha
	pdf.SetFont("Arial", "", 10)
	hoy := time.Now()
	fechaformato := hoy.Format("02/01/2002 15:15:15")
	pdf.MultiCell(0, 10, fechaformato, "", "R", false)
	pdf.Ln(12)

	// Nombre del activo
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 10, ticker)
	pdf.Ln(12)

	// Secciones de gráficos
	//pdf.SetFont("Arial", "B", 14)
	//pdf.Cell(0, 10, "Último mes")
	//pdf.Ln(35) // Espacio para un gráfico

	//pdf.Cell(0, 10, "Últimos 12 meses")
	//pdf.Ln(35) // Espacio para otro gráfico

	// Datos técnicos
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Datos Técnicos")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 10)

	for _, item := range info {
		pdf.Cell(0, 8, string(item))
		pdf.Ln(6)
	}

	// Precio actual
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Precio Actual")
	pdf.Ln(6)
	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(0, 0, 200)
	pdf.Cell(0, 10, fmt.Sprintf("%.2f", cierre))
	pdf.Ln(12)

	// Guardar el PDF
	filename := fmt.Sprintf("reporte_%s.pdf", ticker)
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		log.Fatal(err)
	}
}
