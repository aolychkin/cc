package CLI

//Модель удаленного пользователя
type RemoteUser struct {
	User     string
	Password string
	Host     string
	Port     int
}