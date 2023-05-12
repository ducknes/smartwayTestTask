INSERT INTO airline(iata,name)
VALUES('SU','Аэрофлот'),
('S7','S7'),
('KV','КрасАвиа'),
('U6','Уральские авиалинии'),
('UT','Ютэйр'),
('FZ','Flydubai'),
('JB','JetBlue'),
('SJ','SuperJet'),
('WZ','Wizz Air'),
('N4','Nordwind Airlines'),
('5N','SmartAvia');

INSERT INTO provider(provider_Id,name)
VALUES('AA','AmericanAir'),
('IF','InternationFlights'),
('RS','RedStar');

INSERT INTO schema(name)
VALUES('Основная'),
('Тестовая');

INSERT INTO account(schema_Id, name)
VALUES(2, 'Демо'),
(2, 'Разработка'),
(1, 'Продажи');

INSERT INTO airline_provider(airline_Id,provider_Id)
VALUES(6,1),
(7,1),
(8,1),

(1,2),
(2,2),
(6,2),
(10,2),
(7,2),
(9,2),

(1,3),
(2,3),
(3,3),
(4,3),
(5,3),
(10,3),
(11,3);

INSERT INTO schema_provider(schema_Id,provider_Id)
VALUES(1,1),
(1,2),
(1,3),
(2,2),
(2,3);
