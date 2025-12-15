package types

// QueryRequest represents a SQL query request
type QueryRequest struct {
	EmulatorHost string `json:"emulatorHost"`
	ProjectID    string `json:"projectId"`
	InstanceID   string `json:"instanceId"`
	DatabaseID   string `json:"databaseId"`
	Query        string `json:"query"`
}

// QueryResponse represents the result of a SQL query
type QueryResponse struct {
	Columns      []string                 `json:"columns"`
	Rows         []map[string]interface{} `json:"rows"`
	RowCount     int                      `json:"rowCount"`
	ExecutionTime string                   `json:"executionTime"`
	Error        string                   `json:"error,omitempty"`
}

// TableInfo represents metadata about a table
type TableInfo struct {
	Name       string       `json:"name"`
	RowCount   int64        `json:"rowCount,omitempty"`
	Columns    []ColumnInfo `json:"columns,omitempty"`
}

// ColumnInfo represents metadata about a column
type ColumnInfo struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	IsNullable bool   `json:"isNullable"`
	IsPrimaryKey bool `json:"isPrimaryKey"`
}

// ConnectionRequest represents a connection test request
type ConnectionRequest struct {
	EmulatorHost string `json:"emulatorHost"`
	ProjectID    string `json:"projectId"`
	InstanceID   string `json:"instanceId"`
	DatabaseID   string `json:"databaseId"`
}

// ConnectionResponse represents the result of a connection test
type ConnectionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
