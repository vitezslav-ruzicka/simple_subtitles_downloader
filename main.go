package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	res, err := http.Get("http://titulky.trekdnes.cz/tos.htm")
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	x := 1
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		tits, _ := s.Attr("href")

		if strings.HasSuffix(tits, "cz.srt") || strings.HasSuffix(tits, "cz.sub") {
			//fmt.Println(tits)

			path := "http://titulky.trekdnes.cz/" + tits
			var cmd *exec.Cmd

			if x <= 30 {
				cmd = exec.Command("wget", path, "-P", "/home/vitek/Videos/Star Trek/Season 1")
			} else if x >= 31 && x <= 56 {
				cmd = exec.Command("wget", path, "-P", "/home/vitek/Videos/Star Trek/Season 2")
			} else if x > 56 {
				cmd = exec.Command("wget", path, "-P", "/home/vitek/Videos/Star Trek/Season 3")
			}

			err := cmd.Run()
			if err != nil {
				fmt.Printf("An error has occured while downloading a subtitle %v:", x, err.Error())
			} else {
				fmt.Println(x, path)
			}
			x++
		}
	})
}
