package repository

// TODO: profile_image_id join 추가
const (
	FindByIDQuery    = `SELECT id, email, username, intro, profile_image_id, created_at, updated_at FROM user WHERE id = ? `
	FindByEmailQuery = `SELECT id, email, username, profile_image_id, intro, github_url, created_at, updated_at FROM user WHERE email = ?`
	CreateUserQuery  = `
	INSERT INTO user(email, username, password, intro, profile_image_id, created_at, updated_at)
	VALUES (:email, :username, :password, :intro, :profile_image_id, now(), now())
`
)
