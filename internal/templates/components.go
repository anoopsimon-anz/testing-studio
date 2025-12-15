package templates

const Base64Modal = `<div id="base64Modal" style="display: none; position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.5); z-index: 1000; align-items: center; justify-content: center;">
    <div style="background: white; border-radius: 8px; max-width: 900px; width: 90%; max-height: 90vh; overflow: hidden; display: flex; flex-direction: column;">
        <div style="padding: 20px; border-bottom: 1px solid #dadce0; display: flex; justify-content: space-between; align-items: center;">
            <h2 style="font-size: 20px; font-weight: 500; color: #202124;">Base64 Encoder/Decoder</h2>
            <button onclick="closeBase64Tool()" style="background: none; border: none; font-size: 24px; cursor: pointer; color: #5f6368;">&times;</button>
        </div>
        <div style="display: flex; flex: 1; overflow: hidden;">
            <div style="flex: 1; padding: 20px; border-right: 1px solid #dadce0; display: flex; flex-direction: column;">
                <label style="font-size: 13px; color: #5f6368; font-weight: 500; margin-bottom: 8px;">Input Text:</label>
                <textarea id="base64Input" style="flex: 1; font-family: 'Monaco', monospace; font-size: 13px; border: 1px solid #dadce0; border-radius: 4px; padding: 12px; resize: none;" placeholder="Enter text or Base64 string here"></textarea>
                <div style="display: flex; gap: 8px; margin-top: 12px;">
                    <button onclick="encodeBase64()" style="flex: 1; background: #1a73e8; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; font-weight: 500;">Encode to Base64</button>
                    <button onclick="decodeBase64()" style="flex: 1; background: #188038; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; font-weight: 500;">Decode from Base64</button>
                </div>
            </div>
            <div style="flex: 1; padding: 20px; display: flex; flex-direction: column;">
                <label style="font-size: 13px; color: #5f6368; font-weight: 500; margin-bottom: 8px;">Output:</label>
                <textarea id="base64Output" readonly style="flex: 1; font-family: 'Monaco', monospace; font-size: 13px; border: 1px solid #dadce0; border-radius: 4px; padding: 12px; resize: none; background: #f8f9fa;"></textarea>
                <button onclick="copyOutput()" style="margin-top: 12px; background: white; color: #5f6368; border: 1px solid #dadce0; padding: 10px 20px; border-radius: 4px; cursor: pointer; font-weight: 500;">Copy to Clipboard</button>
            </div>
        </div>
    </div>
</div>`

const Base64ModalJS = `function openBase64Tool() {
    document.getElementById('base64Modal').style.display = 'flex';
}

function closeBase64Tool() {
    document.getElementById('base64Modal').style.display = 'none';
    document.getElementById('base64Input').value = '';
    document.getElementById('base64Output').value = '';
}

function encodeBase64() {
    const input = document.getElementById('base64Input').value;
    const output = document.getElementById('base64Output');

    if (!input) {
        output.value = 'Error: Please enter some text to encode';
        return;
    }

    try {
        const encoded = btoa(unescape(encodeURIComponent(input)));
        output.value = encoded;
    } catch (e) {
        output.value = 'Error: Failed to encode - ' + e.message;
    }
}

function decodeBase64() {
    const input = document.getElementById('base64Input').value;
    const output = document.getElementById('base64Output');

    if (!input) {
        output.value = 'Error: Please enter a Base64 string to decode';
        return;
    }

    try {
        const decoded = decodeURIComponent(escape(atob(input)));
        output.value = decoded;
    } catch (e) {
        output.value = 'Error: Invalid Base64 string - ' + e.message;
    }
}

function copyOutput() {
    const output = document.getElementById('base64Output');
    if (!output.value || output.value.startsWith('Error:')) {
        return;
    }
    output.select();
    document.execCommand('copy');

    const btn = event.target;
    const originalText = btn.textContent;
    btn.textContent = 'Copied!';
    btn.style.background = '#188038';
    btn.style.color = 'white';
    setTimeout(function() {
        btn.textContent = originalText;
        btn.style.background = 'white';
        btn.style.color = '#5f6368';
    }, 2000);
}

// Close modal on outside click
document.getElementById('base64Modal')?.addEventListener('click', function(e) {
    if (e.target === this) {
        closeBase64Tool();
    }
});`

const JWTModal = `<div id="jwtModal" style="display: none; position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.5); z-index: 1000; align-items: center; justify-content: center;">
    <div style="background: white; border-radius: 8px; max-width: 1000px; width: 90%; max-height: 90vh; overflow: hidden; display: flex; flex-direction: column;">
        <div style="padding: 20px; border-bottom: 1px solid #dadce0; display: flex; justify-content: space-between; align-items: center;">
            <h2 style="font-size: 20px; font-weight: 500; color: #202124;">JSON to JWT Token Converter</h2>
            <button onclick="closeJWTTool()" style="background: none; border: none; font-size: 24px; cursor: pointer; color: #5f6368;">&times;</button>
        </div>
        <div style="display: flex; flex: 1; overflow: hidden;">
            <div style="flex: 1; padding: 20px; border-right: 1px solid #dadce0; display: flex; flex-direction: column;">
                <label style="font-size: 13px; color: #5f6368; font-weight: 500; margin-bottom: 8px;">JSON Payload:</label>
                <textarea id="jwtJsonInput" style="flex: 1; font-family: 'Monaco', monospace; font-size: 13px; border: 1px solid #dadce0; border-radius: 4px; padding: 12px; resize: none;" placeholder='{"sub": "1234567890", "name": "John Doe", "iat": 1516239022}'></textarea>
                <div style="margin-top: 12px;">
                    <label style="font-size: 13px; color: #5f6368; font-weight: 500; margin-bottom: 4px; display: block;">Secret Key:</label>
                    <input type="text" id="jwtSecret" value="your-256-bit-secret" style="width: 100%; font-family: 'Monaco', monospace; font-size: 13px; border: 1px solid #dadce0; border-radius: 4px; padding: 8px;">
                </div>
                <div style="display: flex; gap: 8px; margin-top: 12px;">
                    <button onclick="encodeJWT()" style="flex: 1; background: #1a73e8; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; font-weight: 500;">Encode to JWT</button>
                    <button onclick="decodeJWT()" style="flex: 1; background: #188038; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; font-weight: 500;">Decode JWT</button>
                </div>
            </div>
            <div style="flex: 1; padding: 20px; display: flex; flex-direction: column;">
                <label style="font-size: 13px; color: #5f6368; font-weight: 500; margin-bottom: 8px;">JWT Token / Decoded JSON:</label>
                <textarea id="jwtOutput" readonly style="flex: 1; font-family: 'Monaco', monospace; font-size: 13px; border: 1px solid #dadce0; border-radius: 4px; padding: 12px; resize: none; background: #f8f9fa;"></textarea>
                <button onclick="copyJWTOutput()" style="margin-top: 12px; background: white; color: #5f6368; border: 1px solid #dadce0; padding: 10px 20px; border-radius: 4px; cursor: pointer; font-weight: 500;">Copy to Clipboard</button>
                <div style="margin-top: 8px; font-size: 11px; color: #5f6368; line-height: 1.4;">
                    <strong>Note:</strong> This uses HS256 algorithm. For production use, consider RS256 with proper key management.
                </div>
            </div>
        </div>
    </div>
</div>`

const JWTModalJS = `function openJWTTool() {
    document.getElementById('jwtModal').style.display = 'flex';
}

function closeJWTTool() {
    document.getElementById('jwtModal').style.display = 'none';
    document.getElementById('jwtJsonInput').value = '';
    document.getElementById('jwtOutput').value = '';
}

function base64UrlEncode(str) {
    return btoa(str)
        .replace(/\+/g, '-')
        .replace(/\//g, '_')
        .replace(/=+$/, '');
}

function base64UrlDecode(str) {
    str = str.replace(/-/g, '+').replace(/_/g, '/');
    while (str.length % 4) {
        str += '=';
    }
    return atob(str);
}

async function encodeJWT() {
    const jsonInput = document.getElementById('jwtJsonInput').value.trim();
    const secret = document.getElementById('jwtSecret').value;
    const output = document.getElementById('jwtOutput');

    if (!jsonInput) {
        output.value = 'Error: Please enter JSON payload';
        return;
    }

    try {
        // Validate JSON
        const payload = JSON.parse(jsonInput);

        // Create header
        const header = {
            alg: 'HS256',
            typ: 'JWT'
        };

        // Base64Url encode header and payload
        const encodedHeader = base64UrlEncode(JSON.stringify(header));
        const encodedPayload = base64UrlEncode(JSON.stringify(payload));

        // Create signature
        const data = encodedHeader + '.' + encodedPayload;

        // Use Web Crypto API for HMAC SHA256
        const encoder = new TextEncoder();
        const keyData = encoder.encode(secret);
        const messageData = encoder.encode(data);

        const cryptoKey = await crypto.subtle.importKey(
            'raw',
            keyData,
            { name: 'HMAC', hash: 'SHA-256' },
            false,
            ['sign']
        );

        const signature = await crypto.subtle.sign('HMAC', cryptoKey, messageData);
        const encodedSignature = base64UrlEncode(String.fromCharCode(...new Uint8Array(signature)));

        // Combine to create JWT
        const jwt = data + '.' + encodedSignature;
        output.value = jwt;
    } catch (e) {
        output.value = 'Error: ' + e.message;
    }
}

function decodeJWT() {
    const input = document.getElementById('jwtJsonInput').value.trim();
    const output = document.getElementById('jwtOutput');

    if (!input) {
        output.value = 'Error: Please enter a JWT token in the JSON Payload field';
        return;
    }

    try {
        // Split JWT into parts
        const parts = input.split('.');
        if (parts.length !== 3) {
            throw new Error('Invalid JWT format. Expected 3 parts separated by dots.');
        }

        // Decode header and payload
        const header = JSON.parse(base64UrlDecode(parts[0]));
        const payload = JSON.parse(base64UrlDecode(parts[1]));

        // Format output
        const decoded = {
            header: header,
            payload: payload,
            signature: parts[2]
        };

        output.value = JSON.stringify(decoded, null, 2);
    } catch (e) {
        output.value = 'Error: Failed to decode JWT - ' + e.message;
    }
}

function copyJWTOutput() {
    const output = document.getElementById('jwtOutput');
    if (!output.value || output.value.startsWith('Error:')) {
        return;
    }
    output.select();
    document.execCommand('copy');

    const btn = event.target;
    const originalText = btn.textContent;
    btn.textContent = 'Copied!';
    btn.style.background = '#188038';
    btn.style.color = 'white';
    setTimeout(function() {
        btn.textContent = originalText;
        btn.style.background = 'white';
        btn.style.color = '#5f6368';
    }, 2000);
}

// Close modal on outside click
document.getElementById('jwtModal')?.addEventListener('click', function(e) {
    if (e.target === this) {
        closeJWTTool();
    }
});`
