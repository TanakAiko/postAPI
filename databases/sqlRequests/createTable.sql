CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    categorie TEXT NOT NULL,
    content TEXT NOT NULL,
    img TEXT,
    nbrLike INTEGER,
    nbrDislike INTEGER,
    createdAt DATETIME NOT NULL
);