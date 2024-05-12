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
