# Link shortener

## Запуск приложения

Для запуска выполните команды из корневой директории приложения:
```
sudo docker build -t linkshortener .
sudo docker-compose up
```
После этого приложение запущено на localhost:8080

## Примечание:

Приложение разрабатывалось в среде WSL Ubuntu Linux, поэтому команды в Makefile есть как для Desktop Ubuntu (20), так и для WSL Ubuntu Linux.
Запуск приложения через docker-compose протестирован как на wsl, так и на обычной ubuntu, поэтому проблем быть не должно.
База данных не пробрасывается наружу, а запускается в контейнере и инициируется с помощью скрипта mongo-init.js
В приложении есть немного хардкода возвращаемых адресов, потому что я не нашел способ корректного вывода как при простом запуске на машине, так 
и запуске в качестве контейнера.

## Технологии:

 - Сервер написан на Golang
 - Для хранения данных - MongoDB

## Основные методы сервера

Получить сокращенную ссылку (сгенерируется автоматически, длина настраивается в configs/server.toml):
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"longurl": "https://avito.ru/moscow"}' \
  http://localhost:8080/create
```

Ответ:
```
{"shorturl":"http://localhost:8080/68Ad2x9"}
```

Сделать свою кастомную ссылку:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
      "longurl": "https://start.avito.ru/tech",
      "shorturl": "trainee"}' \
  http://localhost:8080/createcustom
```

Ответ:
```
{"shorturl":"http://localhost:8080/trainee"}
```

Перейти по сокращенной ссылке:
```
curl --header "Content-Type: application/json" \
  --request GET \
  http://localhost:8080/{trainee}
```
после этого сервер перенаправит вас на исходную ссылку или выдаст NotFound если ссылка не будет найдена в базе данных

## Тестирование

Для запуска тестов на все приложение запустить:
```
make test
```
- Не все методы работают корректно, поэтому тесты падают.
