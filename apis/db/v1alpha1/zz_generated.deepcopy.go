//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BabelfishConfigObservation) DeepCopyInto(out *BabelfishConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BabelfishConfigObservation.
func (in *BabelfishConfigObservation) DeepCopy() *BabelfishConfigObservation {
	if in == nil {
		return nil
	}
	out := new(BabelfishConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BabelfishConfigParameters) DeepCopyInto(out *BabelfishConfigParameters) {
	*out = *in
	if in.BabelfishEnabled != nil {
		in, out := &in.BabelfishEnabled, &out.BabelfishEnabled
		*out = new(string)
		**out = **in
	}
	if in.MasterUserPassword != nil {
		in, out := &in.MasterUserPassword, &out.MasterUserPassword
		*out = new(string)
		**out = **in
	}
	if in.MasterUsername != nil {
		in, out := &in.MasterUsername, &out.MasterUsername
		*out = new(string)
		**out = **in
	}
	if in.MigrationMode != nil {
		in, out := &in.MigrationMode, &out.MigrationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BabelfishConfigParameters.
func (in *BabelfishConfigParameters) DeepCopy() *BabelfishConfigParameters {
	if in == nil {
		return nil
	}
	out := new(BabelfishConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbInstance) DeepCopyInto(out *DbInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbInstance.
func (in *DbInstance) DeepCopy() *DbInstance {
	if in == nil {
		return nil
	}
	out := new(DbInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbInstanceList) DeepCopyInto(out *DbInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DbInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbInstanceList.
func (in *DbInstanceList) DeepCopy() *DbInstanceList {
	if in == nil {
		return nil
	}
	out := new(DbInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbInstanceObservation) DeepCopyInto(out *DbInstanceObservation) {
	*out = *in
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SSLStatus != nil {
		in, out := &in.SSLStatus, &out.SSLStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbInstanceObservation.
func (in *DbInstanceObservation) DeepCopy() *DbInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(DbInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbInstanceParameters) DeepCopyInto(out *DbInstanceParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(string)
		**out = **in
	}
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
	if in.AutoUpgradeMinorVersion != nil {
		in, out := &in.AutoUpgradeMinorVersion, &out.AutoUpgradeMinorVersion
		*out = new(string)
		**out = **in
	}
	if in.BabelfishConfig != nil {
		in, out := &in.BabelfishConfig, &out.BabelfishConfig
		*out = make([]BabelfishConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BabelfishPort != nil {
		in, out := &in.BabelfishPort, &out.BabelfishPort
		*out = new(string)
		**out = **in
	}
	if in.CAType != nil {
		in, out := &in.CAType, &out.CAType
		*out = new(string)
		**out = **in
	}
	if in.ClientCACert != nil {
		in, out := &in.ClientCACert, &out.ClientCACert
		*out = new(string)
		**out = **in
	}
	if in.ClientCAEnabled != nil {
		in, out := &in.ClientCAEnabled, &out.ClientCAEnabled
		*out = new(float64)
		**out = **in
	}
	if in.ClientCertRevocationList != nil {
		in, out := &in.ClientCertRevocationList, &out.ClientCertRevocationList
		*out = new(string)
		**out = **in
	}
	if in.ClientCrlEnabled != nil {
		in, out := &in.ClientCrlEnabled, &out.ClientCrlEnabled
		*out = new(float64)
		**out = **in
	}
	if in.ConnectionStringPrefix != nil {
		in, out := &in.ConnectionStringPrefix, &out.ConnectionStringPrefix
		*out = new(string)
		**out = **in
	}
	if in.DBInstanceIPArrayAttribute != nil {
		in, out := &in.DBInstanceIPArrayAttribute, &out.DBInstanceIPArrayAttribute
		*out = new(string)
		**out = **in
	}
	if in.DBInstanceIPArrayName != nil {
		in, out := &in.DBInstanceIPArrayName, &out.DBInstanceIPArrayName
		*out = new(string)
		**out = **in
	}
	if in.DBInstanceStorageType != nil {
		in, out := &in.DBInstanceStorageType, &out.DBInstanceStorageType
		*out = new(string)
		**out = **in
	}
	if in.DBIsIgnoreCase != nil {
		in, out := &in.DBIsIgnoreCase, &out.DBIsIgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.DBTimeZone != nil {
		in, out := &in.DBTimeZone, &out.DBTimeZone
		*out = new(string)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = new(string)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.ForceRestart != nil {
		in, out := &in.ForceRestart, &out.ForceRestart
		*out = new(bool)
		**out = **in
	}
	if in.FreshWhiteListReadins != nil {
		in, out := &in.FreshWhiteListReadins, &out.FreshWhiteListReadins
		*out = new(string)
		**out = **in
	}
	if in.HaConfig != nil {
		in, out := &in.HaConfig, &out.HaConfig
		*out = new(string)
		**out = **in
	}
	if in.InstanceChargeType != nil {
		in, out := &in.InstanceChargeType, &out.InstanceChargeType
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.InstanceStorage != nil {
		in, out := &in.InstanceStorage, &out.InstanceStorage
		*out = new(float64)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.MaintainTime != nil {
		in, out := &in.MaintainTime, &out.MaintainTime
		*out = new(string)
		**out = **in
	}
	if in.ManualHaTime != nil {
		in, out := &in.ManualHaTime, &out.ManualHaTime
		*out = new(string)
		**out = **in
	}
	if in.ModifyMode != nil {
		in, out := &in.ModifyMode, &out.ModifyMode
		*out = new(string)
		**out = **in
	}
	if in.MonitoringPeriod != nil {
		in, out := &in.MonitoringPeriod, &out.MonitoringPeriod
		*out = new(float64)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.PgHbaConf != nil {
		in, out := &in.PgHbaConf, &out.PgHbaConf
		*out = make([]PgHbaConfParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
	if in.ReleasedKeepPolicy != nil {
		in, out := &in.ReleasedKeepPolicy, &out.ReleasedKeepPolicy
		*out = new(string)
		**out = **in
	}
	if in.ReplicationACL != nil {
		in, out := &in.ReplicationACL, &out.ReplicationACL
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.SQLCollectorConfigValue != nil {
		in, out := &in.SQLCollectorConfigValue, &out.SQLCollectorConfigValue
		*out = new(float64)
		**out = **in
	}
	if in.SQLCollectorStatus != nil {
		in, out := &in.SQLCollectorStatus, &out.SQLCollectorStatus
		*out = new(string)
		**out = **in
	}
	if in.SSLAction != nil {
		in, out := &in.SSLAction, &out.SSLAction
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityIPMode != nil {
		in, out := &in.SecurityIPMode, &out.SecurityIPMode
		*out = new(string)
		**out = **in
	}
	if in.SecurityIPType != nil {
		in, out := &in.SecurityIPType, &out.SecurityIPType
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
	if in.ServerCert != nil {
		in, out := &in.ServerCert, &out.ServerCert
		*out = new(string)
		**out = **in
	}
	if in.ServerKey != nil {
		in, out := &in.ServerKey, &out.ServerKey
		*out = new(string)
		**out = **in
	}
	if in.StorageAutoScale != nil {
		in, out := &in.StorageAutoScale, &out.StorageAutoScale
		*out = new(string)
		**out = **in
	}
	if in.StorageThreshold != nil {
		in, out := &in.StorageThreshold, &out.StorageThreshold
		*out = new(float64)
		**out = **in
	}
	if in.StorageUpperBound != nil {
		in, out := &in.StorageUpperBound, &out.StorageUpperBound
		*out = new(float64)
		**out = **in
	}
	if in.SwitchTime != nil {
		in, out := &in.SwitchTime, &out.SwitchTime
		*out = new(string)
		**out = **in
	}
	if in.TCPConnectionType != nil {
		in, out := &in.TCPConnectionType, &out.TCPConnectionType
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
	if in.TargetMinorVersion != nil {
		in, out := &in.TargetMinorVersion, &out.TargetMinorVersion
		*out = new(string)
		**out = **in
	}
	if in.TdeStatus != nil {
		in, out := &in.TdeStatus, &out.TdeStatus
		*out = new(string)
		**out = **in
	}
	if in.UpgradeDBInstanceKernelVersion != nil {
		in, out := &in.UpgradeDBInstanceKernelVersion, &out.UpgradeDBInstanceKernelVersion
		*out = new(bool)
		**out = **in
	}
	if in.UpgradeTime != nil {
		in, out := &in.UpgradeTime, &out.UpgradeTime
		*out = new(string)
		**out = **in
	}
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
	if in.WhitelistNetworkType != nil {
		in, out := &in.WhitelistNetworkType, &out.WhitelistNetworkType
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDSlaveA != nil {
		in, out := &in.ZoneIDSlaveA, &out.ZoneIDSlaveA
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDSlaveB != nil {
		in, out := &in.ZoneIDSlaveB, &out.ZoneIDSlaveB
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbInstanceParameters.
func (in *DbInstanceParameters) DeepCopy() *DbInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(DbInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbInstanceSpec) DeepCopyInto(out *DbInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbInstanceSpec.
func (in *DbInstanceSpec) DeepCopy() *DbInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(DbInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbInstanceStatus) DeepCopyInto(out *DbInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbInstanceStatus.
func (in *DbInstanceStatus) DeepCopy() *DbInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(DbInstanceStatus)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgHbaConfObservation) DeepCopyInto(out *PgHbaConfObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgHbaConfObservation.
func (in *PgHbaConfObservation) DeepCopy() *PgHbaConfObservation {
	if in == nil {
		return nil
	}
	out := new(PgHbaConfObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PgHbaConfParameters) DeepCopyInto(out *PgHbaConfParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Mask != nil {
		in, out := &in.Mask, &out.Mask
		*out = new(string)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Option != nil {
		in, out := &in.Option, &out.Option
		*out = new(string)
		**out = **in
	}
	if in.PriorityID != nil {
		in, out := &in.PriorityID, &out.PriorityID
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PgHbaConfParameters.
func (in *PgHbaConfParameters) DeepCopy() *PgHbaConfParameters {
	if in == nil {
		return nil
	}
	out := new(PgHbaConfParameters)
	in.DeepCopyInto(out)
	return out
}
