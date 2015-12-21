package rupert

import "time"

var (
	forum = ForumManager{}
)

type (
	ForumManager struct{}

	Category struct {
		category_id int
		order_idx   int
		name        string
		forums      []Forum
	}

	Forum struct {
		forum_id    int
		category_id int
		name        string
		topics      int
		posts       int
		order_idx   int
		updated_on  time.Time
		threads     []*Thread
		last_thread *Thread
	}

	Thread struct {
		thread_id       int
		forum_id        int
		title           string
		replies         int
		views           int
		sticky          bool
		created_on      time.Time
		updated_on      time.Time
		last_comment_id int
		last_comment    *Comment
	}

	Comment struct {
		comment_id int
		thread_id  int
		message    string
		created_on time.Time
		updated_on time.Time
	}
)

func addCategory(category Category) {

}

func addThread(thread *Thread, comment *Comment) {

}

func addThreadComment(thread *Thread, comment *Comment) {

}

func getThreads(Category_id int, limit int) {

}

func (fm *ForumManager) Initialize() {

}
