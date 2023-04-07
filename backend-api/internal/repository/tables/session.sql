CREATE TABLE IF NOT EXISTS sessions(
    session_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user_id INTEGER,
    expiry TIMESTAMP NOT NULL,
    token TEXT NOT NULL,
    CONSTRAINT fk_users
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
);