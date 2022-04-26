package server

import (
	. "github.com/RafaelRCamargo/go-svg-api/services"
	. "github.com/RafaelRCamargo/go-svg-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/wcharczuk/go-chart/v2"

	"os"
	. "strconv"
)

func statsHandler(c *gin.Context) {
	data := GetData()

	var days []string
	items := make([]chart.Value, 0)

	for _, item := range data {
		date := GetDate(item.Week)
		year, month, day := date.Date()
		localeDate := month.String() + " " + Itoa(day) + ", " + Itoa(year)

		days = append(days, localeDate)
		items = append(items, chart.Value{Value: float64(item.Total), Label: localeDate})
	}

	graph := chart.BarChart{
		Title:      "Test Bar Chart",
		BarSpacing: 12,
		TitleStyle: chart.Hidden(),
		Font:       chart.Hidden().Font,
		Background: chart.Style{
			TextHorizontalAlign: chart.TextHorizontalAlignCenter,
			TextVerticalAlign:   chart.TextVerticalAlignMiddle,
			FontColor:           chart.ColorAlternateGray,
			FontSize:            0.0,
		},
		Height:   512,
		BarWidth: 92,
		Bars:     items,
	}

	c.Header("Content-Type", "image/svg+xml")

	f, _ := os.Create("output.svg")
	defer f.Close()
	graph.Render(chart.SVG, f)
	c.File("output.svg")
}

func Server(port string) {
	println("Starting server...")

	r := gin.Default()

	r.GET("/", statsHandler)

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.ico")
	})

	r.Run(":" + port)

	println("Listing for requests at http://localhost:" + port)
}
