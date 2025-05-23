// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class MdbClickhouseCluster extends pulumi.CustomResource {
    /**
     * Get an existing MdbClickhouseCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MdbClickhouseClusterState, opts?: pulumi.CustomResourceOptions): MdbClickhouseCluster {
        return new MdbClickhouseCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/mdbClickhouseCluster:MdbClickhouseCluster';

    /**
     * Returns true if the given object is an instance of MdbClickhouseCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MdbClickhouseCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MdbClickhouseCluster.__pulumiType;
    }

    /**
     * Access policy to the ClickHouse cluster.
     */
    public readonly access!: pulumi.Output<outputs.MdbClickhouseClusterAccess>;
    /**
     * A password used to authorize as user `admin` when `sqlUserManagement` enabled.
     */
    public readonly adminPassword!: pulumi.Output<string | undefined>;
    /**
     * The period in days during which backups are stored.
     */
    public readonly backupRetainPeriodDays!: pulumi.Output<number | undefined>;
    /**
     * Time to start the daily backup, in the UTC timezone.
     */
    public readonly backupWindowStart!: pulumi.Output<outputs.MdbClickhouseClusterBackupWindowStart>;
    /**
     * Configuration of the ClickHouse subcluster.
     */
    public readonly clickhouse!: pulumi.Output<outputs.MdbClickhouseClusterClickhouse>;
    /**
     * Cloud Storage settings.
     */
    public readonly cloudStorage!: pulumi.Output<outputs.MdbClickhouseClusterCloudStorage>;
    /**
     * The cluster identifier.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Whether to copy schema on new ClickHouse hosts.
     */
    public readonly copySchemaOnNewHosts!: pulumi.Output<boolean | undefined>;
    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A database of the ClickHouse cluster.
     */
    public readonly databases!: pulumi.Output<outputs.MdbClickhouseClusterDatabase[] | undefined>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to use ClickHouse Keeper as a coordination system and place it on the same hosts with ClickHouse. If not, it's
     * used ZooKeeper with placement on separate hosts.
     */
    public readonly embeddedKeeper!: pulumi.Output<boolean>;
    /**
     * Deployment environment of the ClickHouse cluster. Can be either `PRESTABLE` or `PRODUCTION`.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of `protobuf` or `capnproto` format schemas.
     */
    public readonly formatSchemas!: pulumi.Output<outputs.MdbClickhouseClusterFormatSchema[] | undefined>;
    /**
     * Aggregated health of the cluster. Can be `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see
     * `health` field of JSON representation in [the official
     * documentation](https://yandex.cloud/docs/managed-clickhouse/api-ref/Cluster/).
     */
    public /*out*/ readonly health!: pulumi.Output<string>;
    /**
     * A host of the ClickHouse cluster.
     */
    public readonly hosts!: pulumi.Output<outputs.MdbClickhouseClusterHost[]>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly maintenanceWindow!: pulumi.Output<outputs.MdbClickhouseClusterMaintenanceWindow>;
    /**
     * A group of machine learning models.
     */
    public readonly mlModels!: pulumi.Output<outputs.MdbClickhouseClusterMlModel[] | undefined>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The `VPC Network ID` of subnets which resource attached to.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The list of security groups applied to resource or their components.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    public readonly serviceAccountId!: pulumi.Output<string | undefined>;
    /**
     * A group of clickhouse shards.
     */
    public readonly shardGroups!: pulumi.Output<outputs.MdbClickhouseClusterShardGroup[] | undefined>;
    /**
     * A shard of the ClickHouse cluster.
     */
    public readonly shards!: pulumi.Output<outputs.MdbClickhouseClusterShard[]>;
    /**
     * Grants `admin` user database management permission.
     */
    public readonly sqlDatabaseManagement!: pulumi.Output<boolean>;
    /**
     * Enables `admin` user with user management permission.
     */
    public readonly sqlUserManagement!: pulumi.Output<boolean>;
    /**
     * Status of the cluster. Can be `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or
     * `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official
     * documentation](https://yandex.cloud/docs/managed-clickhouse/api-ref/Cluster/).
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A user of the ClickHouse cluster.
     */
    public readonly users!: pulumi.Output<outputs.MdbClickhouseClusterUser[] | undefined>;
    /**
     * Version of the ClickHouse server software.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Configuration of the ZooKeeper subcluster.
     */
    public readonly zookeeper!: pulumi.Output<outputs.MdbClickhouseClusterZookeeper>;

    /**
     * Create a MdbClickhouseCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MdbClickhouseClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MdbClickhouseClusterArgs | MdbClickhouseClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MdbClickhouseClusterState | undefined;
            resourceInputs["access"] = state ? state.access : undefined;
            resourceInputs["adminPassword"] = state ? state.adminPassword : undefined;
            resourceInputs["backupRetainPeriodDays"] = state ? state.backupRetainPeriodDays : undefined;
            resourceInputs["backupWindowStart"] = state ? state.backupWindowStart : undefined;
            resourceInputs["clickhouse"] = state ? state.clickhouse : undefined;
            resourceInputs["cloudStorage"] = state ? state.cloudStorage : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["copySchemaOnNewHosts"] = state ? state.copySchemaOnNewHosts : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["databases"] = state ? state.databases : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["embeddedKeeper"] = state ? state.embeddedKeeper : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["formatSchemas"] = state ? state.formatSchemas : undefined;
            resourceInputs["health"] = state ? state.health : undefined;
            resourceInputs["hosts"] = state ? state.hosts : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            resourceInputs["mlModels"] = state ? state.mlModels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["shardGroups"] = state ? state.shardGroups : undefined;
            resourceInputs["shards"] = state ? state.shards : undefined;
            resourceInputs["sqlDatabaseManagement"] = state ? state.sqlDatabaseManagement : undefined;
            resourceInputs["sqlUserManagement"] = state ? state.sqlUserManagement : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["zookeeper"] = state ? state.zookeeper : undefined;
        } else {
            const args = argsOrState as MdbClickhouseClusterArgs | undefined;
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.hosts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hosts'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            resourceInputs["access"] = args ? args.access : undefined;
            resourceInputs["adminPassword"] = args?.adminPassword ? pulumi.secret(args.adminPassword) : undefined;
            resourceInputs["backupRetainPeriodDays"] = args ? args.backupRetainPeriodDays : undefined;
            resourceInputs["backupWindowStart"] = args ? args.backupWindowStart : undefined;
            resourceInputs["clickhouse"] = args ? args.clickhouse : undefined;
            resourceInputs["cloudStorage"] = args ? args.cloudStorage : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["copySchemaOnNewHosts"] = args ? args.copySchemaOnNewHosts : undefined;
            resourceInputs["databases"] = args ? args.databases : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["embeddedKeeper"] = args ? args.embeddedKeeper : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["formatSchemas"] = args ? args.formatSchemas : undefined;
            resourceInputs["hosts"] = args ? args.hosts : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["mlModels"] = args ? args.mlModels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["shardGroups"] = args ? args.shardGroups : undefined;
            resourceInputs["shards"] = args ? args.shards : undefined;
            resourceInputs["sqlDatabaseManagement"] = args ? args.sqlDatabaseManagement : undefined;
            resourceInputs["sqlUserManagement"] = args ? args.sqlUserManagement : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["zookeeper"] = args ? args.zookeeper : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MdbClickhouseCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MdbClickhouseCluster resources.
 */
export interface MdbClickhouseClusterState {
    /**
     * Access policy to the ClickHouse cluster.
     */
    access?: pulumi.Input<inputs.MdbClickhouseClusterAccess>;
    /**
     * A password used to authorize as user `admin` when `sqlUserManagement` enabled.
     */
    adminPassword?: pulumi.Input<string>;
    /**
     * The period in days during which backups are stored.
     */
    backupRetainPeriodDays?: pulumi.Input<number>;
    /**
     * Time to start the daily backup, in the UTC timezone.
     */
    backupWindowStart?: pulumi.Input<inputs.MdbClickhouseClusterBackupWindowStart>;
    /**
     * Configuration of the ClickHouse subcluster.
     */
    clickhouse?: pulumi.Input<inputs.MdbClickhouseClusterClickhouse>;
    /**
     * Cloud Storage settings.
     */
    cloudStorage?: pulumi.Input<inputs.MdbClickhouseClusterCloudStorage>;
    /**
     * The cluster identifier.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Whether to copy schema on new ClickHouse hosts.
     */
    copySchemaOnNewHosts?: pulumi.Input<boolean>;
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A database of the ClickHouse cluster.
     */
    databases?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterDatabase>[]>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to use ClickHouse Keeper as a coordination system and place it on the same hosts with ClickHouse. If not, it's
     * used ZooKeeper with placement on separate hosts.
     */
    embeddedKeeper?: pulumi.Input<boolean>;
    /**
     * Deployment environment of the ClickHouse cluster. Can be either `PRESTABLE` or `PRODUCTION`.
     */
    environment?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of `protobuf` or `capnproto` format schemas.
     */
    formatSchemas?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterFormatSchema>[]>;
    /**
     * Aggregated health of the cluster. Can be `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information see
     * `health` field of JSON representation in [the official
     * documentation](https://yandex.cloud/docs/managed-clickhouse/api-ref/Cluster/).
     */
    health?: pulumi.Input<string>;
    /**
     * A host of the ClickHouse cluster.
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterHost>[]>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    maintenanceWindow?: pulumi.Input<inputs.MdbClickhouseClusterMaintenanceWindow>;
    /**
     * A group of machine learning models.
     */
    mlModels?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterMlModel>[]>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The `VPC Network ID` of subnets which resource attached to.
     */
    networkId?: pulumi.Input<string>;
    /**
     * The list of security groups applied to resource or their components.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * A group of clickhouse shards.
     */
    shardGroups?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterShardGroup>[]>;
    /**
     * A shard of the ClickHouse cluster.
     */
    shards?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterShard>[]>;
    /**
     * Grants `admin` user database management permission.
     */
    sqlDatabaseManagement?: pulumi.Input<boolean>;
    /**
     * Enables `admin` user with user management permission.
     */
    sqlUserManagement?: pulumi.Input<boolean>;
    /**
     * Status of the cluster. Can be `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or
     * `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official
     * documentation](https://yandex.cloud/docs/managed-clickhouse/api-ref/Cluster/).
     */
    status?: pulumi.Input<string>;
    /**
     * A user of the ClickHouse cluster.
     */
    users?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterUser>[]>;
    /**
     * Version of the ClickHouse server software.
     */
    version?: pulumi.Input<string>;
    /**
     * Configuration of the ZooKeeper subcluster.
     */
    zookeeper?: pulumi.Input<inputs.MdbClickhouseClusterZookeeper>;
}

/**
 * The set of arguments for constructing a MdbClickhouseCluster resource.
 */
export interface MdbClickhouseClusterArgs {
    /**
     * Access policy to the ClickHouse cluster.
     */
    access?: pulumi.Input<inputs.MdbClickhouseClusterAccess>;
    /**
     * A password used to authorize as user `admin` when `sqlUserManagement` enabled.
     */
    adminPassword?: pulumi.Input<string>;
    /**
     * The period in days during which backups are stored.
     */
    backupRetainPeriodDays?: pulumi.Input<number>;
    /**
     * Time to start the daily backup, in the UTC timezone.
     */
    backupWindowStart?: pulumi.Input<inputs.MdbClickhouseClusterBackupWindowStart>;
    /**
     * Configuration of the ClickHouse subcluster.
     */
    clickhouse?: pulumi.Input<inputs.MdbClickhouseClusterClickhouse>;
    /**
     * Cloud Storage settings.
     */
    cloudStorage?: pulumi.Input<inputs.MdbClickhouseClusterCloudStorage>;
    /**
     * The cluster identifier.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Whether to copy schema on new ClickHouse hosts.
     */
    copySchemaOnNewHosts?: pulumi.Input<boolean>;
    /**
     * A database of the ClickHouse cluster.
     */
    databases?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterDatabase>[]>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to use ClickHouse Keeper as a coordination system and place it on the same hosts with ClickHouse. If not, it's
     * used ZooKeeper with placement on separate hosts.
     */
    embeddedKeeper?: pulumi.Input<boolean>;
    /**
     * Deployment environment of the ClickHouse cluster. Can be either `PRESTABLE` or `PRODUCTION`.
     */
    environment: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of `protobuf` or `capnproto` format schemas.
     */
    formatSchemas?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterFormatSchema>[]>;
    /**
     * A host of the ClickHouse cluster.
     */
    hosts: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterHost>[]>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    maintenanceWindow?: pulumi.Input<inputs.MdbClickhouseClusterMaintenanceWindow>;
    /**
     * A group of machine learning models.
     */
    mlModels?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterMlModel>[]>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The `VPC Network ID` of subnets which resource attached to.
     */
    networkId: pulumi.Input<string>;
    /**
     * The list of security groups applied to resource or their components.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * A group of clickhouse shards.
     */
    shardGroups?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterShardGroup>[]>;
    /**
     * A shard of the ClickHouse cluster.
     */
    shards?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterShard>[]>;
    /**
     * Grants `admin` user database management permission.
     */
    sqlDatabaseManagement?: pulumi.Input<boolean>;
    /**
     * Enables `admin` user with user management permission.
     */
    sqlUserManagement?: pulumi.Input<boolean>;
    /**
     * A user of the ClickHouse cluster.
     */
    users?: pulumi.Input<pulumi.Input<inputs.MdbClickhouseClusterUser>[]>;
    /**
     * Version of the ClickHouse server software.
     */
    version?: pulumi.Input<string>;
    /**
     * Configuration of the ZooKeeper subcluster.
     */
    zookeeper?: pulumi.Input<inputs.MdbClickhouseClusterZookeeper>;
}
