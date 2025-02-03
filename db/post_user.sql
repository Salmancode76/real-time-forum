-- User Table
CREATE TABLE User (
    UserID     INTEGER PRIMARY KEY AUTOINCREMENT,
    Username   TEXT NOT NULL UNIQUE,  
    Age        INTEGER NOT NULL,
    Gender     TEXT NOT NULL,
    First_Name TEXT NOT NULL,
    Last_Name  TEXT NOT NULL,
    Email      TEXT NOT NULL UNIQUE,
    Password   TEXT NOT NULL
);

-- Post Table
CREATE TABLE post (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT NOT NULL, 
    body        TEXT NOT NULL, 
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id     INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(UserID) ON DELETE CASCADE
);

-- Comments Table
CREATE TABLE comments (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id     INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,
    comment     TEXT NOT NULL, 
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES User(UserID) ON DELETE CASCADE
);

-- Category Table
CREATE TABLE category (
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- Post-Category Relationship Table (Many-to-Many)
CREATE TABLE post_category (
    post_id     INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);

-- Comment Likes Table
CREATE TABLE comment_likes (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    comment_id  INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,
    is_like     BOOLEAN NOT NULL DEFAULT 0,  
    FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES User(UserID) ON DELETE CASCADE
);

-- Post Likes Table
CREATE TABLE post_likes (
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id   INTEGER NOT NULL,
    user_id   INTEGER NOT NULL,
    is_like   BOOLEAN NOT NULL DEFAULT 0,  
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES User(UserID) ON DELETE CASCADE
);
