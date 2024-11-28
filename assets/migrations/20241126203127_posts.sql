-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  createdAt DATETIME NOT NULL
);
-- +goose StatementEnd

INSERT INTO posts (title, content, createdAt) VALUES 
('Introduction to SQL', 'Learn the basics of SQL and database design.', '2024-11-26 09:00:00'),
('Understanding Datatypes', 'An in-depth look at SQL data types and their uses.', '2024-11-26 10:00:00'),
('Advanced Query Techniques', 'Explore advanced querying with JOINs and subqueries.', '2024-11-26 11:00:00'),
('Optimizing Queries', 'Tips and tricks for improving query performance.', '2024-11-26 12:00:00'),
('Database Indexing', 'Why and how to use indexes to speed up queries.', '2024-11-26 13:00:00');

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
