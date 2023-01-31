CREATE TABLE users(
    id UUID,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE roles(
    name VARCHAR(255),
    PRIMARY KEY (name)
);

CREATE TABLE scopes(
    name VARCHAR(255),
    PRIMARY KEY (name)
);

CREATE TABLE role_scopes(
    role_name VARCHAR(255),
    scope_name VARCHAR(255),
    PRIMARY KEY (role_name, scope_name),
    CONSTRAINT role_name_fk FOREIGN KEY (role_name) REFERENCES roles(name),
    CONSTRAINT scope_name_fk FOREIGN KEY (scope_name) REFERENCES scopes(name)
);

CREATE TABLE user_roles(
    user_id UUID NOT NULL,
    role_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id, role_name),
    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT role_name_fk FOREIGN KEY (role_name) REFERENCES roles(name)
);