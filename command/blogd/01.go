package main

import (
	"flag"
	"os"
	"text/template"
	"time"
)

var tpls = `---
title: {{.Title}}
date: "{{.Date}}"
draft: false
category: [""]
---
`

var layout = "2006-01-02T15:04:05-0700"

type Contents struct {
	Title string
	Date  string
}

func main() {

	var title string
	flag.StringVar(&title, "t", "", "the post's title")
	flag.Parse()

	// 1. get paramse
	t := time.Now().Format(layout)

	cx := Contents{
		title,
		t,
	}

	var temp = template.Must(template.New("name").Parse(tpls))
	temp.Execute(os.Stdout, cx)

}
