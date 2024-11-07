CREATE TABLE
   users (
      user_id VARCHAR(36) NOT NULL PRIMARY KEY,
      username VARCHAR(100) NOT NULL,
      password VARCHAR(100) NOT NULL,
      full_name VARCHAR(100),
      email VARCHAR(100) NOT NULL,
      role VARCHAR(100) DEFAULT 'user',
      cuti_balance INT DEFAULT 0,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
   );