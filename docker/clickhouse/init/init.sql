CREATE TABLE histories (
    id UInt32,
    project_id UInt32,
    name String,
    description String,
    priority UInt32,
    removed Bool,
    event_time DateTime,
    INDEX id_index (id) TYPE minmax GRANULARITY 1,
    INDEX project_id_index (project_id) TYPE minmax GRANULARITY 1,
    INDEX name_index (name) TYPE minmax GRANULARITY 1
) ENGINE = MergeTree
PRIMARY KEY tuple();

