пакет uniqueStore

импорт (
	"база данных / sql"
	"fmt"
	"github.com/GVG/l3/server/tools"
	_ "github.com/lib/pq"
)

type  ForumStore  struct {
	Db  * sql. БД
}

func  NewForumStore ( db  * sql. DB ) * ForumStore {
	возврат  и ForumStore { Db : db }
}

func ( s  * ForumStore ) ListForums () ( * tools. Forums , error ) {
	строк , err  : =  s . Дб . Запрос ( «ВЫБРАТЬ * ИЗ форумов» )
	if  err  ! =  nil {
		вернуть  ноль , ошибиться
	}

	отложить  строки . Закрыть ()

	var  res [] * tools. Форум
	для  рядов . Next () {
		var  f инструменты. Форум
		если  ошибка  : =  строк . Сканировать ( & f . Идентификатор , & f . Имя , & f . Тема ); err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		res  =  append ( res , & f )
	}

	var  fullForums [] * tools. Форум
	if  res  ==  nil {
		fullForums  =  make ([] * tools. Forum , 0 )
	} else {
		для  i  : =  0 ; я  <  len ( разрешение ); i ++ {
			пользователи , err  : =  s . GetForumUsersByID ( res [ i ]. Id )
			if  err  ! =  nil {
				вернуть  ноль , ошибиться
			}
			fullForum  : = инструменты. Forum {
				Id :     res [ i ]. Id ,
				Имя :   res [ i ]. Имя ,
				Тема : res [ i ]. Тема ,
				Пользователи : users }
			fullForums  =  Append ( fullForums , & fullForum )
		}
	}

	результат  : =  & tools. Форумы { fullForums }
	вернуть  результат , ошибка
}

func ( s  * ForumStore ) FindForumByName ( строка имени  ) ( * tools. Forums , error ) {
	var  textError  строка
	var  err  error
	var  fullForums [] * tools. Форум

	if  len ( name ) ==  0 {
		textError  =  "Название форума не указано"
		err  =  fmt . Errorf ( textError )
		fullForums  =  make ([] * tools. Forum , 0 )
		вернуть  ноль , ошибиться
	}
	строк , err  : =  s . Дб . Запрос ( `ВЫБРАТЬ * ИЗ форумов, где имя = $ 1` , имя )
	if  err  ! =  nil {
		textError  =  "Такого форума нет"
		err  =  fmt . Errorf ( textError )
		вернуть  ноль , ошибиться
	}

	отложить  строки . Закрыть ()

	var  res [] * tools. Форум
	для  рядов . Next () {
		var  f инструменты. Форум
		если  err  =  rows . Сканировать ( & f . Идентификатор , & f . Имя , & f . Тема ); err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		res  =  append ( res , & f )
	}
	if  res  ==  nil {
		textError  =  "Нет такого форума"
		err  =  fmt . Errorf ( textError )
		вернуть  ноль , ошибиться
	}
	для  i  : =  0 ; я  <  len ( разрешение ); i ++ {
		пользователи , err  : =  s . GetForumUsersByID ( res [ i ]. Id )
		if  err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		fullForum  : = инструменты. Forum {
			Id :     res [ i ]. Id ,
			Имя :   res [ i ]. Имя ,
			Тема : res [ i ]. Тема ,
			Пользователи : users }
		fullForums  =  Append ( fullForums , & fullForum )
	}

	результат  : =  & tools. Форумы { fullForums }
	вернуть  результат , ноль
}

func ( s  * ForumStore ) FindForumByTopic ( строка имени  ) ([] * tools. Forum , error ) {
	if  len ( name ) ==  0 {
		возврат  ноль , fmt . Errorf ( «Название темы не указано» )
	}
	строк , err  : =  s . Дб . Запрос ( `SELECT * FROM форумов, где topicKeyword = $ 1` , имя )
	if  err  ! =  nil {
		вернуть  ноль , ошибиться
	}

	отложить  строки . Закрыть ()

	var  res [] * tools. Форум
	для  рядов . Next () {
		var  f инструменты. Форум
		если  ошибка  : =  строк . Сканировать ( & f . Идентификатор , & f . Имя , & f . Тема ); err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		res  =  append ( res , & f )
	}
	if  res  ==  nil {
		res  =  make ([] * tools. Forum , 0 )
		возврат  res , fmt . Errorf ( «такого форума нет» )
	}
	return  res , nil
}

func ( s  * ForumStore ) CreateForum ( name , topicKeyword  string ) error {
	if  len ( name ) ==  0 {
		вернуть  fmt . Errorf ( "Название форума не указано" )
	}
	if  len ( topicKeyword ) ==  0 {
		вернуть  fmt . Errorf ( "Название ключевого слова темы не указано" )
	}
	_ , ошибка  : =  s . Дб . Exec ( `INSERT INTO forum (name, topicKeyword) VALUES ($ 1, $ 2)` , name , topicKeyword )
	if  err  ! =  nil {
		вернуть  fmt . Errorf ( «Форум с таким названием или темой уже существует» )
	}
	форумы , err  : =  s . FindForumByName ( имя )
	_ , ошибка  =  s . Дб . Exec ( `INSERT INTO Список пользователей задается (forumsID) VALUES ($ 1)` , форумы . ForumsArr [ 0 ]. Id )
	вернуть  ошибку
}

func ( s  * ForumStore ) GetForumUsersByID ( id  int ) ([] строка , ошибка ) {
	if  id  <  1 {
		возврат  ноль , fmt . Errorf ( «ID неверный» )
	}
	строк , err  : =  s . Дб . Запрос ( `
	Выбрать
		users.name
	из
		форумы
	оставил присоединиться
		usersList
	на
		usersList.forumsID = forum.id
	оставил присоединиться
		пользователи
	на
		users.id = usersList.userID
	где
		forum.id = 1 доллар США
	ГРУППА ПО
		users.id
	HAVING users.id не равен NULL
	` ,
		id )

	if  err  ! =  nil {
		вернуть  ноль , ошибиться
	}

	отложить  строки . Закрыть ()

	var  res [] строка
	для  рядов . Next () {
		var  u  строка
		если  ошибка  : =  строк . Сканировать ( & u ); err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		if  u  ! =  "" {
			res  =  добавить ( res , u )
		}
	}
	if  res  ==  nil {
		res  =  make ([] строка , 0 )
	}

	return  res , nil
}

func ( s  * ForumStore ) AddUserToForum ( idForum , idUser  int ) error {
	_ , ошибка  : =  s . Дб . Exec ( `INSERT INTO usersList (forumID, userID) VALUES ($ 1, $ 2)` , idForum , idUser )
	вернуть  ошибку
}
