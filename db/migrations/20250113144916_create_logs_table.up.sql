CREATE TABLE logs (
      id SERIAL PRIMARY KEY,
      action VARCHAR(255) NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      user_id INT,
      details TEXT
);
