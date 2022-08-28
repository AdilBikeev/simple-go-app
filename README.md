# Файловая архитектура

```sh
C:.
└───src
    ├───cmd
    │   └───main # Точка входа
    ├───docker # Файлы для Docker-а
    └───pkg
        ├───config # Конфигурационные файлы, необходимые для настройки при запуске сервера
        ├───controllers # Описание endpoint-ов
        ├───models # сущности
        ├───routes # Содержит роутинги endpoint-Ов
        └───utils # Вспомогательные функции
```

# Get started

1. Переходим в каталог ./src где лежит dpcker-compose.yml
2. Выполняем команды 
```sh
docker-compose build
docker-compose up
```
3. Переходим в папку .\src\cmd\main
4. Выполняем команды 
```sh
go run main.go
```