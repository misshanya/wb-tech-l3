# URL Shortener

Сервис сокращения URL с аналитикой

## Установка и запуск

1. Склонируйте репозиторий

    ```shell
    git clone https://github.com/misshanya/wb-tech-l3
    cd wb-tech-l3/url-shortener
    ```

2. Скопируйте .env.example в .env и отредактируйте значения под себя

    ```shell
    cp .env.example .env
    ```

3. Запустите через Docker Compose

    ```shell
    docker compose up -d
    ```

## Использование

- Для сокращения ссылки отправьте `POST` запрос на `/api/v1/shorten` со следующим телом:

    ```json5
    {
      "url": "https://github.com/misshanya" // URL, который надо сократить
    }
    ```

- Для редиректа перейдите на `/api/v1/s/{short}`

- Для получения аналитики отправьте `GET` запрос на `/api/v1/analytics/{short}`

