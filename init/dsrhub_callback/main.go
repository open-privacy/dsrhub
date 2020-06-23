package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/loopfz/gadgeto/zesty"
	"github.com/ovh/utask"
	"github.com/ovh/utask/models/resolution"
	"github.com/ovh/utask/pkg/plugins"
	"github.com/sirupsen/logrus"
	"github.com/wI2L/fizz"
)

var Plugin = NewDSRHubCallbackPlugin() // nolint

type DSRHubCallbackPlugin struct{}

func NewDSRHubCallbackPlugin() plugins.InitializerPlugin { return &DSRHubCallbackPlugin{} }
func (p *DSRHubCallbackPlugin) Description() string      { return "DSRHub Callback Plugin" }

func (p *DSRHubCallbackPlugin) Init(service *plugins.Service) error {
	router, ok := service.Server.Handler(context.Background()).(*fizz.Fizz)
	if !ok {
		return fmt.Errorf("failed to load router in plugin: %s", p.Description())
	}
	router.POST("/dsrhub/callback/:resolution_id/:step_name",
		[]fizz.OperationOption{fizz.Summary("Handle dsrhub webhook callback.")},
		tonic.Handler(p.handleCallbackFunc(), 200))
	return nil
}

type inCallback struct {
	// dsrhub related fields
	ResolutionID string `path:"resolution_id" validate:"required"`
	StepName     string `path:"step_name" validate:"required"`

	// opendsr related fields
	APIVersion         string `json:"api_version"`
	Regulation         string `json:"regulation" validate:"required"`
	ControllerID       string `json:"controller_id" validate:"required"`
	RequestStatus      string `json:"request_status" validate:"required"`
	SubjectRequestID   string `json:"subject_request_id" validate:"required"`
	SubjectRequestType string `json:"subject_request_type" validate:"required"`
	IdentityType       string `json:"identity_type"`
	IdentityFormat     string `json:"identity_format"`
	IdentityValue      string `json:"identity_value"`
}

func (p *DSRHubCallbackPlugin) handleCallbackFunc() func(c *gin.Context, in *inCallback) error {
	return func(c *gin.Context, in *inCallback) error {
		dbp, err := zesty.NewDBProvider(utask.DBName)
		if err != nil {
			return err
		}

		if err := dbp.Tx(); err != nil {
			return err
		}

		r, err := resolution.LoadLockedNoWaitFromPublicID(dbp, in.ResolutionID)
		if err != nil {
			dbp.Rollback()
			return err
		}

		if _, ok := r.Steps[in.StepName]; !ok {
			dbp.Rollback()
			return fmt.Errorf("step %s not found", in.StepName)
		}
		r.SetInput(map[string]interface{}{
			in.StepName: in,
		})

		logrus.WithFields(logrus.Fields{
			"resolution_id": in.ResolutionID,
			"controller_id": in.ControllerID,
		}).Debug("update resolution resolver_input from DSRHubCallbackPlugin")

		if err := r.Update(dbp); err != nil {
			dbp.Rollback()
			return err
		}

		if err := dbp.Commit(); err != nil {
			dbp.Rollback()
			return err
		}
		return nil
	}
}
