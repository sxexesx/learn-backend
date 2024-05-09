## Основные операции по работе с SQL

_**DDL** (data defenition language)_ - язык описания баз данных, чтобы определить, что будет лежать в базе данных
(CREATE, ALTER, DROP)  
_**DML** (data manipulation language)_ - язык оперирования данными (SELECT, INSERT, UPDATE, DELETE)
_**DCL** (data control language)_ - язык управления доступом (GRANT, REVOKE, DENY)
_**TCL** (transaction control language)_ - язык управления транзакциями (COMMIT, ROLLBACK, SAVEPOINT)

**Операции: **

1. Выборка (SELECT)
2. Проекция (PROJECT)
3. Соединения (JOIN)
4. Деление (DIVIDE BY)

5. Объединение (UNION)
6. Пересечение (INTERSECT)
7. Вычитание (DIFFERENCE, MINUS)
8. Декартово произведение (CARTESIAN)

## Операция SELECT

PostreSQL syntax:

```
[ WITH [ RECURSIVE ] with_query [, ...] ]
SELECT [ ALL | DISTINCT [ ON ( expression [, ...] ) ] ]
    [ * | expression [ [ AS ] output_name ] [, ...] ]
    [ FROM from_item [, ...] ]
    [ WHERE condition ]
    [ GROUP BY [ ALL | DISTINCT ] grouping_element [, ...] ]
    [ HAVING condition ]
    [ WINDOW window_name AS ( window_definition ) [, ...] ]
    [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] select ]
    [ ORDER BY expression [ ASC | DESC | USING operator ] [ NULLS { FIRST | LAST } ] [, ...] ]
    [ LIMIT { count | ALL } ]
    [ OFFSET start [ ROW | ROWS ] ]
    [ FETCH { FIRST | NEXT } [ count ] { ROW | ROWS } { ONLY | WITH TIES } ]
    [ FOR { UPDATE | NO KEY UPDATE | SHARE | KEY SHARE } [ OF table_name [, ...] ] [ NOWAIT | SKIP LOCKED ] [...] ]

where from_item can be one of:

    [ ONLY ] table_name [ * ] [ [ AS ] alias [ ( column_alias [, ...] ) ] ]
                [ TABLESAMPLE sampling_method ( argument [, ...] ) [ REPEATABLE ( seed ) ] ]
    [ LATERAL ] ( select ) [ [ AS ] alias [ ( column_alias [, ...] ) ] ]
    with_query_name [ [ AS ] alias [ ( column_alias [, ...] ) ] ]
    [ LATERAL ] function_name ( [ argument [, ...] ] )
                [ WITH ORDINALITY ] [ [ AS ] alias [ ( column_alias [, ...] ) ] ]
    [ LATERAL ] function_name ( [ argument [, ...] ] ) [ AS ] alias ( column_definition [, ...] )
    [ LATERAL ] function_name ( [ argument [, ...] ] ) AS ( column_definition [, ...] )
    [ LATERAL ] ROWS FROM( function_name ( [ argument [, ...] ] ) [ AS ( column_definition [, ...] ) ] [, ...] )
                [ WITH ORDINALITY ] [ [ AS ] alias [ ( column_alias [, ...] ) ] ]
    from_item join_type from_item { ON join_condition | USING ( join_column [, ...] ) [ AS join_using_alias ] }
    from_item NATURAL join_type from_item
    from_item CROSS JOIN from_item

and grouping_element can be one of:

    ( )
    expression
    ( expression [, ...] )
    ROLLUP ( { expression | ( expression [, ...] ) } [, ...] )
    CUBE ( { expression | ( expression [, ...] ) } [, ...] )
    GROUPING SETS ( grouping_element [, ...] )

and with_query is:

    with_query_name [ ( column_name [, ...] ) ] AS [ [ NOT ] MATERIALIZED ] ( select | values | insert | update | delete )
        [ SEARCH { BREADTH | DEPTH } FIRST BY column_name [, ...] SET search_seq_col_name ]
        [ CYCLE column_name [, ...] SET cycle_mark_col_name [ TO cycle_mark_value DEFAULT cycle_mark_default ] USING cycle_path_col_name ]

TABLE [ ONLY ] table_name [ * ]
```

Примеры: [селекты](part_3.sql)

### Фильтрация WHERE

Операции с приоритетами:

| Level | Оператор                                                                           |
|-------|------------------------------------------------------------------------------------|
| 1     | `~` (побитовое НЕ)                                                                 |
| 2     | `*` (умножение), `/` (деление), `%` (остаток)                                      |
| 3     | `+`, `-`, `&` (побитовое И), `^` (побитовое исключающее ИЛИ), `\|` (побитовое ИЛИ) |
| 4     | операторы сравнения - `=, >, <, >=, <=, !=, !>, !<`                                |
| 5     | `NOT`                                                                              |
| 6     | `AND`                                                                              |
| 7     | `ALL, ANY, BETWEEN, IN, LIKE, OR, SOME`                                            |
| 8     | `=` (присваивание )                                                                |

### BETWEEN

`WHERE x BETWEEN x1 AND x2` -> `x >= x1 AND x <= x2`

### LIKE

`%` - любая строка, содержащая 0 и больше символов  
`_` - любой одиночный символ  
`[]` - любой одиночный символ содержащийся в наборе `[a-f]` или наборе `[abcdef]`  
`[^]` - любой одиночный символ НЕ содержащийся в наборе  

### NULL

1. NULL не имеет типа
2. NULL может записываться в поля любого типа
3. Любая операция с NULL дает в результате NULL
4. Любое сравнение с NULL дает в результате UNKNOWN