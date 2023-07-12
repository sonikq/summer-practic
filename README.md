```
1)Настройте переменные окружения: configs/app/.env
```
```
2) Для запуска приложения: go run cmd/app/main.go или make run
```

```
3) Для сборки приложения: go build cmd/app/main.go  
```

```
4)В данном приложении реализованы 4 эндпоинта:
 1)/item/add - для добавления новых записей в таблицу
 2)/item/update - для обновления данных в таблице
 3)/item/delete - для удаления конкретной записи в таблице
 4)/table/edit - для добавления новых колонок в таблице
```
```
5)Примеры:
    5.1)JSON для эндпоинта /item/add:
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

    5.2)JSON для эндпоинта /item/update:
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

    5.3)JSON для эндпоинта /item/delete:
        {
        "item_id": 3
        }
    
    5.4)JSON для эндпоинта /table/edit:
        {
        "new_col": "column",
        "type": "string"
        }
```