package rupert

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
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
	queryForumByName = `
		SELECT
			forum_id, category_id, name, topics, posts, order_idx, created_on, updated_on
		FROM
			forum
		WHERE
			name=$1
		LIMIT 1
	`
)

type (
	ForumCategory struct {
		CategoryID int       `db:"category_id"`
		OrderIdx   int       `db:"order_idx"`
		Name       string    `db:"name"`
		Forums     []Forum   `db:"-"`
		CreatedOn  time.Time `db:"created_on"`
		UpdatedOn  time.Time `db:"updated_on"`
	}

	Forum struct {
		ForumID    int            `db:"forum_id"`
		CategoryID int            `db:"category_id"`
		Name       string         `db:"name"`
		Topics     int            `db:"topics"`
		Posts      int            `db:"posts"`
		OrderIdx   int            `db:"order_idx"`
		CreatedOn  time.Time      `db:"created_on"`
		UpdatedOn  time.Time      `db:"updated_on"`
		Threads    []*ForumThread `db:"-"`
		LastThread *ForumThread   `db:"-"`
	}

	ForumThread struct {
		ThreadID      int           `db:"thread_id"`
		ForumID       int           `db:"forum_id"`
		Title         string        `db:"title"`
		Replies       int           `db:"replies"`
		Views         int           `db:"views"`
		Sticky        bool          `db:"sticky"`
		CreatedOn     time.Time     `db:"created_on"`
		UpdatedOn     time.Time     `db:"updated_on"`
		LastCommentID int           `db:"last_comment_id"`
		LastComment   *ForumComment `db:"-"`
	}

	ForumComment struct {
		CommentID int       `db:"comment_id"`
		ThreadID  int       `db:"thread_id"`
		Message   string    `db:"message"`
		UserID    int       `db:"user_id"`
		CreatedOn time.Time `db:"created_on"`
		UpdatedOn time.Time `db:"updated_on"`
	}
)

func CreateForum(db *sqlx.DB, category_id int, name string, order_idx int) (*Forum, error) {
	if category_id <= 0 {
		return nil, errors.New("Invalid category ID, must be positive integer")
	}
	if name == "" {
		return nil, errors.New("Invalid forum name, must be non-empty string")
	}
	q := `INSERT INTO forum (category_id, name, order_idx) VALUES ($1, $2, $3)`
	_, err := db.Exec(q, category_id, name, order_idx)
	if err != nil {
		return nil, err
	}
	return ForumGetByName(db, name)
}

func ForumGetByName(db *sqlx.DB, name string) (*Forum, error) {
	var f Forum
	err := db.Get(&f, queryForumByName, name)
	return &f, err
}

func CreateForumThread(db *sqlx.DB, forum_id int, title string, sticky bool, comment string, user_id int) (*ForumThread, error) {
	var thread_id int
	rows, err := db.Query(`INSERT INTO forum_thread (forum_id, title, sticky) VALUES ($1, $2, $3) RETURNING thread_id`, forum_id, title, sticky)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&thread_id)
	}
	_, err = db.Query(`INSERT INTO forum_comment (thread_id, message, user_id`)
	return &ForumThread{ThreadID: thread_id, ForumID: forum_id, Sticky: sticky, Title: title}, err
}

func addThreadComment(thread *ForumThread, comment *ForumComment) {

}

func getThreads(Category_id int, limit int) {

}

func NewForumThread() {

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
