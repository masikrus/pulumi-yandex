// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIamServiceAccount(args?: GetIamServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetIamServiceAccountResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("yandex:index/getIamServiceAccount:getIamServiceAccount", {
        "folderId": args.folderId,
        "name": args.name,
        "serviceAccountId": args.serviceAccountId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamServiceAccount.
 */
export interface GetIamServiceAccountArgs {
    folderId?: string;
    name?: string;
    serviceAccountId?: string;
}

/**
 * A collection of values returned by getIamServiceAccount.
 */
export interface GetIamServiceAccountResult {
    readonly createdAt: string;
    readonly description: string;
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly serviceAccountId: string;
}
export function getIamServiceAccountOutput(args?: GetIamServiceAccountOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIamServiceAccountResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("yandex:index/getIamServiceAccount:getIamServiceAccount", {
        "folderId": args.folderId,
        "name": args.name,
        "serviceAccountId": args.serviceAccountId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamServiceAccount.
 */
export interface GetIamServiceAccountOutputArgs {
    folderId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    serviceAccountId?: pulumi.Input<string>;
}
