# JSON RPC
JSON RPC обеспечивают взаимодействие системы АРМ c интерфейсом пользователя.

Взаимодействие осуществляется по протоколу JSON-RPC 2.0 ([about JsonRpc protocol](http://example.com/ "codes and error formats"))


Все запросы к JSON-RPC выполняются к `http://host-url:port/api`<br>
В ответе на запрос всегда содержится объект.<br>
Успешный ответ всегда содержит ключ `result`, если произошла какая либо ошибка, информация о ней будет
содержаться в поле `error`<br>

------------------------------------
### Краткое описание протокола.
Все запросы передаются методом POST и должны иметь заголовок `content-type: application/json`

Внутри тела запроса передается json-объект с описанием вызываемого метода

_Пример:_
```json
POST 'http://localhost:4601/api' HTTP/1.1
content-type: application/json

{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "exploration.list",
    "params": {
        "offset": 0,
        "limit": 1,
        "orders": [{
            "column": "dateAdd",
            "type": "DESC"
        }],
        "filter": {
            "name": "11"
        }
    }
}
```
В случае успешного выполнения приходит ответ вида:
```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 17 Dec 2021 11:04:27 GMT
Content-Length: 363
Connection: close

{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "list": [
        {
            "accessType": "all",
            "authorName": "",
            "authorUUID": "00000000-0000-0000-0000-000000000000",
            "comment": "2rr",
            "dateAdd": "2021-11-22T10:42:24.071774+03:00",
            "dateEdit": "2021-11-29T15:32:53.868258+03:00",
            "deleted": false,
            "name": "1111",
            "tags": [],
            "uuid": "8d0ea168-0ae3-45ea-a406-ca996a4d5da7",
            "views": 0
        }
        ],
        "total": 21
    },
    "error": null
}
```
<br>

# Запросы к API:

<pre>
Exploration

POST http://localhost:4601/api

Доступные методы - {create, list, edit, details, deleted}
</pre>

Системные элементы:
- `authorUUID`&nbsp; - &emsp;идентификатор автора
- `authorName`&nbsp; - &emsp;имя автора (пока не используется)
- `dateAdd`&nbsp; &emsp;&nbsp; - &emsp;дата создания записи
- `dateEdit`&nbsp;&emsp; - &emsp;дата редактирования
- `deleted`&nbsp; &emsp;&nbsp;  - &emsp;флаг "удаленный", по умолчанию `false`

[comment]: <> ()

[comment]: <> ()


---
##  ***( Create )*** ##
### Метод `exploration.create` - позволяет создать исследование: ###

\
___Доступные параметры:___&emsp;`name`, `comment`, `tags`, `accessType`

Запрос на создание исследования: 
```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "exploration.create",
    "params": {
        "name": "test",
        "comment": "test",
        "tags": [
            "1",
            "2"
        ],
        "accessType": "test"
    }
}
```
Успешный ответ возвращает все параметры созданного исследования, `uuid` и системные параметры присваиваются автоматически:

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "accessType": "test",
        "authorName": "",
        "authorUUID": "00000000-0000-0000-0000-000000000000",
        "comment": "test",
        "dateAdd": "2021-12-21T12:22:36.487173+03:00",
        "dateEdit": "2021-12-21T12:22:36.487173+03:00",
        "deleted": false,
        "name": "test",
        "tags": [
            "1",
            "2"
        ],
        "uuid": "a3db6830-d48e-41a6-8354-191570e634cf",
        "views": 0
    },
    "error": null
}
```


##  ***( List )*** ##
### Метод `exploration.list` - позволяет получить список элементов исследования: ###

    По умолчанию сортировка набора результатов запроса производится по дате создания и в порядке убывания ( `ORDER BY dateAdd DESC` )

Доступные параметры:
- `offset`&emsp;-&nbsp; пропускает указанное число строк
- `limit`&emsp;&nbsp; -&nbsp; выводит указанное число строк
- `orders`&emsp;-&nbsp; для изменений параметров сортировки результата
- `filter`&emsp;-&nbsp; осуществляет фильтрацию по имени
<br>
<br>

Пример запроса:
```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "exploration.list",
    "params": {
        "orders": [{
            "column": "dateAdd",
            "type": "DESC"
        }],
      
        "filter": {
            "name": ""
        }
    }
}
```

Успешный ответ возвращает исследования с параметром ` "deleted" : false ` и количество исследований:

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "list": [
            {
              "accessType": "",
              "authorName": "",
              "authorUUID": "00000000-0000-0000-0000-000000000000",
              "comment": "",
              "dateAdd": "2021-12-20T16:39:10.435635+03:00",
              "dateEdit": "2021-12-20T16:39:10.435635+03:00",
              "deleted": false,
              "name": "",
              "tags": null,
              "uuid": "9a6ed1d4-86d2-4706-b406-2910219ca198",
              "views": 0
            }
        ],

        "total": 21
    },
    "error": null
}
```

##  ***( Edit )*** ##
### Метод `exploration.edit` - позволяет изменить параметры исследования: ###
\
`uuid` - обязательный параметр, содержит идентификатор изменяемого исследования

___Изменяемые параметры:___&emsp;`name`, `comment`, `tags`, `accessType`
  <br>
  <br>

Запрос на редактирование исследования с параметром `uuid : "a3db6830-d48e-41a6-8354-191570e634cf"`

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "exploration.edit",
    "params": {
        "uuid": "a3db6830-d48e-41a6-8354-191570e634cf",
        "name": "edit",
        "tags": [
            "111",
            "222"
        ],
        "accessType": "testEdit",
        "comment": "testEdit"
    }
}
```

Успешный ответ:

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "accessType": "testEdit",
        "authorName": "",
        "authorUUID": "00000000-0000-0000-0000-000000000000",
        "comment": "testEdit",
        "dateAdd": "2021-12-21T12:22:36.487173+03:00",
        "dateEdit": "2021-12-22T10:01:29.003063+03:00",
        "deleted": false,
        "name": "edit",
        "tags": [
            "111",
            "222"
        ],
        "uuid": "a3db6830-d48e-41a6-8354-191570e634cf",
        "views": 0
    },
    "error": null
}
```

##  ***( Delete )*** ##
### Метод `exploration.delete` - позволяет удалить исследования: ###
\
`uuid` - обязательный параметр, содержит идентификатор удаляемого исследования

\
Пример:
```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "exploration.delete",
    "params": {
        "uuid": "a3db6830-d48e-41a6-8354-191570e634cf"
    }
}
```

В указанном исследовании будет установлен параметр `удаленный ( "deleted" : true )`:

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "result": {
        "accessType": "testEdit",
        "authorName": "",
        "authorUUID": "00000000-0000-0000-0000-000000000000",
        "comment": "testEdit",
        "dateAdd": "2021-12-21T12:22:36.487173+03:00",
        "dateEdit": "2021-12-22T10:01:29.003063+03:00",
        "deleted": true,
        "name": "edit",
        "tags": [
            "111",
            "222"
        ],
        "uuid": "a3db6830-d48e-41a6-8354-191570e634cf",
        "views": 0
    },
    "error": null
}
```

##  ***( Details )*** ##
### Метод `exploration.details` - позволяет получить параметры не удаленного исследования: ###
\
`uuid` - обязательный параметр, содержит идентификатор интересующего исследования

\
Метод работает только с `не удаленными` исследованиям ( ` "deleted" : false ` )

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "exploration.details",
    "params": {
        "uuid": "a3db6830-d48e-41a6-8354-191570e634cf"
    }
}
```




\
\
\
\
<br>


####  ***( END )*** ####
