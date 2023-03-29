
-- to create the database : psql -h localhost -U myuser -d mydb -f BDD_generator.sql

-- Create the tables
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE groups (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  parent_group_id INTEGER,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (parent_group_id) REFERENCES groups (id)
);

CREATE TABLE auth_tokens (
  id SERIAL PRIMARY KEY,
  token VARCHAR(255) NOT NULL,
  expires_at TIMESTAMP NOT NULL
);

CREATE TABLE refresh_tokens (
  id SERIAL PRIMARY KEY,
  token VARCHAR(255) NOT NULL,
  expires_at TIMESTAMP NOT NULL
);

-- Create the relationship tables
CREATE TABLE user_roles (
  user_id INTEGER NOT NULL,
  role_id INTEGER NOT NULL,
  PRIMARY KEY (user_id, role_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (role_id) REFERENCES roles (id)
);

CREATE TABLE user_groups (
  user_id INTEGER NOT NULL,
  group_id INTEGER NOT NULL,
  PRIMARY KEY (user_id, group_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (group_id) REFERENCES groups (id)
);

-- Insert random data
INSERT INTO roles (name, description, created_at) VALUES
  ('Admin', 'Administrator role', NOW()),
  ('User', 'Regular user role', NOW());

INSERT INTO users (name, email, password, created_at) VALUES
  ('John Doe', 'john.doe@example.com', 'password1', NOW()),
  ('Jane Smith', 'jane.smith@example.com', 'password2', NOW());

INSERT INTO user_roles (user_id, role_id) VALUES
  (1, 1),
  (2, 2);

INSERT INTO groups (name, created_at) VALUES
  ('Group A', NOW()),
  ('Group B', NOW());

INSERT INTO user_groups (user_id, group_id) VALUES
  (1, 1),
  (2, 2);