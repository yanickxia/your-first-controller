package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type PodValidator struct {
	Client  client.Client
	Decoder *admission.Decoder
}

func (v *PodValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	return v.validate(ctx, obj)
}

func (v *PodValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (warnings admission.Warnings, err error) {
	return v.validate(ctx, newObj)
}

func (v *PodValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	return v.validate(ctx, obj)
}

func (v *PodValidator) validate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = logf.FromContext(ctx)

	return nil, nil
}
