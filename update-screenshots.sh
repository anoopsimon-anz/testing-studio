#!/bin/bash
set -e

echo "📸 Updating README screenshots..."

# Step 1: Run screenshot capture test
echo "1️⃣ Capturing screenshots..."
go test ./tests -run TestCaptureScreenshots -v

# Step 2: Update README with screenshots
echo "2️⃣ Updating README.md..."
cat > README.md << 'EOF'
# Testing Studio

A web-based tool for testing local cloud services - PubSub, Kafka, REST APIs, GCS, Spanner, and trace journey tracking.

## UI Screenshots

### Homepage
![Homepage](screenshots/ui-homepage.png)

### Settings / Config Editor
![Settings](screenshots/ui-settings.png)

### REST Client (Postman-style)
![REST Client](screenshots/ui-rest-client.png)

## Features

- **Google PubSub** - Pull and view CloudEvents from subscriptions
- **Kafka / EventMesh** - Consume and publish Avro messages
- **REST Client** - Send HTTP requests with collections (Postman-style), TLS certs, and JSON syntax highlighting
- **GCS Browser** - Browse buckets, preview files, and download
- **Spanner Explorer** - Query databases and browse tables
- **Trace Journey Viewer** - Track requests across containers with trace IDs

## Prerequisites

**Requires TMS Suncorp devstack to be running**

```bash
make devstack.start
```

## How to Run

```bash
# Build and run
go build -o testing-studio cmd/server/main.go
./testing-studio

# Or run directly
go run cmd/server/main.go
```

Open **http://localhost:8888** in your browser.

---

Made for TMS Suncorp local development
EOF

# Step 3: Commit changes
echo "3️⃣ Committing changes..."
git add screenshots/*.png README.md
if git diff --cached --quiet; then
    echo "⚠️  No changes to commit (screenshots and README already up to date)"
else
    git commit -m "$(cat <<'COMMIT_EOF'
Update README with UI screenshots

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude Sonnet 4.5 <noreply@anthropic.com>
COMMIT_EOF
)"
    echo "✅ Changes committed!"
fi

echo ""
echo "✅ Done! Screenshots captured and README updated."
echo "To push to remote, run: git push origin main"
