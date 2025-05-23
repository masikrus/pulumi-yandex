// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getLbNetworkLoadBalancer(args?: GetLbNetworkLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLbNetworkLoadBalancerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("yandex:index/getLbNetworkLoadBalancer:getLbNetworkLoadBalancer", {
        "folderId": args.folderId,
        "name": args.name,
        "networkLoadBalancerId": args.networkLoadBalancerId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbNetworkLoadBalancer.
 */
export interface GetLbNetworkLoadBalancerArgs {
    folderId?: string;
    name?: string;
    networkLoadBalancerId?: string;
}

/**
 * A collection of values returned by getLbNetworkLoadBalancer.
 */
export interface GetLbNetworkLoadBalancerResult {
    readonly allowZonalShift: boolean;
    readonly attachedTargetGroups: outputs.GetLbNetworkLoadBalancerAttachedTargetGroup[];
    readonly createdAt: string;
    readonly deletionProtection: boolean;
    readonly description: string;
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly listeners: outputs.GetLbNetworkLoadBalancerListener[];
    readonly name: string;
    readonly networkLoadBalancerId: string;
    readonly regionId: string;
    readonly type: string;
}
export function getLbNetworkLoadBalancerOutput(args?: GetLbNetworkLoadBalancerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLbNetworkLoadBalancerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("yandex:index/getLbNetworkLoadBalancer:getLbNetworkLoadBalancer", {
        "folderId": args.folderId,
        "name": args.name,
        "networkLoadBalancerId": args.networkLoadBalancerId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbNetworkLoadBalancer.
 */
export interface GetLbNetworkLoadBalancerOutputArgs {
    folderId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    networkLoadBalancerId?: pulumi.Input<string>;
}
