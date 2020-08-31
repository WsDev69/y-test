package query

const (
	Schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
user_id    VARCHAR(40) PRIMARY KEY,
first_name VARCHAR(80)  DEFAULT '',
last_name  VARCHAR(80)  DEFAULT '',
email      VARCHAR(250) DEFAULT '',
password   VARCHAR(250) DEFAULT NULL,
avatar_l   VARCHAR(250) DEFAULT ''
);
`
	InsertNewUser = `INSERT INTO user(user_id,email, password) 
								VALUES (:user_id, :email, :password)`
	SelectUserById           = `SELECT * FROM user WHERE user_id = $1`
	SelectUserByPassAndEmail = `SELECT * FROM user 
								WHERE email = $1 
								AND password = $2`
	SelectPageable = `SELECT * FROM user LIMIT $1 OFFSET $2`
	UpdateUser     = `UPDATE user SET %s WHERE user_id=:userId`
)
