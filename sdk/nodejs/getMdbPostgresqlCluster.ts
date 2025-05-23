// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getMdbPostgresqlCluster(args?: GetMdbPostgresqlClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetMdbPostgresqlClusterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("yandex:index/getMdbPostgresqlCluster:getMdbPostgresqlCluster", {
        "clusterId": args.clusterId,
        "deletionProtection": args.deletionProtection,
        "description": args.description,
        "folderId": args.folderId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbPostgresqlCluster.
 */
export interface GetMdbPostgresqlClusterArgs {
    clusterId?: string;
    deletionProtection?: boolean;
    description?: string;
    folderId?: string;
    name?: string;
}

/**
 * A collection of values returned by getMdbPostgresqlCluster.
 */
export interface GetMdbPostgresqlClusterResult {
    readonly clusterId: string;
    readonly configs: outputs.GetMdbPostgresqlClusterConfig[];
    readonly createdAt: string;
    readonly databases: outputs.GetMdbPostgresqlClusterDatabase[];
    readonly deletionProtection: boolean;
    readonly description?: string;
    readonly environment: string;
    readonly folderId: string;
    readonly health: string;
    readonly hostGroupIds: string[];
    readonly hosts: outputs.GetMdbPostgresqlClusterHost[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly maintenanceWindows: outputs.GetMdbPostgresqlClusterMaintenanceWindow[];
    readonly name: string;
    readonly networkId: string;
    readonly securityGroupIds: string[];
    readonly status: string;
    readonly users: outputs.GetMdbPostgresqlClusterUser[];
}
export function getMdbPostgresqlClusterOutput(args?: GetMdbPostgresqlClusterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMdbPostgresqlClusterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("yandex:index/getMdbPostgresqlCluster:getMdbPostgresqlCluster", {
        "clusterId": args.clusterId,
        "deletionProtection": args.deletionProtection,
        "description": args.description,
        "folderId": args.folderId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getMdbPostgresqlCluster.
 */
export interface GetMdbPostgresqlClusterOutputArgs {
    clusterId?: pulumi.Input<string>;
    deletionProtection?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    folderId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
