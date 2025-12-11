# Testing Studio - Playwright Tests

This directory contains end-to-end tests for the Testing Studio application using Playwright for Go.

## Setup

1. Install Playwright for Go:
```bash
go get github.com/playwright-community/playwright-go
```

2. Install Playwright browsers:
```bash
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps chromium
```

## Running Tests

1. Start the Testing Studio server:
```bash
go run cmd/server/main.go
```

Or use the built binary:
```bash
./start.sh
```

2. Run the basic tests:
```bash
cd tests
go test -v
```

3. Run tests with HTML report and screenshots:
```bash
cd tests
go test -v -run TestStudioWithReport
```

This will generate:
- `test-report.html` - Beautiful HTML report with test results
- `screenshots/` - Directory with screenshots of each test step

## Test Coverage

### TestStudioHomePage
Tests the main landing page of Testing Studio:

- ✅ Verifies page title is "Testing Studio"
- ✅ Verifies subtitle text
- ✅ Verifies all 5 option cards are visible:
  - Google PubSub
  - Kafka / EventMesh
  - REST Client
  - GCS Browser
  - Trace Journey Viewer
- ✅ Verifies Tools button is present
- ✅ Verifies Docker status indicator is present
- ✅ Verifies GCloud status indicator is present
- ✅ Tests Tools menu interaction

## UI Element IDs

All interactive and readable UI elements have been assigned unique IDs for testing:

### Main Page Elements
- `pageTitle` - Main heading "Testing Studio"
- `pageSubtitle` - Subtitle text
- `toolsButton` - Tools dropdown button
- `toolsMenu` - Tools menu container
- `dockerStatus` - Docker status indicator
- `gcloudStatus` - GCloud status indicator

### Option Cards
- `cardPubsub` - PubSub card
- `cardKafka` - Kafka card
- `cardRestClient` - REST Client card
- `cardGCS` - GCS Browser card
- `cardTraceJourney` - Trace Journey card

### Status Indicators
- `dockerStatusDot` - Docker status dot
- `dockerStatusText` - Docker status text
- `gcloudStatusDot` - GCloud status dot
- `gcloudText` - GCloud status text

### Menu Links
- `linkConfigEditor` - Configuration Editor link
- `linkFlowDiagram` - Flow Diagram link
- `linkBase64Tool` - Base64 Tool link

### TestStudioWithReport
Advanced test with HTML report generation and screenshots:

- ✅ Page Title Verification
- ✅ Subtitle Verification
- ✅ All 5 Option Cards Verification (with individual screenshots)
- ✅ PubSub Card Details
- ✅ Status Indicators Verification
- ✅ Tools Menu Interaction (before and after screenshots)
- ✅ Individual Card Analysis (5 separate screenshots)

**Report Features:**
- Beautiful HTML report with gradient header
- Summary statistics (Total/Passed/Failed)
- Individual test results with status icons
- Full-page screenshots for each test
- Click-to-enlarge screenshots
- Responsive design

## Sample Output

```
=== RUN   TestStudioWithReport
    ✅ HTML report generated: test-report.html
    ✅ Test completed! Screenshots saved in screenshots/
    📊 Total tests: 11, Passed: 10, Failed: 1
--- PASS: TestStudioWithReport (5.68s)
```

## Generated Files

After running `TestStudioWithReport`:

```
screenshots/
├── 01-homepage.png              (76KB) - Full homepage
├── 02-pubsub-card.png          (80KB) - PubSub card highlighted
├── 03-status-indicators.png    (80KB) - Status indicators
├── 04-tools-menu-closed.png    (80KB) - Tools menu closed
├── 05-tools-menu-open.png      (77KB) - Tools menu open
├── 06-card-1-cardPubsub.png    (80KB) - PubSub card detail
├── 06-card-2-cardKafka.png     (82KB) - Kafka card detail
├── 06-card-3-cardRestClient.png (82KB) - REST Client card detail
├── 06-card-4-cardGCS.png       (82KB) - GCS card detail
└── 06-card-5-cardTraceJourney.png (83KB) - Trace Journey card detail

test-report.html                 - Interactive HTML report (301 lines)
```

## Viewing the Report

Open the HTML report in your browser:
```bash
open test-report.html
```

Or on Linux:
```bash
xdg-open test-report.html
```

The report includes:
- Color-coded test results (green for pass, red for fail)
- Detailed test messages
- Embedded screenshots that can be clicked to view full-size
- Professional styling with gradient headers

## Notes

- Tests run in headless mode by default
- Server must be running on `localhost:8888` before running tests
- Tests use a 30-second timeout for page operations
- Screenshots are in PNG format at 1920x1080 resolution
- Report HTML is self-contained with inline CSS
