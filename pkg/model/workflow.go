package model

import "github.com/qor/transition"

const (
	WorkflowStateCreated   = "@created"
	WorkflowStateSuccess   = "@success"
	WorkflowStateFailed    = "@failed"
	WorkflowStateTimeout   = "@timeout"
	WorkflowStateCancelled = "@cancelled"
)

type Workflow struct {
	BaseModel
	transition.Transition

	Name        string `gorm:"index"`
	Description string
	Timeout     int // timeout in seconds

	// -------- yaml ignored --------------------------------------------------
	TasksList []*Task      `yaml:"-"`
	ScopeID   string       `gorm:"index" yaml:"-"`
	Output    EncryptionKV `gorm:"type:varchar(1000000)" yaml:"-"`
	// ------------------------------------------------------------------------

	// -------- gorm ignored --------------------------------------------------
	Tasks map[string]*Task `gorm:"-"`
	// ------------------------------------------------------------------------

	// ---------gorm and yaml ignored -----------------------------------------
	FSM *transition.StateMachine `gorm:"-" yaml:"-"`
	// ------------------------------------------------------------------------
}

func (w *Workflow) Prepare() {
	// prepare fsm
	fsm := transition.New(&Workflow{})
	fsm.State(WorkflowStateCreated)
	fsm.State(WorkflowStateSuccess)
	fsm.State(WorkflowStateFailed)
	fsm.State(WorkflowStateTimeout)
	fsm.State(WorkflowStateCancelled)
	fsm.Initial(WorkflowStateCreated)
	w.FSM = fsm

	// TODO: prepare tasks and their dependencies
}

func (w *Workflow) Run() error {
	return nil
}
