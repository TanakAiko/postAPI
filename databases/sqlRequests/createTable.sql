CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    categorie TEXT NOT NULL,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL
);