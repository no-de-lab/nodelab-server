package repository

const (
	FindAccountByEmailQuery = `SELECT email FROM account WHERE email = ?`
	CreateUserByEmailQuery  = `
	INSERT INTO account(email, password, created_at, updated_at)
	VALUES (:email, :password, now(), now())`
	CreateUserBySocialQuery = `
	INSERT INTO account(email, created_at, updated_at)
	VALUES (:email, now(), now())`
)
