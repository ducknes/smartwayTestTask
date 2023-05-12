-- +goose Up
INSERT INTO schema_provider(schema_Id,provider_Id)
VALUES(1,1),
(1,2),
(1,3),
(2,2),
(2,3);