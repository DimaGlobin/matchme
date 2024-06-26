# 1. Цель проекта

Создать приложение для знакомств с ИИ помощником. Пользователь может просматривать анкеты, оценивать их, общаться в чате с пользователями, с которыми у него случился мэтч.

# 2. Описание системы

Система состоит из следующих функциональных блоков:

1. Регистрация, аутентификация, авторизация
2. Функционал мэтчинга, оценки анкет
3. Функицонал оплаты подписки
4. Возможность общения в чате
5. Чат-бот
6. Система уведомлений

## 2.1 Регистрация

Классический алгоритм регистрации. У пользователя должны быть запрошены следующие поля:

* email - обязательное поле
* телефон - обязательное поле
* имя - обязательное поле
* дата рождения - обязательное поле
* описание профиля

Так же во время регистрации пользователю необходимо добатвь хотя бы ОДНУ свою фотографию.

## 2.2 Аутентификация

Аутентификация будет проходить по номеру телефона, на который будет проходить одноразовый код.

## 2.3 Редактирование профиля

У пользователя есть возможность редактировать профиль: добавлять/удалять фотографии, изменить пароль, почту, номер телефона.

Пользователи без подписки имеют право загружать только фото. Пользователи с подпиской смогут загружать ещё и короткие видео (до 15 секунд)

## 2.4 Функционал мэтчинга

Пользователь свайпает анкеты влево и вправо (дизлайк и лайк соответственно), так же у пользователя есть возможность отправить лайк с текстом (суперлайк) и вернуть предыдущий профиль при ошибочной оценке. Пользователи, у которых нет платной подписки будут иметь ограничения:

* лайки - 50 в день
* возврат назад - 5 в день
* суперлайки - 5 в день

## 2.5 Общение в чате

Пользователи, у которых случился мэтч, должны иметь возможность переписываться в чате, для пользователей без подписки должны быть ограничение на количество чатов (максимум - 5)

## 2.6 Чат-бот

Пользователи должны уметь обращаться за помощью к чат-боту, который должен уметь давать советы по общению. Переписка должна сохраняться.

## 2.7 Уведомления

Пользователи должны получать пуш-уведомления о том, что им поставили лайк, у них случился мэтч или пришло какое-то сообщение. Также пользователям можно периодически напоминать о том, что в приложение нужно войти и что-то поделать

# 3. Предлагаемый стек технологий

Для реализации системы предлагается следующий стек технологий:

* Бэкенд:
    - Язык Go
    - БД PostgresQL
    - БД Redis

* Фронтенд:
    - Язык Dart
    - Фреймворк flutter

Хранение фото и видео будет осуществляться в S3-совместимом хранилище.