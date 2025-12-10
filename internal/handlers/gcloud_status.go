package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

type GCloudStatusResponse struct {
	Authenticated bool   `json:"authenticated"`
	LastLoginTime string `json:"lastLoginTime"`
	Account       string `json:"account"`
}

func HandleGCloudStatus(w http.ResponseWriter, r *http.Request) {
	response := GCloudStatusResponse{
		Authenticated: false,
		LastLoginTime: "",
		Account:       "",
	}

	// Check if user is authenticated by getting active account
	cmd := exec.Command("gcloud", "auth", "list", "--filter=status:ACTIVE", "--format=value(account)")
	output, err := cmd.Output()
	if err == nil && len(output) > 0 {
		response.Authenticated = true
		response.Account = strings.TrimSpace(string(output))
	}

	// Get token expiry time to determine last login
	if response.Authenticated {
		cmd = exec.Command("gcloud", "auth", "print-access-token", "--format=json")
		tokenOutput, err := cmd.Output()
		if err == nil {
			// Try to get credential info
			cmd = exec.Command("gcloud", "auth", "application-default", "print-access-token")
			_, tokenErr := cmd.Output()

			if tokenErr == nil && len(tokenOutput) > 0 {
				// Token is valid - get the credential file modification time
				// Try macOS stat first, then Linux stat
				cmd = exec.Command("bash", "-c", "stat -f '%Sm' -t '%Y-%m-%d %H:%M' ~/.config/gcloud/access_tokens.db 2>/dev/null || stat -c '%y' ~/.config/gcloud/access_tokens.db 2>/dev/null | cut -d'.' -f1 | cut -d' ' -f1,2")
				timeOutput, timeErr := cmd.Output()
				if timeErr == nil && len(timeOutput) > 0 {
					response.LastLoginTime = strings.TrimSpace(string(timeOutput))
				} else {
					// Fallback: just show "Active"
					response.LastLoginTime = "Active"
				}
			}
		}
	}

	// If we couldn't get the exact time, try checking config modification time
	if response.Authenticated && response.LastLoginTime == "" {
		cmd = exec.Command("bash", "-c", "stat -f '%Sm' -t '%Y-%m-%d %H:%M' ~/.config/gcloud/configurations/config_default 2>/dev/null || stat -c '%y' ~/.config/gcloud/configurations/config_default 2>/dev/null | cut -d'.' -f1 | cut -d' ' -f1,2")
		output, err := cmd.Output()
		if err == nil && len(output) > 0 {
			response.LastLoginTime = strings.TrimSpace(string(output))
		} else {
			// Final fallback - show Active
			response.LastLoginTime = "Active"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
