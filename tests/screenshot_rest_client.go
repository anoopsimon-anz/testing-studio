package tests

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

func TestRestClientScreenshot(t *testing.T) {
	// Start Playwright
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	// Launch browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	// Create new page
	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}

	// Navigate to REST client
	if _, err = page.Goto("http://localhost:8888/rest-client"); err != nil {
		t.Fatalf("could not goto rest-client: %v", err)
	}

	// Wait for page to load
	page.WaitForTimeout(1000)

	// Take screenshot
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("screenshots/rest-client-new.png"),
	}); err != nil {
		t.Fatalf("could not take screenshot: %v", err)
	}

	t.Log("Screenshot saved to screenshots/rest-client-new.png")
}
