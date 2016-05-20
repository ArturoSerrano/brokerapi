package brokerapi

import "errors"

type ServiceBroker interface {
	Services() []Service

	Provision(instanceID string, details map[string]interface{}, asyncAllowed bool) (ProvisionedServiceSpec, error)
	Deprovision(instanceID string, details DeprovisionDetails, asyncAllowed bool) (IsAsync, error)

	Bind(instanceID, bindingID string, details BindDetails) (Binding, error)
	Unbind(instanceID, bindingID string, details UnbindDetails) error

	Update(instanceID string, details map[string]interface{}, asyncAllowed bool) (IsAsync, error)

	LastOperation(instanceID string) (LastOperation, error)
}

type IsAsync bool

type ProvisionedServiceSpec struct {
	IsAsync      bool
	DashboardURL string
}

type BindDetails struct {
	AppGUID      string                 `json:"app_guid"`
	PlanID       string                 `json:"plan_id"`
	ServiceID    string                 `json:"service_id"`
	BindResource *BindResource          `json:"bind_resource,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
}

type BindResource struct {
	AppGuid string `json:"app_guid,omitempty"`
	Route   string `json:"route,omitempty"`
}

type UnbindDetails struct {
	PlanID    string `json:"plan_id"`
	ServiceID string `json:"service_id"`
}

type DeprovisionDetails struct {
	PlanID    string `json:"plan_id"`
	ServiceID string `json:"service_id"`
}

type PreviousValues struct {
	PlanID    string `json:"plan_id"`
	ServiceID string `json:"service_id"`
	OrgID     string `json:"organization_id"`
	SpaceID   string `json:"space_id"`
}

type LastOperation struct {
	State       LastOperationState
	Description string
}

type LastOperationState string

const (
	InProgress LastOperationState = "in progress"
	Succeeded  LastOperationState = "succeeded"
	Failed     LastOperationState = "failed"
)

type Binding struct {
	Credentials     interface{} `json:"credentials"`
	SyslogDrainURL  string      `json:"syslog_drain_url,omitempty"`
	RouteServiceURL string      `json:"route_service_url,omitempty"`
}

var (
	ErrInstanceAlreadyExists  = errors.New("instance already exists")
	ErrInstanceDoesNotExist   = errors.New("instance does not exist")
	ErrInstanceLimitMet       = errors.New("instance limit for this service has been reached")
	ErrPlanQuotaExceeded      = errors.New("The quota for this service plan has been exceeded. Please contact your Operator for help.")
	ErrBindingAlreadyExists   = errors.New("binding already exists")
	ErrBindingDoesNotExist    = errors.New("binding does not exist")
	ErrAsyncRequired          = errors.New("This service plan requires client support for asynchronous service operations.")
	ErrPlanChangeNotSupported = errors.New("The requested plan migration cannot be performed")
	ErrRawParamsInvalid       = errors.New("The format of the parameters is not valid JSON")
)
