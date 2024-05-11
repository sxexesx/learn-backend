### JOIN
Аналог горизонтального соединения

<div>
  <img width="500" height="500" src="src/img09.png" alt="">
</div>

`CROSS JOIN` - декартово произведение(каждой строчке одной таблицы сопоставляем каждую строчку другой)  
Пример:
`FROM t1 CROSS JOIN t2`  
Старый синтаксис (SQL-89): `FROM t1, t2`

<div>
  <img width="300" height="320" src="src/img05.png" alt="">
</div>

`INNER JOIN` - декартово произведение + фильтрация
Пример:
`FROM t1 [INNER] JOIN t2 on t1.id = t2.id`

<div>
  <img width="300" height="320" src="src/img06.png" alt="">
</div>

`LEFT/RIGHT JOIN` - INNER JOIN + внешние строки
Пример:
`FROM t1 LEFT [OUTER] JOIN t2 on t1.id = t2.id`

<div>
  <img width="300" height="320" src="src/img07.png" alt="">
</div>

`FULL JOIN` - LEFT + RIGHT JOIN
Пример:
`FROM t1 FULL [OUTER] JOIN t2 on t1.id = t2.id`

Рекомендуется делать индекс по тем колонкам, по которым делаем JOIN.

### UNION
Аналог вертикального соединения.  
`UNION ALL` дублирует одинаковые строки.

`UNION` - это `UNION ALL` + `DISTINCT`

<div>
  <img width="500" height="240" src="src/img08.png" alt="">
</div>

