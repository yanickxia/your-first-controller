package controller

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	quotav1 "tutorial.controller.io/quota-limit/api/v1"
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
	log := logf.FromContext(ctx)
	pod, ok := obj.(*v1.Pod)
	if !ok {
		return nil, fmt.Errorf("expected a Pod but got a %T", obj)
	}

	log.Info("Validating Pod")
	key := "user"

	// 获得用户的 "user" 信息
	userName, found := pod.Labels[key]
	if !found {
		return nil, fmt.Errorf("missing label %s", key)
	}

	// 获得用户配置的 Quota
	quota := &quotav1.Quota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      userName,
			Namespace: "default", // 我们假设用户的 Quota 都放在 default，和用户名同名
		},
		Spec: quotav1.QuotaSpec{
			Limits: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse("2"),
				v1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
	}
	err := v.Client.Get(ctx, client.ObjectKeyFromObject(quota), quota)
	// 没有配置 Quota，就是采用默认值
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, fmt.Errorf("get quota failed")
	}

	cpuResource := resource.MustParse("0")
	memResource := resource.MustParse("0Mi")

	// 获得用户的所有的 Pod
	podList := &v1.PodList{}
	if err := v.Client.List(ctx, podList, client.MatchingLabels{
		"user": userName,
	}); err != nil {
		return nil, err
	}

	// 累计当前所有的
	for i := range podList.Items {
		pod := podList.Items[i]
		for j := range pod.Spec.Containers {
			container := pod.Spec.Containers[j]
			cpu := container.Resources.Limits.Cpu().DeepCopy()
			cpuResource.Add(cpu)
			mem := container.Resources.Limits.Memory().DeepCopy()
			memResource.Add(mem)
		}
	}

	if cpuResource.Cmp(quota.Spec.Limits.Cpu().DeepCopy()) < 0 ||
		memResource.Cmp(quota.Spec.Limits.Memory().DeepCopy()) < 0 {
		return nil, fmt.Errorf("user %s, limit out of quota", userName)
	}

	return nil, nil
}
