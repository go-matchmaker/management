SET TIME ZONE 'Europe/Istanbul';

-- Create the users table
CREATE TABLE departments (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    role VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(20),
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    department_id VARCHAR(255) NOT NULL REFERENCES departments(id)
);

-- Add a constraint to check the role values
ALTER TABLE users
ADD CONSTRAINT role_check
CHECK (role IN ('super_admin', 'admin', 'employee', 'user'));


-- Create the user_attributes table
CREATE TABLE attributes (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    attribute VARCHAR(255) NOT NULL,
    view BOOLEAN DEFAULT FALSE,
    search BOOLEAN DEFAULT FALSE,
    detail BOOLEAN DEFAULT FALSE,
    add BOOLEAN DEFAULT FALSE,
    update BOOLEAN DEFAULT FALSE,
    delete BOOLEAN DEFAULT FALSE,
    export BOOLEAN DEFAULT FALSE,
    import BOOLEAN DEFAULT FALSE,
    can_see_price BOOLEAN DEFAULT FALSE
);

INSERT INTO departments (id, name)
VALUES (1, 'Super Admin');

INSERT INTO users (id,role, name, surname, email, password,department_id)
VALUES ('202406171','super_admin', 'Super', 'Admin', 'superadmin@superadmin.com', '$2a$10$kiiG/e86xezIkgp5QifAFu6ipYX8J5RPQvtP9bcA7xyu8wKvKPaMW',1);