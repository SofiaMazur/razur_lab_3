пакет generalStore

импорт (
	"база данных / sql"
	"github.com/GVG/l3/server/tools"
	_ "github.com/lib/pq"
)

type  UniqueStore  struct {
	FStore  * ForumStore
	UStore  * UserStore
}

func  NewUniqueStore ( db  * sql. DB ) * GeneralStore {
	fstore  : =  NewForumStore ( db )
	ustore  : =  NewUserStore ( db )
	return  & UniqueStore { FStore : fstore , UStore : ustore }
}

func ( gs  * UniqueStore ) ListForums () ( * tools. Forums , error ) {
	вернуть  gs . FStore . ListForums ()
}

func ( gs  * UniqueStore ) FindForumByName ( строка имени  ) ( * tools. Forums , error ) {
	вернуть  gs . FStore . FindForumByName ( имя )
}

func ( gs  * UniqueStore ) FindUserByName ( строка имени  ) ( * tools. Users , error ) {
	вернуть  gs . UStore . FindUserByName ( имя )
}

func ( gs  * UniqueStore ) CreateForum ( name , topicKeyword  string ) error {
	вернуть  gs . FStore . CreateForum ( название , тема , ключевое слово )
}

func ( gs  * UniqueStore ) ListUsers () ( * tools. Users , error ) {
	вернуть  gs . UStore . ListUsers ()
}

функ ( гс  * UniqueStore ) CreateUser ( имя пользователя  струнные , интересы [] строка ) ошибки {
	вернуть  gs . UStore . CreateUser ( имя пользователя , интересы )
}
