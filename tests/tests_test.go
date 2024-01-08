package tests

import (
	"encoding/json"
	"goServer/db"
	"goServer/routers"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	db.InitDb()
	router = routers.SetupRouter()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setupTest() {
	db.InitTestData()
}
func MakeRequest(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	router.ServeHTTP(w, req)
	return w
}

func TestGetTodos(t *testing.T) {
	setupTest()
	w := MakeRequest("GET", "/todos", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseItems []db.Todo
	err := json.Unmarshal(w.Body.Bytes(), &responseItems)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "test title", responseItems[0].Title)
	assert.Equal(t, "something", responseItems[1].Title)
	assert.Equal(t, "Another", responseItems[2].Title)
	assert.Equal(t, "Hello", responseItems[3].Title)

	assert.Equal(t, db.BoolPointer(true), responseItems[0].Completed)
	assert.Equal(t, db.BoolPointer(false), responseItems[1].Completed)
	assert.Equal(t, db.BoolPointer(false), responseItems[2].Completed)
	assert.Equal(t, db.BoolPointer(true), responseItems[3].Completed)

}

func TestDeleteTodo(t *testing.T) {
	setupTest()
	w := MakeRequest("GET", "/todos", nil)
	var responseItems []db.Todo
	err := json.Unmarshal(w.Body.Bytes(), &responseItems)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 4, len(responseItems))
	url := "/todos/" + string(responseItems[3].ID)
	w2 := MakeRequest("DELETE", url, nil)
	assert.Equal(t, http.StatusNoContent, w2.Code)
	assert.Equal(t, "", w2.Body.String())

	w = MakeRequest("GET", "/todos", nil)
	err = json.Unmarshal(w.Body.Bytes(), &responseItems)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(responseItems))
}
