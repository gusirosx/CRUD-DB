package usecases

// type Service struct {
// 	repo Repository
// }

// func NewService(r Repository) *Service {
// 	return &Service{
// 		repo: r,
// 	}
// }

// type Repository interface {
// 	Reader
// 	//Writer
// }

// type Reader interface {
// 	Search(query string) error
// }

// type UserRepoImpl struct {
// 	personRepo person.PersonRepo
// }

// func CreatePersonUsecase(personRepo person.PersonRepo) person.PersonUsecase {
// 	return &PersonUsecaseImpl{personRepo}
// }

// func (e *PersonUsecaseImpl) Create(person *model.Person) (*model.Person, error) {
// 	return e.personRepo.Create(person)
// }

// func (e *PersonUsecaseImpl) ReadAll() (*[]model.Person, error) {
// 	return e.personRepo.ReadAll()
// }
