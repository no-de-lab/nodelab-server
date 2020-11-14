package repository

const FindByIdQuery = `SELECT id, email, username, intro, profile_image, created_at, updated_at FROM user WHERE id = ? `
const FindByEmailQuery = `SELECT email, password FROM user WHERE email = ?`
const CreateUserQuery = `
	INSERT INTO user(email, username, password, intro, profile_image, created_at, updated_at) 
	VALUES(:email, :username, :password, :intro, :profile_image, now(), now())
`
