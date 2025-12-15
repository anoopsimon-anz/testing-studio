package templates

const RestClient = `
<div style="background: #e8f0fe; border: 1px solid #d2e3fc; border-radius: 4px; padding: 10px 16px; margin-bottom: 16px; font-size: 13px; color: #1967d2;">
    💡 <strong>Tip:</strong> You can manage saved requests and environments in <a href="/config-editor" style="color: #1a73e8; text-decoration: underline; font-weight: 500;">Global Settings</a>
</div>

<div style="background: white; border-radius: 8px; border: 1px solid #dadce0; overflow: hidden;">
    <!-- Top Bar: Request Name -->
    <div style="padding: 16px 20px; border-bottom: 1px solid #e8eaed; display: flex; align-items: center; gap: 12px;">
        <span style="color: #5f6368; font-size: 13px; font-weight: 500;">GET</span>
        <input type="text" id="requestName" value="Untitled Request"
               style="border: none; font-size: 15px; font-weight: 500; color: #202124; flex: 1; outline: none;" />
        <button onclick="saveRequest()" style="padding: 6px 16px; background: white; border: 1px solid #dadce0; border-radius: 4px; cursor: pointer; font-size: 13px; color: #5f6368;">Save</button>
    </div>

    <!-- URL Bar -->
    <div style="padding: 16px 20px; border-bottom: 1px solid #e8eaed; display: flex; gap: 8px; align-items: center;">
        <select id="httpMethod"
                style="padding: 10px 16px; border: 1px solid #dadce0; border-radius: 4px; font-size: 14px; font-weight: 500; color: #202124; background: white; cursor: pointer;">
            <option value="GET">GET</option>
            <option value="POST">POST</option>
            <option value="PUT">PUT</option>
            <option value="PATCH">PATCH</option>
            <option value="DELETE">DELETE</option>
            <option value="HEAD">HEAD</option>
            <option value="OPTIONS">OPTIONS</option>
        </select>

        <input type="text" id="requestUrl" value="http://localhost:8888/api/configs" placeholder="http://localhost:8888/api/configs"
               style="flex: 1; padding: 10px 16px; border: 1px solid #dadce0; border-radius: 4px; font-size: 14px; font-family: 'Monaco', monospace;" />

        <button onclick="sendRequest()" id="sendBtn"
                style="padding: 10px 32px; background: #7c3aed; color: white; border: none; border-radius: 4px; cursor: pointer; font-size: 14px; font-weight: 500; transition: background 0.2s;">
            Send
        </button>
    </div>

    <!-- Tabs -->
    <div style="display: flex; border-bottom: 1px solid #e8eaed; background: #f8f9fa;">
        <div class="rest-tab active" data-tab="params" onclick="switchRestTab('params')" style="padding: 12px 24px; cursor: pointer; font-size: 13px; font-weight: 500; color: #5f6368; border-bottom: 2px solid transparent; transition: all 0.2s;">
            Parameters
        </div>
        <div class="rest-tab" data-tab="body" onclick="switchRestTab('body')" style="padding: 12px 24px; cursor: pointer; font-size: 13px; font-weight: 500; color: #5f6368; border-bottom: 2px solid transparent; transition: all 0.2s;">
            Body
        </div>
        <div class="rest-tab" data-tab="headers" onclick="switchRestTab('headers')" style="padding: 12px 24px; cursor: pointer; font-size: 13px; font-weight: 500; color: #5f6368; border-bottom: 2px solid transparent; transition: all 0.2s;">
            Headers
        </div>
        <div class="rest-tab" data-tab="auth" onclick="switchRestTab('auth')" style="padding: 12px 24px; cursor: pointer; font-size: 13px; font-weight: 500; color: #5f6368; border-bottom: 2px solid transparent; transition: all 0.2s;">
            Authorization
        </div>
    </div>

    <!-- Tab Content -->
    <div style="min-height: 300px;">
        <!-- Parameters Tab -->
        <div id="tab-params" class="rest-tab-content" style="padding: 20px; display: block;">
            <div style="margin-bottom: 12px; color: #5f6368; font-size: 13px; font-weight: 500;">Query Parameters</div>
            <table style="width: 100%; border-collapse: collapse; font-size: 13px;">
                <thead>
                    <tr style="background: #f8f9fa;">
                        <th style="padding: 8px 12px; text-align: left; font-weight: 500; color: #5f6368; border: 1px solid #e8eaed; width: 35%;">Key</th>
                        <th style="padding: 8px 12px; text-align: left; font-weight: 500; color: #5f6368; border: 1px solid #e8eaed; width: 35%;">Value</th>
                        <th style="padding: 8px 12px; text-align: left; font-weight: 500; color: #5f6368; border: 1px solid #e8eaed; width: 25%;">Description</th>
                        <th style="padding: 8px 12px; text-align: center; font-weight: 500; color: #5f6368; border: 1px solid #e8eaed; width: 5%;"></th>
                    </tr>
                </thead>
                <tbody id="paramsTable">
                    <tr>
                        <td style="padding: 8px 12px; border: 1px solid #e8eaed;">
                            <input type="text" placeholder="key" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" />
                        </td>
                        <td style="padding: 8px 12px; border: 1px solid #e8eaed;">
                            <input type="text" placeholder="value" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" />
                        </td>
                        <td style="padding: 8px 12px; border: 1px solid #e8eaed;">
                            <input type="text" placeholder="description" style="width: 100%; border: none; padding: 4px; font-size: 13px;" />
                        </td>
                        <td style="padding: 8px 12px; border: 1px solid #e8eaed; text-align: center;">
                            <button onclick="removeParamRow(this)" style="background: none; border: none; cursor: pointer; color: #d93025; font-size: 16px;">×</button>
                        </td>
                    </tr>
                </tbody>
            </table>
            <button onclick="addParamRow()" style="margin-top: 12px; padding: 6px 16px; background: white; border: 1px solid #dadce0; border-radius: 4px; cursor: pointer; font-size: 13px; color: #5f6368;">
                + Add Parameter
            </button>
        </div>

        <!-- Body Tab -->
        <div id="tab-body" class="rest-tab-content" style="padding: 20px; display: none;">
            <div style="margin-bottom: 12px; color: #5f6368; font-size: 13px; font-weight: 500;">Request Body (JSON)</div>
            <textarea id="requestBody" rows="15" placeholder='{\n  "key": "value"\n}'
                      style="width: 100%; padding: 12px; border: 1px solid #dadce0; border-radius: 4px; font-family: 'Monaco', monospace; font-size: 13px; resize: vertical;"></textarea>
            <div style="margin-top: 8px; font-size: 12px; color: #5f6368;">
                Format: JSON only. Body is sent for POST, PUT, PATCH requests.
            </div>
        </div>

        <!-- Headers Tab -->
        <div id="tab-headers" class="rest-tab-content" style="padding: 20px; display: none;">
            <div style="margin-bottom: 12px; color: #5f6368; font-size: 13px; font-weight: 500;">Request Headers</div>
            <table style="width: 100%; border-collapse: collapse; font-size: 13px;">
                <thead>
                    <tr style="background: #f8f9fa;">
                        <th style="padding: 8px 12px; text-align: left; font-weight: 500; color: #5f6368; border: 1px solid #e8eaed; width: 35%;">Key</th>
                        <th style="padding: 8px 12px; text-align: left; font-weight: 500; color: #5f6368; border: 1px solid #e8eaed; width: 55%;">Value</th>
                        <th style="padding: 8px 12px; text-align: center; font-weight: 500; color: #5f6368; border: 1px solid #e8eaed; width: 10%;"></th>
                    </tr>
                </thead>
                <tbody id="headersTable">
                    <tr>
                        <td style="padding: 8px 12px; border: 1px solid #e8eaed;">
                            <input type="text" value="Content-Type" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" />
                        </td>
                        <td style="padding: 8px 12px; border: 1px solid #e8eaed;">
                            <input type="text" value="application/json" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" />
                        </td>
                        <td style="padding: 8px 12px; border: 1px solid #e8eaed; text-align: center;">
                            <button onclick="removeHeaderRow(this)" style="background: none; border: none; cursor: pointer; color: #d93025; font-size: 16px;">×</button>
                        </td>
                    </tr>
                </tbody>
            </table>
            <button onclick="addHeaderRow()" style="margin-top: 12px; padding: 6px 16px; background: white; border: 1px solid #dadce0; border-radius: 4px; cursor: pointer; font-size: 13px; color: #5f6368;">
                + Add Header
            </button>
        </div>

        <!-- Authorization Tab -->
        <div id="tab-auth" class="rest-tab-content" style="padding: 20px; display: none;">
            <div style="margin-bottom: 12px; color: #5f6368; font-size: 13px; font-weight: 500;">Client Certificates (Optional)</div>
            <div style="margin-bottom: 20px;">
                <label style="display: block; margin-bottom: 8px; font-size: 13px; color: #5f6368; font-weight: 500;">TLS Certificate (PEM format)</label>
                <div style="position: relative;">
                    <input type="file" id="tlsCertFile" accept=".pem,.crt,.cer" style="display: none;" onchange="handleCertFileUpload(this)" />
                    <button onclick="document.getElementById('tlsCertFile').click()"
                            style="padding: 10px 16px; background: white; border: 1px solid #dadce0; border-radius: 4px; cursor: pointer; font-size: 13px; color: #5f6368; display: flex; align-items: center; gap: 8px;">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                            <polyline points="17 8 12 3 7 8"></polyline>
                            <line x1="12" y1="3" x2="12" y2="15"></line>
                        </svg>
                        <span id="tlsCertFileName">Choose Certificate File</span>
                    </button>
                    <div id="tlsCertPreview" style="margin-top: 8px; padding: 8px; background: #f8f9fa; border-radius: 4px; font-size: 11px; color: #5f6368; display: none; font-family: Monaco, monospace;"></div>
                </div>
            </div>
            <div>
                <label style="display: block; margin-bottom: 8px; font-size: 13px; color: #5f6368; font-weight: 500;">TLS Private Key (PEM format)</label>
                <div style="position: relative;">
                    <input type="file" id="tlsKeyFile" accept=".pem,.key" style="display: none;" onchange="handleKeyFileUpload(this)" />
                    <button onclick="document.getElementById('tlsKeyFile').click()"
                            style="padding: 10px 16px; background: white; border: 1px solid #dadce0; border-radius: 4px; cursor: pointer; font-size: 13px; color: #5f6368; display: flex; align-items: center; gap: 8px;">
                        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                            <polyline points="17 8 12 3 7 8"></polyline>
                            <line x1="12" y1="3" x2="12" y2="15"></line>
                        </svg>
                        <span id="tlsKeyFileName">Choose Private Key File</span>
                    </button>
                    <div id="tlsKeyPreview" style="margin-top: 8px; padding: 8px; background: #f8f9fa; border-radius: 4px; font-size: 11px; color: #5f6368; display: none; font-family: Monaco, monospace;"></div>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Response Panel -->
<div id="responsePanel" style="margin-top: 24px; background: white; border-radius: 8px; border: 1px solid #dadce0; display: none;">
    <div style="padding: 16px 20px; border-bottom: 1px solid #e8eaed; display: flex; justify-content: space-between; align-items: center;">
        <div style="font-size: 16px; font-weight: 500; color: #202124;">Response</div>
        <div style="display: flex; gap: 16px; align-items: center; font-size: 13px; color: #5f6368;">
            <span>Status: <span id="responseStatus" style="font-weight: 500;"></span></span>
            <span>Time: <span id="responseTime" style="font-weight: 500;"></span></span>
            <span>Size: <span id="responseSize" style="font-weight: 500;"></span></span>
        </div>
    </div>

    <div style="display: flex; border-bottom: 1px solid #e8eaed; background: #f8f9fa;">
        <div class="resp-tab active" data-tab="resp-body" onclick="switchRespTab('resp-body')" style="padding: 12px 24px; cursor: pointer; font-size: 13px; font-weight: 500; color: #5f6368; border-bottom: 2px solid transparent;">
            Body
        </div>
        <div class="resp-tab" data-tab="resp-headers" onclick="switchRespTab('resp-headers')" style="padding: 12px 24px; cursor: pointer; font-size: 13px; font-weight: 500; color: #5f6368; border-bottom: 2px solid transparent;">
            Headers
        </div>
    </div>

    <div style="min-height: 200px;">
        <div id="tab-resp-body" class="resp-tab-content" style="padding: 20px; display: block;">
            <pre id="responseBody" style="margin: 0; padding: 16px; background: #f8f9fa; border-radius: 4px; font-family: 'Monaco', monospace; font-size: 12px; white-space: pre-wrap; word-wrap: break-word; max-height: 500px; overflow: auto;"></pre>
        </div>
        <div id="tab-resp-headers" class="resp-tab-content" style="padding: 20px; display: none;">
            <table id="responseHeaders" style="width: 100%; border-collapse: collapse; font-size: 13px;"></table>
        </div>
    </div>
</div>

<style>
.rest-tab:hover {
    background: #f1f3f4;
}
.rest-tab.active {
    color: #7c3aed !important;
    border-bottom-color: #7c3aed !important;
    background: white;
}
.resp-tab:hover {
    background: #f1f3f4;
}
.resp-tab.active {
    color: #7c3aed !important;
    border-bottom-color: #7c3aed !important;
    background: white;
}
#sendBtn:hover {
    background: #6d28d9;
}

/* JSON Syntax Highlighting */
.json-key { color: #881391; }
.json-string { color: #1a1aa6; }
.json-number { color: #1c00cf; }
.json-boolean { color: #0d22aa; }
.json-null { color: #808080; }
.json-punctuation { color: #303030; }
</style>
`

const RestClientJS = `
// Global variables to store uploaded cert/key content
let tlsCertContent = '';
let tlsKeyContent = '';

function handleCertFileUpload(input) {
    const file = input.files[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = function(e) {
        tlsCertContent = e.target.result;
        document.getElementById('tlsCertFileName').textContent = file.name;

        // Show preview
        const preview = document.getElementById('tlsCertPreview');
        const lines = tlsCertContent.split('\n');
        preview.textContent = lines.slice(0, 3).join('\n') + '\n... (' + lines.length + ' lines total)';
        preview.style.display = 'block';
    };
    reader.readAsText(file);
}

function handleKeyFileUpload(input) {
    const file = input.files[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = function(e) {
        tlsKeyContent = e.target.result;
        document.getElementById('tlsKeyFileName').textContent = file.name;

        // Show preview
        const preview = document.getElementById('tlsKeyPreview');
        const lines = tlsKeyContent.split('\n');
        preview.textContent = lines.slice(0, 3).join('\n') + '\n... (' + lines.length + ' lines total)';
        preview.style.display = 'block';
    };
    reader.readAsText(file);
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

function switchRestTab(tabName) {
    // Hide all tab contents
    document.querySelectorAll('.rest-tab-content').forEach(el => el.style.display = 'none');
    document.querySelectorAll('.rest-tab').forEach(el => el.classList.remove('active'));

    // Show selected tab
    document.getElementById('tab-' + tabName).style.display = 'block';
    document.querySelector('[data-tab="' + tabName + '"]').classList.add('active');
}

function switchRespTab(tabName) {
    // Hide all response tab contents
    document.querySelectorAll('.resp-tab-content').forEach(el => el.style.display = 'none');
    document.querySelectorAll('.resp-tab').forEach(el => el.classList.remove('active'));

    // Show selected tab
    document.getElementById('tab-' + tabName).style.display = 'block';
    document.querySelector('[data-tab="' + tabName + '"]').classList.add('active');
}

function addParamRow() {
    const table = document.getElementById('paramsTable');
    const row = table.insertRow();
    row.innerHTML = '<td style="padding: 8px 12px; border: 1px solid #e8eaed;"><input type="text" placeholder="key" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" /></td>' +
                    '<td style="padding: 8px 12px; border: 1px solid #e8eaed;"><input type="text" placeholder="value" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" /></td>' +
                    '<td style="padding: 8px 12px; border: 1px solid #e8eaed;"><input type="text" placeholder="description" style="width: 100%; border: none; padding: 4px; font-size: 13px;" /></td>' +
                    '<td style="padding: 8px 12px; border: 1px solid #e8eaed; text-align: center;"><button onclick="removeParamRow(this)" style="background: none; border: none; cursor: pointer; color: #d93025; font-size: 16px;">×</button></td>';
}

function removeParamRow(btn) {
    btn.closest('tr').remove();
}

function addHeaderRow() {
    const table = document.getElementById('headersTable');
    const row = table.insertRow();
    row.innerHTML = '<td style="padding: 8px 12px; border: 1px solid #e8eaed;"><input type="text" placeholder="Header-Name" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" /></td>' +
                    '<td style="padding: 8px 12px; border: 1px solid #e8eaed;"><input type="text" placeholder="value" style="width: 100%; border: none; padding: 4px; font-size: 13px; font-family: Monaco, monospace;" /></td>' +
                    '<td style="padding: 8px 12px; border: 1px solid #e8eaed; text-align: center;"><button onclick="removeHeaderRow(this)" style="background: none; border: none; cursor: pointer; color: #d93025; font-size: 16px;">×</button></td>';
}

function removeHeaderRow(btn) {
    btn.closest('tr').remove();
}

function saveRequest() {
    alert('Request saved! (Feature coming soon)');
}

async function sendRequest() {
    const method = document.getElementById('httpMethod').value;
    let url = document.getElementById('requestUrl').value.trim();

    if (!url) {
        alert('Please enter a URL');
        return;
    }

    // Build query parameters
    const params = [];
    document.querySelectorAll('#paramsTable tr').forEach(row => {
        const inputs = row.querySelectorAll('input');
        const key = inputs[0]?.value.trim();
        const value = inputs[1]?.value.trim();
        if (key && value) {
            params.push(encodeURIComponent(key) + '=' + encodeURIComponent(value));
        }
    });

    if (params.length > 0) {
        url += (url.includes('?') ? '&' : '?') + params.join('&');
    }

    // Build headers
    const headers = {};
    document.querySelectorAll('#headersTable tr').forEach(row => {
        const inputs = row.querySelectorAll('input');
        const key = inputs[0]?.value.trim();
        const value = inputs[1]?.value.trim();
        if (key && value) {
            headers[key] = value;
        }
    });

    // Get body
    let body = null;
    if (['POST', 'PUT', 'PATCH'].includes(method)) {
        const bodyText = document.getElementById('requestBody').value.trim();
        if (bodyText) {
            try {
                body = JSON.parse(bodyText);
            } catch (e) {
                alert('Invalid JSON in body: ' + e.message);
                return;
            }
        }
    }

    // Get TLS certs from uploaded files
    const tlsCert = tlsCertContent.trim();
    const tlsKey = tlsKeyContent.trim();

    const requestData = {
        method: method,
        url: url,
        headers: headers,
        body: body,
        tlsCert: tlsCert || null,
        tlsKey: tlsKey || null
    };

    const startTime = Date.now();
    document.getElementById('sendBtn').textContent = 'Sending...';
    document.getElementById('sendBtn').disabled = true;

    try {
        const response = await fetch('/api/rest/send', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });

        const endTime = Date.now();
        const duration = endTime - startTime;

        const result = await response.json();

        // Show response panel
        document.getElementById('responsePanel').style.display = 'block';

        // Update response info
        const statusColor = result.statusCode >= 200 && result.statusCode < 300 ? '#188038' :
                           result.statusCode >= 400 ? '#d93025' : '#f9ab00';
        document.getElementById('responseStatus').innerHTML = '<span style="color: ' + statusColor + ';">' + result.statusCode + '</span>';
        document.getElementById('responseTime').textContent = duration + ' ms';

        if (result.error) {
            document.getElementById('responseBody').textContent = 'Error: ' + result.error;
            document.getElementById('responseSize').textContent = '-';
        } else {
            const bodyStr = typeof result.body === 'string' ? result.body : JSON.stringify(result.body, null, 2);

            // Try to parse and highlight as JSON
            try {
                const jsonObj = typeof result.body === 'object' ? result.body : JSON.parse(result.body);
                document.getElementById('responseBody').innerHTML = syntaxHighlightJSON(jsonObj);
            } catch (e) {
                // Not JSON, display as plain text
                document.getElementById('responseBody').textContent = bodyStr;
            }

            document.getElementById('responseSize').textContent = (bodyStr.length / 1024).toFixed(2) + ' KB';

            // Update response headers
            const headersTable = document.getElementById('responseHeaders');
            headersTable.innerHTML = '';
            if (result.headers) {
                Object.keys(result.headers).forEach(key => {
                    const row = headersTable.insertRow();
                    row.innerHTML = '<td style="padding: 8px 12px; border: 1px solid #e8eaed; font-weight: 500; color: #5f6368; font-family: Monaco, monospace; font-size: 12px;">' + key + '</td>' +
                                   '<td style="padding: 8px 12px; border: 1px solid #e8eaed; font-family: Monaco, monospace; font-size: 12px;">' + result.headers[key] + '</td>';
                });
            }
        }

        // Scroll to response
        document.getElementById('responsePanel').scrollIntoView({ behavior: 'smooth', block: 'nearest' });

    } catch (error) {
        document.getElementById('responsePanel').style.display = 'block';
        document.getElementById('responseStatus').innerHTML = '<span style="color: #d93025;">Error</span>';
        document.getElementById('responseTime').textContent = '-';
        document.getElementById('responseSize').textContent = '-';
        document.getElementById('responseBody').textContent = 'Request failed: ' + error.message;
    } finally {
        document.getElementById('sendBtn').textContent = 'Send';
        document.getElementById('sendBtn').disabled = false;
    }
}
`

func GetRestClientHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>REST Client - Testing Studio</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", sans-serif;
            background: #f5f5f5;
            color: #202124;
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
        .container { max-width: 1400px; margin: 0 auto; padding: 24px; }
    </style>
</head>
<body>
    <div class="topbar">
        <a href="/" class="logo">Testing Studio</a>
        <a href="/" class="back-btn">← Back</a>
    </div>

    <div class="container">
        ` + RestClient + `
    </div>

    <script>
        ` + RestClientJS + `
    </script>
</body>
</html>`
}
