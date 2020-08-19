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

		//filter trough subtitles, only cz.srt or cz.sub will pass
		if strings.HasSuffix(tits, "cz.srt") || strings.HasSuffix(tits, "cz.sub") {
			//fmt.Println(tits)

			//download file with wget
			path := "http://titulky.trekdnes.cz/" + tits
			cmd := exec.Command("wget", path, "-P", "/home/vitek/Videos/Star Trek/subtitles")
			err := cmd.Run()
			if err != nil {
				fmt.Printf("An error has occured while downloading a subtitle %v: %v\n", x, err.Error())
			} else {
				fmt.Printf("%v.) %v -> ", x, path)
			}

			//using recode to recode file encoding from windows-1250 to UTF-8
			fileName := strings.TrimPrefix(tits, "titulky/")
			cmd = exec.Command("recode", "windows-1250..utf-8", "/home/vitek/Videos/Star Trek/subtitles/"+fileName)
			err = cmd.Run()
			if err != nil {
				fmt.Printf("An error has occured while recoding a subtitle %v: %v\n", x, err.Error())
			} else {
				fmt.Printf("%v recoded into UTF-8\n", fileName)
			}
			x++
		}
	})
}
