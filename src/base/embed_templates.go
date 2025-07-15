package base

import "embed"

//go:embed templates/**/*.tmpl
var TmplFS embed.FS
