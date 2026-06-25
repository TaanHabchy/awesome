package html_link_parser

import (
	"os"
	"strconv"

	"golang.org/x/net/html"
)

/*
Goal: Return a Links object that contains the Link objects with href and text
*/

type Links []Link

type Link struct {
	href string
	text string
}

func FullText(node *html.Node) string {
	rv := node.Data
	if node.NextSibling == nil {
		return rv
	}
	if node.NextSibling.FirstChild == nil {
		return rv
	}
	return rv + FullText(node.NextSibling.FirstChild)
}

func HtmlLinkParser(path int) Links {

	var rv Links = make([]Link, 0)
	file, fileError := os.Open("./html_link_parser/ex" + strconv.Itoa(path) + ".html")
	if fileError != nil {
		return nil
	}

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// want to rip through the descendants, and when i find a type of a, search for href
	// if href and text, then add those to the node list
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {

			for _, a := range n.Attr {
				if a.Key == "href" {

					rv = append(rv, Link{a.Val, FullText(n.FirstChild)})
				}
			}
		}
	}
	for i := range rv {
		println(rv[i].href, " ", rv[i].text)

	}

	return rv
}
