<div align="center">
  <img width="325" height="281" src="https://github.com/sxexesx/learn-backend/blob/main/misc/gopher.png">
  <h3>Как стать настоящим гофером</h3>
</div>

Здесь я собираю ответы на вопросы с собеседований на вакансию backend-инженер/разработчик Golang, а также 
краткие конспекты по разным темам в области IT. Некоторые темы разобраны с примерами и небольшими проектами  

---

## Содержание 

1. [Общие вопросы](https://github.com/sxexesx/learn-backend/blob/main/common/about.md)
   - Принципы ООП
   - Подходы в программировании (SOLID, KISS)
   - Требования к транзакционным системам (ACID)
   - [В чем отличие протоколов TCP и UDP? В каком случае UDP предпочтительнее?](https://github.com/sxexesx/learn-backend/blob/main/common/about.md#%D0%B2-%D1%87%D0%B5%D0%BC-%D0%BE%D1%82%D0%BB%D0%B8%D1%87%D0%B8%D0%B5-%D0%BF%D1%80%D0%BE%D1%82%D0%BE%D0%BA%D0%BE%D0%BB%D0%BE%D0%B2-tcp-%D0%B8-udp-%D0%B2-%D0%BA%D0%B0%D0%BA%D0%BE%D0%BC-%D1%81%D0%BB%D1%83%D1%87%D0%B0%D0%B5-udp-%D0%BF%D1%80%D0%B5%D0%B4%D0%BF%D0%BE%D1%87%D1%82%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%B5%D0%B5)
   - [Что такое NAT?](https://github.com/sxexesx/learn-backend/blob/main/common/about.md#%D0%B2-%D1%87%D0%B5%D0%BC-%D0%BE%D1%82%D0%BB%D0%B8%D1%87%D0%B8%D0%B5-%D0%BF%D1%80%D0%BE%D1%82%D0%BE%D0%BA%D0%BE%D0%BB%D0%BE%D0%B2-tcp-%D0%B8-udp-%D0%B2-%D0%BA%D0%B0%D0%BA%D0%BE%D0%BC-%D1%81%D0%BB%D1%83%D1%87%D0%B0%D0%B5-udp-%D0%BF%D1%80%D0%B5%D0%B4%D0%BF%D0%BE%D1%87%D1%82%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%B5%D0%B5)
   - [Что такое HTTP и HTTPS, в чем их отличия?](https://github.com/sxexesx/learn-backend/blob/main/common/about.md#%D0%B2-%D1%87%D0%B5%D0%BC-%D0%BE%D1%82%D0%BB%D0%B8%D1%87%D0%B8%D0%B5-%D0%BF%D1%80%D0%BE%D1%82%D0%BE%D0%BA%D0%BE%D0%BB%D0%BE%D0%B2-tcp-%D0%B8-udp-%D0%B2-%D0%BA%D0%B0%D0%BA%D0%BE%D0%BC-%D1%81%D0%BB%D1%83%D1%87%D0%B0%D0%B5-udp-%D0%BF%D1%80%D0%B5%D0%B4%D0%BF%D0%BE%D1%87%D1%82%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%B5%D0%B5)
   - [Что такое SSL и TLS, есть ли между ними отличия?](https://github.com/sxexesx/learn-backend/blob/main/common/about.md#%D0%B2-%D1%87%D0%B5%D0%BC-%D0%BE%D1%82%D0%BB%D0%B8%D1%87%D0%B8%D0%B5-%D0%BF%D1%80%D0%BE%D1%82%D0%BE%D0%BA%D0%BE%D0%BB%D0%BE%D0%B2-tcp-%D0%B8-udp-%D0%B2-%D0%BA%D0%B0%D0%BA%D0%BE%D0%BC-%D1%81%D0%BB%D1%83%D1%87%D0%B0%D0%B5-udp-%D0%BF%D1%80%D0%B5%D0%B4%D0%BF%D0%BE%D1%87%D1%82%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%B5%D0%B5)
   - [Что такое cors?](https://github.com/sxexesx/learn-backend/blob/main/common/about.md#%D0%B2-%D1%87%D0%B5%D0%BC-%D0%BE%D1%82%D0%BB%D0%B8%D1%87%D0%B8%D0%B5-%D0%BF%D1%80%D0%BE%D1%82%D0%BE%D0%BA%D0%BE%D0%BB%D0%BE%D0%B2-tcp-%D0%B8-udp-%D0%B2-%D0%BA%D0%B0%D0%BA%D0%BE%D0%BC-%D1%81%D0%BB%D1%83%D1%87%D0%B0%D0%B5-udp-%D0%BF%D1%80%D0%B5%D0%B4%D0%BF%D0%BE%D1%87%D1%82%D0%B8%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%B5%D0%B5)
   - [Типы HTTP-запросов](https://github.com/sxexesx/learn-backend/blob/main/common/about.md#типы-http-запросов)

2. [Компьютерные сети](https://github.com/sxexesx/learn-backend/blob/main/networking/about.md)
   
3. [Паттерны проектирования](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md)
   - [Порождающие](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md#%D0%BF%D0%BE%D1%80%D0%BE%D0%B6%D0%B4%D0%B0%D1%8E%D1%89%D0%B8%D0%B5--creational-patterns)
   - [Структурные](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md#%D1%81%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80%D0%BD%D1%8B%D0%B5--structural-patterns)
   - [Поведенческие](https://github.com/sxexesx/learn-backend/blob/main/patterns/about.md#%D0%BF%D0%BE%D0%B2%D0%B5%D0%B4%D0%B5%D0%BD%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B5--behavioral-patterns)
  
4. [Микросервисная архитектура](https://github.com/sxexesx/learn-backend/blob/main/microsvc/common.md)
   - [Что такое микросервисы? Плюсы и минусы микросервисной архитектуры](https://github.com/sxexesx/learn-backend/blob/main/microsvc/about.md)
   - [Общие понятия микросервисной архитектуры. Принципы декомпозиции](https://github.com/sxexesx/learn-backend/blob/main/microsvc/decomposition.md) 
   - [Паттерны микросервисной архитектуры, которые обязательно знать](https://github.com/sxexesx/learn-backend/blob/main/microsvc/most_known.md)
