package main

import (
	. "github.com/RafaelRCamargo/go-svg-api/server"
	. "github.com/RafaelRCamargo/go-svg-api/services"
	. "github.com/RafaelRCamargo/go-svg-api/utils"

	"github.com/joho/godotenv"
	"github.com/wcharczuk/go-chart/v2"

	"net/http"
	"os"
	. "strconv"
)

func httpserver(w http.ResponseWriter, _ *http.Request) {
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

	w.Header().Set("Content-Type", "image/svg+xml")

	f, _ := os.Create("output.svg")
	defer f.Close()
	graph.Render(chart.SVG, w)
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	Server(port)
	/*http.HandleFunc("/", httpserver)
	http.ListenAndServe(os.Getenv("PORT"), nil) */
}
