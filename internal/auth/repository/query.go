package repository

const (
	findAccountByEmailQuery = `SELECT email FROM account WHERE email = ?`
	createUserByEmailQuery  = `
	INSERT INTO account(email, password, created_at, updated_at)
	VALUES (:email, :password, now(), now())`
	createUserBySocialQuery = `
	INSERT INTO account(email, created_at, updated_at)
	VALUES (:email, now(), now())`
)
