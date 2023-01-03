package main

import (
	"flag"
	"fmt"
	"os"
	"raider/app"
	"runtime"
)

func help() {
	if runtime.GOOS == "windows" {
		fmt.Println("Usage:\nraider.exe --target https://vk.me/join... --duration 50 ")
	} else {
		fmt.Println("Usage:\n./raider --target https://vk.me/join... --duration 50")
	}
}

func main() {

	if len(os.Args) < 1 {
		help()
		os.Exit(0)
	}

	flag.Usage = func() {
		help()
	}

	chat := flag.String("target", "", "")
	duration := flag.Int("duration", 0, "")
	flag.Parse()

	if *duration == 0 || *chat == "" {
		help()
		os.Exit(0)
	}

	app.Launch(*chat, *duration)
}
