package repository

// Query for User Domain
const (
	findAccountByEmailQuery = `SELECT email, password FROM account WHERE email = ?`
	createUserByEmailQuery  = `
	INSERT INTO account(email, password, created_at, updated_at)
	VALUES (:email, :password, now(), now())`
	createUserBySocialQuery = `
	INSERT INTO account(email, provider, provider_id, access_token, created_at, updated_at)
	VALUES (:email, :provider, :provider_id, :access_token, now(), now())`
)
