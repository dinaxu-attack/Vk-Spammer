package app

func Namechanger(first, last string) {
	tokens := ReadTokens()
	for _, token := range tokens {
		var url = "https://api.vk.com/method/account.saveProfileInfo?first_name=" + first + "&last_name=" + last + "&relation=7&v=5.95&access_token=" + token
		Request(url)

	}
}
