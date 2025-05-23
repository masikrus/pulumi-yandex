// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getAlbBackendGroup(args?: GetAlbBackendGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetAlbBackendGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("yandex:index/getAlbBackendGroup:getAlbBackendGroup", {
        "backendGroupId": args.backendGroupId,
        "description": args.description,
        "folderId": args.folderId,
        "grpcBackends": args.grpcBackends,
        "httpBackends": args.httpBackends,
        "labels": args.labels,
        "name": args.name,
        "sessionAffinity": args.sessionAffinity,
        "streamBackends": args.streamBackends,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlbBackendGroup.
 */
export interface GetAlbBackendGroupArgs {
    backendGroupId?: string;
    description?: string;
    folderId?: string;
    grpcBackends?: inputs.GetAlbBackendGroupGrpcBackend[];
    httpBackends?: inputs.GetAlbBackendGroupHttpBackend[];
    labels?: {[key: string]: string};
    name?: string;
    sessionAffinity?: inputs.GetAlbBackendGroupSessionAffinity;
    streamBackends?: inputs.GetAlbBackendGroupStreamBackend[];
}

/**
 * A collection of values returned by getAlbBackendGroup.
 */
export interface GetAlbBackendGroupResult {
    readonly backendGroupId: string;
    readonly createdAt: string;
    readonly description: string;
    readonly folderId: string;
    readonly grpcBackends: outputs.GetAlbBackendGroupGrpcBackend[];
    readonly httpBackends: outputs.GetAlbBackendGroupHttpBackend[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels: {[key: string]: string};
    readonly name: string;
    readonly sessionAffinity: outputs.GetAlbBackendGroupSessionAffinity;
    readonly streamBackends: outputs.GetAlbBackendGroupStreamBackend[];
}
export function getAlbBackendGroupOutput(args?: GetAlbBackendGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAlbBackendGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("yandex:index/getAlbBackendGroup:getAlbBackendGroup", {
        "backendGroupId": args.backendGroupId,
        "description": args.description,
        "folderId": args.folderId,
        "grpcBackends": args.grpcBackends,
        "httpBackends": args.httpBackends,
        "labels": args.labels,
        "name": args.name,
        "sessionAffinity": args.sessionAffinity,
        "streamBackends": args.streamBackends,
    }, opts);
}

/**
 * A collection of arguments for invoking getAlbBackendGroup.
 */
export interface GetAlbBackendGroupOutputArgs {
    backendGroupId?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    folderId?: pulumi.Input<string>;
    grpcBackends?: pulumi.Input<pulumi.Input<inputs.GetAlbBackendGroupGrpcBackendArgs>[]>;
    httpBackends?: pulumi.Input<pulumi.Input<inputs.GetAlbBackendGroupHttpBackendArgs>[]>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    sessionAffinity?: pulumi.Input<inputs.GetAlbBackendGroupSessionAffinityArgs>;
    streamBackends?: pulumi.Input<pulumi.Input<inputs.GetAlbBackendGroupStreamBackendArgs>[]>;
}
