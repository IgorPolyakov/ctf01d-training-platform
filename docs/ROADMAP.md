# MVP - версия #1

## Общее

- [ ] Авторизация
    - [x] Сессии в БД
    - [ ] Закрыть странички за авторизацию
- [x] Регистрация
    - [x] Через админку будет добавление новых пользователей
- [x] База данных (Postgres)
    - [x] Переехать с MySQL на PostgreSQL
    - [x] Добавить seed
    - [x] Версионирование для облегчения обновлений сервера
    - [x] Сделать изменения через миграции, а не init.sql
- [x] CI/CD
    - [x] Линтер
    - [x] Проверка зависимостей Nancy
    - [x] Сборка контейнера и отправка в registry

## Админка

- [x] Пользователи
    - [x] Создание
    - [x] Редактирование
    - [x] Удаление
- [x] Команды
    - [x] Создание
    - [x] Редактирование
    - [x] Удаление
- [x] Сервисы
    - [x] Создание
    - [x] Редактирование
    - [x] Удаление
- [ ] Игры
    - [x] Создание
    - [ ] По расписанию (какие сервисы, какие команды и когда)
    - [ ] При загрузке zip для сервиса необходимо проверить корректность работы чекера (проверить все коды и состояния)
- [ ] Результаты
    - [x] Создание
    - [ ] Редактирование
    - [x] Удаление

## Профиль пользователя

- [x] Информация о команде (командах), в которой находится пользователь
- [x] История по командам (в каких был и сколько)

## Сервисы

- [x] Отображение списка сервисов (которые разрешено показывать публике)
- [x] Автор, логотип
- [ ] Возможность скачать zip, райтапы

## Команды

- [x] Отображение списка команд
- [x] По команде - аватарка, университет, ссылки на социальные сети, краткое описание
- [ ] Текущий капитан (и его зам и зам зам капитана)
- [ ] Управление списком членов команды владельцем команды
- [ ] Возможность владельца команды приглашать пользователей в команду
- [x] Возможность пользователей подавать заявки на вступление в команду
- [ ] Ведение истории команды (создание, изменения состава, слияния и т.д.)

## Игры

- [ ] Список кто с кем и когда
- [ ] Когда игру открывают - доступ (например, через OpenVPN) для соответствующих команд и пользователей
- [ ] Должен быть scoreboard
- [ ] По завершению игры - сохранение результата scoreboard

## Дополнительные требования

- [ ] Процесс принятия в команду через подтверждение
- [ ] Возможность обновления аватарки и информации о себе пользователем
- [ ] Управление ролями в командах
- [ ] Возможность админам подтверждать изменения
- [ ] Ведение исторической информации и метаданных (события, действия, состояния)
- [ ] Страница для владельца команды с управлением членами команды и приглашениями


# MVP - версия #2

## Профиль пользователя

- [ ] Обновление аватарки
- [ ] Возможность обновления информации о себе: навыки, описание, контактная информация (без даты рождения и номера телефона)