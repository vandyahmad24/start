package students

type Service interface {
	CreateStudent(input StudentsInput) (Students, error)
	GetStudentById(ID int) (Students, error)
	UpdateStudentById(ID int, input StudentsInput) (Students, error)
	DeleteStudentByID(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateStudent(input StudentsInput) (Students, error) {
	students := Students{
		Name: input.Name,
		Age:  input.Age,
	}
	newStudent, err := s.repository.Store(students)

	if err != nil {
		return newStudent, err
	}
	return newStudent, nil
}

func (s *service) GetStudentById(ID int) (Students, error) {
	student, err := s.repository.FindByID(ID)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (s *service) UpdateStudentById(ID int, input StudentsInput) (Students, error) {
	student, err := s.repository.FindByID(ID)
	if err != nil {
		return student, err
	}
	student.Name = input.Name
	student.Age = input.Age

	updatedStudent, err := s.repository.Update(student)
	if err != nil {
		return updatedStudent, err
	}

	return updatedStudent, nil
}

func (s *service) DeleteStudentByID(ID int) error {
	student, err := s.repository.FindByID(ID)
	if err != nil {
		return err
	}
	err = s.repository.Delete(student)
	if err != nil {
		return err
	}
	return nil
}
