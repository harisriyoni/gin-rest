package book

type ServiceBook interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
}

type servicebookrepository struct {
	repository RepositoryBookInterface
}

func NewBookService(repository RepositoryBookInterface) *servicebookrepository {
	return &servicebookrepository{repository}
}

func (s *servicebookrepository) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	// return s.repository.FindAll()
	return books, err
}
func (s *servicebookrepository) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	// return s.repository.FindAll()
	return book, err
}

func (s *servicebookrepository) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
	}
	newBook, err := s.repository.Create(book)
	// return s.repository.FindAll()
	return newBook, err
}
