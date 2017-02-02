// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package batch

import (
	v1 "github.com/lavalamp/client-go-flat/apimachinery/pkg/apis/meta/v1"
	conversion "github.com/lavalamp/client-go-flat/apimachinery/pkg/conversion"
	runtime "github.com/lavalamp/client-go-flat/apimachinery/pkg/runtime"
	api "github.com/lavalamp/client-go-flat/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJob, InType: reflect.TypeOf(&CronJob{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJobList, InType: reflect.TypeOf(&CronJobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJobSpec, InType: reflect.TypeOf(&CronJobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJobStatus, InType: reflect.TypeOf(&CronJobStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_Job, InType: reflect.TypeOf(&Job{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobCondition, InType: reflect.TypeOf(&JobCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobList, InType: reflect.TypeOf(&JobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobSpec, InType: reflect.TypeOf(&JobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobStatus, InType: reflect.TypeOf(&JobStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobTemplate, InType: reflect.TypeOf(&JobTemplate{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobTemplateSpec, InType: reflect.TypeOf(&JobTemplateSpec{})},
	)
}

func DeepCopy_batch_CronJob(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJob)
		out := out.(*CronJob)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_batch_CronJobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_CronJobStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_CronJobList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobList)
		out := out.(*CronJobList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CronJob, len(*in))
			for i := range *in {
				if err := DeepCopy_batch_CronJob(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_batch_CronJobSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobSpec)
		out := out.(*CronJobSpec)
		*out = *in
		if in.StartingDeadlineSeconds != nil {
			in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
			*out = new(int64)
			**out = **in
		}
		if in.Suspend != nil {
			in, out := &in.Suspend, &out.Suspend
			*out = new(bool)
			**out = **in
		}
		if err := DeepCopy_batch_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_CronJobStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobStatus)
		out := out.(*CronJobStatus)
		*out = *in
		if in.Active != nil {
			in, out := &in.Active, &out.Active
			*out = make([]api.ObjectReference, len(*in))
			copy(*out, *in)
		}
		if in.LastScheduleTime != nil {
			in, out := &in.LastScheduleTime, &out.LastScheduleTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}

func DeepCopy_batch_Job(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Job)
		out := out.(*Job)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_batch_JobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_JobStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_JobCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobCondition)
		out := out.(*JobCondition)
		*out = *in
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

func DeepCopy_batch_JobList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobList)
		out := out.(*JobList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Job, len(*in))
			for i := range *in {
				if err := DeepCopy_batch_Job(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_batch_JobSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobSpec)
		out := out.(*JobSpec)
		*out = *in
		if in.Parallelism != nil {
			in, out := &in.Parallelism, &out.Parallelism
			*out = new(int32)
			**out = **in
		}
		if in.Completions != nil {
			in, out := &in.Completions, &out.Completions
			*out = new(int32)
			**out = **in
		}
		if in.ActiveDeadlineSeconds != nil {
			in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
			*out = new(int64)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if in.ManualSelector != nil {
			in, out := &in.ManualSelector, &out.ManualSelector
			*out = new(bool)
			**out = **in
		}
		if err := api.DeepCopy_api_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_JobStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobStatus)
		out := out.(*JobStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]JobCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_batch_JobCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.StartTime != nil {
			in, out := &in.StartTime, &out.StartTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.CompletionTime != nil {
			in, out := &in.CompletionTime, &out.CompletionTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}

func DeepCopy_batch_JobTemplate(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobTemplate)
		out := out.(*JobTemplate)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_batch_JobTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_JobTemplateSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobTemplateSpec)
		out := out.(*JobTemplateSpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_batch_JobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}
