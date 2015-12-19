package rupert

var (
	forum = ForumManager{}
)

type (
	ForumManager struct{}

	Category struct {
		category_id int
		order       int
		name        string
		forums      []Forum
	}

	Forum struct {
		forum_id    int
		category_id int
		title       string
		topics      int
		posts       int
		order       int
		threads     []*Thread
		last_thread *Thread
	}

	Thread struct {
		thread_id    int
		category_id  int
		title        string
		replies      int
		views        int
		sticky       bool
		last_comment *Comment
	}

	Comment struct {
		comment_id int
		message    string
		created_on int
		updated_on int
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
