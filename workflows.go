package slack

import (
	"context"
	"encoding/json"
	"net/url"
)

// https://api.slack.com/reference/workflows/workflow_step#input
type WorkflowInput struct {
	Value                   string            `json:"value,omitempty"`
	Variables               map[string]string `json:"variables,omitempty"`
	SkipVariableReplacement bool              `json:"skip_variable_replacement"`
}

// https://api.slack.com/reference/workflows/workflow_step#output
type WorkflowOutput struct {
	Name  string
	Type  string
	Label string
}

type WorkflowRef struct {
	StepName     string
	StepImageURL string
	Outputs      []*WorkflowOutput
	Inputs       map[string]*WorkflowInput
}

func (api *Client) WorkflowUpdateStep(channel string, item ItemRef) error {
	return api.AddStarContext(context.Background(), channel, item)
}

// https://api.slack.com/methods/workflows.updateStep
func (api *Client) WorkflowUpdateStepContext(ctx context.Context, workflowStepEditID string, wf WorkflowRef) error {
	values := url.Values{
		"workflow_step_edit_id": {workflowStepEditID},
		"token":                 {api.token},
	}
	if wf.StepName != "" {
		values.Set("step_name", wf.StepName)
	}
	if wf.StepImageURL != "" {
		values.Set("step_image_url", wf.StepImageURL)
	}
	if wf.Inputs != nil {
		data, err := json.Marshal(wf.Inputs)
		if err != nil {
			return err
		}
		values.Set("inputs", (string)(data))
	}
	if wf.Outputs != nil {
		data, err := json.Marshal(wf.Outputs)
		if err != nil {
			return err
		}
		values.Set("outputs", (string)(data))
	}

	response := &SlackResponse{}
	if err := api.postMethod(ctx, "workflows.updateStep", values, response); err != nil {
		return err
	}

	return response.Err()
}
