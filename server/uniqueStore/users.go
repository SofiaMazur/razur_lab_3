пакет generalStore

импорт (
	"база данных / sql"
	"fmt"
	"github.com/GVG/l3/server/tools"
	_ "github.com/lib/pq"
)

type  UserStore  struct {
	Db  * sql. БД
}

func  NewUserStore ( db  * sql. DB ) * UserStore {
	return  & UserStore { Db : db }
}

func ( s  * UserStore ) ListUsers () ( * tools. Users , error ) {
	строк , err  : =  s . Дб . Запрос ( «ВЫБРАТЬ * ОТ пользователей» )
	if  err  ! =  nil {
		вернуть  ноль , ошибиться
	}

	отложить  строки . Закрыть ()

	var  res [] * tools. Пользователь
	для  рядов . Next () {
		 инструменты var u . Пользователь
		если  ошибка  : =  строк . Сканировать ( & u . Id , & u . Name ); err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		res  =  append ( res , & u )
	}

	var  fullUsers [] * tools. Пользователь
	if  res  ==  nil {
		fullUsers  =  make ([] * tools. Пользователь , 0 )
	} else {
		для  i  : =  0 ; я  <  len ( разрешение ); i ++ {
			интересует , err  : =  s . GetUsersInterestByID ( res [ i ]. Id )
			if  err  ! =  nil {
				вернуть  ноль , ошибиться
			}
			fullUser  : = инструменты. Пользователь { Id : res [ i ]. Id , Name : res [ i ]. Имя , интересы : интересы }
			fullUsers  =  append ( fullUsers , & fullUser )
		}
	}

	результат  : =  & tools. Пользователи { fullUsers }
	вернуть  результат , ноль
}

func ( s  * UserStore ) FindUserByName ( строка имени  ) ( * tools. Users , error ) {
	var  textError  строка
	var  err  error
	var  fullUsers [] * tools. Пользователь

	if  len ( name ) ==  0 {
		textError  =  "Имя пользователя не указано"
		err  =  fmt . Errorf ( textError )
		вернуть  ноль , ошибиться
	}
	строк , err  : =  s . Дб . Запрос ( `SELECT * FROM users where name = $ 1` , name )
	if  err  ! =  nil {
		textError  =  "Такого пользователя нет"
		err  =  fmt . Errorf ( textError )
		вернуть  ноль , ошибиться
	}

	отложить  строки . Закрыть ()

	var  res [] * tools. Пользователь
	для  рядов . Next () {
		 инструменты var u . Пользователь
		если  ошибка  : =  строк . Сканировать ( & u . Id , & u . Name ); err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		res  =  append ( res , & u )
	}

	if  res  ==  nil {
		textError  =  "Нет такого пользователя"
		err  =  fmt . Errorf ( textError )
		fullUsers  =  make ([] * tools. Пользователь , 0 )
		вернуть  ноль , ошибиться
	} 
	для  i  : =  0 ; я  <  len ( разрешение ); i ++ {
		интересует , err  : =  s . GetUsersInterestByID ( res [ i ]. Id )
		if  err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		fullUser  : = инструменты. Пользователь { Id : res [ i ]. Id , Name : res [ i ]. Имя , интересы : интересы }
		fullUsers  =  append ( fullUsers , & fullUser )
	}
	err  =  nil
	результат  : =  & tools. Пользователи { fullUsers }
	вернуть  результат , ошибка
}

функ ( s  * UserStore ) CreateUser ( имя пользователя  струнные , интересы [] строка ) ошибки {
	store  : =  NewForumStore ( s . Db )
	if  len ( имя пользователя ) ==  0 {
		вернуть  fmt . Errorf ( «Имя пользователя не указано» )
	}
	if  len ( интересы ) ==  0 {
		вернуть  fmt . Errorf ( "Интересы не указаны" )
	}
	для  _ , интерес  : =  диапазон  интересов {
		if  len ( проценты ) ==  0 {
			вернуть  fmt . Errorf ( «Проценты не могут быть пустыми» )
		}
	}
	_ , ошибка  : =  s . Дб . Exec ( `INSERT INTO users (name) VALUES ($ 1)` , username )
	if  err  ! =  nil {
		вернуть  fmt . Errorf ( «Пользователь с таким именем уже существует» )
	}
	пользователи , err  : =  s . FindUserByName ( имя пользователя )
	для  i  : =  0 ; i  <  len ( интересы ); i ++ {
		_ , ошибка  =  s . Дб . Exec ( `INSERT INTO InterestList (Interest, userID) VALUES ($ 1, $ 2)` ,
			интересы [ i ], пользователи . UsersArr [ 0 ]. Id )
		форум , укажите  : =  магазин . FindForumByTopic ( интересы [ i ])
		если  указано  ==  nil {
			err  =  магазин . AddUserToForum ( форум [ 0 ]. Id , users . UsersArr [ 0 ]. Id )
		}
	}
	вернуть  ошибку
}

func ( s  * UserStore ) GetUsersInterestByID ( id  int ) ([] строка , ошибка ) {
	if  id  <  1 {
		возврат  ноль , fmt . Errorf ( «ID неверный» )
	}
	строк , err  : =  s . Дб . Запрос ( `
	Выбрать
		InterestList.interest
	из
		пользователи, InterestList
	где
		InterestList.userID = users.id
	и
		users.id = $ 1` ,
		id )

	if  err  ! =  nil {
		вернуть  ноль , ошибиться
	}

	отложить  строки . Закрыть ()

	var  res [] строка
	для  рядов . Next () {
		var  i  строка
		если  ошибка  : =  строк . Сканировать ( & i ); err  ! =  nil {
			вернуть  ноль , ошибиться
		}
		res  =  добавить ( res , i )
	}
	if  res  ==  nil {
		res  =  make ([] строка , 0 )
	}

	return  res , nil
}
