инструменты пакета

импорт (
	"кодировка / json"
	"журнал"
	"net / http"
)

// WriteJsonOk отправляет ответ об ошибке 400 с объектом JSON, описывающим причину ошибки.
func  WriteJsonBadRequest ( rw http. ResponseWriter , строка сообщения  ) {
	writeJson ( rw , http . StatusBadRequest , & errorObject { Сообщение : сообщение })
}

// WriteJsonOk отправляет ответ с ошибкой 500.
func  WriteJsonInternalError ( rw http. ResponseWriter , строка сообщения  ) {
	writeJson ( rw , http . StatusInternalServerError , & errorObject { Сообщение : сообщение })
}

// WriteJsonOk отправляет ответ 200 клиенту, сериализуя входной объект в формате JSON.

func  WriteJsonOk ( rw http. ResponseWriter , res  interface {}) {
	writeJson ( rw , http . StatusOK , res )
}

func  writeJson ( rw http. ResponseWriter , status  int , res  interface {}) {
	rw . Заголовок (). Установить ( "тип содержимого" , "приложение / json" )
	rw . WriteHeader ( статус )
	ошибка  : =  json . NewEncoder ( rw ). Кодировать ( разрешение )
	if  err  ! =  nil {
		журнал . Printf ( "Ошибка записи ответа:% s" , ERR )
	}
}
