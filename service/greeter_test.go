package service

import (
	"testify-mock/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// manual mock

type dbMock struct {
	mock.Mock
}

func (d *dbMock) FetchMessage(lang string) (string, error) {
	args := d.Called(lang)
	return args.String(0), args.Error(1)
}

func (d *dbMock) FetchDefaultMessage() (string, error) {
	args := d.Called()
	return args.String(0), args.Error(1)
}

func TestMockMethodWithoutArgs(t *testing.T) {
	theDBMock := dbMock{}
	theDBMock.On("FetchDefaultMessage").Return("foofofofof", nil)
	g := greeter{&theDBMock, "en"}
	assert.Equal(t, "Message is: foofofofof", g.GreetInDefaultMsg())
	theDBMock.AssertNumberOfCalls(t, "FetchDefaultMessage", 1)
	theDBMock.AssertExpectations(t)
}

func TestMockMethodWithArgs(t *testing.T) {
	theDBMock := dbMock{}
	theDBMock.On("FetchMessage", "sg").Return("lah", nil)
	g := greeter{&theDBMock, "sg"}
	assert.Equal(t, "Message is: lah", g.Greet())
	theDBMock.AssertExpectations(t)
}

func TestMockMethodWithArgsIgnoreArgs(t *testing.T) {
	theDBMock := dbMock{}
	theDBMock.On("FetchMessage", mock.Anything).Return("lah", nil)
	g := greeter{&theDBMock, "in"}
	assert.Equal(t, "Message is: lah", g.Greet())
	theDBMock.AssertCalled(t, "FetchMessage", "in")
	theDBMock.AssertNotCalled(t, "FetchMessage", "ch")
	theDBMock.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, &theDBMock)
}

func TestMockMethodWithArgsIgnoreArgsType(t *testing.T) {
	theDBMock := dbMock{}
	theDBMock.On("FetchMessage", mock.AnythingOfTypeArgument("string")).Return("lah", nil)
	g := greeter{&theDBMock, "in"}
	assert.Equal(t, "Message is: lah", g.Greet())
	theDBMock.AssertCalled(t, "FetchMessage", "in")
	theDBMock.AssertNotCalled(t, "FetchMessage", "ch")
	theDBMock.AssertExpectations(t)
	mock.AssertExpectationsForObjects(t, &theDBMock)
}

func TestMatchedBy(t *testing.T) {
	theDBMock := dbMock{}
	theDBMock.On("FetchMessage", mock.MatchedBy(func(lang string) bool { return lang == "in" })).Return("bzzzz", nil)
	g := greeter{&theDBMock, "in"}
	msg := g.Greet()
	assert.Equal(t, "Message is: bzzzz", msg)
	theDBMock.AssertExpectations(t)
}

func TestMockMethodWithoutArgsUsingMockery(t *testing.T) {
	theDBMock := mocks.DB{}
	theDBMock.On("FetchDefaultMessage").Return("foofofofof", nil)
	g := greeter{&theDBMock, "en"}
	assert.Equal(t, "Message is: foofofofof", g.GreetInDefaultMsg())
	theDBMock.AssertNumberOfCalls(t, "FetchDefaultMessage", 1)
	theDBMock.AssertExpectations(t)
}
