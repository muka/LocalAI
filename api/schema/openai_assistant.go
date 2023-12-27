package schema

// Assistant represents an assistant that can call the model and use tools.
type Assistant struct {
	ID           string            `json:"id" yaml:"id"`
	Object       string            `json:"object" yaml:"object"`
	CreatedAt    int64             `json:"created_at" yaml:"created_at"`
	Name         string            `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	Model        string            `json:"model" yaml:"model"`
	Instructions string            `json:"instructions,omitempty" yaml:"instructions,omitempty"`
	Tools        []Tool            `json:"tools" yaml:"tools"`
	FileIDs      []string          `json:"file_ids" yaml:"file_ids"`
	Metadata     map[string]string `json:"metadata" yaml:"metadata"`
}

// Tool represents a tool enabled on the assistant.
type Tool struct {
	Type string `json:"type" yaml:"type"`
	// Add more fields specific to each tool type as needed.
}

type File struct {
	ID          string `json:"id" yaml:"id"`
	Object      string `json:"object" yaml:"object"`
	CreatedAt   int64  `json:"created_at" yaml:"created_at"`
	AssistantID string `json:"assistant_id" yaml:"assistant_id"`
}

type FilesList struct {
	Files []File `json:"files" yaml:"files"`
}

// AssistantCreateRequest represents the request object for creating an assistant
type AssistantCreateRequest struct {
	Model        string            `json:"model" yaml:"model"`
	Name         string            `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	Instructions string            `json:"instructions,omitempty" yaml:"instructions,omitempty"`
	Tools        []string          `json:"tools,omitempty" yaml:"tools,omitempty"`
	FileIDs      []string          `json:"file_ids,omitempty" yaml:"file_ids,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// AssistantFileCreateRequest represents the request object for creating an assistant file
type AssistantFileCreateRequest struct {
	AssistantID string `json:"assistant_id" yaml:"assistant_id"`
	FileID      string `json:"file_id" yaml:"file_id"`
}

// AssistantListResponse represents the response object for listing assistants
type AssistantListResponse struct {
	Object  string      `json:"object"`
	Data    []Assistant `json:"data"`
	FirstID string      `json:"first_id"`
	LastID  string      `json:"last_id"`
	HasMore bool        `json:"has_more"`
}

// AssistantFileListResponse represents the response object for listing assistant files
type AssistantFileListResponse struct {
	Object  string `json:"object"`
	Data    []File `json:"data"`
	FirstID string `json:"first_id"`
	LastID  string `json:"last_id"`
	HasMore bool   `json:"has_more"`
}

// AssistantModifyRequest represents the request object for modifying an assistant
type AssistantModifyRequest struct {
	Model        string            `json:"model,omitempty" yaml:"model,omitempty"`
	Name         string            `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	Instructions string            `json:"instructions,omitempty" yaml:"instructions,omitempty"`
	Tools        []string          `json:"tools,omitempty" yaml:"tools,omitempty"`
	FileIDs      []string          `json:"file_ids,omitempty" yaml:"file_ids,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// AssistantFileDeleteResponse represents the response object for deleting an assistant file
type AssistantFileDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}
