package pages

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/isacikgoz/tldr/pkg/config"
)

var (
	sep = string(os.PathSeparator)
	// a page should have md file extension
	ext = ".md"

	bold  = color.New(color.Bold)
	blue  = color.New(color.FgBlue)
	red   = color.New(color.FgRed)
	cyan  = color.New(color.FgCyan)
	white = color.New(color.FgWhite)
)

// Read finds and creates the Page, if it does not find, simply returns abstract
// contribution guide
func Read(seq []string) (p *Page, err error) {
	page := ""
	for i, l := range seq {
		if len(seq)-1 == i {
			page = page + l
			break
		} else {
			page = page + l + "-"
		}
	}
	// Common pages are more, so we have better luck there
	p, err = queryCommon(page)
	if err != nil {
		p, err = queryOS(page)
		if err != nil {
			return p, errors.New("This page (" + page + ") doesn't exist yet!\n" +
				"Submit new pages here: https://github.com/tldr-pages/tldr")
		}
	}
	return p, nil
}

// Queries from common folder
func queryCommon(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + "common" + sep
	b, err := ioutil.ReadFile(d + page + ".md")
	if err != nil {
		return p, err
	}
	p = ParsePage(string(b))
	return p, nil
}

// Queries from os specific folder
func queryOS(page string) (p *Page, err error) {
	d := config.SourceDir + sep + "pages" + sep + config.OSName() + sep
	b, err := ioutil.ReadFile(d + page + ".md")
	if err != nil {
		return p, err
	}
	p = ParsePage(string(b))
	return p, nil
}