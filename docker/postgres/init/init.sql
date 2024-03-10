CREATE TABLE projects
(
    id         SERIAL       NOT NULL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX projects_id_index ON projects (id);

CREATE TABLE goods
(
    id          SERIAL       NOT NULL PRIMARY KEY,
    project_id  INT          NOT NULL REFERENCES projects (id),
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    priority    INT,
    removed     BOOLEAN   DEFAULT FALSE,
    created_at  TIMESTAMP DEFAULT now()
);

CREATE INDEX goods_id_index on goods (id);
CREATE INDEX goods_project_id_index on goods (project_id);
CREATE INDEX goods_name_index on goods (name);




