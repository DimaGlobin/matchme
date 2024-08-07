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

### Функциональность чат-бота
#### Начальная версия
 - **История переписки**: Пользователь может в любой момент обратиться к чат-боту за советом. При этом чат-бот получает в качестве промпта всю историю переписки.
 - **Анализ и генерация ответа**: Языковая модель (например, Llama3) тщательно анализирует историю переписки. На основе этого анализа чат-бот генерирует сообщение, которое пользователь может скопировать и отправить собеседнику.
#### Полная версия
Чат-бот станет еще более функциональным, превратившись в нечто наподобие второго пилота в вашей переписке:

 - **Генерация и корректировка сообщений**: Чат-бот не только будет генерировать сообщения, но и сможет самостоятельно вносить небольшие корректировки и отправлять их собеседнику. Это сделает общение более плавным и естественным.
 - **Персонализация и адаптация**: С течением времени чат-бот будет анализировать все сообщения пользователя, создавая его уникальный портрет. Он перенимет манеру общения пользователя, что сделает его советы еще более персонализированными и точными.
#### Подписки и лимиты
В зависимости от уровня подписки пользователи будут иметь разные возможности:

 - **Лимит сообщений**: Базовая подписка будет предоставлять ограниченное количество сообщений, которые можно отправить через чат-бота. Более продвинутые подписки увеличат этот лимит.
 - **Лимит токенов**: Ограничение на количество токенов, используемых в генерации сообщений, также будет варьироваться в зависимости от подписки.

#### Преимущества
 - **Эффективное общение**: Чат-бот поможет пользователям избежать неловких ситуаций, предлагая грамотно составленные ответы.
 - **Время и усилия**: Пользователи смогут сэкономить время и усилия, доверив составление сообщений искусственному интеллекту.
 - **Персонализация**: Анализируя стиль общения пользователя, чат-бот будет становиться всё более точным и полезным.

### Заключение
Чат-бот помощник по общению представляет собой мощный инструмент для улучшения коммуникации. Он будет всегда на вашей стороне, готовый помочь в любой ситуации, предложить наилучший ответ и даже отправить его за вас. С течением времени он станет вашим персональным помощником, который понимает вас с полуслова и помогает вести беседы на высшем уровне.
### Стек технологий для реализации ML/DS составляющей
#### 1. Языковые модели и NLP

- **Llama3 или аналогичные языковые модели**: Основной инструмент для обработки и генерации текста.
- **Hugging Face Transformers**: Библиотека для работы с предобученными моделями и их дообучением.
- **NLTK, SpaCy**: Библиотеки для предобработки текста и синтаксического анализа.

#### 2. Обучение и дообучение моделей

- **PyTorch**: Фреймворки для обучения и дообучения моделей машинного обучения.
- **Hugging Face Datasets**: Инструмент для работы с наборами данных для обучения моделей.

#### 3. Обработка текста и векторизация

- **TF-IDF, Word2Vec, GloVe**: Методы векторизации текста для анализа сходства между текстами.
- **Sentence Transformers**: Модели для получения эмбеддингов предложений.




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