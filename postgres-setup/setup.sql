-- Create the tables
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL
);

-- Création de la table Role
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

-- Création de la table Group
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    parent_group_id INT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (parent_group_id) REFERENCES groups(id)
);

-- Création de la table AuthToken
CREATE TABLE auth_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Création de la table RefreshToken
CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Création de la table UserRole
CREATE TABLE user_roles (
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- Création de la table UserGroup
CREATE TABLE user_groups (
    user_id INT NOT NULL,
    group_id INT NOT NULL,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- Insert sample data into the users table
INSERT INTO users (name, email, password, created_at) VALUES
('Alice', 'alice@example.com', 'alice_password', NOW()),
('Bob', 'bob@example.com', 'bob_password', NOW()),
('Carol', 'carol@example.com', 'carol_password', NOW());

-- Insert sample data into the roles table
INSERT INTO roles (name, description, created_at) VALUES
('Admin', 'Administrator with full access', NOW()),
('Editor', 'Can edit and manage content', NOW()),
('Viewer', 'Can view content only', NOW());

-- Insert sample data into the groups table
INSERT INTO groups (name, parent_group_id, created_at) VALUES
    ('Management', NULL, NOW()),
    ('Marketing', NULL, NOW()),
    ('Sales', NULL, NOW());

INSERT INTO user_roles (user_id, role_id) VALUES
(1, 1), -- Alice has the Admin role
(2, 2), -- Bob has the Editor role
(3, 3); -- Carol has the Viewer role

-- Add user-group relationships
INSERT INTO user_groups (user_id, group_id) VALUES
(1, 1), -- Alice is in the Management group
(2, 2), -- Bob is in the Marketing group
(3, 3); -- Carol is in the Sales group

INSERT INTO auth_tokens (user_id, token, expires_at) VALUES
(1, 'token1', '2023-04-30 00:00:00'),
(2, 'token2', '2023-04-30 00:00:00');

INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES
(1, 'refreshtoken1', '2023-06-30 00:00:00'),
(2, 'refreshtoken2', '2023-06-30 00:00:00');