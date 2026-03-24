<div align="center">
  <img width="325" height="281" src="misc/nerd.png">
  <h3>Как стать настоящим бэкэндером</h3>
</div>

Здесь я собираю ответы на вопросы с собеседований, а также краткие конспекты по разным темам в области IT.

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

2. ОС
    - [syscall](os/syscall.md#syscall) 


3. [Программирование. Общие вопросы](prog/part_1.md)
    - Принципы ООП
    - Как реализованы принципы ООП
    - [Подходы в программировании (SOLID, KISS)](prog/part_1.md#Подходы-в-программировании)
    - Метрики и оценка качества
        - [DORA метрики](prog/part_2.md#dora-метрики)
        - [SLI, SLO, SLA](prog/part_2.md#sla-slo-sli)


4. [git](git/part_1.md)


5. Kafka
    - [Кафка. Общие понятия](kafka/note_1.md)
        - [Почему Kafka?](kafka/note_1#почему-kafka)
        - [Топики и партиции](kafka/note_1#топики-и-партиции)
        - [Консьюмер группа](kafka/note_1#консьюмер-группа)
        - [Гарантирован ли порядок?](kafka/note_1.md#гарантирован-ли-порядок-сообщений-в-кафке)

    - [Гарантии доставки](kafka/note_2.md#гарантии-доставки)
        - [Transactional outbox](kafka/note_2#transactional-outbox)
        - [Dead Leter Queue](kafka/note_2#dead-letter-queue)


6. [Микросервисная архитектура](microsvc/common.md)
    - [Паттерны микросервисной архитектуры, которые обязательно знать](microsvc/most_known.md)


7. System design
    - [Что такое микросервисы? Плюсы и минусы микросервисной архитектуры](system-design/about.md)
    - [Общие понятия микросервисной архитектуры. Принципы декомпозиции](microsvc/decomposition.md)
    - [Что такое распределенная система. Общие понятия](system-design/note_1.md)
    
    - [Система апдейтов операционной системы мобильного устройства](system-design/note_4.md)
    - [Система статусов пользователя](system-design/note_5.md)
    - [Система email оповещений](system-design/note_6.md)
    - [Мэссенджер](system-design/note_7.md)

    - [Бонус](system-design/note_n.md)

9. Машинное обучение. Классические методы
    - [Что такое машинное обучение?](ml/_part_1.md)
    - [Линейная регрессия](ml/_part_2.md)
    - [Обобщающая способность](ml/_part_3.md)
    - [Методы преобразования](ml/_part_a.md)

10. [Благодарности, курсы, полезное и прочее](misc/acknowledgements.md)
