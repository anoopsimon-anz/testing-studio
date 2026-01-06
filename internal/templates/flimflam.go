package templates

func GetFlimFlamHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>FlimFlam Explorer - Testing Studio</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", sans-serif;
            background: #f5f5f5;
            color: #202124;
            height: 100vh;
            overflow: hidden;
        }
        .topbar {
            background: white;
            border-bottom: 1px solid #dadce0;
            padding: 16px 24px;
            display: flex;
            align-items: center;
            gap: 16px;
        }
        .logo {
            font-size: 18px;
            font-weight: 500;
            color: #202124;
            text-decoration: none;
        }
        .back-btn {
            color: #1a73e8;
            padding: 6px 12px;
            border-radius: 4px;
            text-decoration: none;
            font-size: 14px;
            transition: background 0.2s;
        }
        .back-btn:hover { background: #f1f3f4; }

        .container {
            display: flex;
            height: calc(100vh - 60px);
            overflow: hidden;
        }

        /* Left Panel - API List */
        .api-list-panel {
            width: 300px;
            background: #f8f9fa;
            border-right: 1px solid #dadce0;
            display: flex;
            flex-direction: column;
            overflow: hidden;
        }

        .api-list-header {
            padding: 16px;
            border-bottom: 1px solid #dadce0;
            background: white;
        }

        .api-list-title {
            font-size: 14px;
            font-weight: 500;
            color: #202124;
            margin-bottom: 12px;
        }

        .search-box {
            width: 100%;
            padding: 8px 12px;
            border: 1px solid #dadce0;
            border-radius: 4px;
            font-size: 13px;
        }

        .api-list {
            flex: 1;
            overflow-y: auto;
            padding: 8px;
        }

        .api-item {
            padding: 10px 12px;
            margin-bottom: 4px;
            background: white;
            border: 1px solid #dadce0;
            border-radius: 4px;
            cursor: pointer;
            transition: all 0.2s;
            font-size: 13px;
        }

        .api-item:hover {
            background: #e8f0fe;
            border-color: #1a73e8;
        }

        .api-item.active {
            background: #e8f0fe;
            border-color: #1a73e8;
        }

        .api-item-name {
            font-weight: 500;
            color: #202124;
            margin-bottom: 4px;
        }

        .api-item-path {
            font-size: 11px;
            color: #5f6368;
            font-family: Monaco, monospace;
            word-break: break-all;
        }

        .api-item-badge {
            display: inline-block;
            padding: 2px 6px;
            border-radius: 3px;
            font-size: 10px;
            font-weight: 600;
            margin-top: 4px;
        }

        .badge-rest {
            background: #e8f5e9;
            color: #1e8e3e;
        }

        .badge-grpc {
            background: #e8f0fe;
            color: #1967d2;
        }

        /* Right Panel - Editor & Response */
        .main-panel {
            flex: 1;
            display: flex;
            flex-direction: column;
            overflow: hidden;
        }

        .tip-banner {
            background: #e8f0fe;
            border-bottom: 1px solid #d2e3fc;
            padding: 10px 16px;
            font-size: 13px;
            color: #1967d2;
        }

        .editor-section {
            background: white;
            border-bottom: 1px solid #dadce0;
            padding: 16px 20px;
            display: flex;
            flex-direction: column;
            gap: 12px;
        }

        .section-title {
            font-size: 14px;
            font-weight: 500;
            color: #202124;
        }

        .api-info {
            padding: 12px;
            background: #f8f9fa;
            border-radius: 4px;
            font-size: 13px;
            font-family: Monaco, monospace;
            color: #5f6368;
            word-break: break-all;
        }

        .json-editor {
            width: 100%;
            height: 200px;
            padding: 12px;
            border: 1px solid #dadce0;
            border-radius: 4px;
            font-family: Monaco, monospace;
            font-size: 13px;
            resize: vertical;
        }

        .action-buttons {
            display: flex;
            gap: 8px;
        }

        .btn {
            padding: 8px 20px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 13px;
            font-weight: 500;
            transition: background 0.2s;
        }

        .btn-primary {
            background: #1a73e8;
            color: white;
        }

        .btn-primary:hover {
            background: #1557b0;
        }

        .btn-secondary {
            background: white;
            border: 1px solid #dadce0;
            color: #5f6368;
        }

        .btn-secondary:hover {
            background: #f5f5f5;
        }

        .response-section {
            flex: 1;
            display: flex;
            flex-direction: column;
            overflow: hidden;
            background: white;
        }

        .response-header {
            padding: 12px 20px;
            border-bottom: 1px solid #dadce0;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .response-title {
            font-size: 14px;
            font-weight: 500;
            color: #202124;
        }

        .response-stats {
            font-size: 13px;
            color: #5f6368;
        }

        .status-badge {
            display: inline-block;
            padding: 4px 8px;
            border-radius: 3px;
            font-size: 12px;
            font-weight: 600;
            margin-left: 8px;
        }

        .status-200 {
            background: #e8f5e9;
            color: #188038;
        }

        .status-error {
            background: #fce8e6;
            color: #d93025;
        }

        .response-body {
            flex: 1;
            overflow: auto;
            padding: 20px;
        }

        .response-json {
            background: #f8f9fa;
            padding: 16px;
            border-radius: 4px;
            font-family: Monaco, monospace;
            font-size: 12px;
            white-space: pre-wrap;
            word-wrap: break-word;
        }

        .empty-state {
            display: flex;
            align-items: center;
            justify-content: center;
            height: 100%;
            color: #5f6368;
            font-size: 13px;
        }

        .loading {
            color: #5f6368;
            font-size: 13px;
        }

        /* JSON Syntax Highlighting */
        .json-key { color: #881391; }
        .json-string { color: #1a1aa6; }
        .json-number { color: #1c00cf; }
        .json-boolean { color: #0d22aa; }
        .json-null { color: #808080; }
    </style>
</head>
<body>
    <div class="topbar">
        <a href="/" class="logo">Testing Studio</a>
        <a href="/" class="back-btn">← Back</a>
    </div>

    <div class="container">
        <!-- Left Panel: API List -->
        <div class="api-list-panel">
            <div class="api-list-header">
                <div class="api-list-title">FlimFlam APIs</div>
                <input type="text" id="searchBox" class="search-box" placeholder="Search APIs..." onkeyup="filterAPIs()">
            </div>
            <div class="api-list" id="apiList">
                <div class="empty-state">Loading APIs...</div>
            </div>
        </div>

        <!-- Right Panel: Editor & Response -->
        <div class="main-panel">
            <div class="tip-banner">
                💡 <strong>Tip:</strong> FlimFlam mock server running on localhost:9999
            </div>

            <!-- Status Banner -->
            <div id="statusBanner" style="display: none; padding: 12px 16px; font-size: 13px; font-weight: 500; border-bottom: 1px solid;">
                <span id="statusIcon" style="margin-right: 8px; font-size: 16px;"></span>
                <span id="statusMessage"></span>
                <span id="statusEnv" style="margin-left: 8px; padding: 2px 8px; border-radius: 3px; font-size: 11px; font-weight: 600;"></span>
            </div>

            <div class="editor-section">
                <div class="section-title">Request</div>

                <!-- Method + URL Row -->
                <div style="display: flex; gap: 8px; margin-bottom: 12px;">
                    <select id="httpMethod" style="padding: 10px 16px; border: 1px solid #dadce0; border-radius: 4px; font-size: 13px; font-weight: 500; background: white; cursor: pointer; min-width: 100px;">
                        <option value="POST">POST</option>
                        <option value="GET">GET</option>
                        <option value="PUT">PUT</option>
                        <option value="PATCH">PATCH</option>
                        <option value="DELETE">DELETE</option>
                    </select>
                    <input type="text" id="apiUrl" placeholder="http://localhost:9999/path/to/api" style="flex: 1; padding: 10px 12px; border: 1px solid #dadce0; border-radius: 4px; font-size: 13px; font-family: Monaco, monospace;">
                </div>

                <div style="margin-bottom: 8px; font-size: 12px; color: #5f6368; font-weight: 500;">Request Body (JSON)</div>
                <textarea id="requestBody" class="json-editor" placeholder="Enter JSON request body (optional)">{}
</textarea>
                <div class="action-buttons">
                    <button class="btn btn-primary" onclick="sendRequest()">Send Request</button>
                    <button class="btn btn-secondary" onclick="clearRequest()">Clear</button>
                </div>
            </div>

            <div class="response-section">
                <div class="response-header">
                    <div class="response-title">Response</div>
                    <div class="response-stats">
                        <span id="responseStatus"></span>
                        <span id="responseTime"></span>
                    </div>
                </div>
                <div class="response-body" id="responseBody">
                    <div class="empty-state">Send a request to see the response</div>
                </div>
            </div>
        </div>
    </div>

    <script>
        let allAPIs = [];
        let selectedAPI = null;

        // Check FlimFlam status (APP_ENV)
        async function checkFlimFlamStatus() {
            try {
                const response = await fetch('/api/flimflam/status');
                const data = await response.json();

                const banner = document.getElementById('statusBanner');
                const icon = document.getElementById('statusIcon');
                const message = document.getElementById('statusMessage');
                const envBadge = document.getElementById('statusEnv');

                if (data.error) {
                    banner.style.display = 'block';
                    banner.style.background = '#fff3cd';
                    banner.style.borderColor = '#ffc107';
                    banner.style.color = '#856404';
                    icon.textContent = '⚠️';
                    message.textContent = data.error;
                    envBadge.textContent = 'UNKNOWN';
                    envBadge.style.background = '#ffc107';
                    envBadge.style.color = '#856404';
                } else if (data.isLocal) {
                    banner.style.display = 'block';
                    banner.style.background = '#d4edda';
                    banner.style.borderColor = '#c3e6cb';
                    banner.style.color = '#155724';
                    icon.textContent = '✅';
                    message.textContent = data.message;
                    envBadge.textContent = 'APP_ENV=' + data.appEnv.toUpperCase();
                    envBadge.style.background = '#28a745';
                    envBadge.style.color = 'white';
                } else {
                    banner.style.display = 'block';
                    banner.style.background = '#f8d7da';
                    banner.style.borderColor = '#f5c6cb';
                    banner.style.color = '#721c24';
                    icon.textContent = '❌';
                    message.textContent = data.message;
                    envBadge.textContent = 'APP_ENV=' + (data.appEnv || 'NOT SET').toUpperCase();
                    envBadge.style.background = '#dc3545';
                    envBadge.style.color = 'white';
                }
            } catch (error) {
                console.error('Failed to check FlimFlam status:', error);
            }
        }

        // Load APIs on page load
        async function loadAPIs() {
            try {
                const response = await fetch('/api/flimflam/apis');
                const data = await response.json();

                if (data.error) {
                    document.getElementById('apiList').innerHTML = '<div class="empty-state">Error: ' + data.error + '</div>';
                    return;
                }

                allAPIs = data.apis || [];
                renderAPIs(allAPIs);
            } catch (error) {
                document.getElementById('apiList').innerHTML = '<div class="empty-state">Error loading APIs: ' + error.message + '</div>';
            }
        }

        function renderAPIs(apis) {
            const apiList = document.getElementById('apiList');

            if (apis.length === 0) {
                apiList.innerHTML = '<div class="empty-state">No APIs found</div>';
                return;
            }

            apiList.innerHTML = '';

            apis.forEach(api => {
                const item = document.createElement('div');
                item.className = 'api-item';
                item.onclick = () => selectAPI(api);

                const badge = api.isGrpc ? 'badge-grpc' : 'badge-rest';
                const badgeText = api.isGrpc ? 'gRPC' : 'REST';

                item.innerHTML = '<div class="api-item-name">' + api.name + '</div>' +
                                '<div class="api-item-path">' + api.path + '</div>' +
                                '<span class="api-item-badge ' + badge + '">' + badgeText + '</span>';

                apiList.appendChild(item);
            });
        }

        function filterAPIs() {
            const searchTerm = document.getElementById('searchBox').value.toLowerCase();
            const filtered = allAPIs.filter(api =>
                api.name.toLowerCase().includes(searchTerm) ||
                api.path.toLowerCase().includes(searchTerm)
            );
            renderAPIs(filtered);
        }

        function selectAPI(api) {
            selectedAPI = api;

            // Update active state
            document.querySelectorAll('.api-item').forEach(item => item.classList.remove('active'));
            event.currentTarget.classList.add('active');

            // Update method dropdown
            document.getElementById('httpMethod').value = api.method || 'POST';

            // Update URL text box
            document.getElementById('apiUrl').value = 'http://localhost:9999' + api.path;

            // Set default body
            document.getElementById('requestBody').value = '{}';

            // Clear response
            document.getElementById('responseBody').innerHTML = '<div class="empty-state">Send a request to see the response</div>';
            document.getElementById('responseStatus').innerHTML = '';
            document.getElementById('responseTime').textContent = '';
        }

        function syntaxHighlightJSON(json) {
            if (typeof json !== 'string') {
                json = JSON.stringify(json, null, 2);
            }

            json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');

            return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
                let cls = 'json-number';
                if (/^"/.test(match)) {
                    if (/:$/.test(match)) {
                        cls = 'json-key';
                    } else {
                        cls = 'json-string';
                    }
                } else if (/true|false/.test(match)) {
                    cls = 'json-boolean';
                } else if (/null/.test(match)) {
                    cls = 'json-null';
                }
                return '<span class="' + cls + '">' + match + '</span>';
            });
        }

        async function sendRequest() {
            const method = document.getElementById('httpMethod').value;
            const url = document.getElementById('apiUrl').value.trim();

            if (!url) {
                alert('Please enter a URL');
                return;
            }

            const bodyText = document.getElementById('requestBody').value.trim();
            let body = {};

            if (bodyText) {
                try {
                    body = JSON.parse(bodyText);
                } catch (e) {
                    alert('Invalid JSON in request body: ' + e.message);
                    return;
                }
            }

            // Extract path from URL
            const path = url.replace('http://localhost:9999', '');

            const startTime = Date.now();

            try {
                const response = await fetch('/api/flimflam/send', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        path: path,
                        body: body,
                        method: method,
                        contentType: selectedAPI?.contentType || 'application/json'
                    })
                });

                const endTime = Date.now();
                const duration = endTime - startTime;

                const result = await response.json();

                // Update response status
                const statusCode = result.statusCode || 0;
                const statusClass = statusCode >= 200 && statusCode < 300 ? 'status-200' : 'status-error';
                document.getElementById('responseStatus').innerHTML = '<span class="status-badge ' + statusClass + '">' + statusCode + '</span>';
                document.getElementById('responseTime').textContent = duration + ' ms';

                // Update response body with syntax highlighting
                const responseBody = document.getElementById('responseBody');
                if (result.error) {
                    responseBody.innerHTML = '<div class="response-json" style="color: #d93025;">Error: ' + result.error + '</div>';
                } else {
                    const highlighted = syntaxHighlightJSON(result.body);
                    responseBody.innerHTML = '<pre class="response-json">' + highlighted + '</pre>';
                }

            } catch (error) {
                document.getElementById('responseStatus').innerHTML = '<span class="status-badge status-error">ERROR</span>';
                document.getElementById('responseTime').textContent = '';
                document.getElementById('responseBody').innerHTML = '<div class="response-json" style="color: #d93025;">Request failed: ' + error.message + '</div>';
            }
        }

        function clearRequest() {
            document.getElementById('requestBody').value = '{}';
        }

        // Load status and APIs on page load
        checkFlimFlamStatus();
        loadAPIs();
    </script>
</body>
</html>`
}
