# Calculator Service

Это проект калькулятора, который принимает арифметические выражения и вычисляет их с помощью агентов.

## Установка

1. Клонируйте репозиторий:
   ```bash
   git clone <url>
   cd calculator
2. Установите зависимости:

go mod tidy
3. Запустите сервер:

go run ./cmd/calc_service/...

Управление
Для работы через терминал на Curl настоятельно рекомендую использовать Git Bash.
Но можно использовать и Postman - приложение для отправки HTTP-запросов, по-моему в нём работать проще.

Добавление вычисления арифметического выражения
Добавление выражения для вычисления на API.

Curl
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": <выражение>
}'
Postman
URL localhost/api/v1/calculate;
Запрос POST;
Body RAW {"expression": <выражение>};
Нажать на SEND.
Коды ответа:
201 - выражение принято для вычисления
422 - невалидные данные
500 - что-то пошло не так
Тело ответа
{
    "id": <уникальный идентификатор выражения> // его ID
}
Получение списка выражений
Curl
curl --location 'localhost/api/v1/expressions'
Postman
URL localhost/api/v1/expressions;
Запрос GET;
Body NONE;
Нажать на SEND.
Тело ответа:
Получение всех сохранённых выражений(ID не нужен).

{
    "expressions": [
        {
            "id": 8251431,
            "status": "OK",
            "result": 3>
        },
        {
            "id": 34942763,
            "status": "Wait",
            "result": 0
        }
    ]
}
Коды ответа:
200 - успешно получен список выражений
500 - что-то пошло не так
Получение выражения по его идентификатору
Получение выражения по его идентификатору.

Примечание: Для того, чтобы получить выражение по его ID, необходимо сохранить полученный при Добавление вычисления арифметического выражения индитефикатор.

Curl
curl --location 'localhost/api/v1/expressions/<id выражения>'
Postman:
URL localhost/api/v1/expressions/<id выражения>;
Запрос GET;
Тело NONE;
Нажать на SEND.
Тело ответа:
{
    "expression":
        {
            "id": <идентификатор выражения>,
            "status": <статус вычисления выражения>,
            "result": <результат выражения>
        }
}
Коды ответа:
200 - успешно получено выражение
404 - нет такого выражения
500 - что-то пошло не так
Примеры работы с API
Простой пример
Делаем запрос на вычисление выражения
Curl
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2/2"
}'
Postman
URL localhost/api/v1/calculate;
Запрос POST;
Body RAW {"expression": "2+2/2"};
Нажать на SEND.
Ответ
Статус 201(успешно создано);

{
    "id": 12345 // пример
}
Получаем наше выражение
Curl
curl --location 'localhost/api/v1/expressions/12345' // 12345 - это ID выше.
Postman
URL localhost/api/v1/expressions/12345;
Запрос GET;
Тело NONE;
Нажать на SEND
Ответ
Статус 200(успешно получено);

{
    "expression":
        {
            "id": 12345,
            "status": "OK",
            "result": 321
        }
}
Получаем все выражения
Curl
curl --location 'localhost/api/v1/expressions'
Postman
URL localhost/api/v1/expressions;
Запрос GET;
Body NONE;
Нажать на SEND.
Ответ
Статус 200(успешно получены);

{
    "expressions": [
        {
            "id": 12345,
            "status": "OK",
            "result": 321
        },
    ]
}
Пример с ошибкой в запросе №1
Делаем неправильный запрос на вычисление выражения
Curl
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "radhgsags": "2+2/2"
}'
Postman
URL localhost/api/v1/calculate;
Запрос POST;
Body RAW {"radhgsags": "2+2/2"};
Нажать на SEND.
