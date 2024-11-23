package service

import(
	"context"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/entity"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister UserRegister
type TaskAdder interface{
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface{
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}

type UserRegister interface{
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}

type UserGetter interface{
	GetUser(ctx context.Context, db store.Queryer, name string)(*entity.User, error)
}

type TokenGenerator interface{
	GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
}