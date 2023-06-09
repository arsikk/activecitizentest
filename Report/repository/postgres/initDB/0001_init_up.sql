var schema = `
CREATE TABLE IF NOT EXISTS report (
    reportID serial,
    title text,
    description text,
    userID text,
    photourl text
);
CREATE TABLE IF NOT EXISTS users (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	role integer   ,
	created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
);

CREATE TABLE IF NOT EXISTS refresh (
    refresh_id serial ,
    token VARCHAR (100) NOT NULL,
    user_id integer not null references users(user_id),
    expires_at timestamp not null
);

CREATE TABLE IF NOT EXISTS roles (
    role_id serial,
    role_name varchar(15)

);



CREATE TABLE IF NOT EXISTS account_roles (
  user_id INT NOT NULL,
  role_id INT NOT NULL,
  grant_date TIMESTAMP,
  PRIMARY KEY (user_id, role_id),
  FOREIGN KEY (role_id)
  REFERENCES roles (role_id),
  FOREIGN KEY (user_id)
  REFERENCES users (user_id)
);`