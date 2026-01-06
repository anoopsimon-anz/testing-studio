package templates

const Index = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Testing Studio</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Google Sans', 'Product Sans', -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", sans-serif;
            background: #fafafa;
            color: #202124;
            min-height: 100vh;
            padding: 80px 24px 24px 24px;
        }
        .landing {
            max-width: 1000px;
            margin: 0 auto;
        }
        .hero {
            text-align: center;
            margin-bottom: 64px;
        }
        h1 {
            font-size: 48px;
            color: #202124;
            margin-bottom: 12px;
            font-weight: 400;
            letter-spacing: -0.5px;
        }
        .subtitle {
            font-size: 15px;
            color: #5f6368;
            font-weight: 400;
        }
        .options {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
            gap: 24px;
            margin-top: 48px;
        }
        .option-card {
            background: white;
            border: 1px solid #e8eaed;
            border-radius: 12px;
            padding: 32px 24px;
            cursor: pointer;
            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
            text-decoration: none;
            color: inherit;
            display: flex;
            flex-direction: column;
            position: relative;
            overflow: hidden;
        }
        .option-card::before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            height: 3px;
            background: linear-gradient(90deg, #4285f4, #34a853, #fbbc04, #ea4335);
            opacity: 0;
            transition: opacity 0.3s;
        }
        .option-card:hover {
            transform: translateY(-4px);
            box-shadow: 0 8px 24px rgba(0,0,0,0.12);
            border-color: #dadce0;
        }
        .option-card:hover::before {
            opacity: 1;
        }
        .option-title {
            font-size: 18px;
            font-weight: 500;
            color: #202124;
            margin-bottom: 8px;
        }
        .option-desc {
            font-size: 14px;
            color: #5f6368;
            line-height: 1.6;
            flex: 1;
        }
        .badge {
            display: inline-block;
            background: #e8f0fe;
            color: #1967d2;
            padding: 4px 8px;
            border-radius: 4px;
            font-size: 11px;
            font-weight: 500;
            margin-top: 12px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        .status-indicators {
            position: absolute;
            top: 20px;
            right: 20px;
            display: flex;
            flex-direction: column;
            gap: 8px;
        }
        .status-indicator {
            display: flex;
            align-items: center;
            gap: 6px;
            font-size: 12px;
            color: #5f6368;
            font-weight: 500;
        }
        .status-dot {
            width: 10px;
            height: 10px;
            border-radius: 50%;
            box-shadow: 0 0 0 2px rgba(0,0,0,0.1);
        }
        .status-dot.green {
            background: #34a853;
        }
        .status-dot.red {
            background: #ea4335;
        }
        .tools-wrapper {
            position: absolute;
            top: 20px;
            left: 20px;
            z-index: 50;
        }
        .tools-btn {
            background: white;
            border: 1px solid #dadce0;
            padding: 8px 16px;
            border-radius: 20px;
            cursor: pointer;
            color: #5f6368;
            font-size: 14px;
            font-weight: 500;
            transition: all 0.2s;
        }
        .tools-btn:hover {
            background: #f8f9fa;
            border-color: #dadce0;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
        }
        .tools-menu {
            display: none;
            position: absolute;
            left: 0;
            top: 48px;
            background: white;
            border: 1px solid #dadce0;
            border-radius: 8px;
            box-shadow: 0 2px 8px rgba(0,0,0,0.1);
            min-width: 200px;
            overflow: hidden;
        }
        .tools-menu a {
            display: block;
            padding: 12px 16px;
            color: #202124;
            text-decoration: none;
            font-size: 14px;
            transition: background 0.2s;
        }
        .tools-menu a:hover {
            background: #f8f9fa;
        }
        .tools-menu a:not(:last-child) {
            border-bottom: 1px solid #e8eaed;
        }
        .settings-icon {
            position: fixed;
            top: 80px;
            right: 20px;
            z-index: 40;
            background: white;
            border: 2px solid #dadce0;
            border-radius: 50%;
            width: 56px;
            height: 56px;
            display: flex;
            align-items: center;
            justify-content: center;
            cursor: pointer;
            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
            box-shadow: 0 2px 8px rgba(0,0,0,0.1);
            color: #5f6368;
            text-decoration: none;
        }
        .settings-icon:hover {
            background: #1a73e8;
            border-color: #1a73e8;
            color: white;
            box-shadow: 0 4px 16px rgba(26,115,232,0.3);
            transform: rotate(90deg);
        }
        .settings-icon svg {
            transition: transform 0.3s;
        }
    </style>
</head>
<body>
    <div class="tools-wrapper" id="toolsWrapper">
        <button class="tools-btn" id="toolsButton" onclick="toggleToolsMenu()">Tools ▼</button>
        <div class="tools-menu" id="toolsMenu">
            <a href="/flow-diagram" id="linkFlowDiagram">Communications - Event Handling</a>
            <a href="#" id="linkBase64Tool" onclick="openBase64Tool(); toggleToolsMenu(); return false;">Base64 Encoder/Decoder</a>
            <a href="#" id="linkTOONTool" onclick="openTOONTool(); toggleToolsMenu(); return false;">JSON to TOON Converter</a>
        </div>
    </div>
    <div class="status-indicators" id="statusIndicators">
        <div class="status-indicator" id="dockerStatus">
            <div class="status-dot red" id="dockerStatusDot"></div>
            <span id="dockerStatusText">Docker</span>
        </div>
        <div class="status-indicator" id="gcloudStatus">
            <div class="status-dot red" id="gcloudStatusDot"></div>
            <span id="gcloudText">GCloud</span>
        </div>
    </div>
    <a href="/config-editor" class="settings-icon" id="settingsIcon" title="Global Settings - Configure all connections">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
            <path d="M19.14,12.94c0.04-0.3,0.06-0.61,0.06-0.94c0-0.32-0.02-0.64-0.07-0.94l2.03-1.58c0.18-0.14,0.23-0.41,0.12-0.61 l-1.92-3.32c-0.12-0.22-0.37-0.29-0.59-0.22l-2.39,0.96c-0.5-0.38-1.03-0.7-1.62-0.94L14.4,2.81c-0.04-0.24-0.24-0.41-0.48-0.41 h-3.84c-0.24,0-0.43,0.17-0.47,0.41L9.25,5.35C8.66,5.59,8.12,5.92,7.63,6.29L5.24,5.33c-0.22-0.08-0.47,0-0.59,0.22L2.74,8.87 C2.62,9.08,2.66,9.34,2.86,9.48l2.03,1.58C4.84,11.36,4.8,11.69,4.8,12s0.02,0.64,0.07,0.94l-2.03,1.58 c-0.18,0.14-0.23,0.41-0.12,0.61l1.92,3.32c0.12,0.22,0.37,0.29,0.59,0.22l2.39-0.96c0.5,0.38,1.03,0.7,1.62,0.94l0.36,2.54 c0.05,0.24,0.24,0.41,0.48,0.41h3.84c0.24,0,0.44-0.17,0.47-0.41l0.36-2.54c0.59-0.24,1.13-0.56,1.62-0.94l2.39,0.96 c0.22,0.08,0.47,0,0.59-0.22l1.92-3.32c0.12-0.22,0.07-0.47-0.12-0.61L19.14,12.94z M12,15.6c-1.98,0-3.6-1.62-3.6-3.6 s1.62-3.6,3.6-3.6s3.6,1.62,3.6,3.6S13.98,15.6,12,15.6z"/>
        </svg>
    </a>
    <div class="landing" id="landing">
        <div class="hero" id="hero">
            <h1 id="pageTitle">Testing Studio</h1>
            <p class="subtitle" id="pageSubtitle">Requires TMS Suncorp devstack to be running</p>
        </div>

        <div class="options" id="optionsGrid">
            <a href="/pubsub" class="option-card" id="cardPubsub">
                <div class="option-title" id="titlePubsub">Google PubSub</div>
                <div class="option-desc" id="descPubsub">View CloudEvents from Google Cloud PubSub subscriptions</div>
                <span class="badge" id="badgePubsub">CloudEvents</span>
            </a>

            <a href="/kafka" class="option-card" id="cardKafka">
                <div class="option-title" id="titleKafka">Kafka / EventMesh</div>
                <div class="option-desc" id="descKafka">Consume Avro messages from Kafka topics</div>
                <span class="badge" id="badgeKafka">Avro Schema</span>
            </a>

            <a href="/rest-client" class="option-card" id="cardRestClient">
                <div class="option-title" id="titleRestClient">REST Client</div>
                <div class="option-desc" id="descRestClient">Send HTTP requests with custom headers, body, and TLS certificates</div>
                <span class="badge" id="badgeRestClient">API Testing</span>
            </a>

            <a href="/gcs" class="option-card" id="cardGCS">
                <div class="option-title" id="titleGCS">GCS Browser</div>
                <div class="option-desc" id="descGCS">Browse Google Cloud Storage buckets and files with preview and download</div>
                <span class="badge" id="badgeGCS">Storage</span>
            </a>

            <a href="/trace-journey" class="option-card" id="cardTraceJourney">
                <div class="option-title" id="titleTraceJourney">Trace Journey Viewer</div>
                <div class="option-desc" id="descTraceJourney">Track requests across containers with trace IDs and visualize the journey</div>
                <span class="badge" id="badgeTraceJourney">Debugging</span>
            </a>

            <a href="/spanner" class="option-card" id="cardSpanner">
                <div class="option-title" id="titleSpanner">Spanner Explorer</div>
                <div class="option-desc" id="descSpanner">Browse tables, run SQL queries, test local Spanner emulator</div>
                <span class="badge" id="badgeSpanner">Database</span>
            </a>

            <a href="/flimflam-explorer" class="option-card" id="cardFlimFlam">
                <div class="option-title" id="titleFlimFlam">FlimFlam Explorer</div>
                <div class="option-desc" id="descFlimFlam">Test FlimFlam mock APIs (REST & gRPC) running on localhost:9999</div>
                <span class="badge" id="badgeFlimFlam">Mock APIs</span>
            </a>
        </div>
    </div>

    ` + Base64Modal + `
    ` + TOONModal + `

    <script>
        function toggleToolsMenu() {
            const menu = document.getElementById('toolsMenu');
            menu.style.display = menu.style.display === 'none' ? 'block' : 'none';
        }

        // Close menu when clicking outside
        document.addEventListener('click', function(e) {
            const menu = document.getElementById('toolsMenu');
            const button = e.target.closest('button');
            if (menu && !menu.contains(e.target) && (!button || button.textContent.indexOf('Tools') === -1)) {
                menu.style.display = 'none';
            }
        });

        // Check Docker status
        async function checkDockerStatus() {
            try {
                const response = await fetch('/api/docker/status');
                const data = await response.json();
                const dot = document.getElementById('dockerStatusDot');

                if (data.running) {
                    dot.classList.remove('red');
                    dot.classList.add('green');
                } else {
                    dot.classList.remove('green');
                    dot.classList.add('red');
                }
            } catch (error) {
                console.error('Failed to check Docker status:', error);
            }
        }

        // Check Docker status on page load
        checkDockerStatus();

        // Check Docker status every 10 seconds
        setInterval(checkDockerStatus, 10000);

        // Check GCloud authentication status
        async function checkGCloudStatus() {
            try {
                const response = await fetch('/api/gcloud/status');
                const data = await response.json();
                const dot = document.getElementById('gcloudStatusDot');
                const textSpan = document.getElementById('gcloudText');

                if (data.authenticated) {
                    dot.classList.remove('red');
                    dot.classList.add('green');

                    // Show last login time if available
                    if (data.lastLoginTime) {
                        textSpan.textContent = 'GCloud: ' + data.lastLoginTime;
                        textSpan.title = 'Account: ' + data.account;
                    } else {
                        textSpan.textContent = 'GCloud: Active';
                        textSpan.title = 'Account: ' + data.account;
                    }
                } else {
                    dot.classList.remove('green');
                    dot.classList.add('red');
                    textSpan.textContent = 'GCloud: Not authenticated';
                    textSpan.title = 'Run: gcloud auth login';
                }
            } catch (error) {
                console.error('Failed to check GCloud status:', error);
            }
        }

        // Check GCloud status on page load
        checkGCloudStatus();

        // Check GCloud status every 30 seconds
        setInterval(checkGCloudStatus, 30000);

        ` + Base64ModalJS + `
        ` + TOONModalJS + `
    </script>
</body>
</html>`
