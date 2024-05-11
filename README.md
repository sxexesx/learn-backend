<div align="center">
  <img width="325" height="281" src="misc/gopher.png">
  <h3>Как стать настоящим гофером</h3>
</div>

Здесь я собираю ответы на вопросы с собеседований на вакансию backend-инженер/разработчик Golang, а также
краткие конспекты по разным темам в области IT. Некоторые темы разобраны с примерами и небольшими проектами

---

## Содержание

1. [Компьютерные сети](networking/part_1.md)
    - [Стандарты, классификация, топология и модели](networking/part_1.md)
    - [Физический и канальный уровни](networking/part_2.md):
        - [Ethernet](networking/part_3.md) _(коммутируемый, некоммутируемый, MAC адреса, кадры, VLAN)_
        - [Wi-Fi](networking/part_4.md) _(режимы доступа, формат кадра)_
    - [Сетевой уровень](networking/part_5.md) _(IP адреса, передача пакетов, протоколы IP, ICMP, ARP)_
    - [Транспортный уровень](networking/part_7.md) _(протоколы UDP, TCP, подтверждение доставки, метод соединения,
      управление перегрузкой, интерфейс сокетов, NAT)_
    - [Прикладной уровень](networking/part_8.md) _(протоколы DNS, FTP, SMTP, POP3, IMAP)_
    - [HTTP, HTTPs](networking/part_9.md)
    - [Сетевое взаимодействие приложений](networking/part_10.md)
        - [SSL/TLS](networking/part_11.md)
        - [REST](networking/part_10.md#rest)
        - [CORs](networking/part_10.md#что-такое-cors)
    - [Вопросы на интервью](networking/part_x.md)

2. [Программирование. Общие вопросы]()
    - Принципы ООП
    - Как реализованы принципы ООП в Go
    - Подходы в программировании (SOLID, KISS)
    - Требования к транзакционным системам (ACID)

3. [Паттерны проектирования](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md)
    - [Порождающие](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md#%D0%BF%D0%BE%D1%80%D0%BE%D0%B6%D0%B4%D0%B0%D1%8E%D1%89%D0%B8%D0%B5--creational-patterns)
    - [Структурные](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md#%D1%81%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80%D0%BD%D1%8B%D0%B5--structural-patterns)
    - [Поведенческие](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md#%D0%BF%D0%BE%D0%B2%D0%B5%D0%B4%D0%B5%D0%BD%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B5--behavioral-patterns)

4. Язык GO

5. SQL
    - [Основнаые понятия СУБД (определение, классификация, основные компоненты СУБД, теорема CAP)](db/part_1.md)
    - [Основные понятия реляционной модели (ключи, типы данных, миграции)](db/part_2.md)
    - [DML: select, insert, update, delete](db/part_3.md)
    - [JOIN, UNION](db/part_4.md)

6. Kafka

7. [Микросервисная архитектура](microsvc/common.md)
    - [Что такое микросервисы? Плюсы и минусы микросервисной архитектуры](microsvc/about.md)
    - [Общие понятия микросервисной архитектуры. Принципы декомпозиции](microsvc/decomposition.md)
    - [Паттерны микросервисной архитектуры, которые обязательно знать](microsvc/most_known.md)

10. [Благодарности, курсы, полезное и прочее](misc/acknowledgements.md)
