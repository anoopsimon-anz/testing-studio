package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"cloudevents-explorer/internal/templates"
)

type FlimFlamAPI struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	ContentType string `json:"contentType"`
	IsGRPC      bool   `json:"isGrpc"`
}

func FlimFlamExplorerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(templates.GetFlimFlamHTML()))
}

func FlimFlamAPIsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Docker logs from flimflam container
	cmd := exec.Command("docker", "logs", "devstack-dep_flimflam-1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Failed to get Docker logs: " + err.Error(),
			"apis":  []FlimFlamAPI{},
		})
		return
	}

	apis := parseFlimFlamFixtures(string(output))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"apis": apis,
	})
}

func FlimFlamStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get APP_ENV from flimflam container
	cmd := exec.Command("docker", "inspect", "devstack-dep_flimflam-1", "--format", "{{range .Config.Env}}{{println .}}{{end}}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":      "Failed to inspect container: " + err.Error(),
			"appEnv":     "",
			"isLocal":    false,
			"apiEnabled": false,
		})
		return
	}

	// Parse APP_ENV from output
	appEnv := ""
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "APP_ENV=") {
			appEnv = strings.TrimPrefix(line, "APP_ENV=")
			break
		}
	}

	isLocal := appEnv == "local"

	json.NewEncoder(w).Encode(map[string]interface{}{
		"appEnv":     appEnv,
		"isLocal":    isLocal,
		"apiEnabled": isLocal,
		"message":    getMessage(appEnv, isLocal),
	})
}

func getMessage(appEnv string, isLocal bool) string {
	if isLocal {
		return "FlimFlam is running in LOCAL mode - All mock APIs are registered and will return 200 OK responses"
	}
	if appEnv == "" {
		return "Unable to determine FlimFlam environment - Mock APIs may not work"
	}
	return "FlimFlam is running in " + strings.ToUpper(appEnv) + " mode - Mock fixtures are NOT registered. APIs will fail!"
}

func FlimFlamProxyHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Path        string                 `json:"path"`
		Body        map[string]interface{} `json:"body"`
		Method      string                 `json:"method"`
		ContentType string                 `json:"contentType"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Default to POST if not specified
	if requestData.Method == "" {
		requestData.Method = "POST"
	}

	// Default content type
	if requestData.ContentType == "" {
		requestData.ContentType = "application/json"
	}

	// Build URL
	url := "http://localhost:9999" + requestData.Path

	// Marshal body to JSON
	bodyJSON, err := json.Marshal(requestData.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create request
	req, err := http.NewRequest(requestData.Method, url, strings.NewReader(string(bodyJSON)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", requestData.ContentType)

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":      err.Error(),
			"statusCode": 0,
		})
		return
	}
	defer resp.Body.Close()

	// Read response
	var responseBody interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		// If not JSON, return as string
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":      "Failed to parse response: " + err.Error(),
			"statusCode": resp.StatusCode,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": resp.StatusCode,
		"body":       responseBody,
	})
}

func parseFlimFlamFixtures(logOutput string) []FlimFlamAPI {
	var apis []FlimFlamAPI

	// Regex to match: storing global fixture: {test:"" trace_id:"" path:"PATH" ...}
	re := regexp.MustCompile(`storing global fixture: \{test:"" trace_id:"" path:"([^"]+)"(?:\s+content_type:"([^"]+)")?(?:\s+method:"([^"]+)")?\}`)

	matches := re.FindAllStringSubmatch(logOutput, -1)

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		path := match[1]
		contentType := ""
		method := "POST"

		if len(match) > 2 && match[2] != "" {
			contentType = match[2]
		}
		if len(match) > 3 && match[3] != "" {
			method = match[3]
		}

		isGRPC := strings.Contains(contentType, "grpc")

		// Generate friendly name from path
		name := generateFriendlyName(path, isGRPC)

		apis = append(apis, FlimFlamAPI{
			Name:        name,
			Path:        path,
			Method:      method,
			ContentType: contentType,
			IsGRPC:      isGRPC,
		})
	}

	return apis
}

func generateFriendlyName(path string, isGRPC bool) string {
	if isGRPC {
		// Extract service/method from gRPC path
		// e.g., /fabric.service.entitlements.v1.EntitlementsControlAPI/RegisterPersonaToParty
		parts := strings.Split(path, "/")
		if len(parts) >= 2 {
			return parts[len(parts)-1] // Just the method name
		}
		return path
	}

	// For REST APIs, use the last meaningful part of the path
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		// Get last non-empty part
		for i := len(parts) - 1; i >= 0; i-- {
			if parts[i] != "" {
				return parts[i]
			}
		}
	}

	return path
}
