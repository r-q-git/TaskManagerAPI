CREATE TYPE task_status AS ENUM (
    'Completed',
    'Pending',
    'UnderProcess',
    'Skipped',
    'NotDone'
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status task_status NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);