# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
from . import _utilities
import typing
# Export this package's modules as members:
from .alb_backend_group import *
from .alb_http_router import *
from .alb_load_balancer import *
from .alb_target_group import *
from .alb_virtual_host import *
from .api_gateway import *
from .audit_trails_trail import *
from .backup_policy import *
from .backup_policy_bindings import *
from .cdn_origin_group import *
from .cdn_resource import *
from .cm_certificate import *
from .cm_certificate_iam_binding import *
from .cm_certificate_iam_member import *
from .compute_disk import *
from .compute_disk_placement_group import *
from .compute_filesystem import *
from .compute_gpu_cluster import *
from .compute_image import *
from .compute_instance import *
from .compute_instance_group import *
from .compute_placement_group import *
from .compute_snapshot import *
from .compute_snapshot_schedule import *
from .container_registry import *
from .container_registry_iam_binding import *
from .container_registry_ip_permission import *
from .container_repository import *
from .container_repository_iam_binding import *
from .container_repository_lifecycle_policy import *
from .dataproc_cluster import *
from .datatransfer_endpoint import *
from .datatransfer_transfer import *
from .dns_recordset import *
from .dns_zone import *
from .dns_zone_iam_binding import *
from .function import *
from .function_iam_binding import *
from .function_scaling_policy import *
from .function_trigger import *
from .get_alb_backend_group import *
from .get_alb_http_router import *
from .get_alb_load_balancer import *
from .get_alb_target_group import *
from .get_alb_virtual_host import *
from .get_api_gateway import *
from .get_audit_trails_trail import *
from .get_backup_policy import *
from .get_cdn_origin_group import *
from .get_cdn_resource import *
from .get_client_config import *
from .get_cm_certificate import *
from .get_cm_certificate_content import *
from .get_compute_disk import *
from .get_compute_disk_placement_group import *
from .get_compute_filesystem import *
from .get_compute_gpu_cluster import *
from .get_compute_image import *
from .get_compute_instance import *
from .get_compute_instance_group import *
from .get_compute_placement_group import *
from .get_compute_snapshot import *
from .get_compute_snapshot_schedule import *
from .get_container_registry import *
from .get_container_registry_ip_permission import *
from .get_container_repository import *
from .get_container_repository_lifecycle_policy import *
from .get_dataproc_cluster import *
from .get_dns_zone import *
from .get_function import *
from .get_function_scaling_policy import *
from .get_function_trigger import *
from .get_iam_policy import *
from .get_iam_role import *
from .get_iam_service_account import *
from .get_iam_service_agent import *
from .get_iam_user import *
from .get_iam_workload_identity_federated_credential import *
from .get_iam_workload_identity_oidc_federation import *
from .get_iot_core_broker import *
from .get_iot_core_device import *
from .get_iot_core_registry import *
from .get_kms_asymmetric_encryption_key import *
from .get_kms_asymmetric_signature_key import *
from .get_kms_symmetric_key import *
from .get_kubernetes_cluster import *
from .get_kubernetes_node_group import *
from .get_lb_network_load_balancer import *
from .get_lb_target_group import *
from .get_loadtesting_agent import *
from .get_lockbox_secret import *
from .get_lockbox_secret_version import *
from .get_logging_group import *
from .get_mdb_clickhouse_cluster import *
from .get_mdb_elasticsearch_cluster import *
from .get_mdb_greenplum_cluster import *
from .get_mdb_kafka_cluster import *
from .get_mdb_kafka_connector import *
from .get_mdb_kafka_topic import *
from .get_mdb_kafka_user import *
from .get_mdb_mongodb_cluster import *
from .get_mdb_mysql_cluster import *
from .get_mdb_mysql_database import *
from .get_mdb_mysql_user import *
from .get_mdb_postgresql_cluster import *
from .get_mdb_postgresql_database import *
from .get_mdb_postgresql_user import *
from .get_mdb_redis_cluster import *
from .get_mdb_sqlserver_cluster import *
from .get_message_queue import *
from .get_monitoring_dashboard import *
from .get_organizationmanager_group import *
from .get_organizationmanager_os_login_settings import *
from .get_organizationmanager_saml_federation import *
from .get_organizationmanager_saml_federation_user_account import *
from .get_organizationmanager_user_ssh_key import *
from .get_resourcemanager_cloud import *
from .get_resourcemanager_folder import *
from .get_serverless_container import *
from .get_serverless_eventrouter_bus import *
from .get_serverless_eventrouter_connector import *
from .get_serverless_eventrouter_rule import *
from .get_smartcaptcha_captcha import *
from .get_sws_advanced_rate_limiter_profile import *
from .get_sws_security_profile import *
from .get_sws_waf_profile import *
from .get_sws_waf_rule_set_descriptor import *
from .get_vpc_address import *
from .get_vpc_gateway import *
from .get_vpc_network import *
from .get_vpc_private_endpoint import *
from .get_vpc_route_table import *
from .get_vpc_security_group import *
from .get_vpc_subnet import *
from .get_ydb_database_dedicated import *
from .get_ydb_database_serverless import *
from .iam_service_account import *
from .iam_service_account_api_key import *
from .iam_service_account_iam_binding import *
from .iam_service_account_iam_member import *
from .iam_service_account_iam_policy import *
from .iam_service_account_key import *
from .iam_service_account_static_access_key import *
from .iam_workload_identity_federated_credential import *
from .iam_workload_identity_oidc_federation import *
from .iam_workload_identity_oidc_federation_iam_binding import *
from .iot_core_broker import *
from .iot_core_device import *
from .iot_core_registry import *
from .kms_asymmetric_encryption_key import *
from .kms_asymmetric_encryption_key_iam_binding import *
from .kms_asymmetric_encryption_key_iam_member import *
from .kms_asymmetric_signature_key import *
from .kms_asymmetric_signature_key_iam_binding import *
from .kms_asymmetric_signature_key_iam_member import *
from .kms_secret_ciphertext import *
from .kms_symmetric_key import *
from .kms_symmetric_key_iam_binding import *
from .kms_symmetric_key_iam_member import *
from .kubernetes_cluster import *
from .kubernetes_node_group import *
from .lb_network_load_balancer import *
from .lb_target_group import *
from .loadtesting_agent import *
from .lockbox_secret import *
from .lockbox_secret_iam_binding import *
from .lockbox_secret_iam_member import *
from .lockbox_secret_version import *
from .lockbox_secret_version_hashed import *
from .logging_group import *
from .mdb_clickhouse_cluster import *
from .mdb_elasticsearch_cluster import *
from .mdb_greenplum_cluster import *
from .mdb_kafka_cluster import *
from .mdb_kafka_connector import *
from .mdb_kafka_topic import *
from .mdb_kafka_user import *
from .mdb_mongodb_cluster import *
from .mdb_mysql_cluster import *
from .mdb_mysql_database import *
from .mdb_mysql_user import *
from .mdb_postgresql_cluster import *
from .mdb_postgresql_database import *
from .mdb_postgresql_user import *
from .mdb_redis_cluster import *
from .mdb_sqlserver_cluster import *
from .message_queue import *
from .monitoring_dashboard import *
from .organizationmanager_group import *
from .organizationmanager_group_iam_member import *
from .organizationmanager_group_mapping import *
from .organizationmanager_group_mapping_item import *
from .organizationmanager_group_membership import *
from .organizationmanager_organization_iam_binding import *
from .organizationmanager_organization_iam_member import *
from .organizationmanager_os_login_settings import *
from .organizationmanager_saml_federation import *
from .organizationmanager_saml_federation_user_account import *
from .organizationmanager_user_ssh_key import *
from .provider import *
from .resourcemanager_cloud import *
from .resourcemanager_cloud_iam_binding import *
from .resourcemanager_cloud_iam_member import *
from .resourcemanager_folder import *
from .resourcemanager_folder_iam_binding import *
from .resourcemanager_folder_iam_member import *
from .resourcemanager_folder_iam_policy import *
from .serverless_container import *
from .serverless_container_iam_binding import *
from .serverless_eventrouter_bus import *
from .serverless_eventrouter_connector import *
from .serverless_eventrouter_rule import *
from .smartcaptcha_captcha import *
from .storage_bucket import *
from .storage_object import *
from .sws_advanced_rate_limiter_profile import *
from .sws_security_profile import *
from .sws_waf_profile import *
from .vpc_address import *
from .vpc_default_security_group import *
from .vpc_gateway import *
from .vpc_network import *
from .vpc_private_endpoint import *
from .vpc_route_table import *
from .vpc_security_group import *
from .vpc_subnet import *
from .ydb_database_dedicated import *
from .ydb_database_iam_binding import *
from .ydb_database_serverless import *
from .ydb_table import *
from .ydb_table_changefeed import *
from .ydb_table_index import *
from .ydb_topic import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_yandex.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_yandex.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "yandex",
  "mod": "index/albBackendGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albBackendGroup:AlbBackendGroup": "AlbBackendGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albHttpRouter",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albHttpRouter:AlbHttpRouter": "AlbHttpRouter"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albLoadBalancer",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albLoadBalancer:AlbLoadBalancer": "AlbLoadBalancer"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albTargetGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albTargetGroup:AlbTargetGroup": "AlbTargetGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/albVirtualHost",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/albVirtualHost:AlbVirtualHost": "AlbVirtualHost"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/apiGateway",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/apiGateway:ApiGateway": "ApiGateway"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/auditTrailsTrail",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/auditTrailsTrail:AuditTrailsTrail": "AuditTrailsTrail"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/backupPolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/backupPolicy:BackupPolicy": "BackupPolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/backupPolicyBindings",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/backupPolicyBindings:BackupPolicyBindings": "BackupPolicyBindings"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/cdnOriginGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/cdnOriginGroup:CdnOriginGroup": "CdnOriginGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/cdnResource",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/cdnResource:CdnResource": "CdnResource"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/cmCertificate",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/cmCertificate:CmCertificate": "CmCertificate"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/cmCertificateIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/cmCertificateIamBinding:CmCertificateIamBinding": "CmCertificateIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/cmCertificateIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/cmCertificateIamMember:CmCertificateIamMember": "CmCertificateIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeDisk",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeDisk:ComputeDisk": "ComputeDisk"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeDiskPlacementGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeDiskPlacementGroup:ComputeDiskPlacementGroup": "ComputeDiskPlacementGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeFilesystem",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeFilesystem:ComputeFilesystem": "ComputeFilesystem"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeGpuCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeGpuCluster:ComputeGpuCluster": "ComputeGpuCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeImage",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeImage:ComputeImage": "ComputeImage"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeInstance",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeInstance:ComputeInstance": "ComputeInstance"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeInstanceGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeInstanceGroup:ComputeInstanceGroup": "ComputeInstanceGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computePlacementGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computePlacementGroup:ComputePlacementGroup": "ComputePlacementGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeSnapshot",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeSnapshot:ComputeSnapshot": "ComputeSnapshot"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/computeSnapshotSchedule",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/computeSnapshotSchedule:ComputeSnapshotSchedule": "ComputeSnapshotSchedule"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRegistry",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRegistry:ContainerRegistry": "ContainerRegistry"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRegistryIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRegistryIamBinding:ContainerRegistryIamBinding": "ContainerRegistryIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRegistryIpPermission",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRegistryIpPermission:ContainerRegistryIpPermission": "ContainerRegistryIpPermission"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRepository",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRepository:ContainerRepository": "ContainerRepository"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRepositoryIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRepositoryIamBinding:ContainerRepositoryIamBinding": "ContainerRepositoryIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/containerRepositoryLifecyclePolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/containerRepositoryLifecyclePolicy:ContainerRepositoryLifecyclePolicy": "ContainerRepositoryLifecyclePolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/dataprocCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/dataprocCluster:DataprocCluster": "DataprocCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/datatransferEndpoint",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/datatransferEndpoint:DatatransferEndpoint": "DatatransferEndpoint"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/datatransferTransfer",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/datatransferTransfer:DatatransferTransfer": "DatatransferTransfer"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/dnsRecordset",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/dnsRecordset:DnsRecordset": "DnsRecordset"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/dnsZone",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/dnsZone:DnsZone": "DnsZone"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/dnsZoneIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/dnsZoneIamBinding:DnsZoneIamBinding": "DnsZoneIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/function",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/function:Function": "Function"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/functionIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/functionIamBinding:FunctionIamBinding": "FunctionIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/functionScalingPolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/functionScalingPolicy:FunctionScalingPolicy": "FunctionScalingPolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/functionTrigger",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/functionTrigger:FunctionTrigger": "FunctionTrigger"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccount",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccount:IamServiceAccount": "IamServiceAccount"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountApiKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountApiKey:IamServiceAccountApiKey": "IamServiceAccountApiKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountIamBinding:IamServiceAccountIamBinding": "IamServiceAccountIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountIamMember:IamServiceAccountIamMember": "IamServiceAccountIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountIamPolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy": "IamServiceAccountIamPolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountKey:IamServiceAccountKey": "IamServiceAccountKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamServiceAccountStaticAccessKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamServiceAccountStaticAccessKey:IamServiceAccountStaticAccessKey": "IamServiceAccountStaticAccessKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamWorkloadIdentityFederatedCredential",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamWorkloadIdentityFederatedCredential:IamWorkloadIdentityFederatedCredential": "IamWorkloadIdentityFederatedCredential"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamWorkloadIdentityOidcFederation",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamWorkloadIdentityOidcFederation:IamWorkloadIdentityOidcFederation": "IamWorkloadIdentityOidcFederation"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iamWorkloadIdentityOidcFederationIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iamWorkloadIdentityOidcFederationIamBinding:IamWorkloadIdentityOidcFederationIamBinding": "IamWorkloadIdentityOidcFederationIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iotCoreBroker",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iotCoreBroker:IotCoreBroker": "IotCoreBroker"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iotCoreDevice",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iotCoreDevice:IotCoreDevice": "IotCoreDevice"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/iotCoreRegistry",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/iotCoreRegistry:IotCoreRegistry": "IotCoreRegistry"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsAsymmetricEncryptionKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsAsymmetricEncryptionKey:KmsAsymmetricEncryptionKey": "KmsAsymmetricEncryptionKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsAsymmetricEncryptionKeyIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsAsymmetricEncryptionKeyIamBinding:KmsAsymmetricEncryptionKeyIamBinding": "KmsAsymmetricEncryptionKeyIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsAsymmetricEncryptionKeyIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsAsymmetricEncryptionKeyIamMember:KmsAsymmetricEncryptionKeyIamMember": "KmsAsymmetricEncryptionKeyIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsAsymmetricSignatureKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsAsymmetricSignatureKey:KmsAsymmetricSignatureKey": "KmsAsymmetricSignatureKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsAsymmetricSignatureKeyIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsAsymmetricSignatureKeyIamBinding:KmsAsymmetricSignatureKeyIamBinding": "KmsAsymmetricSignatureKeyIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsAsymmetricSignatureKeyIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsAsymmetricSignatureKeyIamMember:KmsAsymmetricSignatureKeyIamMember": "KmsAsymmetricSignatureKeyIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsSecretCiphertext",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsSecretCiphertext:KmsSecretCiphertext": "KmsSecretCiphertext"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsSymmetricKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsSymmetricKey:KmsSymmetricKey": "KmsSymmetricKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsSymmetricKeyIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsSymmetricKeyIamBinding:KmsSymmetricKeyIamBinding": "KmsSymmetricKeyIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kmsSymmetricKeyIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kmsSymmetricKeyIamMember:KmsSymmetricKeyIamMember": "KmsSymmetricKeyIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kubernetesCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kubernetesCluster:KubernetesCluster": "KubernetesCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/kubernetesNodeGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/kubernetesNodeGroup:KubernetesNodeGroup": "KubernetesNodeGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lbNetworkLoadBalancer",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lbNetworkLoadBalancer:LbNetworkLoadBalancer": "LbNetworkLoadBalancer"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lbTargetGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lbTargetGroup:LbTargetGroup": "LbTargetGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/loadtestingAgent",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/loadtestingAgent:LoadtestingAgent": "LoadtestingAgent"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lockboxSecret",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lockboxSecret:LockboxSecret": "LockboxSecret"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lockboxSecretIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lockboxSecretIamBinding:LockboxSecretIamBinding": "LockboxSecretIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lockboxSecretIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lockboxSecretIamMember:LockboxSecretIamMember": "LockboxSecretIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lockboxSecretVersion",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lockboxSecretVersion:LockboxSecretVersion": "LockboxSecretVersion"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/lockboxSecretVersionHashed",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/lockboxSecretVersionHashed:LockboxSecretVersionHashed": "LockboxSecretVersionHashed"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/loggingGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/loggingGroup:LoggingGroup": "LoggingGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbClickhouseCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbClickhouseCluster:MdbClickhouseCluster": "MdbClickhouseCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbElasticsearchCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbElasticsearchCluster:MdbElasticsearchCluster": "MdbElasticsearchCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbGreenplumCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbGreenplumCluster:MdbGreenplumCluster": "MdbGreenplumCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbKafkaCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbKafkaCluster:MdbKafkaCluster": "MdbKafkaCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbKafkaConnector",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbKafkaConnector:MdbKafkaConnector": "MdbKafkaConnector"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbKafkaTopic",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbKafkaTopic:MdbKafkaTopic": "MdbKafkaTopic"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbKafkaUser",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbKafkaUser:MdbKafkaUser": "MdbKafkaUser"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbMongodbCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbMongodbCluster:MdbMongodbCluster": "MdbMongodbCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbMysqlCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbMysqlCluster:MdbMysqlCluster": "MdbMysqlCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbMysqlDatabase",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbMysqlDatabase:MdbMysqlDatabase": "MdbMysqlDatabase"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbMysqlUser",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbMysqlUser:MdbMysqlUser": "MdbMysqlUser"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbPostgresqlCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbPostgresqlCluster:MdbPostgresqlCluster": "MdbPostgresqlCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbPostgresqlDatabase",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbPostgresqlDatabase:MdbPostgresqlDatabase": "MdbPostgresqlDatabase"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbPostgresqlUser",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbPostgresqlUser:MdbPostgresqlUser": "MdbPostgresqlUser"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbRedisCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbRedisCluster:MdbRedisCluster": "MdbRedisCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/mdbSqlserverCluster",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/mdbSqlserverCluster:MdbSqlserverCluster": "MdbSqlserverCluster"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/messageQueue",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/messageQueue:MessageQueue": "MessageQueue"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/monitoringDashboard",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/monitoringDashboard:MonitoringDashboard": "MonitoringDashboard"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerGroup:OrganizationmanagerGroup": "OrganizationmanagerGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerGroupIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerGroupIamMember:OrganizationmanagerGroupIamMember": "OrganizationmanagerGroupIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerGroupMapping",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerGroupMapping:OrganizationmanagerGroupMapping": "OrganizationmanagerGroupMapping"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerGroupMappingItem",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerGroupMappingItem:OrganizationmanagerGroupMappingItem": "OrganizationmanagerGroupMappingItem"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerGroupMembership",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerGroupMembership:OrganizationmanagerGroupMembership": "OrganizationmanagerGroupMembership"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerOrganizationIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerOrganizationIamBinding:OrganizationmanagerOrganizationIamBinding": "OrganizationmanagerOrganizationIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerOrganizationIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerOrganizationIamMember:OrganizationmanagerOrganizationIamMember": "OrganizationmanagerOrganizationIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerOsLoginSettings",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerOsLoginSettings:OrganizationmanagerOsLoginSettings": "OrganizationmanagerOsLoginSettings"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerSamlFederation",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerSamlFederation:OrganizationmanagerSamlFederation": "OrganizationmanagerSamlFederation"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerSamlFederationUserAccount",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerSamlFederationUserAccount:OrganizationmanagerSamlFederationUserAccount": "OrganizationmanagerSamlFederationUserAccount"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/organizationmanagerUserSshKey",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/organizationmanagerUserSshKey:OrganizationmanagerUserSshKey": "OrganizationmanagerUserSshKey"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerCloud",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerCloud:ResourcemanagerCloud": "ResourcemanagerCloud"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerCloudIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerCloudIamBinding:ResourcemanagerCloudIamBinding": "ResourcemanagerCloudIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerCloudIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerCloudIamMember:ResourcemanagerCloudIamMember": "ResourcemanagerCloudIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolder",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolder:ResourcemanagerFolder": "ResourcemanagerFolder"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolderIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolderIamBinding:ResourcemanagerFolderIamBinding": "ResourcemanagerFolderIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolderIamMember",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolderIamMember:ResourcemanagerFolderIamMember": "ResourcemanagerFolderIamMember"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/resourcemanagerFolderIamPolicy",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/resourcemanagerFolderIamPolicy:ResourcemanagerFolderIamPolicy": "ResourcemanagerFolderIamPolicy"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/serverlessContainer",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/serverlessContainer:ServerlessContainer": "ServerlessContainer"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/serverlessContainerIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/serverlessContainerIamBinding:ServerlessContainerIamBinding": "ServerlessContainerIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/serverlessEventrouterBus",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/serverlessEventrouterBus:ServerlessEventrouterBus": "ServerlessEventrouterBus"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/serverlessEventrouterConnector",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/serverlessEventrouterConnector:ServerlessEventrouterConnector": "ServerlessEventrouterConnector"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/serverlessEventrouterRule",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/serverlessEventrouterRule:ServerlessEventrouterRule": "ServerlessEventrouterRule"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/smartcaptchaCaptcha",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/smartcaptchaCaptcha:SmartcaptchaCaptcha": "SmartcaptchaCaptcha"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/storageBucket",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/storageBucket:StorageBucket": "StorageBucket"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/storageObject",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/storageObject:StorageObject": "StorageObject"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/swsAdvancedRateLimiterProfile",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/swsAdvancedRateLimiterProfile:SwsAdvancedRateLimiterProfile": "SwsAdvancedRateLimiterProfile"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/swsSecurityProfile",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/swsSecurityProfile:SwsSecurityProfile": "SwsSecurityProfile"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/swsWafProfile",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/swsWafProfile:SwsWafProfile": "SwsWafProfile"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcAddress",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcAddress:VpcAddress": "VpcAddress"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcDefaultSecurityGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcDefaultSecurityGroup:VpcDefaultSecurityGroup": "VpcDefaultSecurityGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcGateway",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcGateway:VpcGateway": "VpcGateway"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcNetwork",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcNetwork:VpcNetwork": "VpcNetwork"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcPrivateEndpoint",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcPrivateEndpoint:VpcPrivateEndpoint": "VpcPrivateEndpoint"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcRouteTable",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcRouteTable:VpcRouteTable": "VpcRouteTable"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcSecurityGroup",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcSecurityGroup:VpcSecurityGroup": "VpcSecurityGroup"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/vpcSubnet",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/vpcSubnet:VpcSubnet": "VpcSubnet"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbDatabaseDedicated",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbDatabaseDedicated:YdbDatabaseDedicated": "YdbDatabaseDedicated"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbDatabaseIamBinding",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbDatabaseIamBinding:YdbDatabaseIamBinding": "YdbDatabaseIamBinding"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbDatabaseServerless",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbDatabaseServerless:YdbDatabaseServerless": "YdbDatabaseServerless"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbTable",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbTable:YdbTable": "YdbTable"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbTableChangefeed",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbTableChangefeed:YdbTableChangefeed": "YdbTableChangefeed"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbTableIndex",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbTableIndex:YdbTableIndex": "YdbTableIndex"
  }
 },
 {
  "pkg": "yandex",
  "mod": "index/ydbTopic",
  "fqn": "pulumi_yandex",
  "classes": {
   "yandex:index/ydbTopic:YdbTopic": "YdbTopic"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "yandex",
  "token": "pulumi:providers:yandex",
  "fqn": "pulumi_yandex",
  "class": "Provider"
 }
]
"""
)
