package store

import(
	"context"
	"testing"

	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/clock"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/entity"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/testutil"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/testutil/fixture" //추가
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"	
)

func prepareUser(ctx context.Context, t *testing.T, db Execer) entity.UserID{
	t.Helper()
	u := fixture.User(nil)
	result, err := db.ExecContext(ctx, "INSERT INTO user (name, password, role,
	created, modified) VALUES (?, ?, ?, ?, ?);", u.Name, u.Password, u.Role, u.Created, u.Modified)
	if err != nil{
		t.Fatalf("insert user: %v", err)
	}
	id, err := result.LastInserId()
	if err != nil{
		t.Fatalf("got user_id: %v", err)
	}
	return entity.UserID(id)
}

func TestRepository_ListTasks(t *testing.T){
	//추가
	t.Parallel()

	ctx := context.Background()
	//entity.Task를 작성하는 다른 테스트 케이스와 섞이면 테스트가 실패한다. 
	//하지만 트랜잭션을 적용해서 이 테스트 케이스 내에서는 테이블 상태가 유지된다.
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	//이 테스트 케이스가 완료되면 원래대로 되돌린다.
	t.Cleanup(func(){ _ = tx.Rollback() })
	if err != nil{
		t.Fatal(err)
	}
	//수정
	wantUserID, wantss := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	//wantUserID 추가
	gots, err := sut.ListTasks(ctx, tx, wantUserID)
	if err != nil{
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0{
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) (entity.UserID, entity.Tasks) {// UserID 추가
	t.Helper()
	//추가
	userID := prepareUser(ctx, t, con)
	otherUserID := prepareUser(ctx, t, con)

	/* 삭제
	if _, err := con.ExecContext(ctx, "DELETE FROM task;"); err != nil{
		t.Logf("failed to intialize task: %v", err)
	}
	*/
	c := clock.FixedClocker{}
	//UserID 추가
	wants := entity.Tasks{
		{
			UserID: userID,
			Title: "want task 1", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			UserID: userID,
			Title: "want task 2", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		/*삭제
		{
			Title: "want task 3", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		*/
	}
	//추가
	tasks := entity.Tasks{
		wants[0],
		{
			UserID: otherUserID,
			Title: "not want task", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		wants[1],
	}
	//여기까지
	result, err := con.ExecContext(ctx, //user_id 추가
	`INSERT INTO task (user_id, title, status, created, modified)
		VALUES
			(?, ?, ?, ?, ?),
			(?, ?, ?, ?, ?),
			(?, ?, ?, ?, ?);`,
		//수정
		tasks[0].UserID, tasks[0].Title, tasks[0].Status, tasks[0].Created, tasks[0].Modified,
		tasks[1].UserID, tasks[1].Title, tasks[1].Status, tasks[1].Created, tasks[1].Modified,
		tasks[2].UserID, tasks[2].Title, tasks[2].Status, tasks[2].Created, tasks[2].Modified,
		//여기까지
	)
	if err != nil{
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil{
		t.Fatal(err)
	}
	//수정
	tasks[0].ID = entity.TaskID(id)
	tasks[1].ID = entity.TaskID(id + 1)
	tasks[2].ID = entity.TaskID(id + 2)
	return userID, wants
	//여기까지
}

func TestRepository_AddTask(t *testing.T){
	t.Parallel()
	ctx := context.Background()

	c := clock.FixedClocker{}
	var wantID int64 = 20
	okTask := &entity.Task{
		Title:	"ok task",
		Status:	"todo",
		Created:	c.Now(),
		Modified:	c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil{
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	mock.ExpectExec(
		
		`INSERT INTO task \(title, status, created, modified\) VALUES \(\?, \?, \?, \?\)`,
	).WithArgs(okTask.Title, okTask.Status, okTask.Created, okTask.Modified).
		WillReturnResult(sqlmock.NewResult(wantID, 1))

	xdb := sqlx.NewDb(db, "mysql")
	r := &Repository{Clocker: c}
	if err := r.AddTask(ctx, xdb, okTask); err != nil{
		t.Errorf("want no error, but got %v", err)
	}
}