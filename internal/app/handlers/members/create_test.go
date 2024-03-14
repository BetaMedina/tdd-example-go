package members

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMemberRepository struct {
	mock.Mock
}

func (m *MockMemberRepository) Create(email string, name string) error {
	payload := map[string]interface{}{"email": email, "name": name}
	args := m.Called(payload)
	return args.Error(1)
}

func preset() (*fiber.App, CreateMember, *MockMemberRepository) {
	app := fiber.New()
	mockRepo := new(MockMemberRepository)
	handler := NewCreateMember(mockRepo)
	return app, handler, mockRepo
}

func TestCreateSuccessMember(t *testing.T) {
	app, handler, mockRepo := preset()
	app.Post("/members", handler.CreateMember)
	payload := `{
		"name": "John Doe",
		"email": "jonhdoe@gmail.com"
	}`

	req := httptest.NewRequest("POST", "/members", strings.NewReader(payload))
	req.Header.Add("content-type", "application/json; charset=utf-8")
	mockRepo.On("Create", mock.Anything).Return(nil, nil)

	resp, _ := app.Test(req)

	body := make([]byte, resp.ContentLength)

	resp.Body.Read(body)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "They should be equal")
	assert.Contains(t, string(body), "Member created")
	assert.Contains(t, string(body), "success")
}

func TestEmptyPayloadMember(t *testing.T) {
	app, handler, mockRepo := preset()
	app.Post("/members", handler.CreateMember)
	payload := `{}`

	req := httptest.NewRequest("POST", "/members", strings.NewReader(payload))
	mockRepo.On("Create", mock.Anything).Return(true, nil)

	resp, _ := app.Test(req)

	body := make([]byte, resp.ContentLength)

	resp.Body.Read(body)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "They should be equal")
	assert.Contains(t, string(body), "Payload invalid")
	assert.Contains(t, string(body), "error")
}

func TestErrorOnSaveDatabaseMember(t *testing.T) {
	app, handler, mockRepo := preset()
	app.Post("/members", handler.CreateMember)
	payload := `{
		"name": "John Doe",
		"email": "jonhdoe@gmail.com"
	}`

	mockRepo.On("Create", mock.Anything).Return("", errors.New("database error"))
	req := httptest.NewRequest("POST", "/members", strings.NewReader(payload))
	req.Header.Add("content-type", "application/json; charset=utf-8")

	resp, _ := app.Test(req)

	body := make([]byte, resp.ContentLength)

	resp.Body.Read(body)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "They should be equal")
	assert.Contains(t, string(body), "Internal server error")
	assert.Contains(t, string(body), "error")
}
