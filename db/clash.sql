-- Supprime les tables existantes si elles existent
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS clash;
DROP TABLE IF EXISTS vote;
DROP TABLE IF EXISTS comments_clashs;

-- Table user
CREATE TABLE IF NOT EXISTS user (
  id INTEGER PRIMARY KEY,
  username TEXT,
  email TEXT,
  password TEXT
);

-- Table clash
CREATE TABLE IF NOT EXISTS clash (
  id INTEGER PRIMARY KEY,
  techno1 TEXT,
  techno2 TEXT,
  user_id INTEGER,
  created_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES user (id)
);

-- Table comment
CREATE TABLE IF NOT EXISTS comment (
  id INTEGER PRIMARY KEY,
  body TEXT,
  created_at TIMESTAMP,
  user_id INTEGER,
  clash_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES user (id),
  FOREIGN KEY (clash_id) REFERENCES clash (id)
);

-- Table vote
CREATE TABLE IF NOT EXISTS vote (
  id INTEGER PRIMARY KEY,
  user_id INTEGER,
  clash_id INTEGER,
  techno1_id INTEGER,
  techno2_id INTEGER,
  FOREIGN KEY (clash_id) REFERENCES clash (id)
  FOREIGN KEY (user_id) REFERENCES user (id)
);