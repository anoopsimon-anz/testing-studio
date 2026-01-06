package templates

const ConfigEditor = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Config Editor - Testing Studio</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Google Sans', 'Product Sans', -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", sans-serif;
            background: #f5f5f5;
            color: #202124;
        }
        .header {
            background: white;
            border-bottom: 1px solid #dadce0;
            padding: 16px 24px;
            display: flex;
            align-items: center;
            justify-content: space-between;
        }
        .header h1 {
            font-size: 20px;
            font-weight: 500;
            color: #202124;
        }
        .back-link {
            color: #1a73e8;
            text-decoration: none;
            font-size: 14px;
        }
        .back-link:hover {
            text-decoration: underline;
        }
        .container {
            display: flex;
            height: calc(100vh - 60px);
            overflow: hidden;
        }
        .sidebar {
            width: 280px;
            background: white;
            border-right: 1px solid #dadce0;
            overflow-y: auto;
        }
        .sidebar-item {
            padding: 16px 24px;
            cursor: pointer;
            border-bottom: 1px solid #e8eaed;
            transition: background 0.2s;
            display: flex;
            align-items: center;
            gap: 12px;
        }
        .sidebar-item:hover {
            background: #f8f9fa;
        }
        .sidebar-item.active {
            background: #e8f0fe;
            border-left: 4px solid #1a73e8;
            padding-left: 20px;
            font-weight: 500;
            color: #1a73e8;
        }
        .sidebar-item-icon {
            font-size: 20px;
        }
        .content-panel {
            flex: 1;
            overflow-y: auto;
            padding: 24px;
            background: #f5f5f5;
        }
        .section {
            background: white;
            border: 1px solid #dadce0;
            border-radius: 8px;
            padding: 24px;
            display: none;
        }
        .section.active {
            display: block;
        }
        .section-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 20px;
        }
        .section-title {
            font-size: 18px;
            font-weight: 500;
            color: #202124;
        }
        .add-btn {
            background: #1a73e8;
            color: white;
            border: none;
            padding: 8px 16px;
            border-radius: 4px;
            font-size: 14px;
            font-weight: 500;
            cursor: pointer;
            transition: background 0.2s;
        }
        .add-btn:hover {
            background: #1765cc;
        }
        .config-item {
            border: 1px solid #e8eaed;
            border-radius: 8px;
            padding: 16px;
            margin-bottom: 16px;
            background: #f8f9fa;
        }
        .config-item-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 12px;
        }
        .config-name {
            font-size: 16px;
            font-weight: 500;
            color: #202124;
        }
        .config-actions {
            display: flex;
            gap: 8px;
        }
        .edit-btn, .delete-btn, .save-btn, .cancel-btn {
            padding: 6px 12px;
            border-radius: 4px;
            font-size: 13px;
            font-weight: 500;
            cursor: pointer;
            border: none;
            transition: all 0.2s;
        }
        .edit-btn {
            background: #e8f0fe;
            color: #1967d2;
        }
        .edit-btn:hover {
            background: #d2e3fc;
        }
        .delete-btn {
            background: #fce8e6;
            color: #d93025;
        }
        .delete-btn:hover {
            background: #fad2cf;
        }
        .save-btn {
            background: #1a73e8;
            color: white;
        }
        .save-btn:hover {
            background: #1765cc;
        }
        .cancel-btn {
            background: #f1f3f4;
            color: #5f6368;
        }
        .cancel-btn:hover {
            background: #e8eaed;
        }
        .config-details {
            font-size: 13px;
            color: #5f6368;
            line-height: 1.6;
        }
        .config-details div {
            margin-bottom: 4px;
        }
        .form-group {
            margin-bottom: 16px;
        }
        .form-label {
            display: block;
            font-size: 13px;
            font-weight: 500;
            color: #5f6368;
            margin-bottom: 6px;
        }
        .form-input {
            width: 100%;
            padding: 10px 12px;
            border: 1px solid #dadce0;
            border-radius: 4px;
            font-size: 14px;
            font-family: monospace;
        }
        .form-input:focus {
            outline: none;
            border-color: #1a73e8;
        }
        .hidden {
            display: none;
        }
        .new-config-form {
            display: none;
            border: 2px dashed #dadce0;
            border-radius: 8px;
            padding: 20px;
            margin-top: 16px;
            background: white;
        }
        .new-config-form.visible {
            display: block;
        }
        .success-message {
            background: #e6f4ea;
            border: 1px solid #34a853;
            color: #137333;
            padding: 12px 16px;
            border-radius: 4px;
            margin-bottom: 16px;
            font-size: 14px;
            display: none;
        }
        .success-message.visible {
            display: block;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>⚙️ Configuration Editor</h1>
        <a href="/" class="back-link">← Back to Home</a>
    </div>

    <div class="container">
        <div class="sidebar">
            <div class="sidebar-item active" onclick="switchTab('pubsub')" id="tab-pubsub">
                <span class="sidebar-item-icon">📮</span>
                <span>Google PubSub</span>
            </div>
            <div class="sidebar-item" onclick="switchTab('kafka')" id="tab-kafka">
                <span class="sidebar-item-icon">📡</span>
                <span>Kafka / EventMesh</span>
            </div>
            <div class="sidebar-item" onclick="switchTab('spanner')" id="tab-spanner">
                <span class="sidebar-item-icon">🗄️</span>
                <span>Spanner Database</span>
            </div>
            <div class="sidebar-item" onclick="switchTab('gcs')" id="tab-gcs">
                <span class="sidebar-item-icon">☁️</span>
                <span>GCS Storage</span>
            </div>
        </div>

        <div class="content-panel">
            <div id="successMessage" class="success-message"></div>

            <!-- PubSub Configs -->
        <div class="section active" id="section-pubsub">
            <div class="section-header">
                <div class="section-title">Google PubSub Configurations</div>
                <button class="add-btn" onclick="showNewConfigForm('pubsub')">+ Add New</button>
            </div>
            <div id="pubsubConfigs"></div>
            <div id="newPubSubForm" class="new-config-form">
                <div class="form-group">
                    <label class="form-label">Configuration Name</label>
                    <input type="text" class="form-input" id="newPubSubName" placeholder="e.g., TMS PubSub">
                </div>
                <div class="form-group">
                    <label class="form-label">Emulator Host</label>
                    <input type="text" class="form-input" id="newPubSubHost" placeholder="e.g., localhost:8086">
                </div>
                <div class="form-group">
                    <label class="form-label">Project ID</label>
                    <input type="text" class="form-input" id="newPubSubProject" placeholder="e.g., tms-suncorp-local">
                </div>
                <div class="form-group">
                    <label class="form-label">Subscription ID</label>
                    <input type="text" class="form-input" id="newPubSubSub" placeholder="e.g., cloudevents.subscription">
                </div>
                <div style="display: flex; gap: 8px; justify-content: flex-end;">
                    <button class="cancel-btn" onclick="hideNewConfigForm('pubsub')">Cancel</button>
                    <button class="save-btn" onclick="saveNewPubSubConfig()">Save Configuration</button>
                </div>
            </div>
        </div>

        <!-- Kafka Configs -->
        <div class="section" id="section-kafka">
            <div class="section-header">
                <div class="section-title">Kafka / EventMesh Configurations</div>
                <button class="add-btn" onclick="showNewConfigForm('kafka')">+ Add New</button>
            </div>
            <div id="kafkaConfigs"></div>
            <div id="newKafkaForm" class="new-config-form">
                <div class="form-group">
                    <label class="form-label">Configuration Name</label>
                    <input type="text" class="form-input" id="newKafkaName" placeholder="e.g., Unica Events">
                </div>
                <div class="form-group">
                    <label class="form-label">Brokers</label>
                    <input type="text" class="form-input" id="newKafkaBrokers" placeholder="e.g., localhost:19092">
                </div>
                <div class="form-group">
                    <label class="form-label">Topic</label>
                    <input type="text" class="form-input" id="newKafkaTopic" placeholder="e.g., unica.marketing.response.events">
                </div>
                <div class="form-group">
                    <label class="form-label">Consumer Group</label>
                    <input type="text" class="form-input" id="newKafkaGroup" placeholder="e.g., testing-studio-consumer">
                </div>
                <div class="form-group">
                    <label class="form-label">Schema Registry URL</label>
                    <input type="text" class="form-input" id="newKafkaSchema" placeholder="e.g., http://localhost:18081">
                </div>
                <div style="display: flex; gap: 8px; justify-content: flex-end;">
                    <button class="cancel-btn" onclick="hideNewConfigForm('kafka')">Cancel</button>
                    <button class="save-btn" onclick="saveNewKafkaConfig()">Save Configuration</button>
                </div>
            </div>
        </div>

        <!-- Spanner Configs -->
        <div class="section" id="section-spanner">
            <div class="section-header">
                <div class="section-title">Spanner Database Configurations</div>
                <button class="add-btn" onclick="showNewConfigForm('spanner')">+ Add New</button>
            </div>
            <div id="spannerConfigs"></div>
            <div id="newSpannerForm" class="new-config-form">
                <div class="form-group">
                    <label class="form-label">Configuration Name</label>
                    <input type="text" class="form-input" id="newSpannerName" placeholder="e.g., TMS Local">
                </div>
                <div class="form-group">
                    <label class="form-label">Emulator Host</label>
                    <input type="text" class="form-input" id="newSpannerHost" placeholder="e.g., localhost:9010">
                </div>
                <div class="form-group">
                    <label class="form-label">Project ID</label>
                    <input type="text" class="form-input" id="newSpannerProject" placeholder="e.g., tms-suncorp-local">
                </div>
                <div class="form-group">
                    <label class="form-label">Instance ID</label>
                    <input type="text" class="form-input" id="newSpannerInstance" placeholder="e.g., tms-suncorp-local">
                </div>
                <div class="form-group">
                    <label class="form-label">Database ID</label>
                    <input type="text" class="form-input" id="newSpannerDatabase" placeholder="e.g., tms-suncorp-db">
                </div>
                <div style="display: flex; gap: 8px; justify-content: flex-end;">
                    <button class="cancel-btn" onclick="hideNewConfigForm('spanner')">Cancel</button>
                    <button class="save-btn" onclick="saveNewSpannerConfig()">Save Configuration</button>
                </div>
            </div>
        </div>

        <!-- GCS Configs -->
        <div class="section" id="section-gcs">
            <div class="section-header">
                <div class="section-title">GCS Storage Configurations</div>
                <button class="add-btn" onclick="showNewConfigForm('gcs')">+ Add New</button>
            </div>
            <div id="gcsConfigs"></div>
            <div id="newGCSForm" class="new-config-form">
                <div class="form-group">
                    <label class="form-label">Configuration Name</label>
                    <input type="text" class="form-input" id="newGCSName" placeholder="e.g., TMS GCS Local">
                </div>
                <div class="form-group">
                    <label class="form-label">Emulator Host</label>
                    <input type="text" class="form-input" id="newGCSHost" placeholder="e.g., localhost:4443">
                </div>
                <div class="form-group">
                    <label class="form-label">Project ID</label>
                    <input type="text" class="form-input" id="newGCSProject" placeholder="e.g., tms-suncorp-local">
                </div>
                <div style="display: flex; gap: 8px; justify-content: flex-end;">
                    <button class="cancel-btn" onclick="hideNewConfigForm('gcs')">Cancel</button>
                    <button class="save-btn" onclick="saveNewGCSConfig()">Save Configuration</button>
                </div>
            </div>
        </div>
        </div>
    </div>

    <script>
        let configs = { pubsubConfigs: [], kafkaConfigs: [], spannerConfigs: [], gcsConfigs: [] };

        function switchTab(tabName) {
            // Remove active from all tabs
            document.querySelectorAll('.sidebar-item').forEach(item => item.classList.remove('active'));
            document.querySelectorAll('.section').forEach(section => section.classList.remove('active'));

            // Add active to selected tab
            document.getElementById('tab-' + tabName).classList.add('active');
            document.getElementById('section-' + tabName).classList.add('active');
        }

        async function loadConfigs() {
            const response = await fetch('/api/configs');
            configs = await response.json();
            renderConfigs();
        }

        function renderConfigs() {
            renderPubSubConfigs();
            renderKafkaConfigs();
            renderSpannerConfigs();
            renderGCSConfigs();
        }

        function renderPubSubConfigs() {
            const container = document.getElementById('pubsubConfigs');
            container.innerHTML = configs.pubsubConfigs.map((config, index) => ` + "`" + `
                <div class="config-item" id="pubsub-${index}">
                    <div class="config-item-header">
                        <div class="config-name">${config.name}</div>
                        <div class="config-actions">
                            <button class="edit-btn" onclick="editPubSubConfig(${index})">Edit</button>
                            <button class="delete-btn" onclick="deletePubSubConfig(${index})">Delete</button>
                        </div>
                    </div>
                    <div class="config-details" id="pubsub-details-${index}">
                        <div><strong>Emulator Host:</strong> ${config.emulatorHost}</div>
                        <div><strong>Project ID:</strong> ${config.projectId}</div>
                        <div><strong>Subscription ID:</strong> ${config.subscriptionId}</div>
                    </div>
                    <div class="hidden" id="pubsub-form-${index}">
                        <div class="form-group">
                            <label class="form-label">Configuration Name</label>
                            <input type="text" class="form-input" id="edit-pubsub-name-${index}" value="${config.name}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Emulator Host</label>
                            <input type="text" class="form-input" id="edit-pubsub-host-${index}" value="${config.emulatorHost}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Project ID</label>
                            <input type="text" class="form-input" id="edit-pubsub-project-${index}" value="${config.projectId}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Subscription ID</label>
                            <input type="text" class="form-input" id="edit-pubsub-sub-${index}" value="${config.subscriptionId}">
                        </div>
                        <div style="display: flex; gap: 8px; justify-content: flex-end;">
                            <button class="cancel-btn" onclick="cancelEditPubSub(${index})">Cancel</button>
                            <button class="save-btn" onclick="savePubSubConfig(${index})">Save</button>
                        </div>
                    </div>
                </div>
            ` + "`" + `).join('');
        }

        function renderKafkaConfigs() {
            const container = document.getElementById('kafkaConfigs');
            container.innerHTML = configs.kafkaConfigs.map((config, index) => ` + "`" + `
                <div class="config-item" id="kafka-${index}">
                    <div class="config-item-header">
                        <div class="config-name">${config.name}</div>
                        <div class="config-actions">
                            <button class="edit-btn" onclick="editKafkaConfig(${index})">Edit</button>
                            <button class="delete-btn" onclick="deleteKafkaConfig(${index})">Delete</button>
                        </div>
                    </div>
                    <div class="config-details" id="kafka-details-${index}">
                        <div><strong>Brokers:</strong> ${config.brokers}</div>
                        <div><strong>Topic:</strong> ${config.topic}</div>
                        <div><strong>Consumer Group:</strong> ${config.consumerGroup}</div>
                        <div><strong>Schema Registry:</strong> ${config.schemaRegistry}</div>
                    </div>
                    <div class="hidden" id="kafka-form-${index}">
                        <div class="form-group">
                            <label class="form-label">Configuration Name</label>
                            <input type="text" class="form-input" id="edit-kafka-name-${index}" value="${config.name}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Brokers</label>
                            <input type="text" class="form-input" id="edit-kafka-brokers-${index}" value="${config.brokers}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Topic</label>
                            <input type="text" class="form-input" id="edit-kafka-topic-${index}" value="${config.topic}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Consumer Group</label>
                            <input type="text" class="form-input" id="edit-kafka-group-${index}" value="${config.consumerGroup}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Schema Registry URL</label>
                            <input type="text" class="form-input" id="edit-kafka-schema-${index}" value="${config.schemaRegistry}">
                        </div>
                        <div style="display: flex; gap: 8px; justify-content: flex-end;">
                            <button class="cancel-btn" onclick="cancelEditKafka(${index})">Cancel</button>
                            <button class="save-btn" onclick="saveKafkaConfig(${index})">Save</button>
                        </div>
                    </div>
                </div>
            ` + "`" + `).join('');
        }

        function renderSpannerConfigs() {
            const container = document.getElementById('spannerConfigs');
            if (!configs.spannerConfigs) configs.spannerConfigs = [];
            container.innerHTML = configs.spannerConfigs.map((config, index) => ` + "`" + `
                <div class="config-item" id="spanner-${index}">
                    <div class="config-item-header">
                        <div class="config-name">${config.name}</div>
                        <div class="config-actions">
                            <button class="edit-btn" onclick="editSpannerConfig(${index})">Edit</button>
                            <button class="delete-btn" onclick="deleteSpannerConfig(${index})">Delete</button>
                        </div>
                    </div>
                    <div class="config-details" id="spanner-details-${index}">
                        <div><strong>Emulator Host:</strong> ${config.emulatorHost}</div>
                        <div><strong>Project ID:</strong> ${config.projectId}</div>
                        <div><strong>Instance ID:</strong> ${config.instanceId}</div>
                        <div><strong>Database ID:</strong> ${config.databaseId}</div>
                    </div>
                    <div class="hidden" id="spanner-form-${index}">
                        <div class="form-group">
                            <label class="form-label">Configuration Name</label>
                            <input type="text" class="form-input" id="edit-spanner-name-${index}" value="${config.name}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Emulator Host</label>
                            <input type="text" class="form-input" id="edit-spanner-host-${index}" value="${config.emulatorHost}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Project ID</label>
                            <input type="text" class="form-input" id="edit-spanner-project-${index}" value="${config.projectId}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Instance ID</label>
                            <input type="text" class="form-input" id="edit-spanner-instance-${index}" value="${config.instanceId}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Database ID</label>
                            <input type="text" class="form-input" id="edit-spanner-database-${index}" value="${config.databaseId}">
                        </div>
                        <div style="display: flex; gap: 8px; justify-content: flex-end;">
                            <button class="cancel-btn" onclick="cancelEditSpanner(${index})">Cancel</button>
                            <button class="save-btn" onclick="saveSpannerConfig(${index})">Save</button>
                        </div>
                    </div>
                </div>
            ` + "`" + `).join('');
        }

        function editPubSubConfig(index) {
            document.getElementById(` + "`pubsub-details-${index}`" + `).classList.add('hidden');
            document.getElementById(` + "`pubsub-form-${index}`" + `).classList.remove('hidden');
        }

        function cancelEditPubSub(index) {
            document.getElementById(` + "`pubsub-details-${index}`" + `).classList.remove('hidden');
            document.getElementById(` + "`pubsub-form-${index}`" + `).classList.add('hidden');
        }

        async function savePubSubConfig(index) {
            const updatedConfig = {
                name: document.getElementById(` + "`edit-pubsub-name-${index}`" + `).value,
                emulatorHost: document.getElementById(` + "`edit-pubsub-host-${index}`" + `).value,
                projectId: document.getElementById(` + "`edit-pubsub-project-${index}`" + `).value,
                subscriptionId: document.getElementById(` + "`edit-pubsub-sub-${index}`" + `).value
            };

            configs.pubsubConfigs[index] = updatedConfig;
            await saveConfigs();
            showSuccess('PubSub configuration updated successfully!');
        }

        function editKafkaConfig(index) {
            document.getElementById(` + "`kafka-details-${index}`" + `).classList.add('hidden');
            document.getElementById(` + "`kafka-form-${index}`" + `).classList.remove('hidden');
        }

        function cancelEditKafka(index) {
            document.getElementById(` + "`kafka-details-${index}`" + `).classList.remove('hidden');
            document.getElementById(` + "`kafka-form-${index}`" + `).classList.add('hidden');
        }

        async function saveKafkaConfig(index) {
            const updatedConfig = {
                name: document.getElementById(` + "`edit-kafka-name-${index}`" + `).value,
                brokers: document.getElementById(` + "`edit-kafka-brokers-${index}`" + `).value,
                topic: document.getElementById(` + "`edit-kafka-topic-${index}`" + `).value,
                consumerGroup: document.getElementById(` + "`edit-kafka-group-${index}`" + `).value,
                schemaRegistry: document.getElementById(` + "`edit-kafka-schema-${index}`" + `).value
            };

            configs.kafkaConfigs[index] = updatedConfig;
            await saveConfigs();
            showSuccess('Kafka configuration updated successfully!');
        }

        async function deletePubSubConfig(index) {
            if (!confirm('Are you sure you want to delete this configuration?')) return;
            configs.pubsubConfigs.splice(index, 1);
            await saveConfigs();
            showSuccess('PubSub configuration deleted successfully!');
        }

        async function deleteKafkaConfig(index) {
            if (!confirm('Are you sure you want to delete this configuration?')) return;
            configs.kafkaConfigs.splice(index, 1);
            await saveConfigs();
            showSuccess('Kafka configuration deleted successfully!');
        }

        function editSpannerConfig(index) {
            document.getElementById(` + "`spanner-details-${index}`" + `).classList.add('hidden');
            document.getElementById(` + "`spanner-form-${index}`" + `).classList.remove('hidden');
        }

        function cancelEditSpanner(index) {
            document.getElementById(` + "`spanner-details-${index}`" + `).classList.remove('hidden');
            document.getElementById(` + "`spanner-form-${index}`" + `).classList.add('hidden');
        }

        async function saveSpannerConfig(index) {
            const updatedConfig = {
                name: document.getElementById(` + "`edit-spanner-name-${index}`" + `).value,
                emulatorHost: document.getElementById(` + "`edit-spanner-host-${index}`" + `).value,
                projectId: document.getElementById(` + "`edit-spanner-project-${index}`" + `).value,
                instanceId: document.getElementById(` + "`edit-spanner-instance-${index}`" + `).value,
                databaseId: document.getElementById(` + "`edit-spanner-database-${index}`" + `).value
            };

            configs.spannerConfigs[index] = updatedConfig;
            await saveConfigs();
            showSuccess('Spanner configuration updated successfully!');
        }

        async function deleteSpannerConfig(index) {
            if (!confirm('Are you sure you want to delete this configuration?')) return;
            configs.spannerConfigs.splice(index, 1);
            await saveConfigs();
            showSuccess('Spanner configuration deleted successfully!');
        }

        function showNewConfigForm(type) {
            if (type === 'pubsub') {
                document.getElementById('newPubSubForm').classList.add('visible');
            } else if (type === 'kafka') {
                document.getElementById('newKafkaForm').classList.add('visible');
            } else if (type === 'spanner') {
                document.getElementById('newSpannerForm').classList.add('visible');
            }
        }

        function hideNewConfigForm(type) {
            if (type === 'pubsub') {
                document.getElementById('newPubSubForm').classList.remove('visible');
                clearPubSubForm();
            } else if (type === 'kafka') {
                document.getElementById('newKafkaForm').classList.remove('visible');
                clearKafkaForm();
            } else if (type === 'spanner') {
                document.getElementById('newSpannerForm').classList.remove('visible');
                clearSpannerForm();
            }
        }

        function clearPubSubForm() {
            document.getElementById('newPubSubName').value = '';
            document.getElementById('newPubSubHost').value = '';
            document.getElementById('newPubSubProject').value = '';
            document.getElementById('newPubSubSub').value = '';
        }

        function clearKafkaForm() {
            document.getElementById('newKafkaName').value = '';
            document.getElementById('newKafkaBrokers').value = '';
            document.getElementById('newKafkaTopic').value = '';
            document.getElementById('newKafkaGroup').value = '';
            document.getElementById('newKafkaSchema').value = '';
        }

        function clearSpannerForm() {
            document.getElementById('newSpannerName').value = '';
            document.getElementById('newSpannerHost').value = '';
            document.getElementById('newSpannerProject').value = '';
            document.getElementById('newSpannerInstance').value = '';
            document.getElementById('newSpannerDatabase').value = '';
        }

        async function saveNewPubSubConfig() {
            const newConfig = {
                name: document.getElementById('newPubSubName').value,
                emulatorHost: document.getElementById('newPubSubHost').value,
                projectId: document.getElementById('newPubSubProject').value,
                subscriptionId: document.getElementById('newPubSubSub').value
            };

            if (!newConfig.name || !newConfig.emulatorHost || !newConfig.projectId || !newConfig.subscriptionId) {
                alert('Please fill in all fields');
                return;
            }

            configs.pubsubConfigs.push(newConfig);
            await saveConfigs();
            hideNewConfigForm('pubsub');
            showSuccess('New PubSub configuration added successfully!');
        }

        async function saveNewKafkaConfig() {
            const newConfig = {
                name: document.getElementById('newKafkaName').value,
                brokers: document.getElementById('newKafkaBrokers').value,
                topic: document.getElementById('newKafkaTopic').value,
                consumerGroup: document.getElementById('newKafkaGroup').value,
                schemaRegistry: document.getElementById('newKafkaSchema').value
            };

            if (!newConfig.name || !newConfig.brokers || !newConfig.topic || !newConfig.consumerGroup || !newConfig.schemaRegistry) {
                alert('Please fill in all fields');
                return;
            }

            configs.kafkaConfigs.push(newConfig);
            await saveConfigs();
            hideNewConfigForm('kafka');
            showSuccess('New Kafka configuration added successfully!');
        }

        async function saveNewSpannerConfig() {
            const newConfig = {
                name: document.getElementById('newSpannerName').value,
                emulatorHost: document.getElementById('newSpannerHost').value,
                projectId: document.getElementById('newSpannerProject').value,
                instanceId: document.getElementById('newSpannerInstance').value,
                databaseId: document.getElementById('newSpannerDatabase').value
            };

            if (!newConfig.name || !newConfig.emulatorHost || !newConfig.projectId || !newConfig.instanceId || !newConfig.databaseId) {
                alert('Please fill in all fields');
                return;
            }

            configs.spannerConfigs.push(newConfig);
            await saveConfigs();
            hideNewConfigForm('spanner');
            showSuccess('New Spanner configuration added successfully!');
        }

        function renderGCSConfigs() {
            const container = document.getElementById('gcsConfigs');
            if (!configs.gcsConfigs) configs.gcsConfigs = [];
            container.innerHTML = configs.gcsConfigs.map((config, index) => ` + "`" + `
                <div class="config-item" id="gcs-${index}">
                    <div class="config-item-header">
                        <div class="config-name">${config.name}</div>
                        <div class="config-actions">
                            <button class="edit-btn" onclick="editGCSConfig(${index})">Edit</button>
                            <button class="delete-btn" onclick="deleteGCSConfig(${index})">Delete</button>
                        </div>
                    </div>
                    <div class="config-details" id="gcs-details-${index}">
                        <div><strong>Emulator Host:</strong> ${config.emulatorHost}</div>
                        <div><strong>Project ID:</strong> ${config.projectId}</div>
                    </div>
                    <div class="hidden" id="gcs-form-${index}">
                        <div class="form-group">
                            <label class="form-label">Configuration Name</label>
                            <input type="text" class="form-input" id="edit-gcs-name-${index}" value="${config.name}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Emulator Host</label>
                            <input type="text" class="form-input" id="edit-gcs-host-${index}" value="${config.emulatorHost}">
                        </div>
                        <div class="form-group">
                            <label class="form-label">Project ID</label>
                            <input type="text" class="form-input" id="edit-gcs-project-${index}" value="${config.projectId}">
                        </div>
                        <div style="display: flex; gap: 8px; justify-content: flex-end;">
                            <button class="cancel-btn" onclick="cancelEditGCS(${index})">Cancel</button>
                            <button class="save-btn" onclick="saveGCSConfig(${index})">Save</button>
                        </div>
                    </div>
                </div>
            ` + "`" + `).join('');
        }

        function editGCSConfig(index) {
            document.getElementById('gcs-details-' + index).classList.add('hidden');
            document.getElementById('gcs-form-' + index).classList.remove('hidden');
        }

        function cancelEditGCS(index) {
            document.getElementById('gcs-details-' + index).classList.remove('hidden');
            document.getElementById('gcs-form-' + index).classList.add('hidden');
        }

        async function saveGCSConfig(index) {
            configs.gcsConfigs[index] = {
                name: document.getElementById('edit-gcs-name-' + index).value,
                emulatorHost: document.getElementById('edit-gcs-host-' + index).value,
                projectId: document.getElementById('edit-gcs-project-' + index).value
            };
            await saveConfigs();
            showSuccess('GCS configuration updated successfully!');
        }

        async function deleteGCSConfig(index) {
            if (confirm('Are you sure you want to delete this GCS configuration?')) {
                configs.gcsConfigs.splice(index, 1);
                await saveConfigs();
                showSuccess('GCS configuration deleted successfully!');
            }
        }

        async function saveNewGCSConfig() {
            const newConfig = {
                name: document.getElementById('newGCSName').value,
                emulatorHost: document.getElementById('newGCSHost').value,
                projectId: document.getElementById('newGCSProject').value
            };

            if (!newConfig.name || !newConfig.emulatorHost || !newConfig.projectId) {
                alert('Please fill in all fields');
                return;
            }

            configs.gcsConfigs.push(newConfig);
            await saveConfigs();
            hideNewConfigForm('gcs');
            showSuccess('New GCS configuration added successfully!');
        }

        async function saveConfigs() {
            const response = await fetch('/api/configs/save', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(configs)
            });

            if (response.ok) {
                await loadConfigs();
            } else {
                alert('Failed to save configurations');
            }
        }

        function showSuccess(message) {
            const msgEl = document.getElementById('successMessage');
            msgEl.textContent = message;
            msgEl.classList.add('visible');
            setTimeout(() => {
                msgEl.classList.remove('visible');
            }, 3000);
        }

        loadConfigs();
    </script>
</body>
</html>`
