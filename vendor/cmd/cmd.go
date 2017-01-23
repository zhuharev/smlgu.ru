package cmd

import (
	// /"fmt"
	"html/template"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/session"
	"github.com/urfave/cli"
	"gopkg.in/macaron.v1"

	"controllers"
	"controllers/api/v1"
	"controllers/teachers"
	"models"
	"modules/middleware"
)

var (
	CmdWeb = cli.Command{
		Name:        "web",
		Usage:       "Preform admin operations on command line",
		Description: `Run web server`,
		Action:      runWeb,
	}
)

func runWeb(ctx *cli.Context) error {
	models.NewContext()

	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{Layout: "layout", Funcs: []template.FuncMap{
		template.FuncMap{
			"declofnum": func(co int64, s ...string) string {
				cases := []int{2, 0, 1, 1, 1, 2}
				ind := 0
				if co%100 > 4 && co%100 < 20 {
					ind = 2
				} else {
					if co%10 < 5 {
						ind = cases[co%10]
					} else {
						ind = cases[5]
					}
				}
				return s[ind]
			},
			"gurufeature": func(i interface{}) string {
				var (
					a = 0
				)
				if str, ok := i.(string); ok {
					switch str {
					case "humor":
						a = 1
					case "goodwill":
						a = 2
					case "understandability":
						a = 3
					}
				} else {
					a = i.(int)
				}
				switch a {
				case 1:
					return "Чувство юмора"
				case 2:
					return "Доброжелательность"
				case 3:
					return "Понятность объяснений"
				}
				return ""
			},
		},
	}}))
	m.Use(cache.Cacher())
	m.Use(session.Sessioner())
	m.Use(middleware.Contexter())

	m.Get("/test", controllers.Test)

	m.Group("/rating", func() {
		m.Get("/:slug", teachers.Show)
	})

	m.Group("/teachers", func() {
		m.Get("/", teachers.List)
		m.Get("/:slug", teachers.Get)
		m.Get("/avatars/:slug", teachers.Avatar)
	})

	m.Group("/api/v1", func() {
		m.Group("/teachers", func() {
			m.Get("/", v1.TeachersList)
			m.Get("/:id", v1.TeachersGet)
			m.Get("/:id/votes", v1.TeachersGetVotes)
		})
	})

	m.Run()

	return nil
}
