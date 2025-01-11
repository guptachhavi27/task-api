-- Create the task_status enum type
DO $$ BEGIN
    CREATE TYPE task_status AS ENUM ('pending', 'in-progress', 'completed');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Create the tasks table
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status task_status DEFAULT 'pending'::task_status,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
