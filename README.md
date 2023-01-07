# Vk-Spammer
Спамер для ВК, который позволяет рейдить чаты

# Установка
```
Перейдите https://github.com/dinaxu-attack/Vk-Spammer/releases

windows: нажмите raider.exe

linux: нажмите raider

Дальше создайте папку с любым названием, в нее засуньте файл и создайте (в той же папке) новую папку с названием assets, в этой папке создайте tokens.txt

В config.json укажите сообщение (message)
В config.json укажите прокси (proxies)
В config.json укажите апи ключ от анти капчи (anticaptcha)
В assets/tokens.txt напишите вк токены

Перейдите к след. шагу

```

# Использование

Windows:

```
raider.exe --target https://vk.me/join... --duration 50 --delay 3 --namechanger (необяз.) --firstname Имя --lastname Фамилия
target - Ссылка на беседу
duration - Длительность спама
delay - Задержка 
namechanger - Меняет ники (Для этого нужно будет указать firstname и lastname)
```

Linux:
```
raider --target https://vk.me/join... --duration 50 --namechanger (необяз.) --firstname Имя --lastname Фамилия

target - Ссылка на беседу
duration - Длительность спама
delay - Задержка 
namechanger - Меняет ники (Для этого нужно будет указать firstname и lastname)
```
