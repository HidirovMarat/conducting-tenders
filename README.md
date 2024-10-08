# Conducting tenders

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](http://forthebadge.com/images/badges/built-with-love.svg)](http://forthebadge.com)

Микросервис, который поможет бизнесу организовать тендер на оказание услуг. Участники тендера — другие бизнесы — смогут предложить свои выгодные условия для победы в конкурсе.

Используемые технологии:
- PostgreSQL (в качестве хранилища данных)
- Docker (для запуска сервиса)
- Echo (веб фреймворк)
- golang-migrate/migrate (для миграций БД)
- pgx (драйвер для работы с PostgreSQL)
- golang/mock, testify (для тестирования)

Сервис разработан с использованием Clean Architecture, что обеспечивает простоту расширения его функциональности и удобство тестирования.
Также был реализован Graceful Shutdown для корректного завершения работы сервиса

# Перед началом работы
Для начала надо склонировать к себе репозиторий и настроить .env файл(можно оставить значения по умолчанию). Также написан init.sql файл для DB чтобы выполнялись условия прописаные в задание "Сущности пользователя и организации уже созданы и представлены в базе данных следующим образом:", если это не нужно, тогда просто очистить файл init.sql

# Сборка и Запуск
Сборка и запуск осуществляется через docker и docker-compose. Команда `docker-compose up` запускает проект, a опция `-d` в фоновом режиме.     

с портом 8080 по умолчанию

## Examples

Некоторые примеры запросов
  //TODO

# Решения <a name="decisions"></a>

В ходе разработки был сомнения по тем или иным вопросам, которые были решены следующим образом:

//TODO
