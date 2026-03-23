package resources

import (
	"context"

	"github.com/rewritetoday/golang/api"
)

// OTPManager provides OTP resource operations.
type OTPManager struct {
	Base
}

// OTP is kept as a compatibility alias for OTPManager.
type OTP = OTPManager

// SendOTPMessageOptions carries OTP send input and optional idempotency metadata.
type SendOTPMessageOptions struct {
	IdempotencyKey string `json:"-"`
	api.RESTPostSendOTPMessageBody
}

// CreateOTPOptions is kept as a compatibility alias for SendOTPMessageOptions.
type CreateOTPOptions = SendOTPMessageOptions

// VerifyOTPOptions carries OTP verification input.
type VerifyOTPOptions struct {
	ID api.Snowflake `json:"-"`
	api.RESTPostVerifyOTPCodeBody
}

// Send creates a new OTP challenge.
func (r *OTPManager) Send(ctx context.Context, options SendOTPMessageOptions) (api.RESTPostSendOTPMessageData, error) {
	var out api.RESTPostSendOTPMessageData
	err := r.Rest.Post(
		ctx,
		api.Routes.OTP.Send(),
		options.RESTPostSendOTPMessageBody,
		&out,
		withIdempotencyKey(options.IdempotencyKey),
	)
	return out, err
}

// Create is a compatibility alias for Send.
func (r *OTPManager) Create(ctx context.Context, options CreateOTPOptions) (api.RESTPostSendOTPMessageData, error) {
	return r.Send(ctx, options)
}

// Verify verifies an OTP challenge.
func (r *OTPManager) Verify(ctx context.Context, options VerifyOTPOptions) (api.RESTPostVerifyOTPCodeData, error) {
	var out api.RESTPostVerifyOTPCodeData
	err := r.Rest.Post(ctx, api.Routes.OTP.Verify(string(options.ID)), options.RESTPostVerifyOTPCodeBody, &out, nil)
	return out, err
}
