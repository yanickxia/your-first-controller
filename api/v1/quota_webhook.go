/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var quotalog = logf.Log.WithName("quota-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Quota) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-resource-tutorial-controller-io-v1-quota,mutating=true,failurePolicy=fail,sideEffects=None,groups=resource.tutorial.controller.io,resources=quotas,verbs=create;update,versions=v1,name=mquota.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Quota{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Quota) Default() {
	quotalog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-resource-tutorial-controller-io-v1-quota,mutating=false,failurePolicy=fail,sideEffects=None,groups=resource.tutorial.controller.io,resources=quotas,verbs=create;update,versions=v1,name=vquota.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Quota{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Quota) ValidateCreate() (admission.Warnings, error) {
	quotalog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Quota) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	quotalog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Quota) ValidateDelete() (admission.Warnings, error) {
	quotalog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
