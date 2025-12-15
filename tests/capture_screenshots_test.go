package tests

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

func TestCaptureScreenshots(t *testing.T) {
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

	// Create new page with viewport
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		Viewport: &playwright.Size{
			Width:  1400,
			Height: 900,
		},
	})
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}

	// Capture main page
	if _, err = page.Goto("http://localhost:8888"); err != nil {
		t.Fatalf("could not goto homepage: %v", err)
	}
	page.WaitForTimeout(1000)
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("../screenshots/ui-homepage.png"),
	}); err != nil {
		t.Fatalf("could not take homepage screenshot: %v", err)
	}
	t.Log("✅ Captured homepage screenshot")

	// Capture settings/config page
	if _, err = page.Goto("http://localhost:8888/config-editor"); err != nil {
		t.Fatalf("could not goto config page: %v", err)
	}
	page.WaitForTimeout(1000)
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("../screenshots/ui-settings.png"),
	}); err != nil {
		t.Fatalf("could not take settings screenshot: %v", err)
	}
	t.Log("✅ Captured settings screenshot")

	// Capture REST client
	if _, err = page.Goto("http://localhost:8888/rest-client"); err != nil {
		t.Fatalf("could not goto rest-client: %v", err)
	}
	page.WaitForTimeout(1000)
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("../screenshots/ui-rest-client.png"),
	}); err != nil {
		t.Fatalf("could not take rest-client screenshot: %v", err)
	}
	t.Log("✅ Captured REST client screenshot")

	t.Log("🎉 All screenshots captured successfully!")
}
