package router

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Renderer struct {
	template *template.Template
	debug bool
	location string
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render (
	w io.Writer, 
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}
	return t.template.ExecuteTemplate(w, name, data)
}

func NewRenderer(location string, debug bool) *Renderer  {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func Config() *echo.Echo {
	e := echo.New()
	e.Static("/", "public/static")

	e.Renderer = NewRenderer("./app/views/*.go.tpl", true)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}