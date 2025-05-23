// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getBackupPolicy(args?: GetBackupPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupPolicyResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("yandex:index/getBackupPolicy:getBackupPolicy", {
        "name": args.name,
        "policyId": args.policyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupPolicy.
 */
export interface GetBackupPolicyArgs {
    name?: string;
    policyId?: string;
}

/**
 * A collection of values returned by getBackupPolicy.
 */
export interface GetBackupPolicyResult {
    readonly archiveName: string;
    readonly cbt: string;
    readonly compression: string;
    readonly createdAt: string;
    readonly enabled: boolean;
    readonly fastBackupEnabled: boolean;
    readonly folderId: string;
    readonly format: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly multiVolumeSnapshottingEnabled: boolean;
    readonly name: string;
    readonly performanceWindowEnabled: boolean;
    readonly policyId: string;
    readonly preserveFileSecuritySettings: boolean;
    readonly quiesceSnapshottingEnabled: boolean;
    readonly reattempts: outputs.GetBackupPolicyReattempt[];
    readonly retentions: outputs.GetBackupPolicyRetention[];
    readonly schedulings: outputs.GetBackupPolicyScheduling[];
    readonly silentModeEnabled: boolean;
    readonly splittingBytes: string;
    readonly updatedAt: string;
    readonly vmSnapshotReattempts: outputs.GetBackupPolicyVmSnapshotReattempt[];
    readonly vssProvider: string;
}
export function getBackupPolicyOutput(args?: GetBackupPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBackupPolicyResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("yandex:index/getBackupPolicy:getBackupPolicy", {
        "name": args.name,
        "policyId": args.policyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupPolicy.
 */
export interface GetBackupPolicyOutputArgs {
    name?: pulumi.Input<string>;
    policyId?: pulumi.Input<string>;
}
