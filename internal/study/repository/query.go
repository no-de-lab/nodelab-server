package repository

// Study query
const (
	// FindByIDQuery study leader profile image 파일 테이블 따로 조인 필요
	FindByIDQuery    = "SELECT id, `name`, created_at, updated_at, `limit`, start_date, finish_date, summary, title, content, notice, leader_id, `status`, thumbnail_url FROM study WHERE id = ? AND deleted_at IS NULL;"
	FindByTitleQuery = `
		SELECT id, 'name', created_at, updated_at, 'limit', start_date, finish_date, summary, title, content, notice, leader_id
		FROM study
		WHERE title like %?%
		AND status = 'OPEN'
		AND deleted_at IS NULL;`
	CreateStudyQuery = "INSERT INTO study(`name`, `limit`, start_date, finish_date, summary, title, content, leader_id, notice, thumbnail_url, `status`) VALUES (:name, :limit, :start_date, :finish_date, :summary, :title, :content, :leader_id, :notice, :thumbnail_url, 'PROGRESS');"
	UpdateStudyQuery = "UPDATE study SET name=COALESCE(:name, `name`), `limit`=COALESCE(:limit, `limit`), start_date=COALESCE(:start_date, start_date), finish_date=COALESCE(:finish_date, finish_date), summary=COALESCE(:summary, summary), title=COALESCE(:title, title), content=COALESCE(:content, content), notice=COALESCE(:notice, notice), thumbnail_url=COALESCE(:thumbnail_url, thumbnail_url), `status`=COALESCE(:status, status) WHERE id=:id AND deleted_at IS NULL;"
	DeleteStudyQuery = `UPDATE study set deleted_at=now() WHERE id=?;`
	FindByEmailQuery = `SELECT id, email, username, profile_image_id, intro, github_url, created_at, updated_at FROM user WHERE email = ?`
)
