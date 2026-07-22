# Kafka

## Для чего нужны очереди?

1. Помогают устранить дисбаланс в потоке входящих задач
2. Планирование исполнения
<img src="./src/img_01.png">
3. Основа для репликации сообщений
4. Создание отказоустойчивых систем
5. Коммуникация микросервисов
6. Event sourcing
7. Стримы

## Что такое очередь?

1. Среда коммуникации посредством сообщений.
2. Put/take
<img src="./src/img_02.png">
3. Pub/Sub
<img src="./src/img_03.png">
4. Request/Responce
5. Протоколы: AMPQ, MQTT, STOMP, NATS, ZeroMQ

## Какие бывают?
1. Managed / Cloud Services
    - Amazon SQS, EventBridge
    - Google Cloud Tasks
    - CloudAMQP
2. Message brokers
    - RabbitMQ
    - Apache Kafka
    - NATS JetStream
    - ActiveMQ
    - NSQ
3. Database-backed queues
    - PgQueue
    - Tarantool
    - Redis
4. Socket on steroids (lightweight or embedded)
    - ZeroMQ
    - NATS (core)
    - nanomsg

## Проблемы и алгоритмы

1. 