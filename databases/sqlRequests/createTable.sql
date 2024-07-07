CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    nickname TEXT NOT NULL,
    categorie TEXT NOT NULL,
    likedBy TEXT NOT NULL,
    dislikedBy TEXT NOT NULL,
    content TEXT NOT NULL,
    img TEXT NOT NULL,
    nbrLike INTEGER,
    nbrDislike INTEGER,
    createdAt DATETIME NOT NULL
);