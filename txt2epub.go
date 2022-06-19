package main

import (
	"flag"
	"github.com/LinxuanZheng/txt2epub/src"
)

var cover = flag.Bool("c", false, "Gen default cover")

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		if args[0] == "init" {
			if *cover {
				src.GenDefaultCover()
			}
			src.InitConfig()
			return
		} else if args[0] == "gen" {
			src.GenBook()
		} else {

		}
	}
}
