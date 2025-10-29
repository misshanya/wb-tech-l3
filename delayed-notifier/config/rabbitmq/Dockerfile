FROM rabbitmq:4-management

RUN apt-get update && \
    apt-get install -y curl

RUN curl -L https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/v4.2.0-rc.1/rabbitmq_delayed_message_exchange-4.2.0-rc.1.ez \
    -o /opt/rabbitmq/plugins/rabbitmq_delayed_message_exchange-4.2.0-rc.1.ez

RUN rabbitmq-plugins enable rabbitmq_delayed_message_exchange