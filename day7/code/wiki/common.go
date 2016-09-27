package main

import (
	"os"
)

const (
	DefaultPagesPath     = "pages"
	DefaultTemplatesPath = "templates"
	envTemplatesPath     = "TEMPLATES_PATH"
	envPagesPath         = "PAGES_PATH"
	StaticPath           = "static"
)

var TemplatesPath = envOrDefault(envTemplatesPath, DefaultTemplatesPath)
var PagesPath = envOrDefault(envPagesPath, DefaultPagesPath)

func envOrDefault(envar, def string) string {
	if s := os.Getenv(envar); len(s) > 0 {
		return s
	}
	return def
}
