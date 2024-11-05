package util

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func GetWorkspaceID(r *http.Request) (string, error) {
	file, err := os.OpenFile("/Users/grant/headers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer file.Close()

		file.WriteString(r.Header.Get("X-GPTScript-Env"))
	}

	for _, pair := range strings.Split(r.Header.Get("X-GPTScript-Env"), ",") {
		key, value, _ := strings.Cut(pair, "=")
		if key == "GPTSCRIPT_WORKSPACE_ID" {
			return value, nil
		}
	}

	return "", fmt.Errorf("GPTSCRIPT_WORKSPACE_ID not found in environment header")
}
