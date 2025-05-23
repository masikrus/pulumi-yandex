// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getComputeImage(args?: GetComputeImageArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("yandex:index/getComputeImage:getComputeImage", {
        "family": args.family,
        "folderId": args.folderId,
        "imageId": args.imageId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeImage.
 */
export interface GetComputeImageArgs {
    family?: string;
    folderId?: string;
    imageId?: string;
    name?: string;
}

/**
 * A collection of values returned by getComputeImage.
 */
export interface GetComputeImageResult {
    readonly createdAt: string;
    readonly description: string;
    readonly family: string;
    readonly folderId: string;
    readonly hardwareGenerations: outputs.GetComputeImageHardwareGeneration[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId: string;
    readonly kmsKeyId: string;
    readonly labels: {[key: string]: string};
    readonly minDiskSize: number;
    readonly name: string;
    readonly osType: string;
    readonly pooled: boolean;
    readonly productIds: string[];
    readonly size: number;
    readonly status: string;
}
export function getComputeImageOutput(args?: GetComputeImageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetComputeImageResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("yandex:index/getComputeImage:getComputeImage", {
        "family": args.family,
        "folderId": args.folderId,
        "imageId": args.imageId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeImage.
 */
export interface GetComputeImageOutputArgs {
    family?: pulumi.Input<string>;
    folderId?: pulumi.Input<string>;
    imageId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
