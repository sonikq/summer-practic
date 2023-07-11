```
1)С самого начала нужно настроить переменные окружения:
 файл для работы с переменными окружения
 находится здесь: configs/app/.env.
 Вы должны настроить параметры сервера, на котором будет
 запускаться и работать приложение(это поля, начинающиеся с HTTP..).
 Так же вы должны настроить параметры подключения к БД
 (это все поля, которые начинаются с POSTGRES).
```
```
2) Для запуска приложения нужно
 в корневой папке в терминале ввеcти
 данную команду: go run cmd/app/main.go
```

```
3) Для сборки приложения в бинарный фаайл введите:
 go build cmd/app/main.go  
```

```
4)В данном приложении реализованы 4 эндпоинта:
 1)/item/add - для добавления новых записей в таблицу
 2)/item/update - для обновления данных в таблице
 3)/item/delete - для удаления конкретной записи в таблице
 4)/table/edit - для добавления новых колонок в таблице
```
```
5)Для корректной работы микро-сервиса
 нужно посылать входные данные в виде JSON`a
5.1)А так же удобно будет скачать приложение Postman
 и работать с микро-сервисом через Postman.
 Как это работает:
    1)Выбираете метод запроса - POST
    2)Вводите сам запрос(например для добавления новых
    данных - localhost:8080/table/edit) и нажимаете кнопку
    "SEND".
```
```
6)Примеры:
    6.1)JSON для эндпоинта /item/add:
        {
        "item_id": 4,
        "name": "dsfsa",
        "designation": "23131",
        "link": "https://...",
        "quantity": 3,
        "placeDIN": 10,
        "placeU": 5,
        "portAIN": 15,
        "portAOT": 7,
        "portDIN": 6,
        "portDOT": 8,
        "portEHT": 25,
        "pci": 3,
        "connRJ": 11,
        "connPIN20": 9,
        "connPIN50": 30,
        "connDB37": 22,
        "connDB62": 70,
        "connSNC55": 41,
        "wireMGTF": 5,
        "wireCAT6": 34,
        "length": 68,
        "braceWasher": 55,
        "braceNut": 77,
        "braceBolt": 88,
        "power": 10.8,
        "voltage": "wqwq",
        "price": 12432.7
        }

    6.2)JSON для эндпоинта /item/update:
        {
        "item_id": 2,
        "name": "ds",
        "designation": "23131",
        "link": "https://.2.",
        "quantity": 3,
        "placeDIN": 10,
        "placeU": 5,
        "portAIN": 15,
        "portAOT": 7,
        "portDIN": 6,
        "portDOT": 8,
        "portEHT": 25,
        "pci": 3,
        "connRJ": 11,
        "connPIN20": 9,
        "connPIN50": 30,
        "connDB37": 22,
        "connDB62": 70,
        "connSNC55": 41,
        "wireMGTF": 5,
        "wireCAT6": 34,
        "length": 68,
        "braceWasher": 55,
        "braceNut": 77,
        "braceBolt": 88,
        "power": 10.8,
        "voltage": "wqwq",
        "price": 12432.7
        }

    6.3)JSON для эндпоинта /item/delete:
        {
        "item_id": 3
        }
    
    6.4)JSON для эндпоинта /table/edit:
        {
        "new_col": "column",
        "type": "string"
        }
```