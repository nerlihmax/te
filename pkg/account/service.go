package account

type Service interface {
	GetAccountById(ID string) (*Account, error)
	CreateAccount(Name string) (*Account, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) GetAccountById(id string) (*Account, error) {
	return s.repo.GetAccountById(id)
}

func (s service) CreateAccount(id string) (*Account, error) {
	return s.repo.CreateAccount(id)
}
