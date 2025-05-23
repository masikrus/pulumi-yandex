// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/mdbMysqlCluster:MdbMysqlCluster")]
    public partial class MdbMysqlCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access policy to the MySQL cluster.
        /// </summary>
        [Output("access")]
        public Output<Outputs.MdbMysqlClusterAccess> Access { get; private set; } = null!;

        [Output("allowRegenerationHost")]
        public Output<bool?> AllowRegenerationHost { get; private set; } = null!;

        /// <summary>
        /// The period in days during which backups are stored.
        /// </summary>
        [Output("backupRetainPeriodDays")]
        public Output<int> BackupRetainPeriodDays { get; private set; } = null!;

        /// <summary>
        /// Time to start the daily backup, in the UTC.
        /// </summary>
        [Output("backupWindowStart")]
        public Output<Outputs.MdbMysqlClusterBackupWindowStart> BackupWindowStart { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// To manage databases, please switch to using a separate resource type `yandex_mdb_mysql_databases`.
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<Outputs.MdbMysqlClusterDatabase>> Databases { get; private set; } = null!;

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// The resource description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Deployment environment of the MySQL cluster.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Aggregated health of the cluster.
        /// </summary>
        [Output("health")]
        public Output<string> Health { get; private set; } = null!;

        [Output("hostGroupIds")]
        public Output<ImmutableArray<string>> HostGroupIds { get; private set; } = null!;

        /// <summary>
        /// A host of the MySQL cluster.
        /// </summary>
        [Output("hosts")]
        public Output<ImmutableArray<Outputs.MdbMysqlClusterHost>> Hosts { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Maintenance policy of the MySQL cluster.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.MdbMysqlClusterMaintenanceWindow> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// MySQL cluster config block.
        /// </summary>
        [Output("mysqlConfig")]
        public Output<ImmutableDictionary<string, string>> MysqlConfig { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The `VPC Network ID` of subnets which resource attached to.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Cluster performance diagnostics settings. [YC
        /// Documentation](https://yandex.cloud/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics).
        /// </summary>
        [Output("performanceDiagnostics")]
        public Output<Outputs.MdbMysqlClusterPerformanceDiagnostics> PerformanceDiagnostics { get; private set; } = null!;

        /// <summary>
        /// Resources allocated to hosts of the MySQL cluster.
        /// </summary>
        [Output("resources")]
        public Output<Outputs.MdbMysqlClusterResources> Resources { get; private set; } = null!;

        /// <summary>
        /// The cluster will be created from the specified backup.
        /// </summary>
        [Output("restore")]
        public Output<Outputs.MdbMysqlClusterRestore?> Restore { get; private set; } = null!;

        /// <summary>
        /// The list of security groups applied to resource or their components.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Status of the cluster.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// To manage users, please switch to using a separate resource type `yandex.MdbMysqlUser`.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.MdbMysqlClusterUser>> Users { get; private set; } = null!;

        /// <summary>
        /// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0).
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a MdbMysqlCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MdbMysqlCluster(string name, MdbMysqlClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/mdbMysqlCluster:MdbMysqlCluster", name, args ?? new MdbMysqlClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MdbMysqlCluster(string name, Input<string> id, MdbMysqlClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/mdbMysqlCluster:MdbMysqlCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MdbMysqlCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MdbMysqlCluster Get(string name, Input<string> id, MdbMysqlClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new MdbMysqlCluster(name, id, state, options);
        }
    }

    public sealed class MdbMysqlClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy to the MySQL cluster.
        /// </summary>
        [Input("access")]
        public Input<Inputs.MdbMysqlClusterAccessArgs>? Access { get; set; }

        [Input("allowRegenerationHost")]
        public Input<bool>? AllowRegenerationHost { get; set; }

        /// <summary>
        /// The period in days during which backups are stored.
        /// </summary>
        [Input("backupRetainPeriodDays")]
        public Input<int>? BackupRetainPeriodDays { get; set; }

        /// <summary>
        /// Time to start the daily backup, in the UTC.
        /// </summary>
        [Input("backupWindowStart")]
        public Input<Inputs.MdbMysqlClusterBackupWindowStartArgs>? BackupWindowStart { get; set; }

        [Input("databases")]
        private InputList<Inputs.MdbMysqlClusterDatabaseArgs>? _databases;

        /// <summary>
        /// To manage databases, please switch to using a separate resource type `yandex_mdb_mysql_databases`.
        /// </summary>
        [Obsolete(@"to manage databases, please switch to using a separate resource type yandex_mdb_mysql_database")]
        public InputList<Inputs.MdbMysqlClusterDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.MdbMysqlClusterDatabaseArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The resource description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deployment environment of the MySQL cluster.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("hostGroupIds")]
        private InputList<string>? _hostGroupIds;
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("hosts", required: true)]
        private InputList<Inputs.MdbMysqlClusterHostArgs>? _hosts;

        /// <summary>
        /// A host of the MySQL cluster.
        /// </summary>
        public InputList<Inputs.MdbMysqlClusterHostArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.MdbMysqlClusterHostArgs>());
            set => _hosts = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maintenance policy of the MySQL cluster.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbMysqlClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        [Input("mysqlConfig")]
        private InputMap<string>? _mysqlConfig;

        /// <summary>
        /// MySQL cluster config block.
        /// </summary>
        public InputMap<string> MysqlConfig
        {
            get => _mysqlConfig ?? (_mysqlConfig = new InputMap<string>());
            set => _mysqlConfig = value;
        }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The `VPC Network ID` of subnets which resource attached to.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Cluster performance diagnostics settings. [YC
        /// Documentation](https://yandex.cloud/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics).
        /// </summary>
        [Input("performanceDiagnostics")]
        public Input<Inputs.MdbMysqlClusterPerformanceDiagnosticsArgs>? PerformanceDiagnostics { get; set; }

        /// <summary>
        /// Resources allocated to hosts of the MySQL cluster.
        /// </summary>
        [Input("resources", required: true)]
        public Input<Inputs.MdbMysqlClusterResourcesArgs> Resources { get; set; } = null!;

        /// <summary>
        /// The cluster will be created from the specified backup.
        /// </summary>
        [Input("restore")]
        public Input<Inputs.MdbMysqlClusterRestoreArgs>? Restore { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The list of security groups applied to resource or their components.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("users")]
        private InputList<Inputs.MdbMysqlClusterUserArgs>? _users;

        /// <summary>
        /// To manage users, please switch to using a separate resource type `yandex.MdbMysqlUser`.
        /// </summary>
        [Obsolete(@"to manage users, please switch to using a separate resource type yandex_mdb_mysql_user")]
        public InputList<Inputs.MdbMysqlClusterUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.MdbMysqlClusterUserArgs>());
            set => _users = value;
        }

        /// <summary>
        /// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0).
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public MdbMysqlClusterArgs()
        {
        }
        public static new MdbMysqlClusterArgs Empty => new MdbMysqlClusterArgs();
    }

    public sealed class MdbMysqlClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy to the MySQL cluster.
        /// </summary>
        [Input("access")]
        public Input<Inputs.MdbMysqlClusterAccessGetArgs>? Access { get; set; }

        [Input("allowRegenerationHost")]
        public Input<bool>? AllowRegenerationHost { get; set; }

        /// <summary>
        /// The period in days during which backups are stored.
        /// </summary>
        [Input("backupRetainPeriodDays")]
        public Input<int>? BackupRetainPeriodDays { get; set; }

        /// <summary>
        /// Time to start the daily backup, in the UTC.
        /// </summary>
        [Input("backupWindowStart")]
        public Input<Inputs.MdbMysqlClusterBackupWindowStartGetArgs>? BackupWindowStart { get; set; }

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("databases")]
        private InputList<Inputs.MdbMysqlClusterDatabaseGetArgs>? _databases;

        /// <summary>
        /// To manage databases, please switch to using a separate resource type `yandex_mdb_mysql_databases`.
        /// </summary>
        [Obsolete(@"to manage databases, please switch to using a separate resource type yandex_mdb_mysql_database")]
        public InputList<Inputs.MdbMysqlClusterDatabaseGetArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.MdbMysqlClusterDatabaseGetArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The resource description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deployment environment of the MySQL cluster.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Aggregated health of the cluster.
        /// </summary>
        [Input("health")]
        public Input<string>? Health { get; set; }

        [Input("hostGroupIds")]
        private InputList<string>? _hostGroupIds;
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("hosts")]
        private InputList<Inputs.MdbMysqlClusterHostGetArgs>? _hosts;

        /// <summary>
        /// A host of the MySQL cluster.
        /// </summary>
        public InputList<Inputs.MdbMysqlClusterHostGetArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.MdbMysqlClusterHostGetArgs>());
            set => _hosts = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maintenance policy of the MySQL cluster.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbMysqlClusterMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

        [Input("mysqlConfig")]
        private InputMap<string>? _mysqlConfig;

        /// <summary>
        /// MySQL cluster config block.
        /// </summary>
        public InputMap<string> MysqlConfig
        {
            get => _mysqlConfig ?? (_mysqlConfig = new InputMap<string>());
            set => _mysqlConfig = value;
        }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The `VPC Network ID` of subnets which resource attached to.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Cluster performance diagnostics settings. [YC
        /// Documentation](https://yandex.cloud/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics).
        /// </summary>
        [Input("performanceDiagnostics")]
        public Input<Inputs.MdbMysqlClusterPerformanceDiagnosticsGetArgs>? PerformanceDiagnostics { get; set; }

        /// <summary>
        /// Resources allocated to hosts of the MySQL cluster.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.MdbMysqlClusterResourcesGetArgs>? Resources { get; set; }

        /// <summary>
        /// The cluster will be created from the specified backup.
        /// </summary>
        [Input("restore")]
        public Input<Inputs.MdbMysqlClusterRestoreGetArgs>? Restore { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The list of security groups applied to resource or their components.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Status of the cluster.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("users")]
        private InputList<Inputs.MdbMysqlClusterUserGetArgs>? _users;

        /// <summary>
        /// To manage users, please switch to using a separate resource type `yandex.MdbMysqlUser`.
        /// </summary>
        [Obsolete(@"to manage users, please switch to using a separate resource type yandex_mdb_mysql_user")]
        public InputList<Inputs.MdbMysqlClusterUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.MdbMysqlClusterUserGetArgs>());
            set => _users = value;
        }

        /// <summary>
        /// Version of the MySQL cluster. (allowed versions are: 5.7, 8.0).
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public MdbMysqlClusterState()
        {
        }
        public static new MdbMysqlClusterState Empty => new MdbMysqlClusterState();
    }
}
