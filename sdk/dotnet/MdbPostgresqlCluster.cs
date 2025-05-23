// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/mdbPostgresqlCluster:MdbPostgresqlCluster")]
    public partial class MdbPostgresqlCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration of the PostgreSQL cluster.
        /// </summary>
        [Output("config")]
        public Output<Outputs.MdbPostgresqlClusterConfig> Config { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// &gt; Deprecated! To manage databases, please switch to using a separate resource type `yandex.MdbPostgresqlDatabase`.
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<Outputs.MdbPostgresqlClusterDatabase>> Databases { get; private set; } = null!;

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
        /// Deployment environment of the PostgreSQL cluster.
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

        /// <summary>
        /// Host Group IDs.
        /// </summary>
        [Output("hostGroupIds")]
        public Output<ImmutableArray<string>> HostGroupIds { get; private set; } = null!;

        [Output("hostMasterName")]
        public Output<string> HostMasterName { get; private set; } = null!;

        /// <summary>
        /// A host of the PostgreSQL cluster.
        /// </summary>
        [Output("hosts")]
        public Output<ImmutableArray<Outputs.MdbPostgresqlClusterHost>> Hosts { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Maintenance policy of the PostgreSQL cluster.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.MdbPostgresqlClusterMaintenanceWindow> MaintenanceWindow { get; private set; } = null!;

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
        /// The cluster will be created from the specified backup.
        /// </summary>
        [Output("restore")]
        public Output<Outputs.MdbPostgresqlClusterRestore?> Restore { get; private set; } = null!;

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
        /// &gt; Deprecated! To manage users, please switch to using a separate resource type `yandex.MdbPostgresqlUser`.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.MdbPostgresqlClusterUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a MdbPostgresqlCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MdbPostgresqlCluster(string name, MdbPostgresqlClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/mdbPostgresqlCluster:MdbPostgresqlCluster", name, args ?? new MdbPostgresqlClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MdbPostgresqlCluster(string name, Input<string> id, MdbPostgresqlClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/mdbPostgresqlCluster:MdbPostgresqlCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MdbPostgresqlCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MdbPostgresqlCluster Get(string name, Input<string> id, MdbPostgresqlClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new MdbPostgresqlCluster(name, id, state, options);
        }
    }

    public sealed class MdbPostgresqlClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the PostgreSQL cluster.
        /// </summary>
        [Input("config", required: true)]
        public Input<Inputs.MdbPostgresqlClusterConfigArgs> Config { get; set; } = null!;

        [Input("databases")]
        private InputList<Inputs.MdbPostgresqlClusterDatabaseArgs>? _databases;

        /// <summary>
        /// &gt; Deprecated! To manage databases, please switch to using a separate resource type `yandex.MdbPostgresqlDatabase`.
        /// </summary>
        [Obsolete(@"to manage databases, please switch to using a separate resource type yandex_mdb_postgresql_database")]
        public InputList<Inputs.MdbPostgresqlClusterDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.MdbPostgresqlClusterDatabaseArgs>());
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
        /// Deployment environment of the PostgreSQL cluster.
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

        /// <summary>
        /// Host Group IDs.
        /// </summary>
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("hostMasterName")]
        public Input<string>? HostMasterName { get; set; }

        [Input("hosts", required: true)]
        private InputList<Inputs.MdbPostgresqlClusterHostArgs>? _hosts;

        /// <summary>
        /// A host of the PostgreSQL cluster.
        /// </summary>
        public InputList<Inputs.MdbPostgresqlClusterHostArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.MdbPostgresqlClusterHostArgs>());
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
        /// Maintenance policy of the PostgreSQL cluster.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbPostgresqlClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

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
        /// The cluster will be created from the specified backup.
        /// </summary>
        [Input("restore")]
        public Input<Inputs.MdbPostgresqlClusterRestoreArgs>? Restore { get; set; }

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
        private InputList<Inputs.MdbPostgresqlClusterUserArgs>? _users;

        /// <summary>
        /// &gt; Deprecated! To manage users, please switch to using a separate resource type `yandex.MdbPostgresqlUser`.
        /// </summary>
        [Obsolete(@"to manage users, please switch to using a separate resource type yandex_mdb_postgresql_user")]
        public InputList<Inputs.MdbPostgresqlClusterUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.MdbPostgresqlClusterUserArgs>());
            set => _users = value;
        }

        public MdbPostgresqlClusterArgs()
        {
        }
        public static new MdbPostgresqlClusterArgs Empty => new MdbPostgresqlClusterArgs();
    }

    public sealed class MdbPostgresqlClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the PostgreSQL cluster.
        /// </summary>
        [Input("config")]
        public Input<Inputs.MdbPostgresqlClusterConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("databases")]
        private InputList<Inputs.MdbPostgresqlClusterDatabaseGetArgs>? _databases;

        /// <summary>
        /// &gt; Deprecated! To manage databases, please switch to using a separate resource type `yandex.MdbPostgresqlDatabase`.
        /// </summary>
        [Obsolete(@"to manage databases, please switch to using a separate resource type yandex_mdb_postgresql_database")]
        public InputList<Inputs.MdbPostgresqlClusterDatabaseGetArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.MdbPostgresqlClusterDatabaseGetArgs>());
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
        /// Deployment environment of the PostgreSQL cluster.
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

        /// <summary>
        /// Host Group IDs.
        /// </summary>
        public InputList<string> HostGroupIds
        {
            get => _hostGroupIds ?? (_hostGroupIds = new InputList<string>());
            set => _hostGroupIds = value;
        }

        [Input("hostMasterName")]
        public Input<string>? HostMasterName { get; set; }

        [Input("hosts")]
        private InputList<Inputs.MdbPostgresqlClusterHostGetArgs>? _hosts;

        /// <summary>
        /// A host of the PostgreSQL cluster.
        /// </summary>
        public InputList<Inputs.MdbPostgresqlClusterHostGetArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.MdbPostgresqlClusterHostGetArgs>());
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
        /// Maintenance policy of the PostgreSQL cluster.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbPostgresqlClusterMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

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
        /// The cluster will be created from the specified backup.
        /// </summary>
        [Input("restore")]
        public Input<Inputs.MdbPostgresqlClusterRestoreGetArgs>? Restore { get; set; }

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
        private InputList<Inputs.MdbPostgresqlClusterUserGetArgs>? _users;

        /// <summary>
        /// &gt; Deprecated! To manage users, please switch to using a separate resource type `yandex.MdbPostgresqlUser`.
        /// </summary>
        [Obsolete(@"to manage users, please switch to using a separate resource type yandex_mdb_postgresql_user")]
        public InputList<Inputs.MdbPostgresqlClusterUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.MdbPostgresqlClusterUserGetArgs>());
            set => _users = value;
        }

        public MdbPostgresqlClusterState()
        {
        }
        public static new MdbPostgresqlClusterState Empty => new MdbPostgresqlClusterState();
    }
}
