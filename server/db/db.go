пакет db

импорт (
	"база данных / sql"
	"сеть / URL"

	_ "github.com/lib/pq"
)

type  Connection  struct {
	         Строка DbName
	Пользователь , строка пароля 
	           Строка хоста
	DisableSSL      bool
}

func ( c  * Connection ) ConnectionURL () string {
	dbUrl  : =  & url. URL {
		Схема : "postgres" ,
		Хост :    c . Хост ,
		Пользователь :    url . UserPassword ( c . Пользователь , c . Пароль ),
		Путь :    c . DbName ,
	}
	если  c . DisableSSL {
		dbUrl . RawQuery  = url. Values {
			"sslmode" : [] строка { "отключить" },
		}. Кодировать ()
	}
	вернуть  dbUrl . Строка ()
}

func ( c  * Connection ) Open () ( * sql. DB , error ) {
	вернуть  sql . Open ( "postgres" , c . ConnectionURL ())
}
