-- нет источника // константы
SELECT 123 AS Col01,
       'A' AS Col02;

-- одна таблица
SELECT
    OrderLineID AS Order_Line_ID,
    StockItemID,
    Price = UnitPrice, -- синоним
    'Batch 1' AS BatchID, -- константа
    (Quantity * UnitPrice) AS Total -- арифметическое выражение
FROM Sales.OrderLines;

-- выборка всего
SELECT *
FROM Sales.OrderLines;

-- ORDER BY
SELECT * FROM Sales.OrderLines l
ORDER BY l.Description; -- asc (по умолчанию) - по возрастания, desc - по убыванию

-- Разбивка на страницы
SELECT * FROM Sales.OrderLines l
ORDER BY l.Description
OFFSET 20 ROWS -- указыаем сколько пропускаем строк
FETCH NEXT 50 ROWS ONLY;

-- DISTINCT
SELECT DISTINCT Contact
FROM Sales.OrderLines;

-- простое условие
SELECT *
FROM Sales.StockItems
WHERE Sales.StockItemName = 'Chocolate sharks 20g';

-- простой LIKE
SELECT *
FROM Sales.StockItems
WHERE Sales.StockItemName like 'Chocolate%';

-- работа с NULL
SELECT *
FROM Sales.StockItems
WHERE Sales.StockItemName IS NULL