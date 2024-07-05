CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    nickname TEXT NOT NULL,
    categorie TEXT NOT NULL,
    likedBy TEXT,
    content TEXT NOT NULL,
    img TEXT,
    nbrLike INTEGER,
    nbrDislike INTEGER,
    createdAt DATETIME NOT NULL
);