package store

import(
	"context"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/entity"
)

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, id entity.UserID,
) (entity.Tasks, error){
	tasks := entity.Tasks{}
	//user_id 추가, WHERE 절 추가
	sql := `SELECT
		id, user_id, title,
		status, created, modified
	FROM task
	WHERE user_id = ?;`
	//id 추가
	if err := db.SelectContext(ctx, &tasks, sql, id); err != nil{ 
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
		//t.UserID 추가
		ctx, sql, t.UserID, t.Title, t.Status, 
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