-- +goose Up
INSERT INTO provider(provider_Id,name)
VALUES('AA','AmericanAir'),
('IF','InternationFlights'),
('RS','RedStar');