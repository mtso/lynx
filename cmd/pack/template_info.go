package main

import (
	"html/template"
)

const (
	INDEX = iota
	POST
)

type templateType int

type templateInfo struct {
	templatetype templateType
	data         string
}

type LynxTemplate struct {
	Templates map[templateType]template.Template
}