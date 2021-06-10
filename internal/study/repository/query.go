package repository

// Study query
const (
	// FindByIDQuery study leader profile image 파일 테이블 따로 조인 필요
	FindByIDQuery    = `SELECT id, studyname, created_at, updated_at, 'limit', start_date, finish_date, summary, title, content, notice, leader_id FROM study WHERE status = "OPEN" AND id = ?;`
	FindByTitleQuery = `
	SELECT id, studyname, created_at, updated_at, 'limit', start_date, finish_date, summary, title, content, notice, leader_id
	FROM study
	WHERE title like %?%
	AND status = 'OPEN';
	`
)
