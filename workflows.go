package slack

import (
	"context"
)

// Workflow interaction events that can be triggered by Slack
const (
	WorkflowEventStepExec    = "workflow_step_execute"
	WorkflowEventStepEdit    = "workflow_step_edit"
	WorkflowEventUnpublished = "workflow_unpublished"
	WorkflowEventStepDeleted = "workflow_step_deleted"
	WorkflowEventPublished   = "workflow_published"
	WorkflowEventDeleted     = "workflow_deleted"
)

// WorkflowInput - https://api.slack.com/reference/workflows/workflow_step#input
type WorkflowInput struct {
	Value                   string            `json:"value"`
	Variables               map[string]string `json:"variables,omitempty"`
	SkipVariableReplacement bool              `json:"skip_variable_replacement,omitempty"`
}

// WorkflowOutput - https://api.slack.com/reference/workflows/workflow_step#output
type WorkflowOutput struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

// WorkflowRequest - https://api.slack.com/methods/workflows.updateStep#args
type WorkflowRequest struct {
	TriggerID    string                    `json:"workflow_step_edit_id"`
	StepName     string                    `json:"step_name,omitempty"`
	StepImageURL string                    `json:"step_image_url,omitempty"`
	Outputs      []*WorkflowOutput         `json:"outputs,omitempty"`
	Inputs       map[string]*WorkflowInput `json:"inputs,omitempty"`
}

//NewWorkflowRequest create a new workflow request object. Requires a trigger_id from a ViewRequest object that was
// sent as part of a `workflow_step_execute` type interaction.
func NewWorkflowRequest(triggerID string) WorkflowRequest {
	return WorkflowRequest{
		TriggerID:    triggerID,
		StepName:     "",
		StepImageURL: "",
		Outputs:      []*WorkflowOutput{},
		Inputs:       map[string]*WorkflowInput{},
	}
}

// WorkflowUpdateStep notifies Slack that the workflow update in question has been handled
func (api *Client) WorkflowUpdateStep(request WorkflowRequest) (*SlackResponse, error) {
	return api.WorkflowUpdateStepContext(context.Background(), request)
}

// WorkflowUpdateStepContext notifies Slack that the workflow update in question has been handled
// https://api.slack.com/methods/workflows.updateStep
func (api *Client) WorkflowUpdateStepContext(ctx context.Context, request WorkflowRequest) (*SlackResponse, error) {
	response := &SlackResponse{}
	if err := api.postJSON(ctx, "workflows.updateStep", request, response); err != nil {
		return response, err
	}
	return response, response.Err()
}
