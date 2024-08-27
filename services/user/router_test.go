package user

import (
	"standardApi/types"
	"testing"
)

func TestUserServiceHandler(t *testing.T) {

	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should failed if the user payload is valid", func(t *testing.T) {

	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}
