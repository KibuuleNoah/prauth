package services

import (
	"fmt"
	"html/template"
)

type TemplatesService struct{}

func (r TemplatesService) LoadStyles(hrefs []any) string {
	var styles string

	for _, href := range hrefs {
		styles += fmt.Sprintf(
			`<link rel="stylesheet" type="text/css" href="/static/%s" />`,
			href,
		)
	}

	return styles
}

func (r TemplatesService) LoadScripts(srcs []any) string {
	var scripts string

	for _, src := range srcs {
		scripts += fmt.Sprintf(
			`<script src="/static/%s"></script>`,
			src,
		)
	}

	return scripts
}

func (r TemplatesService) SafeHTML(s string) template.HTML {
	return template.HTML(s)
}

func (r TemplatesService) CustomSlice(elems ...any) []any {
	return elems
}

func (r TemplatesService) Seq(n int) []int {
	var s []int
	for i := 1; i <= n; i++ {
		s = append(s, i)
	}
	return s
}

func (r TemplatesService) All() template.FuncMap {
	return template.FuncMap{
		"loadStyles":  r.LoadStyles,
		"loadScripts": r.LoadScripts,
		"safeHTML":    r.SafeHTML,
		"seq":         r.Seq,
		// custom datatypes
		"customSlice": r.CustomSlice,
	}
}
