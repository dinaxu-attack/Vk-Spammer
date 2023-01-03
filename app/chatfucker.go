package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func Launch(chat string, duration int) {
	messagee, err := ioutil.ReadFile("./assets/message.txt")
	if err != nil {
		fmt.Println("Не могу найти файл assets/message.txt")
		os.Exit(0)
	}

	_, err = ioutil.ReadFile("./assets/proxies.txt") /* Чтобы избежать ошибок*/
	if err != nil {
		fmt.Println("Не могу найти файл assets/proxies.txt")
		os.Exit(0)
	}

	message := strings.ReplaceAll(string(messagee), " ", "+") /*  Плюс заменяет пробел */

	fmt.Println(time.Now().Format("15:04:05") + " -> Атака запущена")
	sec := time.Now().Unix()
	pot := 0

	for time.Now().Unix() <= sec+int64(duration)-1 /* Это таймер */ {
		if pot < 10 /* Чтобы пк не взорвался от количества потоков*/ {
			pot++
			go attack(chat, message)
		}
		time.Sleep(time.Duration(duration) / 3 /* Чтобы не засрать пк потоками */ * time.Second)
	}
	fmt.Println(time.Now().Format("15:04:05") + " -> Атака завершена. Всего было запущено " + strconv.Itoa(pot) + " потоков")

}

func attack(chat string, message string) {
	tokens := ReadTokens()

	for _, token := range tokens {

		var ress iinfo
		var cid cchatid

		var url = "https://api.vk.com/method/messages.joinChatByInviteLink?link=" + chat + "&v=5.95&access_token=" + token
		bod := Request(url)
		json.Unmarshal(bod, &cid)

		url = "https://api.vk.com/method/messages.send?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&message=" + message + "&random_id=0&v=5.95&access_token=" + token
		Request(url)

		url = "https://api.vk.com/method/messages.send?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&message=" + message + "&random_id=0&v=5.95&access_token=" + token
		Request(url)

		url = "https://api.vk.com/method/messages.send?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&message=" + message + "&random_id=0&v=5.95&access_token=" + token
		Request(url)
		//  ------------------------------------------------------
		url = "https://api.vk.com/method/account.getProfileInfo?&v=5.95&access_token=" + token
		res := Request(url)
		json.Unmarshal(res, &ress)
		//  ------------------------------------------------------
		url = "https://api.vk.com/method/messages.removeChatUser?chat_id=" + strconv.Itoa(cid.Response.Chatid) + "&user_id=" + strconv.Itoa(ress.Response.Id) + "&member_id=" + strconv.Itoa(ress.Response.Id) + "&v=5.95&access_token=" + token
		Request(url)
	}

}
