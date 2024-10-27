CREATE TABLE IF NOT EXISTS Entries (
    id INTEGER NOT NULL PRIMARY KEY,
    title TEXT,
    content TEXT,
    link TEXT,
    author TEXT,
    published_on TEXT,
    collected_on TEXT,
    is_read INTEGER,
    is_starred INTEGER,
    category TEXT,
    original_id TEXT,
    feed_id INTEGER
);

CREATE TABLE IF NOT EXISTS Feeds (
    id INTEGER NOT NULL PRIMARY KEY,
    title TEXT,
    feed_url TEXT,
    website_url TEXT,
    category_id INTEGER,
    is_muted INTEGER,
    visibility INTEGER
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER NOT NULL PRIMARY KEY,
    title TEXT,
    visibility INTEGER
);

INSERT INTO categories (id, title, visibility) VALUES (1, "Uncategorized", 1);
