package xtemplate_gofeed

import (
	"html/template"

	"github.com/infogulch/xtemplate/register"
)

func init() {
	register.RegisterFuncMap("gofeed", template.FuncMap{"fetchFeed": funcFetchFeed})
}
