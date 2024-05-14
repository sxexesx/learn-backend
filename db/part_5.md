# Common Table Expressions и подзапросы

## Подзапросы
Вложенный запрос (подзапрос) - когда один запрос вложенный в другой  
```
SELECT name, surname
FROM actor
WHERE actor_id  in (
    SELECT person_id FROM persons WHERE person_name = 'john'
) 
```
Бывают двух типов:
1. Независимые
```
SELECT 
    StockItemID,
    (SELECT MAX(UnitPrice)
    FROM StockItems) AS MaxPrice
FROM StockItems;
```
2. Зависимые (correlated) - когда подзапрос ссылается на внешнюю таблицу

Важное:  
+ Подзапрос может возвращать одно значение, может несколько, либо таблицу - (derived tables).  
+ Когда данных не много, то проще сделать JOIN.  

Операторы используемые с подзапросами:
* IN или NOT IN
* EXISTS или NOT EXISTS
Быстрее, чем IN
* ANY (SOME) - хотя бы одно значение
= SOME -> IN
```
WHERE UnitPrice <= SOME (SELECT UnitPrice FROM StockItems)
```
* ALL 

## Обобщенные табличные выражения (CTE)
По сути это именованный подзапрос:
```
WITH [название] (поля) 
AS ( подзапрос )
SELECT .....
```

Важное:  
* СTE хранться в памяти
* Можно использовать несколько CTE в одном запросе, но CTE используется только один раз.  
* Можно удалять из CTE 
* Можно делать JOIN нескольких CTE

## Оконные функции

Необходимы для выполнения следующих целей:
- нарастающий итог
- в запросе нужны данные из предыдущей строки
- нужно совместить и сами данные, и агрегаты
- выбрать топ 3 из группы: топ 3 самых дорогих товара в каждой группе
- разделить набор на N групп

```
Function() OVER (
    окно по которому идет расчет
)
```

```
Func() OVER (
    PARTITION BY [список полей]
    ORDER BY
    ROWS/RANGE
)
```

пустое `OVER()` - по всем строкам   

Какие функции бывают:
1. Ранжирующие (`ROW_NUMBER()`, `RANK()`, `DENSE_RANK()`, `NTILE(кол-во групп)`)
2. Функции смещения (`LAG(поле, [смещение], default)`, `LEAD(поле, [смещение], default)`,
`FIRST_VALUE(поле, [смещение], default)`, `LAST_VALUE(поле, [смещение], default)`)
3. Распределения (`CUME_DIST()`, `Percentile_COUNT()`, `Percentile_DISC()`, `Percent_RANK()`)
4. Агрегатные (`SUM()`, `COUNT()`, `COUNT_BIG()`, `AVG()`, `MIN()`, `MAX()`)

