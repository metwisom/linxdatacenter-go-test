# linxdatacenter-go-test

Тестовое задание для Linxdatacenter

## Особенности

```
Для CSV можно обрабатывать огромные файлы т.к. там построчное чтение, тестил на 2кк записей.
Для JSON такое не провернуть, формат не позволяет
```

## Использование

```sh
cd app
go run . ./data/1.json      #Запуск для JSON
go run . ./data/1.csv       #Запуск для CSV
go run . ./data/1.csv ","   #Запуск для CSV со специфичным разделителем
go build -o app             #Собрать пакет
./app ./data/1.csv          #Запустить собранный пакет
```

## Docker 
```sh
sudo docker build . -t linxtest                                                         #Собираем образ
sudo docker run -v $(pwd)/app/data/:/var/data --rm linxtest /var/data/1.csv ","     #Запускам для CSV прокидывая папку с хоста
sudo docker run -v $(pwd)/app/data/:/var/data --rm linxtest /var/data/1.json        #Запускам для JSON прокидывая папку с хоста

#Можно и не прокидывать, данные копируются в образ при билде
```