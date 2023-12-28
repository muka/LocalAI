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

// Thread represents the thread object
type Thread struct {
	ID        string            `json:"id"`
	Object    string            `json:"object"`
	CreatedAt int64             `json:"created_at"`
	Metadata  map[string]string `json:"metadata"`
}

type FilesList struct {
	Files []File `json:"files" yaml:"files"`
}

// ThreadMessage represents the thread message object
type ThreadMessage struct {
	ID          string                 `json:"id" yaml:"id"`
	Object      string                 `json:"object" yaml:"object"`
	CreatedAt   int64                  `json:"created_at" yaml:"created_at"`
	ThreadID    string                 `json:"thread_id" yaml:"thread_id"`
	Role        string                 `json:"role" yaml:"role"`
	Content     []ThreadMessageContent `json:"content" yaml:"content"`
	FileIDs     []string               `json:"file_ids" yaml:"file_ids"`
	AssistantID string                 `json:"assistant_id" yaml:"assistant_id"`
	RunID       string                 `json:"run_id" yaml:"run_id"`
	Metadata    map[string]string      `json:"metadata" yaml:"metadata"`
}

// ThreadMessageFile represents the thread message file object
type ThreadMessageFile struct {
	ID        string `json:"id" yaml:"id"`
	Object    string `json:"object" yaml:"object"`
	CreatedAt int64  `json:"created_at" yaml:"created_at"`
	MessageID string `json:"message_id" yaml:"message_id"`
	FileID    string `json:"file_id" yaml:"file_id"`
}

// ThreadMessageContent represents the content within a thread message
type ThreadMessageContent struct {
	// text or image_file
	Type string `json:"type" yaml:"type"`

	Text struct {
		Value       string                    `json:"value" yaml:"value"`
		Annotations []ThreadMessageAnnotation `json:"annotations" yaml:"annotations"`
	} `json:"text,omitempty" yaml:"text,omitempty"`

	ImageFile struct {
		FileID string `json:"file_id" yaml:"file_id"`
	} `json:"image_file,omitempty" yaml:"image_file,omitempty"`
}

// ThreadMessageAnnotation represents the annotation within a thread message content
type ThreadMessageAnnotation struct {
	Start int    `json:"start" yaml:"start"`
	End   int    `json:"end" yaml:"end"`
	Label string `json:"label" yaml:"label"`
}

// Run represents an execution run on a thread
type Run struct {
	ID             string            `json:"id" yaml:"id"`
	Object         string            `json:"object" yaml:"object"`
	CreatedAt      int64             `json:"created_at" yaml:"created_at"`
	AssistantID    string            `json:"assistant_id" yaml:"assistant_id"`
	ThreadID       string            `json:"thread_id" yaml:"thread_id"`
	Status         string            `json:"status" yaml:"status"`
	RequiredAction *RequiredAction   `json:"required_action,omitempty" yaml:"required_action,omitempty"`
	LastError      *LastError        `json:"last_error,omitempty" yaml:"last_error,omitempty"`
	ExpiresAt      *int64            `json:"expires_at,omitempty" yaml:"expires_at,omitempty"`
	StartedAt      *int64            `json:"started_at,omitempty" yaml:"started_at,omitempty"`
	CancelledAt    *int64            `json:"cancelled_at,omitempty" yaml:"cancelled_at,omitempty"`
	FailedAt       *int64            `json:"failed_at,omitempty" yaml:"failed_at,omitempty"`
	CompletedAt    *int64            `json:"completed_at,omitempty" yaml:"completed_at,omitempty"`
	Model          string            `json:"model" yaml:"model"`
	Instructions   string            `json:"instructions,omitempty" yaml:"instructions,omitempty"`
	Tools          []Tool            `json:"tools,omitempty" yaml:"tools,omitempty"`
	FileIDs        []string          `json:"file_ids,omitempty" yaml:"file_ids,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// RequiredAction represents details on the action required to continue the run
type RequiredAction struct {
	// Add properties as needed based on the actual structure
}

// LastError represents the last error associated with a run
type LastError struct {
	// Add properties as needed based on the actual structure
}

// RunStep represents a step in the execution of a run
type RunStep struct {
	ID          string            `json:"id" yaml:"id"`
	Object      string            `json:"object" yaml:"object"`
	CreatedAt   int64             `json:"created_at" yaml:"created_at"`
	AssistantID string            `json:"assistant_id" yaml:"assistant_id"`
	ThreadID    string            `json:"thread_id" yaml:"thread_id"`
	RunID       string            `json:"run_id" yaml:"run_id"`
	Type        string            `json:"type" yaml:"type"`
	Status      string            `json:"status" yaml:"status"`
	ExpiredAt   *int64            `json:"expired_at,omitempty" yaml:"expired_at,omitempty"`
	CancelledAt *int64            `json:"cancelled_at,omitempty" yaml:"cancelled_at,omitempty"`
	FailedAt    *int64            `json:"failed_at,omitempty" yaml:"failed_at,omitempty"`
	CompletedAt *int64            `json:"completed_at,omitempty" yaml:"completed_at,omitempty"`
	LastError   *LastError        `json:"last_error,omitempty" yaml:"last_error,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	StepDetails *StepDetails      `json:"step_details" yaml:"step_details"`
}

// StepDetails represents the details of the run step
type StepDetails struct {
	Type            string               `json:"type" yaml:"type"`
	MessageCreation *MessageCreationStep `json:"message_creation,omitempty" yaml:"message_creation,omitempty"`
	// Add more properties for other types as needed
}

// MessageCreationStep represents the details of a message creation step
type MessageCreationStep struct {
	MessageID string `json:"message_id" yaml:"message_id"`
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

// CreateThreadRequest represents the request object for creating a thread
type CreateThreadRequest struct {
	Messages []ThreadMessage   `json:"messages,omitempty" yaml:"messages,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// CreateThreadResponse represents the response object for creating a thread
type CreateThreadResponse struct {
	ID        string                 `json:"id" yaml:"id"`
	Object    string                 `json:"object" yaml:"object"`
	CreatedAt int64                  `json:"created_at" yaml:"created_at"`
	Messages  []ThreadMessageContent `json:"messages" yaml:"messages"`
	Metadata  map[string]string      `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// RetrieveThreadResponse represents the response object for retrieving a thread
type RetrieveThreadResponse struct {
	ID        string                 `json:"id" yaml:"id"`
	Object    string                 `json:"object" yaml:"object"`
	CreatedAt int64                  `json:"created_at" yaml:"created_at"`
	Messages  []ThreadMessageContent `json:"messages" yaml:"messages"`
	Metadata  map[string]string      `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// ModifyThreadRequest represents the request object for modifying a thread
type ModifyThreadRequest struct {
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// ModifyThreadResponse represents the response object for modifying a thread
type ModifyThreadResponse struct {
	ID        string                 `json:"id" yaml:"id"`
	Object    string                 `json:"object" yaml:"object"`
	CreatedAt int64                  `json:"created_at" yaml:"created_at"`
	Messages  []ThreadMessageContent `json:"messages" yaml:"messages"`
	Metadata  map[string]string      `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// DeleteThreadResponse represents the response object for deleting a thread
type DeleteThreadResponse struct {
	ID      string `json:"id" yaml:"id"`
	Object  string `json:"object" yaml:"object"`
	Deleted bool   `json:"deleted" yaml:"deleted"`
}

// CreateMessageRequest represents the request object for creating a message
type CreateMessageRequest struct {
	Role     string            `json:"role" yaml:"role"`
	Content  string            `json:"content" yaml:"content"`
	FileIDs  []string          `json:"file_ids,omitempty" yaml:"file_ids,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// ListMessagesResponse represents the response object for listing messages
type ListMessagesResponse struct {
	Object  string          `json:"object" yaml:"object"`
	Data    []ThreadMessage `json:"data" yaml:"data"`
	FirstID string          `json:"first_id" yaml:"first_id"`
	LastID  string          `json:"last_id" yaml:"last_id"`
	HasMore bool            `json:"has_more" yaml:"has_more"`
}

// ListMessageFilesResponse represents the response object for listing message files
type ListMessageFilesResponse struct {
	Object  string              `json:"object" yaml:"object"`
	Data    []ThreadMessageFile `json:"data" yaml:"data"`
	FirstID string              `json:"first_id" yaml:"first_id"`
	LastID  string              `json:"last_id" yaml:"last_id"`
	HasMore bool                `json:"has_more" yaml:"has_more"`
}

// ModifyMessageRequest represents the request object for modifying a message
type ModifyMessageRequest struct {
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// CreateRunRequest represents the request object for creating a run
type CreateRunRequest struct {
	AssistantID            string            `json:"assistant_id" yaml:"assistant_id"`
	Model                  *string           `json:"model,omitempty" yaml:"model,omitempty"`
	Instructions           *string           `json:"instructions,omitempty" yaml:"instructions,omitempty"`
	AdditionalInstructions *string           `json:"additional_instructions,omitempty" yaml:"additional_instructions,omitempty"`
	Tools                  *[]string         `json:"tools,omitempty" yaml:"tools,omitempty"`
	Metadata               map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

type CreateThreadAndRunRequest struct {
	AssistantID string  `json:"assistant_id" yaml:"assistant_id"`
	Thread      *Thread `json:"thread,omitempty" yaml:"thread,omitempty"`
}

// ModifyRunRequest represents the request body for modifying a run
type ModifyRunRequest struct {
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// SubmitToolOutputsRequest represents the request body for submitting tool outputs to a run
type SubmitToolOutputsRequest struct {
	ToolOutputs []ToolOutput `json:"tool_outputs" yaml:"tool_outputs"`
}

// ToolOutput represents the output of a tool call to be submitted to continue the run
type ToolOutput struct {
	ToolCallID string `json:"tool_call_id,omitempty" yaml:"tool_call_id,omitempty"`
	Output     string `json:"output,omitempty" yaml:"output,omitempty"`
}
