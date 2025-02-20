package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	h["Authorization"] = []string{"ApiKey 12345"}

	key, _ := GetAPIKey(h)
	fmt.Println(key)
}
