package model

type DSL struct {
	Version     string
	WorkflowMap map[string]*Workflow
}
