package app

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type iinfo struct {
	Response struct {
		Id int `json:"id"`
	}
}

type cchatid struct {
	Response struct {
		Chatid int `json:"chat_id"`
	}
}

func Launch(chat string, duration int, delay int) {
	messagee := Message

	message := strings.ReplaceAll(string(messagee), " ", "+") /*  Плюс заменяет пробел */

	go Title()
	fmt.Println(time.Now().Format("15:04:05") + " -> Атака запущена")
	sec := time.Now().Unix()
	pot := 0

	for time.Now().Unix() <= sec+int64(duration)-1 /* Это таймер */ {
		if pot < 20 /* Чтобы пк не взорвался от количества потоков*/ {
			pot++
			go attack(chat, message)
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
	fmt.Println(time.Now().Format("15:04:05") + " -> Атака завершена. Всего было запущено " + strconv.Itoa(pot) + " потоков")
	fmt.Println("\nУспешно отправлено сообщений: " + strconv.Itoa(sent) + "\nРешено капч: " + strconv.Itoa(solved) + "\nОшибок: " + strconv.Itoa(errors))

}

func attack(chat string, message string) {
	tokens := ReadTokens()

	for _, token := range tokens {

		var ress iinfo
		var cid cchatid

		var url = "https://api.vk.com/method/messages.joinChatByInviteLink?link=" + chat + "&v=5.95&access_token=" + token
		bod := Request(url)
		json.Unmarshal(bod, &cid)

		url = "https://api.vk.com/method/messages.send?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&message=" + message + "+|+" + RandomString(10) + "&random_id=0&v=5.95&access_token=" + token
		st := Request(url)
		Handler(st, url)

		url = "https://api.vk.com/method/messages.send?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&message=" + message + "+|+" + RandomString(10) + "&random_id=0&v=5.95&access_token=" + token
		st = Request(url)
		Handler(st, url)

		url = "https://api.vk.com/method/messages.send?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&message=" + message + "+|+" + RandomString(10) + "&random_id=0&v=5.95&access_token=" + token
		st = Request(url)
		Handler(st, url)
		//  ------------------------------------------------------
		url = "https://api.vk.com/method/account.getProfileInfo?&v=5.95&access_token=" + token
		res := Request(url)
		json.Unmarshal(res, &ress)
		//  ------------------------------------------------------
		url = "https://api.vk.com/method/messages.removeChatUser?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&user_id=" + strconv.Itoa(ress.Response.Id) + "&member_id=" + strconv.Itoa(ress.Response.Id) + "&v=5.95&access_token=" + token
		Request(url)
	}

}

var sent int
var solved int
var errors int

func Title() {
	for {
		switch runtime.GOOS {
		case "darwin":
			os.Stderr.WriteString("\033]0;SENT: " + strconv.Itoa(sent) + " :: SOLVED: " + strconv.Itoa(solved) + "  :: ERRORS: " + strconv.Itoa(errors) + "\007")
		case "linux":
			os.Stderr.WriteString("\033]0;SENT: " + strconv.Itoa(sent) + " :: SOLVED: " + strconv.Itoa(solved) + "  :: ERRORS: " + strconv.Itoa(errors) + "\007")
		case "windows":
			exec.Command("cmd", "/C", "title", "SENT: "+strconv.Itoa(sent)+" :: SOLVED: "+strconv.Itoa(solved)+"  :: ERRORS: "+strconv.Itoa(errors))
		default:
			os.Stderr.WriteString("\033]0;SENT: " + strconv.Itoa(sent) + " :: SOLVED: " + strconv.Itoa(solved) + "  :: ERRORS: " + strconv.Itoa(errors) + "\007")
		}
		time.Sleep(1 * time.Second)
	}
}
