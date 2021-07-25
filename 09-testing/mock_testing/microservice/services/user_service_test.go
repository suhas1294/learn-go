package services

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/suhas1294/learn-go/09-testing/mock_testing/microservice/models"
)

type MockHelper struct{}

func (m *MockHelper) GetAllUsers() *[]models.User {
	sliceOfUsers := make([]models.User, 0)
	
	geo := models.Geo{"12.9569", "77.7011"}
	address := models.Address{"Kulas Light", "Apt. 556", "Gwenborough", "92998-3874", geo}
	company := models.Company{"Romaguera-Crona", "Multi-layered client-server neural-net", "harness real-time e-markets"}
	user := models.User{1, "Leanne Graham", "Bret", "Sincere@april.biz", "phone", "website", company, address}
	sliceOfUsers = append(sliceOfUsers, user)
	return &sliceOfUsers
}

func TestGetAllUsersWithMock(t *testing.T){
	// previously, 'UserService' was '&UserHelper{}', no we are changing it.
	UserService = &MockHelper{}
	
	allUsers := UserService.GetAllUsers()
	assert.EqualValues(t, len(*(allUsers)), 1) // dereferencing
}