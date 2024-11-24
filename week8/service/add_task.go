package service

import(
	"context"
	"fmt"

	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/auth" // 추가
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/entity"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/store"
)

type AddTask struct{
	DB	store.Execer
	Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error){
	//추가
	id, ok := auth.GetUserID(ctx)
	if !ok{
		return nil, fmt.Errorf("user_id not found")
	}
	//여기까지
	t := &entity.Task{
		UserID: id, //추가
		Title: title,
		Status: entity.TaskStatusTodo,
	}
	err := a.Repo.AddTask(ctx, a.DB, t)
	if err != nil{
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}