package main

import (
	"flag"
	"fmt"
	"os"
	"raider/app"
	"runtime"
	"time"
)

func help() {
	if runtime.GOOS == "windows" {
		fmt.Println("Usage:\nraider.exe --target https://vk.me/join... --duration 50 --delay 4 --namechanger --name Афанасий Князев")
	} else {
		fmt.Println("Usage:\n./raider --target https://vk.me/join... --duration 50 --delay 4 --namechanger --name Афанасий Князев")
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
	delay := flag.Int("delay", 00, "")
	namechanger := flag.Bool("namechanger", false, "")
	fname := flag.String("firstname", "", "")
	lname := flag.String("lastname", "", "")
	flag.Parse()

	if *duration == 0 || *chat == "" || *delay == 00 {
		help()
		os.Exit(0)
	}

	app.Read()
	if *namechanger == true {
		if *fname == "" || *lname == "" {
			help()
			os.Exit(0)
		}
		fmt.Println(time.Now().Format("15:04:05") + " -> Начинаю менять ники")
		app.Namechanger(*fname, *lname)
	}
	app.Launch(*chat, *duration, *delay)
}
