//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvStoreInstance) DeepCopyInto(out *KvStoreInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvStoreInstance.
func (in *KvStoreInstance) DeepCopy() *KvStoreInstance {
	if in == nil {
		return nil
	}
	out := new(KvStoreInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KvStoreInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvStoreInstanceList) DeepCopyInto(out *KvStoreInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KvStoreInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvStoreInstanceList.
func (in *KvStoreInstanceList) DeepCopy() *KvStoreInstanceList {
	if in == nil {
		return nil
	}
	out := new(KvStoreInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KvStoreInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvStoreInstanceObservation) DeepCopyInto(out *KvStoreInstanceObservation) {
	*out = *in
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionDomain != nil {
		in, out := &in.ConnectionDomain, &out.ConnectionDomain
		*out = new(string)
		**out = **in
	}
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.QPS != nil {
		in, out := &in.QPS, &out.QPS
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvStoreInstanceObservation.
func (in *KvStoreInstanceObservation) DeepCopy() *KvStoreInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(KvStoreInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvStoreInstanceParameters) DeepCopyInto(out *KvStoreInstanceParameters) {
	*out = *in
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.AutoRenewPeriod != nil {
		in, out := &in.AutoRenewPeriod, &out.AutoRenewPeriod
		*out = new(float64)
		**out = **in
	}
	if in.AutoUseCoupon != nil {
		in, out := &in.AutoUseCoupon, &out.AutoUseCoupon
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(string)
		**out = **in
	}
	if in.BackupPeriod != nil {
		in, out := &in.BackupPeriod, &out.BackupPeriod
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.BackupTime != nil {
		in, out := &in.BackupTime, &out.BackupTime
		*out = new(string)
		**out = **in
	}
	if in.BusinessInfo != nil {
		in, out := &in.BusinessInfo, &out.BusinessInfo
		*out = new(string)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ConnectionStringPrefix != nil {
		in, out := &in.ConnectionStringPrefix, &out.ConnectionStringPrefix
		*out = new(string)
		**out = **in
	}
	if in.CouponNo != nil {
		in, out := &in.CouponNo, &out.CouponNo
		*out = new(string)
		**out = **in
	}
	if in.DBInstanceName != nil {
		in, out := &in.DBInstanceName, &out.DBInstanceName
		*out = new(string)
		**out = **in
	}
	if in.DedicatedHostGroupID != nil {
		in, out := &in.DedicatedHostGroupID, &out.DedicatedHostGroupID
		*out = new(string)
		**out = **in
	}
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.EnableBackupLog != nil {
		in, out := &in.EnableBackupLog, &out.EnableBackupLog
		*out = new(float64)
		**out = **in
	}
	if in.EnablePublic != nil {
		in, out := &in.EnablePublic, &out.EnablePublic
		*out = new(bool)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.ForceUpgrade != nil {
		in, out := &in.ForceUpgrade, &out.ForceUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.GlobalInstance != nil {
		in, out := &in.GlobalInstance, &out.GlobalInstance
		*out = new(bool)
		**out = **in
	}
	if in.GlobalInstanceID != nil {
		in, out := &in.GlobalInstanceID, &out.GlobalInstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceChargeType != nil {
		in, out := &in.InstanceChargeType, &out.InstanceChargeType
		*out = new(string)
		**out = **in
	}
	if in.InstanceClass != nil {
		in, out := &in.InstanceClass, &out.InstanceClass
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.InstanceReleaseProtection != nil {
		in, out := &in.InstanceReleaseProtection, &out.InstanceReleaseProtection
		*out = new(bool)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.KMSEncryptedPassword != nil {
		in, out := &in.KMSEncryptedPassword, &out.KMSEncryptedPassword
		*out = new(string)
		**out = **in
	}
	if in.KMSEncryptionContext != nil {
		in, out := &in.KMSEncryptionContext, &out.KMSEncryptionContext
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MaintainEndTime != nil {
		in, out := &in.MaintainEndTime, &out.MaintainEndTime
		*out = new(string)
		**out = **in
	}
	if in.MaintainStartTime != nil {
		in, out := &in.MaintainStartTime, &out.MaintainStartTime
		*out = new(string)
		**out = **in
	}
	if in.ModifyMode != nil {
		in, out := &in.ModifyMode, &out.ModifyMode
		*out = new(float64)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.OrderType != nil {
		in, out := &in.OrderType, &out.OrderType
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PaymentType != nil {
		in, out := &in.PaymentType, &out.PaymentType
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.PrivateConnectionPort != nil {
		in, out := &in.PrivateConnectionPort, &out.PrivateConnectionPort
		*out = new(string)
		**out = **in
	}
	if in.PrivateConnectionPrefix != nil {
		in, out := &in.PrivateConnectionPrefix, &out.PrivateConnectionPrefix
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.RestoreTime != nil {
		in, out := &in.RestoreTime, &out.RestoreTime
		*out = new(string)
		**out = **in
	}
	if in.SSLEnable != nil {
		in, out := &in.SSLEnable, &out.SSLEnable
		*out = new(string)
		**out = **in
	}
	if in.SecondaryZoneID != nil {
		in, out := &in.SecondaryZoneID, &out.SecondaryZoneID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityIPGroupAttribute != nil {
		in, out := &in.SecurityIPGroupAttribute, &out.SecurityIPGroupAttribute
		*out = new(string)
		**out = **in
	}
	if in.SecurityIPGroupName != nil {
		in, out := &in.SecurityIPGroupName, &out.SecurityIPGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecurityIps != nil {
		in, out := &in.SecurityIps, &out.SecurityIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SrcdbInstanceID != nil {
		in, out := &in.SrcdbInstanceID, &out.SrcdbInstanceID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCAuthMode != nil {
		in, out := &in.VPCAuthMode, &out.VPCAuthMode
		*out = new(string)
		**out = **in
	}
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvStoreInstanceParameters.
func (in *KvStoreInstanceParameters) DeepCopy() *KvStoreInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(KvStoreInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvStoreInstanceSpec) DeepCopyInto(out *KvStoreInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvStoreInstanceSpec.
func (in *KvStoreInstanceSpec) DeepCopy() *KvStoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(KvStoreInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KvStoreInstanceStatus) DeepCopyInto(out *KvStoreInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KvStoreInstanceStatus.
func (in *KvStoreInstanceStatus) DeepCopy() *KvStoreInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(KvStoreInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersObservation) DeepCopyInto(out *ParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersObservation.
func (in *ParametersObservation) DeepCopy() *ParametersObservation {
	if in == nil {
		return nil
	}
	out := new(ParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersParameters) DeepCopyInto(out *ParametersParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersParameters.
func (in *ParametersParameters) DeepCopy() *ParametersParameters {
	if in == nil {
		return nil
	}
	out := new(ParametersParameters)
	in.DeepCopyInto(out)
	return out
}