# Веб-сервис Calculator

## Описание

Программа представляет собой API-сервис для конкурентного вычисления выражений с веб интерфейсом.

## Предназначение

Пользователь хочет считать арифметические выражения. Он вводит строку `2 + 2 * 2` и хочет получить в ответ `6`.<br>
Но наши операции сложения и умножения (также деления и вычитания) выполняются "очень-очень" долго(*симуляция*).<br>
Поэтому вариант, при котором пользователь делает *http-запрос* и получает в качетсве ответа результат, невозможна.<br>
Более того, вычисление каждой такой операции в нашей "альтернативной реальности" занимает "гигантские" вычислительные мощности.<br>
Соответственно, каждое действие мы должны уметь выполнять отдельно и масштабировать эту систему можем добавлением вычислительных мощностей в нашу систему в виде новых "машин".<br>
Поэтому пользователь может с какой-то периодичностью уточнять у сервера "не посчиталость ли выражение"?<br>
Если выражение наконец будет вычислено - то он получит результат.<br>

## Установка

<<<<<<< HEAD
 - Для установки нужно выбрать директорию проекта:
 - Потом необходимо выполнить эту команду:
```bash
git clone https://github.com/MP5s/calculator
```
 - В выбранной папке появится папка ```calculator``` c проектом.
=======
1. Клонируйте репозиторий:
   ```bash
   git clone <url>
   cd calculator
2. Установите зависимости:
>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5

## Работа с API

<<<<<<< HEAD
### Конфигурация
#### Переменные среды
Сначала необходимо открыть файл ```./config/.env``` и установить параметры:

 - **TIME_ADDITION_MS** - время вычисления сложения(в миллисекундах);

 - **TIME_SUBTRACTION_MS** - время вычисления вычитания;

 - **TIME_MULTIPLICATIONS_MS** - время вычисления умножения;

 - **TIME_DIVISIONS_MS** - время вычисления деления;

 - **COMPUTING_POWER** - максмальное количество *worker*'ов, которые параллельно выполняют арифметические действия.

#### Другие параметры

Потом необходимо открыть файл ```config.json``` в той же папке и установить следущие параметры(**true** - включено, **false** - выключено):

 - ```debug``` - отладка(вывод событий в лог)

 - ```web``` - веб-интерфейс(об его использовании читайте дальше в **Веб-интерфейс**)

По умолчанию и то и другое выключено.

### Запуск
 - Для запуска API необходимо выбрать директорию проекта:
```
cd <путь к папке Calculator_Service>
```
 - Далее надо запустить файл ```cmd/main.go```:
```
go run cmd/main.go
```

### Управление

Для работы через терминал на *Curl* настоятельно рекомендую использовать **Git Bash**.<br> Но можно использовать и **Postman** - приложение для отправки *HTTP-запросов*, по-моему в нём работать проще.

#### Добавление вычисления арифметического выражения

Добавление выражения для вычисления на **API**.

##### Curl
```
=======
go run ./cmd/calc_service/...

Управление
Для работы через терминал на Curl настоятельно рекомендую использовать Git Bash.
Но можно использовать и Postman - приложение для отправки HTTP-запросов, по-моему в нём работать проще.

Добавление вычисления арифметического выражения
Добавление выражения для вычисления на API.

Curl
>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": <выражение>
}'
<<<<<<< HEAD
```
##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"expression": <выражение>};
 - Нажать на ****SEND****.

##### Коды ответа: 
 - 201 - выражение принято для вычисления
 - 422 - невалидные данные
 - 500 - что-то пошло не так

##### Тело ответа

```
{
    "id": <уникальный идентификатор выражения> // его ID
}
```
#### Получение списка выражений
##### Curl
```
curl --location 'localhost/api/v1/expressions'
```
##### Postman
 - **URL** localhost/api/v1/expressions;
 - Запрос **GET**;
 - **Body** **NONE**;
 - Нажать на ****SEND****.

##### Тело ответа:

Получение всех сохранённых выражений(**ID** не нужен).

```
=======
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

>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
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
<<<<<<< HEAD
```
##### Коды ответа:
 - 200 - успешно получен список выражений
 - 500 - что-то пошло не так

#### Получение выражения по его идентификатору

Получение выражения по его идентификатору.

*Примечание:* Для того, чтобы получить выражение по его ID, необходимо сохранить полученный при **Добавление вычисления арифметического выражения** индитефикатор.

##### Curl

```
curl --location 'localhost/api/v1/expressions/<id выражения>'
```

##### Postman:
 - **URL** localhost/api/v1/expressions/<id выражения>;
 - Запрос **GET**;
 - Тело **NONE**;
 - Нажать на ****SEND****.

##### Тело ответа:

```
=======
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
>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
{
    "expression":
        {
            "id": <идентификатор выражения>,
            "status": <статус вычисления выражения>,
            "result": <результат выражения>
        }
}
<<<<<<< HEAD
```

##### Коды ответа:
 - 200 - успешно получено выражение
 - 404 - нет такого выражения
 - 500 - что-то пошло не так

## Примеры работы с API

### Простой пример

#### Делаем запрос на вычисление выражения

##### Curl
```
=======
Коды ответа:
200 - успешно получено выражение
404 - нет такого выражения
500 - что-то пошло не так
Примеры работы с API
Простой пример
Делаем запрос на вычисление выражения
Curl
>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2/2"
}'
<<<<<<< HEAD
```
##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"expression": "2+2/2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 201(успешно создано);
```
{
    "id": 12345 // пример
}
```

#### Получаем наше выражение
##### Curl
```
curl --location 'localhost/api/v1/expressions/12345' // 12345 - это ID выше.
```
##### Postman
 - **URL** localhost/api/v1/expressions/12345;
 - Запрос **GET**;
 - Тело **NONE**;
 - Нажать на ****SEND****

##### Ответ
Статус 200(успешно получено);
```
=======
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

>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
{
    "expression":
        {
            "id": 12345,
            "status": "OK",
            "result": 321
        }
}
<<<<<<< HEAD
```

#### Получаем все выражения
##### Curl
```
curl --location 'localhost/api/v1/expressions'
```
##### Postman
 - **URL** localhost/api/v1/expressions;
 - Запрос **GET**;
 - **Body** **NONE**;
 - Нажать на ****SEND****.

##### Ответ
Статус 200(успешно получены);
```
=======
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

>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
{
    "expressions": [
        {
            "id": 12345,
            "status": "OK",
            "result": 321
        },
    ]
}
<<<<<<< HEAD
```

### Пример с ошибкой в запросе №1

#### Делаем **неправильный** запрос на вычисление выражения

##### Curl

```
=======
Пример с ошибкой в запросе №1
Делаем неправильный запрос на вычисление выражения
Curl
>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "radhgsags": "2+2/2"
}'
<<<<<<< HEAD
```

##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"radhgsags": "2+2/2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 422(**неправильный** запрос);


### Пример с ошибкой в запросе №2

#### Делаем **правильный** запрос на вычисление выражения

##### Curl
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2/2"
}'
```


##### Postman
 - **URL** localhost/api/v1/calculate;
 - Запрос **POST**;
 - **Body** **RAW** {"radhgsags": "2+2/2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 201(успешно создано);
```
{
    "id": 12345 // пример
}
```
#### Далее получаем наше выражение(**неправильный** ID)
##### Curl
```
curl --location 'localhost/api/v1/expressions/45362'
```

##### Postman:
 - **URL** localhost/api/v1/expressions/45362;
 - Запрос **GET**;
 - Тело **NONE**;
 - Нажать на ****SEND****.

##### Ответ
Статус 404(не найдено);


### Пример с ошибкой в запросе №3

#### Делаем запрос с **некорректным** URL на вычисление выражения

##### Curl
```
curl --location 'localhost/api/v1/abc' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "121+2"
}'
```
##### Postman
 - **URL** localhost/api/v1/abc;
 - Запрос **POST**;
 - **Body** **RAW** {"expression": "121+2"};
 - Нажать на ****SEND****.

##### Ответ
Статус 404(**NOT FOUND**);


### Веб-интерфейс

Вот ссылки на веб-страницы:

 - [Главная страница](http://localhost:8080/api/v1/web)
 - [Вычисление выражения](http://localhost:8080/api/v1/web/calculate)
 - [Просмотр всех выражений](http://localhost:8080/api/v1/web/expressions)
 - [Просмотр выражения по его ID](http://localhost:8080/api/v1/web/expression)

****ВАЖНО:**** По умолчанию веб-интерфейс выключен. Чтобы его включить, нужно изменить параметр *Веб интерфейс* в **Конфигурация/Другие Параметры**.
=======
Postman
URL localhost/api/v1/calculate;
Запрос POST;
Body RAW {"radhgsags": "2+2/2"};
Нажать на SEND.
>>>>>>> 74a7658c4def7e05c5584b7120ecc3fec9d58df5
