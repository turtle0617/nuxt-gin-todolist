CREATE TYPE todo_status AS ENUM ('idle', 'completed');

CREATE TABLE IF NOT EXISTS todos
(
    id serial NOT NULL,
    status todo_status NOT NULL DEFAULT 'idle',
    title character varying(500) NOT NULL,
    CONSTRAINT todoinfo_key PRIMARY KEY (id)
)