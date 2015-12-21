package rupert

var Schema = `
	CREATE TABLE forum
(
    forum_id INTEGER PRIMARY KEY NOT NULL,
    category_id INTEGER NOT NULL,
    name VARCHAR(256),
    topics INTEGER DEFAULT 0,
    posts INTEGER DEFAULT 0,
    order_idx INTEGER DEFAULT 100,
    last_thread_id INTEGER,
    updated_on TIMESTAMP WITH TIME ZONE DEFAULT now()
);
CREATE TABLE forum_category
(
    category_id INTEGER PRIMARY KEY NOT NULL,
    order_idx INTEGER DEFAULT 100 NOT NULL,
    name VARCHAR(256) NOT NULL
);
CREATE TABLE forum_comment
(
    comment_id INTEGER PRIMARY KEY NOT NULL,
    thread_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_on TIMESTAMP WITH TIME ZONE DEFAULT now(),
    author_id INTEGER NOT NULL
);
CREATE TABLE forum_thread
(
    thread_id INTEGER PRIMARY KEY NOT NULL,
    forum_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    replies INTEGER DEFAULT 0,
    views INTEGER DEFAULT 0,
    sticky BOOLEAN DEFAULT false,
    last_comment_id INTEGER,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_on TIMESTAMP WITH TIME ZONE DEFAULT now(),
    author_id INTEGER NOT NULL
);
CREATE TABLE users
(
    user_id INTEGER PRIMARY KEY NOT NULL,
    username VARCHAR NOT NULL,
    hash VARCHAR NOT NULL,
    salt VARCHAR NOT NULL,
    enabled BOOLEAN DEFAULT true,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_on TIMESTAMP WITH TIME ZONE DEFAULT now()
);
ALTER TABLE forum ADD FOREIGN KEY (category_id) REFERENCES forum_category (category_id);
CREATE UNIQUE INDEX forum_name_uindex ON forum (name);
CREATE UNIQUE INDEX forum_last_thread_id_uindex ON forum (last_thread_id);
CREATE UNIQUE INDEX forum_category_name_uindex ON forum_category (name);
ALTER TABLE forum_comment ADD FOREIGN KEY (author_id) REFERENCES users (user_id);
ALTER TABLE forum_thread ADD FOREIGN KEY (forum_id) REFERENCES forum (forum_id);
ALTER TABLE forum_thread ADD FOREIGN KEY (last_comment_id) REFERENCES forum_comment (comment_id);
CREATE UNIQUE INDEX forum_thread_title_uindex ON forum_thread (title);
CREATE UNIQUE INDEX user_username_uindex ON users (username);`
