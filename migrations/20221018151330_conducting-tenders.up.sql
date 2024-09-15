CREATE TYPE service_type AS ENUM (
    'Construction',
    'Delivery',
    'Manufacture'
);

CREATE TYPE status_t AS ENUM (
    'Created',
    'Published',
    'Closed'
);

CREATE TABLE tenders (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    type service_type,
    status status_t,
    organization_id UUID,
    version INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tag UUID
);

CREATE TYPE status_b AS ENUM (
    'Created',
    'Published',
    'Canceled'
);

CREATE TYPE author_type AS ENUM (
    'Organization',
    'User'
);

CREATE TABLE bids (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status status_b,
    tender_id UUID,
    author_t author_type,
    author_id UUID,
    version INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tag UUID NOT NULL
);
