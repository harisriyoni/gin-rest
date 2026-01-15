package book

type ServiceBook interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type servicebookrepository struct {
	repository RepositoryBookInterface
}

func NewBookService(repository RepositoryBookInterface) *servicebookrepository {
	return &servicebookrepository{repository}
}

func (s *servicebookrepository) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}
func (s *servicebookrepository) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *servicebookrepository) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *servicebookrepository) Update(ID int, bookRequest BookRequest) (Book, error) {

	book, err := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}
func (s *servicebookrepository) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)

	deleteBook, err := s.repository.Delete(book)
	return deleteBook, err
}
