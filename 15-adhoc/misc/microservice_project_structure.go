package main

import "fmt"

// ------------------------------------------------------------------------
// models/person.go
type Person struct{
	Name string
	age int
}

// ------------------------------------------------------------------------
// dummy database
var(
	persons = map[int]Person {
		123: Person{"bar", 26},
		567: Person{"foo", 45},
	}
)
// ------------------------------------------------------------------------
// controllers/person_controller.go
service.PersonPerson.GetPerson(123)

// ------------------------------------------------------------------------
// services/person_service.go
var (
	PersonService personService
)
type personService struct{}
func (p *personService) GetPerson(id int) *models.Person{
	return models.PersonDao.GetPerson(123)
}

// ------------------------------------------------------------------------
// models/person_dao.go
var (
	PersonDao personDao
)
type PersonInterface interface{
	GetPerson(int) *Person
}
type personDao struct{}
func init(){
	PersonDao = &personDao{}
}
func (p *personDao) GetPerson(id int) *models.Person{
	// query from database and give
	if person, ok := persons[id]; ok {
		return person
	}
	return error.New("No such user found")
}

// ------------------------------------------------------------------------
// services/person_service_test.go
var (
	PersonService personService
	GetPerson = func(id int) *models.Person
)
type personDaoMock struct{}
func (p *personDaoMock) GetPerson(id int) *models.Person {
	return GetPerson(id)
}

func init(){
	models.PersonDao = &personDaoMock{}
	// if we dont do this, models.PersonDao will be &personDao{}
}

func TestGetPersonValid(t *testing.T){
	// step-1:	override 'GetPerson' method which was just declared at start of 'services/person_service_test.go' file.
	GetPerson = func(id int) *models.Person{
		return &Person{
			Name: "bar",
			Age: 26
		}
	}
	// step-2: call service method, which will call
	person := PersonService.GetPerson(123)

	/* 
		When this file('person_service_test.go') get called, init() will be called and models.PersonDao will be 
		initialized to '&personDaoMock{}' and NOT '&personDao{}'.
		Hence when we execute GetPerson() method, GetPerson method attached to type 'personDaoMock'
		will be called and GetPerson method attached to 'personDao' will NOT be called.
		Flow of execution:
		1) init() of services/person_service_test.go
		2) GetPerson variable will initialised with a function.
		3) when we say 'PersonService.GetPerson(123)', person_service will person_dao, 
			person_dao's 'PersonDao' is now pointing to '&personDaoMock{}'
		4) Hence GetPerson() method attached to 'personDaoMock' will get executed
		5) Note that 'GetPerson' variable is already initialised by this time.
			Now inside GetPerson method that is attached to 'personDaoMock', 
			we are calling local variable 'GetPerson' with parameter as '123' and 
			this will return a &Person{"bar", 26}
	*/

	// step-3: validations & assertions
	assert.NotNil(person, "expeced a valid person")
}

func TestGetPersonInvalid(t *testing.T){

}

// ------------------------------------------------------------------------
// main.go
func main(){
	
}