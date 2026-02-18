package visualisation

import (
	"Homework1/models"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// VisualizeStats aggregates and visualizes user stats
func VisualizeStats(users models.Users, usersRepos []models.Repositories, languages models.Languages) {
	for i, user := range users {
		repos := usersRepos[i]

		// Collect languages for this user
		userLangs := models.Languages{}
		for _, lang := range languages {
			for _, repo := range repos {
				if lang.RepoName == repo.FullName {
					userLangs = append(userLangs, lang)
				}
			}
		}

		// Aggregate stats
		totalForks := 0
		yearActivity := make(map[string]int)
		langDistribution := make(map[string]int)

		for _, repo := range repos {
			totalForks += repo.ForksCount

			createdYear := fmt.Sprintf("%d", repo.CreatedAt.Year())
			updatedYear := fmt.Sprintf("%d", repo.UpdatedAt.Year())

			yearActivity[createdYear]++
			if createdYear != updatedYear {
				yearActivity[updatedYear]++
			}
		}

		for _, lang := range userLangs {
			for k, v := range lang.LangStats {
				langDistribution[k] += v
			}
		}

		// Optional: print summary to console
		fmt.Printf("\n🤖 %s (%s)\nFollowers: %d | Public Repos: %d | Total Forks: %d\n",
			user.Login, user.Name, user.Followers, user.PublicRepos, totalForks)

		fmt.Println("\n📊 Language Distribution:")
		for lang, count := range langDistribution {
			fmt.Printf("   %-15s %d bytes\n", lang, count)
		}

		fmt.Println("\n📅 Activity by Year:")
		for year, count := range yearActivity {
			fmt.Printf("   %-6s %d repos\n", year, count)
		}
		fmt.Println("---------------------------------------------------------")

		// Generate HTML dashboard
		createUserDashboard(user.Login, langDistribution, yearActivity)
	}
}

func createUserDashboard(username string, langDist map[string]int, yearly map[string]int) {
	page := components.NewPage()
	page.PageTitle = fmt.Sprintf("%s - GitHub Statistics", username)
	page.SetLayout(components.PageFlexLayout)

	langChart := buildLanguagePie(langDist)
	actChart := buildActivityBar(yearly)

	page.AddCharts(langChart, actChart)

	f, err := os.Create(fmt.Sprintf("%s_dashboard.html", username))
	if err != nil {
		fmt.Println("Error creating dashboard:", err)
		return
	}
	defer f.Close()

	page.Render(f)
	fmt.Printf("✅ Generated %s_dashboard.html\n", username)
}

func buildLanguagePie(langDist map[string]int) *charts.Pie {
	pie := charts.NewPie()
	items := make([]opts.PieData, 0, len(langDist))

	for lang, count := range langDist {
		items = append(items, opts.PieData{Name: lang, Value: count})
	}

	show := true
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Language Usage",
			Subtitle: "Code bytes per language",
			Top:      "5%", // Keep title at top
			Left:     "center",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show:   &show,
			Orient: "vertical",
			Right:  "10%",    // Position legend on the right side
			Top:    "middle", // Center it vertically
		}),
	)

	pie.AddSeries("Languages", items).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      &show,
				Formatter: "{b}: {d}%",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Center: []string{"40%", "55%"},
				Radius: []string{"40%", "70%"},
			}),
		)

	return pie
}

func buildActivityBar(yearly map[string]int) *charts.Bar {
	bar := charts.NewBar()
	years := make([]string, 0, len(yearly))
	for y := range yearly {
		years = append(years, y)
	}
	sort.Strings(years)

	values := make([]opts.BarData, 0, len(years))
	for _, y := range years {
		values = append(values, opts.BarData{Value: yearly[y]})
	}

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Repo Activity by Year",
			Subtitle: fmt.Sprintf("Generated on %s", time.Now().Format("2006-01-02")),
			Top:      "5%",
			Left:     "center",
		}),
		charts.WithGridOpts(opts.Grid{
			Top:    "20%",
			Bottom: "15%",
			Left:   "10%",
			Right:  "10%",
		}),
	)

	bar.SetXAxis(years).AddSeries("Repositories", values)
	return bar
}
