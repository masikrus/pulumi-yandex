// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "yandex:index/albBackendGroup:AlbBackendGroup":
		r = &AlbBackendGroup{}
	case "yandex:index/albHttpRouter:AlbHttpRouter":
		r = &AlbHttpRouter{}
	case "yandex:index/albLoadBalancer:AlbLoadBalancer":
		r = &AlbLoadBalancer{}
	case "yandex:index/albTargetGroup:AlbTargetGroup":
		r = &AlbTargetGroup{}
	case "yandex:index/albVirtualHost:AlbVirtualHost":
		r = &AlbVirtualHost{}
	case "yandex:index/apiGateway:ApiGateway":
		r = &ApiGateway{}
	case "yandex:index/auditTrailsTrail:AuditTrailsTrail":
		r = &AuditTrailsTrail{}
	case "yandex:index/backupPolicy:BackupPolicy":
		r = &BackupPolicy{}
	case "yandex:index/backupPolicyBindings:BackupPolicyBindings":
		r = &BackupPolicyBindings{}
	case "yandex:index/cdnOriginGroup:CdnOriginGroup":
		r = &CdnOriginGroup{}
	case "yandex:index/cdnResource:CdnResource":
		r = &CdnResource{}
	case "yandex:index/cmCertificate:CmCertificate":
		r = &CmCertificate{}
	case "yandex:index/cmCertificateIamBinding:CmCertificateIamBinding":
		r = &CmCertificateIamBinding{}
	case "yandex:index/cmCertificateIamMember:CmCertificateIamMember":
		r = &CmCertificateIamMember{}
	case "yandex:index/computeDisk:ComputeDisk":
		r = &ComputeDisk{}
	case "yandex:index/computeDiskPlacementGroup:ComputeDiskPlacementGroup":
		r = &ComputeDiskPlacementGroup{}
	case "yandex:index/computeFilesystem:ComputeFilesystem":
		r = &ComputeFilesystem{}
	case "yandex:index/computeGpuCluster:ComputeGpuCluster":
		r = &ComputeGpuCluster{}
	case "yandex:index/computeImage:ComputeImage":
		r = &ComputeImage{}
	case "yandex:index/computeInstance:ComputeInstance":
		r = &ComputeInstance{}
	case "yandex:index/computeInstanceGroup:ComputeInstanceGroup":
		r = &ComputeInstanceGroup{}
	case "yandex:index/computePlacementGroup:ComputePlacementGroup":
		r = &ComputePlacementGroup{}
	case "yandex:index/computeSnapshot:ComputeSnapshot":
		r = &ComputeSnapshot{}
	case "yandex:index/computeSnapshotSchedule:ComputeSnapshotSchedule":
		r = &ComputeSnapshotSchedule{}
	case "yandex:index/containerRegistry:ContainerRegistry":
		r = &ContainerRegistry{}
	case "yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding":
		r = &ContainerRegistryIamBinding{}
	case "yandex:index/containerRegistryIpPermission:ContainerRegistryIpPermission":
		r = &ContainerRegistryIpPermission{}
	case "yandex:index/containerRepository:ContainerRepository":
		r = &ContainerRepository{}
	case "yandex:index/containerRepositoryIamBinding:ContainerRepositoryIamBinding":
		r = &ContainerRepositoryIamBinding{}
	case "yandex:index/containerRepositoryLifecyclePolicy:ContainerRepositoryLifecyclePolicy":
		r = &ContainerRepositoryLifecyclePolicy{}
	case "yandex:index/dataprocCluster:DataprocCluster":
		r = &DataprocCluster{}
	case "yandex:index/datatransferEndpoint:DatatransferEndpoint":
		r = &DatatransferEndpoint{}
	case "yandex:index/datatransferTransfer:DatatransferTransfer":
		r = &DatatransferTransfer{}
	case "yandex:index/dnsRecordset:DnsRecordset":
		r = &DnsRecordset{}
	case "yandex:index/dnsZone:DnsZone":
		r = &DnsZone{}
	case "yandex:index/dnsZoneIamBinding:DnsZoneIamBinding":
		r = &DnsZoneIamBinding{}
	case "yandex:index/function:Function":
		r = &Function{}
	case "yandex:index/functionIamBinding:FunctionIamBinding":
		r = &FunctionIamBinding{}
	case "yandex:index/functionScalingPolicy:FunctionScalingPolicy":
		r = &FunctionScalingPolicy{}
	case "yandex:index/functionTrigger:FunctionTrigger":
		r = &FunctionTrigger{}
	case "yandex:index/iamServiceAccount:IamServiceAccount":
		r = &IamServiceAccount{}
	case "yandex:index/iamServiceAccountApiKey:IamServiceAccountApiKey":
		r = &IamServiceAccountApiKey{}
	case "yandex:index/iamServiceAccountIamBinding:IamServiceAccountIamBinding":
		r = &IamServiceAccountIamBinding{}
	case "yandex:index/iamServiceAccountIamMember:IamServiceAccountIamMember":
		r = &IamServiceAccountIamMember{}
	case "yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy":
		r = &IamServiceAccountIamPolicy{}
	case "yandex:index/iamServiceAccountKey:IamServiceAccountKey":
		r = &IamServiceAccountKey{}
	case "yandex:index/iamServiceAccountStaticAccessKey:IamServiceAccountStaticAccessKey":
		r = &IamServiceAccountStaticAccessKey{}
	case "yandex:index/iamWorkloadIdentityFederatedCredential:IamWorkloadIdentityFederatedCredential":
		r = &IamWorkloadIdentityFederatedCredential{}
	case "yandex:index/iamWorkloadIdentityOidcFederation:IamWorkloadIdentityOidcFederation":
		r = &IamWorkloadIdentityOidcFederation{}
	case "yandex:index/iamWorkloadIdentityOidcFederationIamBinding:IamWorkloadIdentityOidcFederationIamBinding":
		r = &IamWorkloadIdentityOidcFederationIamBinding{}
	case "yandex:index/iotCoreBroker:IotCoreBroker":
		r = &IotCoreBroker{}
	case "yandex:index/iotCoreDevice:IotCoreDevice":
		r = &IotCoreDevice{}
	case "yandex:index/iotCoreRegistry:IotCoreRegistry":
		r = &IotCoreRegistry{}
	case "yandex:index/kmsAsymmetricEncryptionKey:KmsAsymmetricEncryptionKey":
		r = &KmsAsymmetricEncryptionKey{}
	case "yandex:index/kmsAsymmetricEncryptionKeyIamBinding:KmsAsymmetricEncryptionKeyIamBinding":
		r = &KmsAsymmetricEncryptionKeyIamBinding{}
	case "yandex:index/kmsAsymmetricEncryptionKeyIamMember:KmsAsymmetricEncryptionKeyIamMember":
		r = &KmsAsymmetricEncryptionKeyIamMember{}
	case "yandex:index/kmsAsymmetricSignatureKey:KmsAsymmetricSignatureKey":
		r = &KmsAsymmetricSignatureKey{}
	case "yandex:index/kmsAsymmetricSignatureKeyIamBinding:KmsAsymmetricSignatureKeyIamBinding":
		r = &KmsAsymmetricSignatureKeyIamBinding{}
	case "yandex:index/kmsAsymmetricSignatureKeyIamMember:KmsAsymmetricSignatureKeyIamMember":
		r = &KmsAsymmetricSignatureKeyIamMember{}
	case "yandex:index/kmsSecretCiphertext:KmsSecretCiphertext":
		r = &KmsSecretCiphertext{}
	case "yandex:index/kmsSymmetricKey:KmsSymmetricKey":
		r = &KmsSymmetricKey{}
	case "yandex:index/kmsSymmetricKeyIamBinding:KmsSymmetricKeyIamBinding":
		r = &KmsSymmetricKeyIamBinding{}
	case "yandex:index/kmsSymmetricKeyIamMember:KmsSymmetricKeyIamMember":
		r = &KmsSymmetricKeyIamMember{}
	case "yandex:index/kubernetesCluster:KubernetesCluster":
		r = &KubernetesCluster{}
	case "yandex:index/kubernetesNodeGroup:KubernetesNodeGroup":
		r = &KubernetesNodeGroup{}
	case "yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer":
		r = &LbNetworkLoadBalancer{}
	case "yandex:index/lbTargetGroup:LbTargetGroup":
		r = &LbTargetGroup{}
	case "yandex:index/loadtestingAgent:LoadtestingAgent":
		r = &LoadtestingAgent{}
	case "yandex:index/lockboxSecret:LockboxSecret":
		r = &LockboxSecret{}
	case "yandex:index/lockboxSecretIamBinding:LockboxSecretIamBinding":
		r = &LockboxSecretIamBinding{}
	case "yandex:index/lockboxSecretIamMember:LockboxSecretIamMember":
		r = &LockboxSecretIamMember{}
	case "yandex:index/lockboxSecretVersion:LockboxSecretVersion":
		r = &LockboxSecretVersion{}
	case "yandex:index/lockboxSecretVersionHashed:LockboxSecretVersionHashed":
		r = &LockboxSecretVersionHashed{}
	case "yandex:index/loggingGroup:LoggingGroup":
		r = &LoggingGroup{}
	case "yandex:index/mdbClickhouseCluster:MdbClickhouseCluster":
		r = &MdbClickhouseCluster{}
	case "yandex:index/mdbElasticsearchCluster:MdbElasticsearchCluster":
		r = &MdbElasticsearchCluster{}
	case "yandex:index/mdbGreenplumCluster:MdbGreenplumCluster":
		r = &MdbGreenplumCluster{}
	case "yandex:index/mdbKafkaCluster:MdbKafkaCluster":
		r = &MdbKafkaCluster{}
	case "yandex:index/mdbKafkaConnector:MdbKafkaConnector":
		r = &MdbKafkaConnector{}
	case "yandex:index/mdbKafkaTopic:MdbKafkaTopic":
		r = &MdbKafkaTopic{}
	case "yandex:index/mdbKafkaUser:MdbKafkaUser":
		r = &MdbKafkaUser{}
	case "yandex:index/mdbMongodbCluster:MdbMongodbCluster":
		r = &MdbMongodbCluster{}
	case "yandex:index/mdbMysqlCluster:MdbMysqlCluster":
		r = &MdbMysqlCluster{}
	case "yandex:index/mdbMysqlDatabase:MdbMysqlDatabase":
		r = &MdbMysqlDatabase{}
	case "yandex:index/mdbMysqlUser:MdbMysqlUser":
		r = &MdbMysqlUser{}
	case "yandex:index/mdbPostgresqlCluster:MdbPostgresqlCluster":
		r = &MdbPostgresqlCluster{}
	case "yandex:index/mdbPostgresqlDatabase:MdbPostgresqlDatabase":
		r = &MdbPostgresqlDatabase{}
	case "yandex:index/mdbPostgresqlUser:MdbPostgresqlUser":
		r = &MdbPostgresqlUser{}
	case "yandex:index/mdbRedisCluster:MdbRedisCluster":
		r = &MdbRedisCluster{}
	case "yandex:index/mdbSqlserverCluster:MdbSqlserverCluster":
		r = &MdbSqlserverCluster{}
	case "yandex:index/messageQueue:MessageQueue":
		r = &MessageQueue{}
	case "yandex:index/monitoringDashboard:MonitoringDashboard":
		r = &MonitoringDashboard{}
	case "yandex:index/organizationmanagerGroup:OrganizationmanagerGroup":
		r = &OrganizationmanagerGroup{}
	case "yandex:index/organizationmanagerGroupIamMember:OrganizationmanagerGroupIamMember":
		r = &OrganizationmanagerGroupIamMember{}
	case "yandex:index/organizationmanagerGroupMapping:OrganizationmanagerGroupMapping":
		r = &OrganizationmanagerGroupMapping{}
	case "yandex:index/organizationmanagerGroupMappingItem:OrganizationmanagerGroupMappingItem":
		r = &OrganizationmanagerGroupMappingItem{}
	case "yandex:index/organizationmanagerGroupMembership:OrganizationmanagerGroupMembership":
		r = &OrganizationmanagerGroupMembership{}
	case "yandex:index/organizationmanagerOrganizationIamBinding:OrganizationmanagerOrganizationIamBinding":
		r = &OrganizationmanagerOrganizationIamBinding{}
	case "yandex:index/organizationmanagerOrganizationIamMember:OrganizationmanagerOrganizationIamMember":
		r = &OrganizationmanagerOrganizationIamMember{}
	case "yandex:index/organizationmanagerOsLoginSettings:OrganizationmanagerOsLoginSettings":
		r = &OrganizationmanagerOsLoginSettings{}
	case "yandex:index/organizationmanagerSamlFederation:OrganizationmanagerSamlFederation":
		r = &OrganizationmanagerSamlFederation{}
	case "yandex:index/organizationmanagerSamlFederationUserAccount:OrganizationmanagerSamlFederationUserAccount":
		r = &OrganizationmanagerSamlFederationUserAccount{}
	case "yandex:index/organizationmanagerUserSshKey:OrganizationmanagerUserSshKey":
		r = &OrganizationmanagerUserSshKey{}
	case "yandex:index/resourcemanagerCloud:ResourcemanagerCloud":
		r = &ResourcemanagerCloud{}
	case "yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding":
		r = &ResourcemanagerCloudIamBinding{}
	case "yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember":
		r = &ResourcemanagerCloudIamMember{}
	case "yandex:index/resourcemanagerFolder:ResourcemanagerFolder":
		r = &ResourcemanagerFolder{}
	case "yandex:index/resourcemanagerFolderIamBinding:ResourcemanagerFolderIamBinding":
		r = &ResourcemanagerFolderIamBinding{}
	case "yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember":
		r = &ResourcemanagerFolderIamMember{}
	case "yandex:index/resourcemanagerFolderIamPolicy:ResourcemanagerFolderIamPolicy":
		r = &ResourcemanagerFolderIamPolicy{}
	case "yandex:index/serverlessContainer:ServerlessContainer":
		r = &ServerlessContainer{}
	case "yandex:index/serverlessContainerIamBinding:ServerlessContainerIamBinding":
		r = &ServerlessContainerIamBinding{}
	case "yandex:index/serverlessEventrouterBus:ServerlessEventrouterBus":
		r = &ServerlessEventrouterBus{}
	case "yandex:index/serverlessEventrouterConnector:ServerlessEventrouterConnector":
		r = &ServerlessEventrouterConnector{}
	case "yandex:index/serverlessEventrouterRule:ServerlessEventrouterRule":
		r = &ServerlessEventrouterRule{}
	case "yandex:index/smartcaptchaCaptcha:SmartcaptchaCaptcha":
		r = &SmartcaptchaCaptcha{}
	case "yandex:index/storageBucket:StorageBucket":
		r = &StorageBucket{}
	case "yandex:index/storageObject:StorageObject":
		r = &StorageObject{}
	case "yandex:index/swsAdvancedRateLimiterProfile:SwsAdvancedRateLimiterProfile":
		r = &SwsAdvancedRateLimiterProfile{}
	case "yandex:index/swsSecurityProfile:SwsSecurityProfile":
		r = &SwsSecurityProfile{}
	case "yandex:index/swsWafProfile:SwsWafProfile":
		r = &SwsWafProfile{}
	case "yandex:index/vpcAddress:VpcAddress":
		r = &VpcAddress{}
	case "yandex:index/vpcDefaultSecurityGroup:VpcDefaultSecurityGroup":
		r = &VpcDefaultSecurityGroup{}
	case "yandex:index/vpcGateway:VpcGateway":
		r = &VpcGateway{}
	case "yandex:index/vpcNetwork:VpcNetwork":
		r = &VpcNetwork{}
	case "yandex:index/vpcPrivateEndpoint:VpcPrivateEndpoint":
		r = &VpcPrivateEndpoint{}
	case "yandex:index/vpcRouteTable:VpcRouteTable":
		r = &VpcRouteTable{}
	case "yandex:index/vpcSecurityGroup:VpcSecurityGroup":
		r = &VpcSecurityGroup{}
	case "yandex:index/vpcSubnet:VpcSubnet":
		r = &VpcSubnet{}
	case "yandex:index/ydbDatabaseDedicated:YdbDatabaseDedicated":
		r = &YdbDatabaseDedicated{}
	case "yandex:index/ydbDatabaseIamBinding:YdbDatabaseIamBinding":
		r = &YdbDatabaseIamBinding{}
	case "yandex:index/ydbDatabaseServerless:YdbDatabaseServerless":
		r = &YdbDatabaseServerless{}
	case "yandex:index/ydbTable:YdbTable":
		r = &YdbTable{}
	case "yandex:index/ydbTableChangefeed:YdbTableChangefeed":
		r = &YdbTableChangefeed{}
	case "yandex:index/ydbTableIndex:YdbTableIndex":
		r = &YdbTableIndex{}
	case "yandex:index/ydbTopic:YdbTopic":
		r = &YdbTopic{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:yandex" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"yandex",
		"index/albBackendGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/albHttpRouter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/albLoadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/albTargetGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/albVirtualHost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/apiGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/auditTrailsTrail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/backupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/backupPolicyBindings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/cdnOriginGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/cdnResource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/cmCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/cmCertificateIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/cmCertificateIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeDisk",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeDiskPlacementGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeFilesystem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeGpuCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeImage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeInstanceGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computePlacementGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/computeSnapshotSchedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/containerRegistry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/containerRegistryIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/containerRegistryIpPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/containerRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/containerRepositoryIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/containerRepositoryLifecyclePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/dataprocCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/datatransferEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/datatransferTransfer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/dnsRecordset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/dnsZone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/dnsZoneIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/functionIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/functionScalingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/functionTrigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamServiceAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamServiceAccountApiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamServiceAccountIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamServiceAccountIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamServiceAccountIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamServiceAccountKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamServiceAccountStaticAccessKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamWorkloadIdentityFederatedCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamWorkloadIdentityOidcFederation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iamWorkloadIdentityOidcFederationIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iotCoreBroker",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iotCoreDevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/iotCoreRegistry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsAsymmetricEncryptionKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsAsymmetricEncryptionKeyIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsAsymmetricEncryptionKeyIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsAsymmetricSignatureKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsAsymmetricSignatureKeyIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsAsymmetricSignatureKeyIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsSecretCiphertext",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsSymmetricKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsSymmetricKeyIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kmsSymmetricKeyIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kubernetesCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/kubernetesNodeGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/lbNetworkLoadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/lbTargetGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/loadtestingAgent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/lockboxSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/lockboxSecretIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/lockboxSecretIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/lockboxSecretVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/lockboxSecretVersionHashed",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/loggingGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbClickhouseCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbElasticsearchCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbGreenplumCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbKafkaCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbKafkaConnector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbKafkaTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbKafkaUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbMongodbCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbMysqlCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbMysqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbMysqlUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbPostgresqlCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbPostgresqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbPostgresqlUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbRedisCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/mdbSqlserverCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/messageQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/monitoringDashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerGroupIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerGroupMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerGroupMappingItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerGroupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerOrganizationIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerOrganizationIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerOsLoginSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerSamlFederation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerSamlFederationUserAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/organizationmanagerUserSshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/resourcemanagerCloud",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/resourcemanagerCloudIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/resourcemanagerCloudIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/resourcemanagerFolder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/resourcemanagerFolderIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/resourcemanagerFolderIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/resourcemanagerFolderIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/serverlessContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/serverlessContainerIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/serverlessEventrouterBus",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/serverlessEventrouterConnector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/serverlessEventrouterRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/smartcaptchaCaptcha",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/storageBucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/storageObject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/swsAdvancedRateLimiterProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/swsSecurityProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/swsWafProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcDefaultSecurityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcPrivateEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcRouteTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcSecurityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/vpcSubnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/ydbDatabaseDedicated",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/ydbDatabaseIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/ydbDatabaseServerless",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/ydbTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/ydbTableChangefeed",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/ydbTableIndex",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"yandex",
		"index/ydbTopic",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"yandex",
		&pkg{version},
	)
}
