package breed

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_All_Should_Return_The_Available_Breeds(t *testing.T) {
	repository := NewRepository(http.DefaultClient)
	breeds, err := repository.GetAll()
	assert.Nil(t, err)
	assert.Contains(t, breeds, "dingo")
}
