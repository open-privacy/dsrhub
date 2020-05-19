package model

import "github.com/qor/transition"

const (
	WORKFLOW_STATE_CREATED   = "workflow_created"
	WORKFLOW_STATE_PENDING   = "workflow_pending"
	WORKFLOW_STATE_COMPLETED = "workflow_completed"
	WORKFLOW_STATE_CANCELLED = "workflow_cancelled"
)

type Workflow struct {
	BaseModel
	transition.Transition

	Name        string
	Description string
	Input       EncryptionKV `gorm:"type:varchar(65532)"`
	Output      EncryptionKV `gorm:"type:varchar(1000000)"`
	ScopeID     string
	Tasks       []*Task
	Timeout     int // timeout in seconds

	FSM *transition.StateMachine `gorm:"-"`
}

func (w *Workflow) Prepare() {
	// prepare fsm
	fsm := transition.New(&Workflow{})
	fsm.State(WORKFLOW_STATE_CREATED)
	fsm.State(WORKFLOW_STATE_PENDING)
	fsm.State(WORKFLOW_STATE_COMPLETED)
	fsm.State(WORKFLOW_STATE_CANCELLED)
	fsm.Initial(WORKFLOW_STATE_CREATED)
	w.FSM = fsm

	// TODO: prepare tasks and their dependencies
}

func (w *Workflow) Run() error {
	return nil
}
