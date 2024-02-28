package services

type (
	UserService struct{}
)

func (u *UserService) UserByID(id string) (data string, err error) {
	return "Hello World", nil
}
