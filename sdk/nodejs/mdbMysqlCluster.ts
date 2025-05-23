// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class MdbMysqlCluster extends pulumi.CustomResource {
    /**
     * Get an existing MdbMysqlCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MdbMysqlClusterState, opts?: pulumi.CustomResourceOptions): MdbMysqlCluster {
        return new MdbMysqlCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/mdbMysqlCluster:MdbMysqlCluster';

    /**
     * Returns true if the given object is an instance of MdbMysqlCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MdbMysqlCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MdbMysqlCluster.__pulumiType;
    }

    /**
     * Access policy to the MySQL cluster.
     */
    public readonly access!: pulumi.Output<outputs.MdbMysqlClusterAccess>;
    /**
     * @deprecated You can safely remove this option. There is no need to recreate host if assignPublicIp is changed.
     */
    public readonly allowRegenerationHost!: pulumi.Output<boolean | undefined>;
    /**
     * The period in days during which backups are stored.
     */
    public readonly backupRetainPeriodDays!: pulumi.Output<number>;
    /**
     * Time to start the daily backup, in the UTC.
     */
    public readonly backupWindowStart!: pulumi.Output<outputs.MdbMysqlClusterBackupWindowStart>;
    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * To manage databases, please switch to using a separate resource type `yandexMdbMysqlDatabases`.
     *
     * @deprecated to manage databases, please switch to using a separate resource type yandex_mdb_mysql_database
     */
    public readonly databases!: pulumi.Output<outputs.MdbMysqlClusterDatabase[] | undefined>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Deployment environment of the MySQL cluster.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Aggregated health of the cluster.
     */
    public /*out*/ readonly health!: pulumi.Output<string>;
    public readonly hostGroupIds!: pulumi.Output<string[]>;
    /**
     * A host of the MySQL cluster.
     */
    public readonly hosts!: pulumi.Output<outputs.MdbMysqlClusterHost[]>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Maintenance policy of the MySQL cluster.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.MdbMysqlClusterMaintenanceWindow>;
    /**
     * MySQL cluster config block.
     */
    public readonly mysqlConfig!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The `VPC Network ID` of subnets which resource attached to.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Cluster performance diagnostics settings. [YC
     * Documentation](https://yandex.cloud/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics).
     */
    public readonly performanceDiagnostics!: pulumi.Output<outputs.MdbMysqlClusterPerformanceDiagnostics>;
    /**
     * Resources allocated to hosts of the MySQL cluster.
     */
    public readonly resources!: pulumi.Output<outputs.MdbMysqlClusterResources>;
    /**
     * The cluster will be created from the specified backup.
     */
    public readonly restore!: pulumi.Output<outputs.MdbMysqlClusterRestore | undefined>;
    /**
     * The list of security groups applied to resource or their components.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Status of the cluster.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * To manage users, please switch to using a separate resource type `yandex.MdbMysqlUser`.
     *
     * @deprecated to manage users, please switch to using a separate resource type yandex_mdb_mysql_user
     */
    public readonly users!: pulumi.Output<outputs.MdbMysqlClusterUser[] | undefined>;
    /**
     * Version of the MySQL cluster. (allowed versions are: 5.7, 8.0).
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a MdbMysqlCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MdbMysqlClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MdbMysqlClusterArgs | MdbMysqlClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MdbMysqlClusterState | undefined;
            resourceInputs["access"] = state ? state.access : undefined;
            resourceInputs["allowRegenerationHost"] = state ? state.allowRegenerationHost : undefined;
            resourceInputs["backupRetainPeriodDays"] = state ? state.backupRetainPeriodDays : undefined;
            resourceInputs["backupWindowStart"] = state ? state.backupWindowStart : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["databases"] = state ? state.databases : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["health"] = state ? state.health : undefined;
            resourceInputs["hostGroupIds"] = state ? state.hostGroupIds : undefined;
            resourceInputs["hosts"] = state ? state.hosts : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            resourceInputs["mysqlConfig"] = state ? state.mysqlConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["performanceDiagnostics"] = state ? state.performanceDiagnostics : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["restore"] = state ? state.restore : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as MdbMysqlClusterArgs | undefined;
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.hosts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hosts'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.resources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resources'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["access"] = args ? args.access : undefined;
            resourceInputs["allowRegenerationHost"] = args ? args.allowRegenerationHost : undefined;
            resourceInputs["backupRetainPeriodDays"] = args ? args.backupRetainPeriodDays : undefined;
            resourceInputs["backupWindowStart"] = args ? args.backupWindowStart : undefined;
            resourceInputs["databases"] = args ? args.databases : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["hostGroupIds"] = args ? args.hostGroupIds : undefined;
            resourceInputs["hosts"] = args ? args.hosts : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["mysqlConfig"] = args ? args.mysqlConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["performanceDiagnostics"] = args ? args.performanceDiagnostics : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["restore"] = args ? args.restore : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MdbMysqlCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MdbMysqlCluster resources.
 */
export interface MdbMysqlClusterState {
    /**
     * Access policy to the MySQL cluster.
     */
    access?: pulumi.Input<inputs.MdbMysqlClusterAccess>;
    /**
     * @deprecated You can safely remove this option. There is no need to recreate host if assignPublicIp is changed.
     */
    allowRegenerationHost?: pulumi.Input<boolean>;
    /**
     * The period in days during which backups are stored.
     */
    backupRetainPeriodDays?: pulumi.Input<number>;
    /**
     * Time to start the daily backup, in the UTC.
     */
    backupWindowStart?: pulumi.Input<inputs.MdbMysqlClusterBackupWindowStart>;
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * To manage databases, please switch to using a separate resource type `yandexMdbMysqlDatabases`.
     *
     * @deprecated to manage databases, please switch to using a separate resource type yandex_mdb_mysql_database
     */
    databases?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterDatabase>[]>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Deployment environment of the MySQL cluster.
     */
    environment?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Aggregated health of the cluster.
     */
    health?: pulumi.Input<string>;
    hostGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A host of the MySQL cluster.
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterHost>[]>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Maintenance policy of the MySQL cluster.
     */
    maintenanceWindow?: pulumi.Input<inputs.MdbMysqlClusterMaintenanceWindow>;
    /**
     * MySQL cluster config block.
     */
    mysqlConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The `VPC Network ID` of subnets which resource attached to.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Cluster performance diagnostics settings. [YC
     * Documentation](https://yandex.cloud/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics).
     */
    performanceDiagnostics?: pulumi.Input<inputs.MdbMysqlClusterPerformanceDiagnostics>;
    /**
     * Resources allocated to hosts of the MySQL cluster.
     */
    resources?: pulumi.Input<inputs.MdbMysqlClusterResources>;
    /**
     * The cluster will be created from the specified backup.
     */
    restore?: pulumi.Input<inputs.MdbMysqlClusterRestore>;
    /**
     * The list of security groups applied to resource or their components.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Status of the cluster.
     */
    status?: pulumi.Input<string>;
    /**
     * To manage users, please switch to using a separate resource type `yandex.MdbMysqlUser`.
     *
     * @deprecated to manage users, please switch to using a separate resource type yandex_mdb_mysql_user
     */
    users?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterUser>[]>;
    /**
     * Version of the MySQL cluster. (allowed versions are: 5.7, 8.0).
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MdbMysqlCluster resource.
 */
export interface MdbMysqlClusterArgs {
    /**
     * Access policy to the MySQL cluster.
     */
    access?: pulumi.Input<inputs.MdbMysqlClusterAccess>;
    /**
     * @deprecated You can safely remove this option. There is no need to recreate host if assignPublicIp is changed.
     */
    allowRegenerationHost?: pulumi.Input<boolean>;
    /**
     * The period in days during which backups are stored.
     */
    backupRetainPeriodDays?: pulumi.Input<number>;
    /**
     * Time to start the daily backup, in the UTC.
     */
    backupWindowStart?: pulumi.Input<inputs.MdbMysqlClusterBackupWindowStart>;
    /**
     * To manage databases, please switch to using a separate resource type `yandexMdbMysqlDatabases`.
     *
     * @deprecated to manage databases, please switch to using a separate resource type yandex_mdb_mysql_database
     */
    databases?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterDatabase>[]>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Deployment environment of the MySQL cluster.
     */
    environment: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    hostGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A host of the MySQL cluster.
     */
    hosts: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterHost>[]>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Maintenance policy of the MySQL cluster.
     */
    maintenanceWindow?: pulumi.Input<inputs.MdbMysqlClusterMaintenanceWindow>;
    /**
     * MySQL cluster config block.
     */
    mysqlConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * The `VPC Network ID` of subnets which resource attached to.
     */
    networkId: pulumi.Input<string>;
    /**
     * Cluster performance diagnostics settings. [YC
     * Documentation](https://yandex.cloud/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics).
     */
    performanceDiagnostics?: pulumi.Input<inputs.MdbMysqlClusterPerformanceDiagnostics>;
    /**
     * Resources allocated to hosts of the MySQL cluster.
     */
    resources: pulumi.Input<inputs.MdbMysqlClusterResources>;
    /**
     * The cluster will be created from the specified backup.
     */
    restore?: pulumi.Input<inputs.MdbMysqlClusterRestore>;
    /**
     * The list of security groups applied to resource or their components.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * To manage users, please switch to using a separate resource type `yandex.MdbMysqlUser`.
     *
     * @deprecated to manage users, please switch to using a separate resource type yandex_mdb_mysql_user
     */
    users?: pulumi.Input<pulumi.Input<inputs.MdbMysqlClusterUser>[]>;
    /**
     * Version of the MySQL cluster. (allowed versions are: 5.7, 8.0).
     */
    version: pulumi.Input<string>;
}
