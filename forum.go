package rupert

import (
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	forum = ForumManager{}

	queryCreateForumCategory = `
		INSERT INTO forum_category
			(name, order_idx)
		VALUES
			($1, $2)
	`

	queryForumCategoryByName = `
		SELECT
			category_id, order_idx, name, created_on, updated_on
		FROM
			forum_category
		WHERE
			username=$1
	`
)

type (
	ForumManager struct{}

	ForumCategory struct {
		CategoryID int       `db:"category_id"`
		OrderIdx   int       `db:"order_idx"`
		Name       string    `db:"name"`
		Forums     []Forum   `db:"-"`
		CreatedOn  time.Time `db:"created_on"`
		UpdatedOn  time.Time `db:"updated_on"`
	}

	Forum struct {
		ForumID    int       `db:"forum_id"`
		CategoryID int       `db:"category_id"`
		Name       string    `db:"name"`
		Topics     int       `db:"topics"`
		Posts      int       `db:"posts"`
		OrderIdx   int       `db:"order_idx"`
		CreatedOn  time.Time `db:"created_on"`
		UpdatedOn  time.Time `db:"updated_on"`
		Threads    []*Thread `db:"-"`
		LastThread *Thread   `db:"-"`
	}

	Thread struct {
		ThreadID      int       `db:"thread_id"`
		Forum_id      int       `db:"forum_id"`
		Title         string    `db:"title"`
		Replies       int       `db:"replies"`
		Views         int       `db:"views"`
		Sticky        bool      `db:"sticky"`
		CreatedOn     time.Time `db:"created_on"`
		UpdatedOn     time.Time `db:"updated_on"`
		LastCommentID int       `db:"last_comment_id"`
		LastComment   *Comment  `db:"-"`
	}

	Comment struct {
		CommentID int       `db:"comment_id"`
		ThreadID  int       `db:"thread_id"`
		Message   string    `db:"message"`
		CreatedOn time.Time `db:"created_on"`
		UpdatedOn time.Time `db:"updated_on"`
	}
)

func addThread(thread *Thread, comment *Comment) {

}

func addThreadComment(thread *Thread, comment *Comment) {

}

func getThreads(Category_id int, limit int) {

}

func (fm *ForumManager) Initialize() {

}

func NewForumCategory(name string, order_idx int) ForumCategory {
	return ForumCategory{OrderIdx: order_idx, Name: name}
}

func CategoryCreate(db *sqlx.DB, name string, order_idx int) (*ForumCategory, error) {
	cat := NewForumCategory(name, order_idx)
	tx := db.MustBegin()
	_, err := tx.NamedExec(queryCreateForumCategory, &cat)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return ForumCategoryGetByName(db, name)
}

func ForumCategoryGetByName(db *sqlx.DB, name string) (*ForumCategory, error) {
	var fc ForumCategory
	err := db.Get(&fc, queryForumCategoryByName, name)
	return &fc, err
}
