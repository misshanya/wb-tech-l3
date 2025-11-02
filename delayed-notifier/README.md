# Delayed Notifier

Сервис для отложенных уведомлений через RabbitMQ.

Поддерживает уведомления через:

- [Telegram](https://t.me)
- [ntfy](https://ntfy.sh)

## Установка и запуск

1. Склонируйте репозиторий

    ```shell
    git clone https://github.com/misshanya/wb-tech-l3
    cd wb-tech-l3/delayed-notifier
    ```

2. Скопируйте .env.example в .env и отредактируйте значения под себя (например, токен Telegram бота)

    ```shell
    cp .env.example .env
    ```

3. Запустите через Docker Compose

    ```shell
    docker compose up -d
    ```

## Использование

- Для создания уведомления отправьте `POST` запрос на `/api/v1/notify` со следующим телом:

    ```json5
    {
        "scheduled_at": "2025-11-01T16:20:30+03:00", // время отправки в формате RFC3339
        "title": "Привет!",
        "content": "Чуть-чуть тестируемся :)",
        "channel": "telegram", // telegram или ntfy
        "receiver": "1234567890" // ID пользователя Telegram или топик ntfy
    }
    ```

- Для просмотра уведомления (и его статуса) отправьте `GET` запрос на `/api/v1/notify/{id}`

- Для отмены уведомления отправьте `DELETE` запрос на `/api/v1/notify/{id}`
