package customrender

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/go-echarts/go-echarts/v2/render"
	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

//go:embed mainpage.tpl
var mainpageTpl string

type customRender struct {
	c      interface{}
	before []func()
}

func NewCustomRender(c interface{}, before ...func()) render.Renderer {
	return &customRender{c: c, before: before}
}

func (cR *customRender) Render(w io.Writer) error {
	for _, fn := range cR.before {
		fn()
	}

	contents := []string{mainpageTpl, tpls.HeaderTpl, tpls.BaseTpl}
	tpl := render.MustTemplate("mainpage", contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "mainpage", cR.c); err != nil {
		return err
	}

	_, err := w.Write(buf.Bytes())
	return err
}
