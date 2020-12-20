инструменты пакета

// Пользовательская структура JSON
type  User  struct {
	Id         int       `json:" - "`
	 Строка  имени `json:" name "`
	Интересы [] строка  `json:" интересы "`
}

// Структура форума JSON
type  Forum  struct {
	Id       int     `json:" - "`
	 Строка  имени `json:" name "`
	 Строка  темы `json:" topic "`
	Users [] строка  `json:" users "`
}

// Структура JSON форумов
type  Forums  struct {
	ForumsArr [] * Форум  `json:" форумы "`
}

// Структура ResponseName JSON
type  ResponseName  struct {
	 Строка  имени `json:" name "`
}

// Массив пользователей
type  Users  struct {
	UsersArr [] * Пользователь  `json:" users "`
}

type  errorObject  struct {
	 Строка  сообщения `json:" message "`
}
