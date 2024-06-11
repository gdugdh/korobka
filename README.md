### Добавление функционала:

1. Создать .proto файл в /api/API_NAME_v1/API_NAME.proto
2. Создать цель в Makefile "generate-API_NAME-api", добавить эту цель в generate: ...
3. Добавить файл /internal/api/API_NAME/service.go.
    - Создать структуру имплементирующую структуру сгенерированную Protobuf converter-ом
    - Добавить методы, сервиса из API_NAME.proto
    - В каждом методе будут вызываться одноименные методы из /internal/service/
4. в /internal/service/service.go добавить interface с методами, которую буду вызываться в 3 шаге
5. Создать в /internal/service/API_NAME/service.go со всеми методами бизнес-логики
    - Все методы здесь

### Запуск

1. Установить Golang зависимости
2. make run
