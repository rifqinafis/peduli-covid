package roles_test

import (
	"errors"
	"os"
	"peduli-covid/businesses/roles"
	role_mock "peduli-covid/businesses/roles/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	roleRepository role_mock.Repository
	roleUsecase    roles.Usecase
)

func setup() {
	roleUsecase = roles.NewRoleUsecase(&roleRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetByID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		roleDomain := roles.Domain{
			ID:   1,
			Code: "admin",
			Name: "admin",
		}
		roleRepository.On("GetByID", mock.AnythingOfType("int")).Return(roleDomain, nil).Once()

		result := roleUsecase.GetByID(1)

		assert.Equal(t, roleDomain.Code, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		roleRepository.On("GetByID", mock.AnythingOfType("int")).Return(roles.Domain{}, errors.New("error")).Once()

		result := roleUsecase.GetByID(1)

		assert.Equal(t, "", result)
	})
}
