package controllers

import "text/template"

var templates = template.Must(template.ParseGlob("templates/*.html"))
