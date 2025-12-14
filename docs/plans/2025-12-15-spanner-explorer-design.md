# Spanner Database Explorer Design

**Date:** 2025-12-15
**Purpose:** Add Spanner database explorer for local emulator testing
**Scope:** Personal testing tool (non-production)

## Overview

Add a Spanner database explorer to Testing Studio that allows browsing tables, running SQL queries (SELECT and DML), and testing connections to the local Spanner emulator. The UI will follow the Spanner Studio design pattern with a table browser sidebar and SQL editor.

## User Interface

### Homepage Integration

**New Card:**
- Add "Spanner Explorer" card on homepage alongside existing tools (PubSub, Kafka, GCS, REST Client)
- Card description: "Browse tables, run queries, test local Spanner emulator"
- Links to `/spanner` route
- Uses same CSS styling as existing tool cards

### Main Page Layout (Spanner Studio Style)

**Top Bar - Connection Settings:**
- Collapsible configuration panel
- Profile selector dropdown (saved configs)
- Editable fields:
  - Emulator Host
  - Project ID
  - Instance ID
  - Database ID
- Buttons: "Connect", "Save Config", "New Config"
- Connection status indicator (green/red dot with text)

**Left Sidebar - Table Browser:**
- Fixed width (~250px), scrollable
- Search/filter box at top
- List of all tables (clickable)
- Shows table count badge
- Clicking a table:
  - Loads schema information
  - Generates SELECT query in editor

**Main Content Area:**

*SQL Editor (Top Half):*
- Large textarea with monospace font
- Basic SQL syntax highlighting
- "Run Query" button
- Query examples dropdown
- Placeholder with sample query

*Results Panel (Bottom Half):*
- HTML table displaying query results
- Row count and execution time
- Scrollable for large datasets
- Error messages displayed in red alert
- Success messages in green banner

## Configuration Management

### configs.json Structure

```json
{
  "spanner": [
    {
      "name": "TMS Local",
      "emulatorHost": "localhost:9010",
      "projectId": "tms-suncorp-local",
      "instanceId": "tms-suncorp-local",
      "databaseId": "tms-suncorp-db"
    }
  ]
}
```

### Environment Variable Auto-Population

On first visit, pre-fill form from environment variables:
- `SPANNER_EMULATOR_HOST` → emulatorHost
- `SPANNER_PROJECT` → projectId
- `SPANNER_INSTANCE` → instanceId
- `SPANNER_DATABASE` → databaseId

If configs exist in configs.json, load first profile by default.

### Configuration Features

- Multiple connection profiles saved in configs.json
- UI form to create/edit/select profiles
- Save/update configurations via "Save Config" button
- Environment variables as fallback/defaults

## Backend Implementation

### File Structure

```
internal/
├── spanner/
│   └── spanner.go          # Spanner client operations
├── handlers/
│   └── spanner.go          # HTTP handlers
├── templates/
│   └── spanner.go          # HTML template
└── types/
    └── spanner.go          # Spanner types
```

### Routes

- `GET /spanner` - Main Spanner explorer page
- `POST /spanner/query` - Execute SQL queries
- `GET /spanner/tables` - List tables in database (JSON)
- `GET /spanner/schema/:table` - Get schema for table (JSON)
- `POST /spanner/execute` - Execute DML statements

### Core Functions (internal/spanner/spanner.go)

- `Connect()` - Create Spanner client with emulator host
- `ListTables()` - Query INFORMATION_SCHEMA for table list
- `GetTableSchema()` - Get column definitions for a table
- `ExecuteQuery()` - Run SELECT queries, return rows
- `ExecuteDML()` - Run INSERT/UPDATE/DELETE statements

### Dependencies

```bash
go get cloud.google.com/go/spanner
go get google.golang.org/api/iterator
```

### Implementation Details

- Use `cloud.google.com/go/spanner` official client
- Set `SPANNER_EMULATOR_HOST` environment variable
- Parse query results into `[]map[string]interface{}` for JSON
- Connection pooling and error handling
- Read configurations from configs.json

## Frontend Implementation

### HTML Template

- Extends `base.go` template for consistency
- Uses existing Material Design CSS classes
- Responsive grid layout (sidebar + main content)
- Same styling as PubSub/Kafka pages

### JavaScript Functionality

- Vanilla JavaScript (no frameworks)
- AJAX calls using fetch API
- Table click handlers
- Form submission for connection
- Query execution with loading states
- Real-time UI updates

### Features

**Table Schema Viewer:**
- Click table → display schema in modal/panel
- Shows: Column name, Type, Nullable, Primary Key
- "Query this table" button generates SELECT query

**Query History:**
- Store last 5 queries in localStorage
- Dropdown to re-run previous queries
- Clear history button

**Export Results:**
- "Copy as JSON" button
- Optional: "Download CSV" for larger datasets

**Sample Queries Dropdown:**
- List all tables
- Count rows in selected table
- Show table schema
- INSERT/UPDATE templates

## Error Handling

### Connection Errors

- Emulator not running → Clear message with suggestion
- Invalid configuration → Specific error details
- Network issues → Display raw error

### Query Errors

- SQL syntax errors → Show Spanner error in red alert
- Permission errors → Display error details
- Timeout → "Query timed out after 30 seconds"

### UI Feedback

- Loading spinners during execution
- Disable buttons during operations
- Toast notifications for actions
- Warning before DELETE/DROP queries

## Testing Checklist

1. Start Spanner emulator: `docker run -p 9010:9010 gcr.io/cloud-spanner-emulator/emulator`
2. Set environment variables
3. Start testing-studio: `go run cmd/server/main.go`
4. Verify homepage card appears
5. Test connection with env var defaults
6. Save config to configs.json
7. Load tables from database
8. Click table, verify schema loads
9. Run SELECT query
10. Run INSERT/UPDATE/DELETE
11. Test error cases

## Integration Points

### Existing Code Modifications

- Add Spanner card to `internal/templates/index.go`
- Register routes in `cmd/server/main.go`
- Add Spanner config to `internal/config/config.go`
- Follow existing handler patterns

### No Breaking Changes

- All changes are additive
- Doesn't modify existing tools
- configs.json extended with new section
- New routes, no conflicts

## Future Enhancements (Not Initial Version)

- Mutation API support
- Transaction management
- Index information display
- Query plan visualization
- Multi-statement execution
- Query performance metrics

## Success Criteria

- Can connect to local Spanner emulator
- Can browse tables and view schemas
- Can run SELECT queries and see results
- Can execute DML (INSERT/UPDATE/DELETE)
- Configurations persist in configs.json
- UI matches existing tool aesthetics
- Error handling works gracefully
