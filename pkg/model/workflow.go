package model

import "github.com/qor/transition"

const (
	WORKFLOW_STATE_CREATED   = "@created"
	WORKFLOW_STATE_COMPLETED = "@completed"
	WORKFLOW_STATE_FAILED    = "@failed"
	WORKFLOW_STATE_CANCELLED = "@cancelled"
)

type Workflow struct {
	BaseModel
	transition.Transition

	Name        string `gorm:"index"`
	Description string
	Input       KV  `gorm:"type:varchar(65532)"`
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
	fsm.State(WORKFLOW_STATE_CREATED)
	fsm.State(WORKFLOW_STATE_COMPLETED)
	fsm.State(WORKFLOW_STATE_FAILED)
	fsm.State(WORKFLOW_STATE_CANCELLED)
	fsm.Initial(WORKFLOW_STATE_CREATED)
	w.FSM = fsm

	// TODO: prepare tasks and their dependencies
}

func (w *Workflow) Run() error {
	return nil
}
