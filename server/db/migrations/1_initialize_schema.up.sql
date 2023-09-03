CREATE TABLE IF NOT EXISTS Articles (
    id INTEGER NOT NULL PRIMARY KEY,
    title TEXT,
    content TEXT,
    link TEXT,
    author TEXT,
    published_on TEXT,
    collected_on TEXT,
    is_read INTEGER,
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
    visibility INTEGER
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER NOT NULL PRIMARY KEY,
    title TEXT,
    visibility INTEGER
);
