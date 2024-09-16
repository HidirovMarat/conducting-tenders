-- init.sql

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

CREATE TABLE organization (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);

-- Вставка данных в employee
INSERT INTO employee (username, first_name, last_name) VALUES
('jdoe', 'John', 'Doe'),
('asmith', 'Alice', 'Smith'),
('bbrown', 'Bob', 'Brown');

-- Вставка данных в organization
INSERT INTO organization (name, description, type) VALUES
('Tech Corp', 'A technology company', 'LLC'),
('Retail Inc', 'A retail company', 'JSC'),
('Consulting LLC', 'A consulting company', 'IE');

-- Вставка данных в organization_responsible
INSERT INTO organization_responsible (organization_id, user_id)
SELECT o.id, e.id
FROM organization o, employee e
WHERE o.name = 'Tech Corp' AND e.username = 'jdoe';

INSERT INTO organization_responsible (organization_id, user_id)
SELECT o.id, e.id
FROM organization o, employee e
WHERE o.name = 'Retail Inc' AND e.username = 'asmith';

INSERT INTO organization_responsible (organization_id, user_id)
SELECT o.id, e.id
FROM organization o, employee e
WHERE o.name = 'Consulting LLC' AND e.username = 'bbrown';
