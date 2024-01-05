# Webscrapper

## note
Проблема:
Площадка/Платформа - не отдают данные по API.

Решение - забирать данные с помощью эмуляции действий пользователя.

Идея:
Веб скраппер с отдачей данных по API.

Суть:

Пользователь прописывать путь до файлов, которые нужно скачивать по API. (через фронт с помощью пресета какого-нибудь)
(Ссылки, креды для авторизации)

На своей стороне через selenium питона мы забираем этот файл и отдаем по апишке пользователю.

Возможные проблемы:
Капчи/Баны подозрительных аккаунтов.

Примерные ручки API приложения:

POST запрос на создание пресета для перехода до файла и его скачивание.
GET запрос на забор данных по пресету.

## Description
This platform resolve problem with getting data from web pages, which don't have API.
Main idea is to get data from web pages with selenium and give it to user with API.

Idea is to create presets for getting data from web pages.
Use can create preset in format like this:
```json
{
    "name": "preset_name",
    "url": "https://www.example.com",
    "login": "login",
    "password": "password",
    "path": [
      "url_one",
      "url_two",
      ...
    ]
}
```

After that user can get data from preset with GET request to API.

For example, user send request to our API with preset unique code and get data from preset:
Write example http request:
```http request
GET http://localhost:8080/api/preset/unique_code/report
```

## Example .env file
```dotenv
# Sharing settings
MODE=prod

# Scrapper app settings
SCRAPPER_PORT=8000
SCRAPPER_HOST=0.0.0.0
SCRAPPER_COUNT_WORKERS=10


# Deployment settings
GITHUB_TOKEN=
```


