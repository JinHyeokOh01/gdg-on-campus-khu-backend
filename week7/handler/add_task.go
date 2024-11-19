package handler

import(
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/entity"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/store"
)

type AddTask struct{
	DB	*sqlx.DB
	Repo	*store.Repository
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	var b struct{
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil{
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := at.Validator.Struct(b); err != nil{
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	t := &entity.Task{
		Title:	b.Title,
		Status: entity.TaskStatusTodo,
	}
	err := at.Repo.AddTask(ctx, at.DB, t) //AddTask는 store에 정의된 method
	if err != nil{
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct{
		ID entity.TaskID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}