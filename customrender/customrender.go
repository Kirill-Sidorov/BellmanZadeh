package customrender

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/render"
	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

var resultpageTpl string

type customRender struct {
	c      interface{}
	before []func()
}

func Init() {
    b, err := os.ReadFile("resources/resultpage.tpl")
    if err != nil {
        log.Fatal(err)
    }

    resultpageTpl = string(b)
}

func NewCustomRender(c interface{}, before ...func()) render.Renderer {
	return &customRender{c: c, before: before}
}

func (cR *customRender) Render(w io.Writer) error {
	for _, fn := range cR.before {
		fn()
	}

	contents := []string{resultpageTpl, tpls.BaseTpl}
	tpl := render.MustTemplate("resultpage", contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "resultpage", cR.c); err != nil {
		return err
	}

	_, err := w.Write(buf.Bytes())
	return err
}
