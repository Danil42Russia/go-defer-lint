package target

import (
	"io"
	"net/http"
	"os"
)

func deferHttpRequest() {
	req, err := http.NewRequest("GET", "http://localhost", http.NoBody) // want `\Qcan rewrite as 'defer req.Body.Close()'`
	if err != nil {
	}
	req.Body.Close()
	_, _ = io.ReadAll(req.Body)
}

func deferHttpGet() {
	resp, err := http.Get("http://localhost") // want `\Qcan rewrite as 'defer resp.Body.Close()'`
	if err != nil {
	}
	resp.Body.Close()
	_, _ = io.ReadAll(resp.Body)
}

func deferPostGet() {
	resp, err := http.Post("http://localhost", "", http.NoBody) // want `\Qcan rewrite as 'defer resp.Body.Close()'`
	if err != nil {
	}
	resp.Body.Close()
	_, _ = io.ReadAll(resp.Body)
}

func deferOsOpen() {
	file, err := os.Open("test") // want `\Qcan rewrite as 'defer file.Close()'`
	if err != nil {
	}

	file.Name()
}

func deferOsCleate() {
	file, err := os.Create("test") // want `\Qcan rewrite as 'defer file.Close()'`
	if err != nil {
	}

	file.Name()
}
