package ecs

import (
	"github.com/denverdino/aliyungo/ecs"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/ecs",

	"AcceptingSide":               ecs.AcceptingSide,
	"Active":                      ecs.Active,
	"Creating":                    ecs.Creating,
	"DefaultTimeout":              ecs.DefaultTimeout,
	"DefaultWaitForInterval":      ecs.DefaultWaitForInterval,
	"Deleted":                     ecs.Deleted,
	"DiskCategoryAll":             ecs.DiskCategoryAll,
	"DiskCategoryCloud":           ecs.DiskCategoryCloud,
	"DiskCategoryCloudEfficiency": ecs.DiskCategoryCloudEfficiency,
	"DiskCategoryCloudSSD":        ecs.DiskCategoryCloudSSD,
	"DiskCategoryEphemeral":       ecs.DiskCategoryEphemeral,
	"DiskCategoryEphemeralSSD":    ecs.DiskCategoryEphemeralSSD,
	"DiskStatusAll":               ecs.DiskStatusAll,
	"DiskStatusAttaching":         ecs.DiskStatusAttaching,
	"DiskStatusAvailable":         ecs.DiskStatusAvailable,
	"DiskStatusCreating":          ecs.DiskStatusCreating,
	"DiskStatusDetaching":         ecs.DiskStatusDetaching,
	"DiskStatusInUse":             ecs.DiskStatusInUse,
	"DiskStatusReIniting":         ecs.DiskStatusReIniting,
	"DiskTypeAll":                 ecs.DiskTypeAll,
	"DiskTypeAllData":             ecs.DiskTypeAllData,
	"DiskTypeAllSystem":           ecs.DiskTypeAllSystem,
	"ECSAPIVersion":               ecs.ECSAPIVersion,
	"ECSDefaultEndpoint":          ecs.ECSDefaultEndpoint,
	"ECSServiceCode":              ecs.ECSServiceCode,
	"EipStatusAssociating":        ecs.EipStatusAssociating,
	"EipStatusAvailable":          ecs.EipStatusAvailable,
	"EipStatusInUse":              ecs.EipStatusInUse,
	"EipStatusUnassociating":      ecs.EipStatusUnassociating,
	"Idl": ecs.Idl,
	"ImageDefaultTimeout":             ecs.ImageDefaultTimeout,
	"ImageOwnerDefault":               ecs.ImageOwnerDefault,
	"ImageOwnerMarketplace":           ecs.ImageOwnerMarketplace,
	"ImageOwnerOthers":                ecs.ImageOwnerOthers,
	"ImageOwnerSelf":                  ecs.ImageOwnerSelf,
	"ImageOwnerSystem":                ecs.ImageOwnerSystem,
	"ImageStatusAvailable":            ecs.ImageStatusAvailable,
	"ImageStatusCreateFailed":         ecs.ImageStatusCreateFailed,
	"ImageStatusCreating":             ecs.ImageStatusCreating,
	"ImageStatusUnAvailable":          ecs.ImageStatusUnAvailable,
	"ImageUsageInstance":              ecs.ImageUsageInstance,
	"ImageUsageNone":                  ecs.ImageUsageNone,
	"Inactive":                        ecs.Inactive,
	"InitiatingSide":                  ecs.InitiatingSide,
	"InstanceDefaultTimeout":          ecs.InstanceDefaultTimeout,
	"IpProtocolAll":                   ecs.IpProtocolAll,
	"IpProtocolGRE":                   ecs.IpProtocolGRE,
	"IpProtocolICMP":                  ecs.IpProtocolICMP,
	"IpProtocolTCP":                   ecs.IpProtocolTCP,
	"IpProtocolUDP":                   ecs.IpProtocolUDP,
	"Large1":                          ecs.Large1,
	"Large2":                          ecs.Large2,
	"LockReasonFinancial":             ecs.LockReasonFinancial,
	"LockReasonSecurity":              ecs.LockReasonSecurity,
	"Middle1":                         ecs.Middle1,
	"Middle2":                         ecs.Middle2,
	"Middle5":                         ecs.Middle5,
	"NatGatewayLargeSpec":             ecs.NatGatewayLargeSpec,
	"NatGatewayMiddleSpec":            ecs.NatGatewayMiddleSpec,
	"NatGatewaySmallSpec":             ecs.NatGatewaySmallSpec,
	"NextHopIntance":                  ecs.NextHopIntance,
	"NextHopTunnel":                   ecs.NextHopTunnel,
	"NicTypeInternet":                 ecs.NicTypeInternet,
	"NicTypeIntranet":                 ecs.NicTypeIntranet,
	"NoSpot":                          ecs.NoSpot,
	"Pending":                         ecs.Pending,
	"PermissionPolicyAccept":          ecs.PermissionPolicyAccept,
	"PermissionPolicyDrop":            ecs.PermissionPolicyDrop,
	"PostPaid":                        ecs.PostPaid,
	"PrePaid":                         ecs.PrePaid,
	"ResourceTypeDisk":                ecs.ResourceTypeDisk,
	"ResourceTypeIOOptimizedInstance": ecs.ResourceTypeIOOptimizedInstance,
	"ResourceTypeInstance":            ecs.ResourceTypeInstance,
	"ResourceTypeVSwitch":             ecs.ResourceTypeVSwitch,
	"RouteEntryStatusAvailable":       ecs.RouteEntryStatusAvailable,
	"RouteEntryStatusModifying":       ecs.RouteEntryStatusModifying,
	"RouteEntryStatusPending":         ecs.RouteEntryStatusPending,
	"RouteTableCustom":                ecs.RouteTableCustom,
	"RouteTableSystem":                ecs.RouteTableSystem,
	"Running":                         ecs.Running,
	"Small1":                          ecs.Small1,
	"Small2":                          ecs.Small2,
	"Small5":                          ecs.Small5,
	"SnapshotDefaultTimeout":          ecs.SnapshotDefaultTimeout,
	"SpotWithPriceLimit":              ecs.SpotWithPriceLimit,
	"Starting":                        ecs.Starting,
	"Stopped":                         ecs.Stopped,
	"Stopping":                        ecs.Stopping,
	"SupportedDataDiskCategory":       ecs.SupportedDataDiskCategory,
	"SupportedInstanceGeneration":     ecs.SupportedInstanceGeneration,
	"SupportedInstanceType":           ecs.SupportedInstanceType,
	"SupportedInstanceTypeFamily":     ecs.SupportedInstanceTypeFamily,
	"SupportedNetworkCategory":        ecs.SupportedNetworkCategory,
	"SupportedSystemDiskCategory":     ecs.SupportedSystemDiskCategory,
	"TagResourceDisk":                 ecs.TagResourceDisk,
	"TagResourceImage":                ecs.TagResourceImage,
	"TagResourceInstance":             ecs.TagResourceInstance,
	"TagResourceSnapshot":             ecs.TagResourceSnapshot,
	"VBR":                    ecs.VBR,
	"VPCAPIVersion":          ecs.VPCAPIVersion,
	"VPCDefaultEndpoint":     ecs.VPCDefaultEndpoint,
	"VPCServiceCode":         ecs.VPCServiceCode,
	"VRouter":                ecs.VRouter,
	"VSwitchStatusAvailable": ecs.VSwitchStatusAvailable,
	"VSwitchStatusPending":   ecs.VSwitchStatusPending,
	"VpcStatusAvailable":     ecs.VpcStatusAvailable,
	"VpcStatusPending":       ecs.VpcStatusPending,

	"IoOptimizedNone":      ecs.IoOptimizedNone,
	"IoOptimizedOptimized": ecs.IoOptimizedOptimized,

	"AccountType":                            qlang.StructOf((*ecs.AccountType)(nil)),
	"AddTagsArgs":                            qlang.StructOf((*ecs.AddTagsArgs)(nil)),
	"AddTagsResponse":                        qlang.StructOf((*ecs.AddTagsResponse)(nil)),
	"AllocateEipAddressArgs":                 qlang.StructOf((*ecs.AllocateEipAddressArgs)(nil)),
	"AllocateEipAddressResponse":             qlang.StructOf((*ecs.AllocateEipAddressResponse)(nil)),
	"AllocatePublicIpAddressArgs":            qlang.StructOf((*ecs.AllocatePublicIpAddressArgs)(nil)),
	"AllocatePublicIpAddressResponse":        qlang.StructOf((*ecs.AllocatePublicIpAddressResponse)(nil)),
	"AssociateEipAddressArgs":                qlang.StructOf((*ecs.AssociateEipAddressArgs)(nil)),
	"AssociateEipAddressResponse":            qlang.StructOf((*ecs.AssociateEipAddressResponse)(nil)),
	"AttachDiskArgs":                         qlang.StructOf((*ecs.AttachDiskArgs)(nil)),
	"AttachDiskResponse":                     qlang.StructOf((*ecs.AttachDiskResponse)(nil)),
	"AttachInstancesArgs":                    qlang.StructOf((*ecs.AttachInstancesArgs)(nil)),
	"AttachKeyPairArgs":                      qlang.StructOf((*ecs.AttachKeyPairArgs)(nil)),
	"AuthorizeSecurityGroupArgs":             qlang.StructOf((*ecs.AuthorizeSecurityGroupArgs)(nil)),
	"AuthorizeSecurityGroupEgressArgs":       qlang.StructOf((*ecs.AuthorizeSecurityGroupEgressArgs)(nil)),
	"AuthorizeSecurityGroupEgressResponse":   qlang.StructOf((*ecs.AuthorizeSecurityGroupEgressResponse)(nil)),
	"AuthorizeSecurityGroupResponse":         qlang.StructOf((*ecs.AuthorizeSecurityGroupResponse)(nil)),
	"AvailableDiskCategoriesType":            qlang.StructOf((*ecs.AvailableDiskCategoriesType)(nil)),
	"AvailableInstanceTypesType":             qlang.StructOf((*ecs.AvailableInstanceTypesType)(nil)),
	"AvailableResourceCreationType":          qlang.StructOf((*ecs.AvailableResourceCreationType)(nil)),
	"AvailableResourcesType":                 qlang.StructOf((*ecs.AvailableResourcesType)(nil)),
	"BandwidthPackageIdType":                 qlang.StructOf((*ecs.BandwidthPackageIdType)(nil)),
	"BandwidthPackageType":                   qlang.StructOf((*ecs.BandwidthPackageType)(nil)),
	"BusinessInfo":                           qlang.StructOf((*ecs.BusinessInfo)(nil)),
	"Client":                                 qlang.StructOf((*ecs.Client)(nil)),
	"newClient":                              ecs.NewClient,
	"newClientWithEndpoint":                  ecs.NewClientWithEndpoint,
	"newClientWithRegion":                    ecs.NewClientWithRegion,
	"newECSClient":                           ecs.NewECSClient,
	"newVPCClient":                           ecs.NewVPCClient,
	"newVPCClientWithRegion":                 ecs.NewVPCClientWithRegion,
	"CopyImageArgs":                          qlang.StructOf((*ecs.CopyImageArgs)(nil)),
	"CopyImageResponse":                      qlang.StructOf((*ecs.CopyImageResponse)(nil)),
	"CreateDiskArgs":                         qlang.StructOf((*ecs.CreateDiskArgs)(nil)),
	"CreateDisksResponse":                    qlang.StructOf((*ecs.CreateDisksResponse)(nil)),
	"CreateForwardEntryArgs":                 qlang.StructOf((*ecs.CreateForwardEntryArgs)(nil)),
	"CreateForwardEntryResponse":             qlang.StructOf((*ecs.CreateForwardEntryResponse)(nil)),
	"CreateImageArgs":                        qlang.StructOf((*ecs.CreateImageArgs)(nil)),
	"CreateImageResponse":                    qlang.StructOf((*ecs.CreateImageResponse)(nil)),
	"CreateInstanceArgs":                     qlang.StructOf((*ecs.CreateInstanceArgs)(nil)),
	"CreateInstanceResponse":                 qlang.StructOf((*ecs.CreateInstanceResponse)(nil)),
	"CreateKeyPairArgs":                      qlang.StructOf((*ecs.CreateKeyPairArgs)(nil)),
	"CreateKeyPairResponse":                  qlang.StructOf((*ecs.CreateKeyPairResponse)(nil)),
	"CreateNatGatewayArgs":                   qlang.StructOf((*ecs.CreateNatGatewayArgs)(nil)),
	"CreateNatGatewayResponse":               qlang.StructOf((*ecs.CreateNatGatewayResponse)(nil)),
	"CreateRouteEntryArgs":                   qlang.StructOf((*ecs.CreateRouteEntryArgs)(nil)),
	"CreateRouteEntryResponse":               qlang.StructOf((*ecs.CreateRouteEntryResponse)(nil)),
	"CreateRouterInterfaceArgs":              qlang.StructOf((*ecs.CreateRouterInterfaceArgs)(nil)),
	"CreateRouterInterfaceResponse":          qlang.StructOf((*ecs.CreateRouterInterfaceResponse)(nil)),
	"CreateSecurityGroupArgs":                qlang.StructOf((*ecs.CreateSecurityGroupArgs)(nil)),
	"CreateSecurityGroupResponse":            qlang.StructOf((*ecs.CreateSecurityGroupResponse)(nil)),
	"CreateSnapshotArgs":                     qlang.StructOf((*ecs.CreateSnapshotArgs)(nil)),
	"CreateSnapshotResponse":                 qlang.StructOf((*ecs.CreateSnapshotResponse)(nil)),
	"CreateSnatEntryArgs":                    qlang.StructOf((*ecs.CreateSnatEntryArgs)(nil)),
	"CreateSnatEntryResponse":                qlang.StructOf((*ecs.CreateSnatEntryResponse)(nil)),
	"CreateVSwitchArgs":                      qlang.StructOf((*ecs.CreateVSwitchArgs)(nil)),
	"CreateVSwitchResponse":                  qlang.StructOf((*ecs.CreateVSwitchResponse)(nil)),
	"CreateVpcArgs":                          qlang.StructOf((*ecs.CreateVpcArgs)(nil)),
	"CreateVpcResponse":                      qlang.StructOf((*ecs.CreateVpcResponse)(nil)),
	"DataDiskType":                           qlang.StructOf((*ecs.DataDiskType)(nil)),
	"DeleteBandwidthPackageArgs":             qlang.StructOf((*ecs.DeleteBandwidthPackageArgs)(nil)),
	"DeleteBandwidthPackageResponse":         qlang.StructOf((*ecs.DeleteBandwidthPackageResponse)(nil)),
	"DeleteDiskArgs":                         qlang.StructOf((*ecs.DeleteDiskArgs)(nil)),
	"DeleteDiskResponse":                     qlang.StructOf((*ecs.DeleteDiskResponse)(nil)),
	"DeleteForwardEntryArgs":                 qlang.StructOf((*ecs.DeleteForwardEntryArgs)(nil)),
	"DeleteForwardEntryResponse":             qlang.StructOf((*ecs.DeleteForwardEntryResponse)(nil)),
	"DeleteImageArgs":                        qlang.StructOf((*ecs.DeleteImageArgs)(nil)),
	"DeleteImageResponse":                    qlang.StructOf((*ecs.DeleteImageResponse)(nil)),
	"DeleteInstanceArgs":                     qlang.StructOf((*ecs.DeleteInstanceArgs)(nil)),
	"DeleteInstanceResponse":                 qlang.StructOf((*ecs.DeleteInstanceResponse)(nil)),
	"DeleteKeyPairsArgs":                     qlang.StructOf((*ecs.DeleteKeyPairsArgs)(nil)),
	"DeleteNatGatewayArgs":                   qlang.StructOf((*ecs.DeleteNatGatewayArgs)(nil)),
	"DeleteNatGatewayResponse":               qlang.StructOf((*ecs.DeleteNatGatewayResponse)(nil)),
	"DeleteRouteEntryArgs":                   qlang.StructOf((*ecs.DeleteRouteEntryArgs)(nil)),
	"DeleteRouteEntryResponse":               qlang.StructOf((*ecs.DeleteRouteEntryResponse)(nil)),
	"DeleteSecurityGroupArgs":                qlang.StructOf((*ecs.DeleteSecurityGroupArgs)(nil)),
	"DeleteSecurityGroupResponse":            qlang.StructOf((*ecs.DeleteSecurityGroupResponse)(nil)),
	"DeleteSnapshotArgs":                     qlang.StructOf((*ecs.DeleteSnapshotArgs)(nil)),
	"DeleteSnapshotResponse":                 qlang.StructOf((*ecs.DeleteSnapshotResponse)(nil)),
	"DeleteSnatEntryArgs":                    qlang.StructOf((*ecs.DeleteSnatEntryArgs)(nil)),
	"DeleteSnatEntryResponse":                qlang.StructOf((*ecs.DeleteSnatEntryResponse)(nil)),
	"DeleteVSwitchArgs":                      qlang.StructOf((*ecs.DeleteVSwitchArgs)(nil)),
	"DeleteVSwitchResponse":                  qlang.StructOf((*ecs.DeleteVSwitchResponse)(nil)),
	"DeleteVpcArgs":                          qlang.StructOf((*ecs.DeleteVpcArgs)(nil)),
	"DeleteVpcResponse":                      qlang.StructOf((*ecs.DeleteVpcResponse)(nil)),
	"DescribeBandwidthPackageType":           qlang.StructOf((*ecs.DescribeBandwidthPackageType)(nil)),
	"DescribeBandwidthPackagesArgs":          qlang.StructOf((*ecs.DescribeBandwidthPackagesArgs)(nil)),
	"DescribeBandwidthPackagesResponse":      qlang.StructOf((*ecs.DescribeBandwidthPackagesResponse)(nil)),
	"DescribeDiskMonitorDataArgs":            qlang.StructOf((*ecs.DescribeDiskMonitorDataArgs)(nil)),
	"DescribeDiskMonitorDataResponse":        qlang.StructOf((*ecs.DescribeDiskMonitorDataResponse)(nil)),
	"DescribeDisksArgs":                      qlang.StructOf((*ecs.DescribeDisksArgs)(nil)),
	"DescribeDisksResponse":                  qlang.StructOf((*ecs.DescribeDisksResponse)(nil)),
	"DescribeEipAddressesArgs":               qlang.StructOf((*ecs.DescribeEipAddressesArgs)(nil)),
	"DescribeEipAddressesResponse":           qlang.StructOf((*ecs.DescribeEipAddressesResponse)(nil)),
	"DescribeEipMonitorDataArgs":             qlang.StructOf((*ecs.DescribeEipMonitorDataArgs)(nil)),
	"DescribeEipMonitorDataResponse":         qlang.StructOf((*ecs.DescribeEipMonitorDataResponse)(nil)),
	"DescribeForwardTableEntriesArgs":        qlang.StructOf((*ecs.DescribeForwardTableEntriesArgs)(nil)),
	"DescribeForwardTableEntriesResponse":    qlang.StructOf((*ecs.DescribeForwardTableEntriesResponse)(nil)),
	"DescribeImagesArgs":                     qlang.StructOf((*ecs.DescribeImagesArgs)(nil)),
	"DescribeImagesResponse":                 qlang.StructOf((*ecs.DescribeImagesResponse)(nil)),
	"DescribeInstanceAttributeArgs":          qlang.StructOf((*ecs.DescribeInstanceAttributeArgs)(nil)),
	"DescribeInstanceAttributeResponse":      qlang.StructOf((*ecs.DescribeInstanceAttributeResponse)(nil)),
	"DescribeInstanceMonitorDataArgs":        qlang.StructOf((*ecs.DescribeInstanceMonitorDataArgs)(nil)),
	"DescribeInstanceMonitorDataResponse":    qlang.StructOf((*ecs.DescribeInstanceMonitorDataResponse)(nil)),
	"DescribeInstanceRamRoleResponse":        qlang.StructOf((*ecs.DescribeInstanceRamRoleResponse)(nil)),
	"DescribeInstanceStatusArgs":             qlang.StructOf((*ecs.DescribeInstanceStatusArgs)(nil)),
	"DescribeInstanceStatusResponse":         qlang.StructOf((*ecs.DescribeInstanceStatusResponse)(nil)),
	"DescribeInstanceTypeFamiliesArgs":       qlang.StructOf((*ecs.DescribeInstanceTypeFamiliesArgs)(nil)),
	"DescribeInstanceTypeFamiliesResponse":   qlang.StructOf((*ecs.DescribeInstanceTypeFamiliesResponse)(nil)),
	"DescribeInstanceTypesArgs":              qlang.StructOf((*ecs.DescribeInstanceTypesArgs)(nil)),
	"DescribeInstanceTypesResponse":          qlang.StructOf((*ecs.DescribeInstanceTypesResponse)(nil)),
	"DescribeInstanceVncUrlArgs":             qlang.StructOf((*ecs.DescribeInstanceVncUrlArgs)(nil)),
	"DescribeInstanceVncUrlResponse":         qlang.StructOf((*ecs.DescribeInstanceVncUrlResponse)(nil)),
	"DescribeInstancesArgs":                  qlang.StructOf((*ecs.DescribeInstancesArgs)(nil)),
	"DescribeInstancesResponse":              qlang.StructOf((*ecs.DescribeInstancesResponse)(nil)),
	"DescribeKeyPairsArgs":                   qlang.StructOf((*ecs.DescribeKeyPairsArgs)(nil)),
	"DescribeKeyPairsResponse":               qlang.StructOf((*ecs.DescribeKeyPairsResponse)(nil)),
	"DescribeNatGatewayResponse":             qlang.StructOf((*ecs.DescribeNatGatewayResponse)(nil)),
	"DescribeNatGatewaysArgs":                qlang.StructOf((*ecs.DescribeNatGatewaysArgs)(nil)),
	"DescribeRegionsArgs":                    qlang.StructOf((*ecs.DescribeRegionsArgs)(nil)),
	"DescribeRegionsResponse":                qlang.StructOf((*ecs.DescribeRegionsResponse)(nil)),
	"DescribeResourceByTagsArgs":             qlang.StructOf((*ecs.DescribeResourceByTagsArgs)(nil)),
	"DescribeResourceByTagsResponse":         qlang.StructOf((*ecs.DescribeResourceByTagsResponse)(nil)),
	"DescribeRouteTablesArgs":                qlang.StructOf((*ecs.DescribeRouteTablesArgs)(nil)),
	"DescribeRouteTablesResponse":            qlang.StructOf((*ecs.DescribeRouteTablesResponse)(nil)),
	"DescribeRouterInterfacesArgs":           qlang.StructOf((*ecs.DescribeRouterInterfacesArgs)(nil)),
	"DescribeRouterInterfacesResponse":       qlang.StructOf((*ecs.DescribeRouterInterfacesResponse)(nil)),
	"DescribeSecurityGroupAttributeArgs":     qlang.StructOf((*ecs.DescribeSecurityGroupAttributeArgs)(nil)),
	"DescribeSecurityGroupAttributeResponse": qlang.StructOf((*ecs.DescribeSecurityGroupAttributeResponse)(nil)),
	"DescribeSecurityGroupsArgs":             qlang.StructOf((*ecs.DescribeSecurityGroupsArgs)(nil)),
	"DescribeSecurityGroupsResponse":         qlang.StructOf((*ecs.DescribeSecurityGroupsResponse)(nil)),
	"DescribeSnapshotsArgs":                  qlang.StructOf((*ecs.DescribeSnapshotsArgs)(nil)),
	"DescribeSnapshotsResponse":              qlang.StructOf((*ecs.DescribeSnapshotsResponse)(nil)),
	"DescribeSnatTableEntriesArgs":           qlang.StructOf((*ecs.DescribeSnatTableEntriesArgs)(nil)),
	"DescribeSnatTableEntriesResponse":       qlang.StructOf((*ecs.DescribeSnatTableEntriesResponse)(nil)),
	"DescribeTagsArgs":                       qlang.StructOf((*ecs.DescribeTagsArgs)(nil)),
	"DescribeTagsResponse":                   qlang.StructOf((*ecs.DescribeTagsResponse)(nil)),
	"DescribeUserdataArgs":                   qlang.StructOf((*ecs.DescribeUserdataArgs)(nil)),
	"DescribeUserdataItemType":               qlang.StructOf((*ecs.DescribeUserdataItemType)(nil)),
	"DescribeUserdataResponse":               qlang.StructOf((*ecs.DescribeUserdataResponse)(nil)),
	"DescribeVRoutersArgs":                   qlang.StructOf((*ecs.DescribeVRoutersArgs)(nil)),
	"DescribeVRoutersResponse":               qlang.StructOf((*ecs.DescribeVRoutersResponse)(nil)),
	"DescribeVSwitchesArgs":                  qlang.StructOf((*ecs.DescribeVSwitchesArgs)(nil)),
	"DescribeVSwitchesResponse":              qlang.StructOf((*ecs.DescribeVSwitchesResponse)(nil)),
	"DescribeVpcsArgs":                       qlang.StructOf((*ecs.DescribeVpcsArgs)(nil)),
	"DescribeVpcsResponse":                   qlang.StructOf((*ecs.DescribeVpcsResponse)(nil)),
	"DescribeZonesArgs":                      qlang.StructOf((*ecs.DescribeZonesArgs)(nil)),
	"DescribeZonesResponse":                  qlang.StructOf((*ecs.DescribeZonesResponse)(nil)),
	"DetachDiskArgs":                         qlang.StructOf((*ecs.DetachDiskArgs)(nil)),
	"DetachDiskResponse":                     qlang.StructOf((*ecs.DetachDiskResponse)(nil)),
	"DetachKeyPairArgs":                      qlang.StructOf((*ecs.DetachKeyPairArgs)(nil)),
	"DiskDeviceMapping":                      qlang.StructOf((*ecs.DiskDeviceMapping)(nil)),
	"DiskItemType":                           qlang.StructOf((*ecs.DiskItemType)(nil)),
	"DiskMonitorDataType":                    qlang.StructOf((*ecs.DiskMonitorDataType)(nil)),
	"EcsCommonResponse":                      qlang.StructOf((*ecs.EcsCommonResponse)(nil)),
	"EipAddressAssociateType":                qlang.StructOf((*ecs.EipAddressAssociateType)(nil)),
	"EipAddressSetType":                      qlang.StructOf((*ecs.EipAddressSetType)(nil)),
	"EipMonitorDataType":                     qlang.StructOf((*ecs.EipMonitorDataType)(nil)),
	"Filter":                                 qlang.StructOf((*ecs.Filter)(nil)),
	"ForwardTableEntrySetType":               qlang.StructOf((*ecs.ForwardTableEntrySetType)(nil)),
	"ForwardTableIdType":                     qlang.StructOf((*ecs.ForwardTableIdType)(nil)),
	"ImageSharePermissionResponse":           qlang.StructOf((*ecs.ImageSharePermissionResponse)(nil)),
	"ImageType":                              qlang.StructOf((*ecs.ImageType)(nil)),
	"ImportImageArgs":                        qlang.StructOf((*ecs.ImportImageArgs)(nil)),
	"ImportImageResponse":                    qlang.StructOf((*ecs.ImportImageResponse)(nil)),
	"ImportKeyPairArgs":                      qlang.StructOf((*ecs.ImportKeyPairArgs)(nil)),
	"ImportKeyPairResponse":                  qlang.StructOf((*ecs.ImportKeyPairResponse)(nil)),
	"InstanceAttributesType":                 qlang.StructOf((*ecs.InstanceAttributesType)(nil)),
	"InstanceIdSets":                         qlang.StructOf((*ecs.InstanceIdSets)(nil)),
	"InstanceMonitorDataType":                qlang.StructOf((*ecs.InstanceMonitorDataType)(nil)),
	"InstanceRamRoleSetType":                 qlang.StructOf((*ecs.InstanceRamRoleSetType)(nil)),
	"InstanceStatusItemType":                 qlang.StructOf((*ecs.InstanceStatusItemType)(nil)),
	"InstanceTypeFamilies":                   qlang.StructOf((*ecs.InstanceTypeFamilies)(nil)),
	"InstanceTypeFamily":                     qlang.StructOf((*ecs.InstanceTypeFamily)(nil)),
	"InstanceTypeItemType":                   qlang.StructOf((*ecs.InstanceTypeItemType)(nil)),
	"IpAddressSetType":                       qlang.StructOf((*ecs.IpAddressSetType)(nil)),
	"KeyPairItemType":                        qlang.StructOf((*ecs.KeyPairItemType)(nil)),
	"LockReasonType":                         qlang.StructOf((*ecs.LockReasonType)(nil)),
	"ModifyDiskAttributeArgs":                qlang.StructOf((*ecs.ModifyDiskAttributeArgs)(nil)),
	"ModifyDiskAttributeResponse":            qlang.StructOf((*ecs.ModifyDiskAttributeResponse)(nil)),
	"ModifyEipAddressAttributeArgs":          qlang.StructOf((*ecs.ModifyEipAddressAttributeArgs)(nil)),
	"ModifyEipAddressAttributeResponse":      qlang.StructOf((*ecs.ModifyEipAddressAttributeResponse)(nil)),
	"ModifyForwardEntryArgs":                 qlang.StructOf((*ecs.ModifyForwardEntryArgs)(nil)),
	"ModifyForwardEntryResponse":             qlang.StructOf((*ecs.ModifyForwardEntryResponse)(nil)),
	"ModifyImageSharePermissionArgs":         qlang.StructOf((*ecs.ModifyImageSharePermissionArgs)(nil)),
	"ModifyInstanceAttributeArgs":            qlang.StructOf((*ecs.ModifyInstanceAttributeArgs)(nil)),
	"ModifyInstanceAttributeResponse":        qlang.StructOf((*ecs.ModifyInstanceAttributeResponse)(nil)),
	"ModifyInstanceNetworkSpec":              qlang.StructOf((*ecs.ModifyInstanceNetworkSpec)(nil)),
	"ModifyInstanceNetworkSpecResponse":      qlang.StructOf((*ecs.ModifyInstanceNetworkSpecResponse)(nil)),
	"ModifyNatGatewayAttributeArgs":          qlang.StructOf((*ecs.ModifyNatGatewayAttributeArgs)(nil)),
	"ModifyNatGatewayAttributeResponse":      qlang.StructOf((*ecs.ModifyNatGatewayAttributeResponse)(nil)),
	"ModifyNatGatewaySpecArgs":               qlang.StructOf((*ecs.ModifyNatGatewaySpecArgs)(nil)),
	"ModifyRouterInterfaceAttributeArgs":     qlang.StructOf((*ecs.ModifyRouterInterfaceAttributeArgs)(nil)),
	"ModifyRouterInterfaceSpecArgs":          qlang.StructOf((*ecs.ModifyRouterInterfaceSpecArgs)(nil)),
	"ModifyRouterInterfaceSpecResponse":      qlang.StructOf((*ecs.ModifyRouterInterfaceSpecResponse)(nil)),
	"ModifySecurityGroupAttributeArgs":       qlang.StructOf((*ecs.ModifySecurityGroupAttributeArgs)(nil)),
	"ModifySecurityGroupAttributeResponse":   qlang.StructOf((*ecs.ModifySecurityGroupAttributeResponse)(nil)),
	"ModifySnatEntryArgs":                    qlang.StructOf((*ecs.ModifySnatEntryArgs)(nil)),
	"ModifySnatEntryResponse":                qlang.StructOf((*ecs.ModifySnatEntryResponse)(nil)),
	"ModifyVRouterAttributeArgs":             qlang.StructOf((*ecs.ModifyVRouterAttributeArgs)(nil)),
	"ModifyVRouterAttributeResponse":         qlang.StructOf((*ecs.ModifyVRouterAttributeResponse)(nil)),
	"ModifyVSwitchAttributeArgs":             qlang.StructOf((*ecs.ModifyVSwitchAttributeArgs)(nil)),
	"ModifyVSwitchAttributeResponse":         qlang.StructOf((*ecs.ModifyVSwitchAttributeResponse)(nil)),
	"ModifyVpcAttributeArgs":                 qlang.StructOf((*ecs.ModifyVpcAttributeArgs)(nil)),
	"ModifyVpcAttributeResponse":             qlang.StructOf((*ecs.ModifyVpcAttributeResponse)(nil)),
	"NatGatewaySetType":                      qlang.StructOf((*ecs.NatGatewaySetType)(nil)),
	"NextHopItemType":                        qlang.StructOf((*ecs.NextHopItemType)(nil)),
	"NextHopListType":                        qlang.StructOf((*ecs.NextHopListType)(nil)),
	"OperateRouterInterfaceArgs":             qlang.StructOf((*ecs.OperateRouterInterfaceArgs)(nil)),
	"OperationLocksType":                     qlang.StructOf((*ecs.OperationLocksType)(nil)),
	"PermissionType":                         qlang.StructOf((*ecs.PermissionType)(nil)),
	"PublicIpAddresseType":                   qlang.StructOf((*ecs.PublicIpAddresseType)(nil)),
	"ReInitDiskArgs":                         qlang.StructOf((*ecs.ReInitDiskArgs)(nil)),
	"ReInitDiskResponse":                     qlang.StructOf((*ecs.ReInitDiskResponse)(nil)),
	"RebootInstanceArgs":                     qlang.StructOf((*ecs.RebootInstanceArgs)(nil)),
	"RebootInstanceResponse":                 qlang.StructOf((*ecs.RebootInstanceResponse)(nil)),
	"RegionType":                             qlang.StructOf((*ecs.RegionType)(nil)),
	"ReleaseEipAddressArgs":                  qlang.StructOf((*ecs.ReleaseEipAddressArgs)(nil)),
	"ReleaseEipAddressResponse":              qlang.StructOf((*ecs.ReleaseEipAddressResponse)(nil)),
	"RemoveTagsArgs":                         qlang.StructOf((*ecs.RemoveTagsArgs)(nil)),
	"RemoveTagsResponse":                     qlang.StructOf((*ecs.RemoveTagsResponse)(nil)),
	"ReplaceSystemDiskArgs":                  qlang.StructOf((*ecs.ReplaceSystemDiskArgs)(nil)),
	"ReplaceSystemDiskResponse":              qlang.StructOf((*ecs.ReplaceSystemDiskResponse)(nil)),
	"ResetDiskArgs":                          qlang.StructOf((*ecs.ResetDiskArgs)(nil)),
	"ResetDiskResponse":                      qlang.StructOf((*ecs.ResetDiskResponse)(nil)),
	"ResourceItemType":                       qlang.StructOf((*ecs.ResourceItemType)(nil)),
	"ResourcesInfoType":                      qlang.StructOf((*ecs.ResourcesInfoType)(nil)),
	"RevokeSecurityGroupArgs":                qlang.StructOf((*ecs.RevokeSecurityGroupArgs)(nil)),
	"RevokeSecurityGroupEgressArgs":          qlang.StructOf((*ecs.RevokeSecurityGroupEgressArgs)(nil)),
	"RevokeSecurityGroupEgressResponse":      qlang.StructOf((*ecs.RevokeSecurityGroupEgressResponse)(nil)),
	"RevokeSecurityGroupResponse":            qlang.StructOf((*ecs.RevokeSecurityGroupResponse)(nil)),
	"RouteEntrySetType":                      qlang.StructOf((*ecs.RouteEntrySetType)(nil)),
	"RouteTableSetType":                      qlang.StructOf((*ecs.RouteTableSetType)(nil)),
	"RouterInterfaceItemType":                qlang.StructOf((*ecs.RouterInterfaceItemType)(nil)),
	"RunInstanceArgs":                        qlang.StructOf((*ecs.RunInstanceArgs)(nil)),
	"RunInstanceResponse":                    qlang.StructOf((*ecs.RunInstanceResponse)(nil)),
	"SecurityGroupArgs":                      qlang.StructOf((*ecs.SecurityGroupArgs)(nil)),
	"SecurityGroupIdSetType":                 qlang.StructOf((*ecs.SecurityGroupIdSetType)(nil)),
	"SecurityGroupItemType":                  qlang.StructOf((*ecs.SecurityGroupItemType)(nil)),
	"SecurityGroupResponse":                  qlang.StructOf((*ecs.SecurityGroupResponse)(nil)),
	"SnapshotType":                           qlang.StructOf((*ecs.SnapshotType)(nil)),
	"SnatEntrySetType":                       qlang.StructOf((*ecs.SnatEntrySetType)(nil)),
	"SnatTableIdType":                        qlang.StructOf((*ecs.SnatTableIdType)(nil)),
	"StartInstanceArgs":                      qlang.StructOf((*ecs.StartInstanceArgs)(nil)),
	"StartInstanceResponse":                  qlang.StructOf((*ecs.StartInstanceResponse)(nil)),
	"StopInstanceArgs":                       qlang.StructOf((*ecs.StopInstanceArgs)(nil)),
	"StopInstanceResponse":                   qlang.StructOf((*ecs.StopInstanceResponse)(nil)),
	"StringOrBool":                           qlang.StructOf((*ecs.StringOrBool)(nil)),
	"SystemDiskType":                         qlang.StructOf((*ecs.SystemDiskType)(nil)),
	"TagItemType":                            qlang.StructOf((*ecs.TagItemType)(nil)),
	"UnallocateEipAddressArgs":               qlang.StructOf((*ecs.UnallocateEipAddressArgs)(nil)),
	"UnallocateEipAddressResponse":           qlang.StructOf((*ecs.UnallocateEipAddressResponse)(nil)),
	"VRouterSetType":                         qlang.StructOf((*ecs.VRouterSetType)(nil)),
	"VSwitchSetType":                         qlang.StructOf((*ecs.VSwitchSetType)(nil)),
	"VpcAttributesType":                      qlang.StructOf((*ecs.VpcAttributesType)(nil)),
	"VpcSetType":                             qlang.StructOf((*ecs.VpcSetType)(nil)),
	"ZoneType":                               qlang.StructOf((*ecs.ZoneType)(nil)),

	"DiskCategory": qlang.StructOf((*ecs.DiskCategory)(nil)),
	"IoOptimized":  qlang.StructOf((*ecs.IoOptimized)(nil)),
}
