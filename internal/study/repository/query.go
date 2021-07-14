package repository

// Study query
const (
	// FindByIDQuery study leader profile image 파일 테이블 따로 조인 필요
	FindByIDQuery    = "SELECT id, `name`, created_at, updated_at, `limit`, start_date, finish_date, summary, title, content, notice, leader_id, `status`, thumbnail_url FROM study WHERE id = ? and deleted_at is null;"
	FindByTitleQuery = `
		SELECT id, 'name', created_at, updated_at, 'limit', start_date, finish_date, summary, title, content, notice, leader_id
		FROM study
		WHERE title like %?%
		AND status = 'OPEN'
		AND deleted_at is null;`
	CreateStudyQuery = "INSERT INTO study(`name`, `limit`, start_date, finish_date, summary, title, content, leader_id, notice, thumbnail_url, `status`) VALUES (:name, :limit, :start_date, :finish_date, :summary, :title, :content, :leader_id, :notice, :thumbnail, 'PROGRESS');"
	UpdateStudyQuery = "UPDATE study SET name=coalesce(:name, `name`), `limit`=coalesce(:limit, `limit`), start_date=coalesce(:start_date, start_date), finish_date=coalesce(:finish_date, finish_date), summary=coalesce(:summary, summary), title=coalesce(:title, title), content=coalesce(:content, content), notice=coalesce(:notice, notice), thumbnail_url=coalesce(:thumbnail_url, thumbnail_url), `status`=coalesce(:status, status) WHERE id=:id AND deleted_at is null;"
	DeleteStudyQuery = `DELETE FROM study where id=?;`
	FindByEmailQuery = `SELECT id, email, username, profile_image_id, intro, github_url, created_at, updated_at FROM user WHERE email = ?`
)
