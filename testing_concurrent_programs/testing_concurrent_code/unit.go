// user.go

package main

type User struct {
	ID    string
	Email string
}

type UserRepository interface {
	FindUserByID(id string) (*User, error)
}

type NotificationService interface {
	SendEmail(to, subject, body string) error
}

type UserService struct {
	repo         UserRepository
	notification NotificationService
}

func NewUserService(repo UserRepository, notification NotificationService) *UserService {
	return &UserService{
		repo:         repo,
		notification: notification,
	}
}

func (s *UserService) NotifyUser(userID, subject, message string) error {
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return err
	}

	err = s.notification.SendEmail(user.Email, subject, message)
	if err != nil {
		return err
	}

	return nil
}

type MockUserRepository struct {
	User *User
	Err  error
}

func (m *MockUserRepository) FindUserByID(id string) (*User, error) {
	return m.User, m.Err
}

type MockNotificationService struct {
	Err error
}

func (m *MockNotificationService) SendEmail(to, subject, body string) error {
	return m.Err
}
