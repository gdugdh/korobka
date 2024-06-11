### Добавление функционала:

1. Создать .proto файл в `/api/API_NAME_v1/API_NAME.proto`
2. Создать цель в Makefile `generate-API_NAME-api`, добавить эту цель в `generate: ...`
3. Добавить файл `/internal/api/API_NAME/service.go`.
    - Создать структуру наследуюшую структуру сгенерированную Protobuf-конвертеромом
    - Добавить методы, сервиса из `API_NAME.proto`
4. Запросы к БД требуемые для сервеиса, записывать в `/internal/repository`

### Запуск

1. Установить Golang зависимости
2. make run

Запуститься два сервера

-   GRPC: `localhost:50051`
-   HTTP зеркало: `localhost:8081`

Открыть `frontend/index.html`
