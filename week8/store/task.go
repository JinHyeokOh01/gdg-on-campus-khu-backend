package store

import(
	"context"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/entity"
)

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, id entity.UserID, // id 추가
) (entity.Tasks, error){
	tasks := entity.Tasks{}
	//user_id 추가, WHERE 절 추가
	sql := `SELECT
		id, user_id, title,
		status, created, modified
	FROM task
	WHERE user_id = ?;`
	if err := db.SelectContext(ctx, &tasks, sql, id); err != nil{ //id 추가
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	//user_id 추가
	sql := `INSERT INTO task
		(user_id, title, status, created, modified)
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.UserID, t.Title, t.Status, //t.UserID 추가
		t.Created, t.Modified,
	)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	if err != nil{
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}