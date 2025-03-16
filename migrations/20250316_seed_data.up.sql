-- Seed data for countries
INSERT INTO countries (name, alpha2_code, alpha3_code, numeric_code)
VALUES
    ('United States', 'US', 'USA', '840'),
    ('Canada', 'CA', 'CAN', '124'),
    ('United Kingdom', 'GB', 'GBR', '826'),
    ('Australia', 'AU', 'AUS', '036'),
    ('Germany', 'DE', 'DEU', '276'),
    ('France', 'FR', 'FRA', '250'),
    ('Italy', 'IT', 'ITA', '380'),
    ('Spain', 'ES', 'ESP', '724'),
    ('Netherlands', 'NL', 'NLD', '528'),
    ('Belgium', 'BE', 'BEL', '056'),
    ('Switzerland', 'CH', 'CHE', '756'),
    ('Austria', 'AT', 'AUT', '040'),
    ('Sweden', 'SE', 'SWE', '752'),
    ('Norway', 'NO', 'NOR', '578'),
    ('Denmark', 'DK', 'DNK', '208'),
    ('Finland', 'FI', 'FIN', '246'),
    ('Ireland', 'IE', 'IRL', '372'),
    ('Portugal', 'PT', 'PRT', '620'),
    ('Greece', 'GR', 'GRC', '300'),
    ('Japan', 'JP', 'JPN', '392');

-- Seed data for cities
INSERT INTO cities (name, country_id)
VALUES
    ('New York', (SELECT id FROM countries WHERE alpha2_code = 'US')),
    ('Toronto', (SELECT id FROM countries WHERE alpha2_code = 'CA')),
    ('London', (SELECT id FROM countries WHERE alpha2_code = 'GB')),
    ('Sydney', (SELECT id FROM countries WHERE alpha2_code = 'AU')),
    ('Berlin', (SELECT id FROM countries WHERE alpha2_code = 'DE')),
    ('Paris', (SELECT id FROM countries WHERE alpha2_code = 'FR')),
    ('Rome', (SELECT id FROM countries WHERE alpha2_code = 'IT')),
    ('Madrid', (SELECT id FROM countries WHERE alpha2_code = 'ES')),
    ('Amsterdam', (SELECT id FROM countries WHERE alpha2_code = 'NL')),
    ('Brussels', (SELECT id FROM countries WHERE alpha2_code = 'BE')),
    ('Zurich', (SELECT id FROM countries WHERE alpha2_code = 'CH')),
    ('Vienna', (SELECT id FROM countries WHERE alpha2_code = 'AT')),
    ('Stockholm', (SELECT id FROM countries WHERE alpha2_code = 'SE')),
    ('Oslo', (SELECT id FROM countries WHERE alpha2_code = 'NO')),
    ('Copenhagen', (SELECT id FROM countries WHERE alpha2_code = 'DK')),
    ('Helsinki', (SELECT id FROM countries WHERE alpha2_code = 'FI')),
    ('Dublin', (SELECT id FROM countries WHERE alpha2_code = 'IE')),
    ('Lisbon', (SELECT id FROM countries WHERE alpha2_code = 'PT')),
    ('Athens', (SELECT id FROM countries WHERE alpha2_code = 'GR')),
    ('Tokyo', (SELECT id FROM countries WHERE alpha2_code = 'JP'));

-- Seed data for addresses
INSERT INTO addresses (street, postal_code, city_id)
VALUES
    ('123 Main St', '10001', (SELECT id FROM cities WHERE name = 'New York')),
    ('456 Queen St', 'M5H 2N2', (SELECT id FROM cities WHERE name = 'Toronto')),
    ('789 King St', 'SW1A 1AA', (SELECT id FROM cities WHERE name = 'London')),
    ('101 George St', '2000', (SELECT id FROM cities WHERE name = 'Sydney')),
    ('202 Alexanderplatz', '10178', (SELECT id FROM cities WHERE name = 'Berlin')),
    ('303 Champs-Élysées', '75008', (SELECT id FROM cities WHERE name = 'Paris')),
    ('404 Via del Corso', '00186', (SELECT id FROM cities WHERE name = 'Rome')),
    ('505 Gran Via', '28013', (SELECT id FROM cities WHERE name = 'Madrid')),
    ('606 Dam Square', '1012 NP', (SELECT id FROM cities WHERE name = 'Amsterdam')),
    ('707 Grand Place', '1000', (SELECT id FROM cities WHERE name = 'Brussels')),
    ('808 Bahnhofstrasse', '8001', (SELECT id FROM cities WHERE name = 'Zurich')),
    ('909 Kärntner Strasse', '1010', (SELECT id FROM cities WHERE name = 'Vienna')),
    ('1010 Drottninggatan', '111 60', (SELECT id FROM cities WHERE name = 'Stockholm')),
    ('1111 Karl Johans gate', '0154', (SELECT id FROM cities WHERE name = 'Oslo')),
    ('1212 Strøget', '1552', (SELECT id FROM cities WHERE name = 'Copenhagen')),
    ('1313 Aleksanterinkatu', '00100', (SELECT id FROM cities WHERE name = 'Helsinki')),
    ('1414 O''Connell Street', 'D01', (SELECT id FROM cities WHERE name = 'Dublin')),
    ('1515 Avenida da Liberdade', '1250-096', (SELECT id FROM cities WHERE name = 'Lisbon')),
    ('1616 Ermou Street', '105 63', (SELECT id FROM cities WHERE name = 'Athens')),
    ('1717 Shibuya Crossing', '150-0002', (SELECT id FROM cities WHERE name = 'Tokyo'));

-- Seed data for contacts
INSERT INTO contacts (id, first_name, last_name, address_id)
VALUES
    (uuid_generate_v4(), 'John', 'Doe', (SELECT id FROM addresses WHERE street = '123 Main St'));

-- Seed data for phones
INSERT INTO phones (contact_id, number, type)
VALUES
    ((SELECT id FROM contacts WHERE first_name = 'John' AND last_name = 'Doe'), '+12025550123', 'mobile');