package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"bosh-softlayer-cpi/softlayer/pool/models"
)

// NewUpdateVMWithStateParams creates a new UpdateVMWithStateParams object
// with the default values initialized.
func NewUpdateVMWithStateParams() *UpdateVMWithStateParams {
	var ()
	return &UpdateVMWithStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMWithStateParamsWithTimeout creates a new UpdateVMWithStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVMWithStateParamsWithTimeout(timeout time.Duration) *UpdateVMWithStateParams {
	var ()
	return &UpdateVMWithStateParams{

		timeout: timeout,
	}
}

// NewUpdateVMWithStateParamsWithContext creates a new UpdateVMWithStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVMWithStateParamsWithContext(ctx context.Context) *UpdateVMWithStateParams {
	var ()
	return &UpdateVMWithStateParams{

		Context: ctx,
	}
}

/*UpdateVMWithStateParams contains all the parameters to send to the API endpoint
for the update Vm with state operation typically these are written to a http.Request
*/
type UpdateVMWithStateParams struct {

	/*Body
	  Vm state that needs to be updated

	*/
	Body *models.VMState
	/*Cid
	  ID of vm that needs to be updated

	*/
	Cid int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the update Vm with state params
func (o *UpdateVMWithStateParams) WithTimeout(timeout time.Duration) *UpdateVMWithStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm with state params
func (o *UpdateVMWithStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm with state params
func (o *UpdateVMWithStateParams) WithContext(ctx context.Context) *UpdateVMWithStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm with state params
func (o *UpdateVMWithStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the update Vm with state params
func (o *UpdateVMWithStateParams) WithBody(body *models.VMState) *UpdateVMWithStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update Vm with state params
func (o *UpdateVMWithStateParams) SetBody(body *models.VMState) {
	o.Body = body
}

// WithCid adds the cid to the update Vm with state params
func (o *UpdateVMWithStateParams) WithCid(cid int32) *UpdateVMWithStateParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the update Vm with state params
func (o *UpdateVMWithStateParams) SetCid(cid int32) {
	o.Cid = cid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMWithStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.VMState)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cid
	if err := r.SetPathParam("cid", swag.FormatInt32(o.Cid)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}