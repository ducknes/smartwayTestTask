-- +goose Up
INSERT INTO account(schema_Id, name)
VALUES(2, 'Демо'),
(2, 'Разработка'),
(1, 'Продажи');