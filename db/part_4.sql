-- простая вставка
INSERT INTO warehouse.colors
    (colorId, colorName, LastEditedBy)
VALUES (NEXTVAL('Sequences.colors_id_seq'), 'Ohra', 1);

-- копирование через SELECT
SELECT first_name, last_name
INTO actor_copy
FROM actor
WHERE actor.actor_id < 10;

-- MSSQL
INSERT INTO category (category_id, name, last_update)
    OUTPUT category_id, name, last_update
    INTO category_copy(category_id, name, last_update)
VALUES (NEXTVAL(category_category_id_seq), 'trash movie', now()),
       (NEXTVAL(category_category_id_seq), 'unknown', now());

-- MSSQL
BEGIN transaction;
INSERT INTO category (category_id, name, last_update)
VALUES (NEXTVAL(category_category_id_seq), 'trash movie', now());
COMMIT;

INSERT INTO category_temp (category_id, name, last_update)
SELECT name, last_update
FROM category
WHERE category_id > 10;

-- простой UPDATE
UPDATE People
SET PhoneNumber = '799922234567',
    FaxNumber   = '799922236677'
WHERE PersonId = 3;

-- UPDATE с подзапросом
UPDATE People
SET FirstSale = (SELECT MIN(Invoice)
                 FROM Invoices AS I
                 WHERE People.PersonId = I.SalesPersonId);

-- UPDATE всей таблицы
UPDATE People
SET PersonSale  = NULL,
    PersonSale2 = NULL;