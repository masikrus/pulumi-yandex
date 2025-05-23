// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ServerlessContainer extends pulumi.CustomResource {
    /**
     * Get an existing ServerlessContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerlessContainerState, opts?: pulumi.CustomResourceOptions): ServerlessContainer {
        return new ServerlessContainer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/serverlessContainer:ServerlessContainer';

    /**
     * Returns true if the given object is an instance of ServerlessContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerlessContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerlessContainer.__pulumiType;
    }

    /**
     * Concurrency of Yandex Cloud Serverless Container.
     */
    public readonly concurrency!: pulumi.Output<number>;
    /**
     * Network access. If specified the revision will be attached to specified network.
     */
    public readonly connectivity!: pulumi.Output<outputs.ServerlessContainerConnectivity | undefined>;
    /**
     * Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
     */
    public readonly coreFraction!: pulumi.Output<number>;
    /**
     * Cores (**1+**) of the Yandex Cloud Serverless Container.
     */
    public readonly cores!: pulumi.Output<number | undefined>;
    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
     */
    public readonly executionTimeout!: pulumi.Output<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Revision deployment image for Yandex Cloud Serverless Container.
     */
    public readonly image!: pulumi.Output<outputs.ServerlessContainerImage>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Options for logging from Yandex Cloud Serverless Container.
     */
    public readonly logOptions!: pulumi.Output<outputs.ServerlessContainerLogOptions | undefined>;
    /**
     * Memory in megabytes (**aligned to 128 MB**).
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * Options set the access mode to revision's metadata endpoints.
     */
    public readonly metadataOptions!: pulumi.Output<outputs.ServerlessContainerMetadataOptions>;
    /**
     * Mounts for Yandex Cloud Serverless Container.
     */
    public readonly mounts!: pulumi.Output<outputs.ServerlessContainerMount[]>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Provision policy. If specified the revision will have prepared instances.
     */
    public readonly provisionPolicy!: pulumi.Output<outputs.ServerlessContainerProvisionPolicy | undefined>;
    /**
     * Last revision ID of the Yandex Cloud Serverless Container.
     */
    public /*out*/ readonly revisionId!: pulumi.Output<string>;
    /**
     * Runtime for Yandex Cloud Serverless Container.
     */
    public readonly runtime!: pulumi.Output<outputs.ServerlessContainerRuntime>;
    /**
     * Secrets for Yandex Cloud Serverless Container.
     */
    public readonly secrets!: pulumi.Output<outputs.ServerlessContainerSecret[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    public readonly serviceAccountId!: pulumi.Output<string | undefined>;
    /**
     * (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
     *
     * @deprecated The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
     */
    public readonly storageMounts!: pulumi.Output<outputs.ServerlessContainerStorageMount[]>;
    /**
     * Invoke URL for the Yandex Cloud Serverless Container.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a ServerlessContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerlessContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerlessContainerArgs | ServerlessContainerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerlessContainerState | undefined;
            resourceInputs["concurrency"] = state ? state.concurrency : undefined;
            resourceInputs["connectivity"] = state ? state.connectivity : undefined;
            resourceInputs["coreFraction"] = state ? state.coreFraction : undefined;
            resourceInputs["cores"] = state ? state.cores : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["executionTimeout"] = state ? state.executionTimeout : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["image"] = state ? state.image : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["logOptions"] = state ? state.logOptions : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["metadataOptions"] = state ? state.metadataOptions : undefined;
            resourceInputs["mounts"] = state ? state.mounts : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["provisionPolicy"] = state ? state.provisionPolicy : undefined;
            resourceInputs["revisionId"] = state ? state.revisionId : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["secrets"] = state ? state.secrets : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["storageMounts"] = state ? state.storageMounts : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ServerlessContainerArgs | undefined;
            if ((!args || args.image === undefined) && !opts.urn) {
                throw new Error("Missing required property 'image'");
            }
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            resourceInputs["concurrency"] = args ? args.concurrency : undefined;
            resourceInputs["connectivity"] = args ? args.connectivity : undefined;
            resourceInputs["coreFraction"] = args ? args.coreFraction : undefined;
            resourceInputs["cores"] = args ? args.cores : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["executionTimeout"] = args ? args.executionTimeout : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["image"] = args ? args.image : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["logOptions"] = args ? args.logOptions : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["metadataOptions"] = args ? args.metadataOptions : undefined;
            resourceInputs["mounts"] = args ? args.mounts : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["provisionPolicy"] = args ? args.provisionPolicy : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["secrets"] = args ? args.secrets : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["storageMounts"] = args ? args.storageMounts : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["revisionId"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerlessContainer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerlessContainer resources.
 */
export interface ServerlessContainerState {
    /**
     * Concurrency of Yandex Cloud Serverless Container.
     */
    concurrency?: pulumi.Input<number>;
    /**
     * Network access. If specified the revision will be attached to specified network.
     */
    connectivity?: pulumi.Input<inputs.ServerlessContainerConnectivity>;
    /**
     * Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
     */
    coreFraction?: pulumi.Input<number>;
    /**
     * Cores (**1+**) of the Yandex Cloud Serverless Container.
     */
    cores?: pulumi.Input<number>;
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
     */
    executionTimeout?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Revision deployment image for Yandex Cloud Serverless Container.
     */
    image?: pulumi.Input<inputs.ServerlessContainerImage>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Options for logging from Yandex Cloud Serverless Container.
     */
    logOptions?: pulumi.Input<inputs.ServerlessContainerLogOptions>;
    /**
     * Memory in megabytes (**aligned to 128 MB**).
     */
    memory?: pulumi.Input<number>;
    /**
     * Options set the access mode to revision's metadata endpoints.
     */
    metadataOptions?: pulumi.Input<inputs.ServerlessContainerMetadataOptions>;
    /**
     * Mounts for Yandex Cloud Serverless Container.
     */
    mounts?: pulumi.Input<pulumi.Input<inputs.ServerlessContainerMount>[]>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Provision policy. If specified the revision will have prepared instances.
     */
    provisionPolicy?: pulumi.Input<inputs.ServerlessContainerProvisionPolicy>;
    /**
     * Last revision ID of the Yandex Cloud Serverless Container.
     */
    revisionId?: pulumi.Input<string>;
    /**
     * Runtime for Yandex Cloud Serverless Container.
     */
    runtime?: pulumi.Input<inputs.ServerlessContainerRuntime>;
    /**
     * Secrets for Yandex Cloud Serverless Container.
     */
    secrets?: pulumi.Input<pulumi.Input<inputs.ServerlessContainerSecret>[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
     *
     * @deprecated The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
     */
    storageMounts?: pulumi.Input<pulumi.Input<inputs.ServerlessContainerStorageMount>[]>;
    /**
     * Invoke URL for the Yandex Cloud Serverless Container.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerlessContainer resource.
 */
export interface ServerlessContainerArgs {
    /**
     * Concurrency of Yandex Cloud Serverless Container.
     */
    concurrency?: pulumi.Input<number>;
    /**
     * Network access. If specified the revision will be attached to specified network.
     */
    connectivity?: pulumi.Input<inputs.ServerlessContainerConnectivity>;
    /**
     * Core fraction (**0...100**) of the Yandex Cloud Serverless Container.
     */
    coreFraction?: pulumi.Input<number>;
    /**
     * Cores (**1+**) of the Yandex Cloud Serverless Container.
     */
    cores?: pulumi.Input<number>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Execution timeout in seconds (**duration format**) for Yandex Cloud Serverless Container.
     */
    executionTimeout?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Revision deployment image for Yandex Cloud Serverless Container.
     */
    image: pulumi.Input<inputs.ServerlessContainerImage>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Options for logging from Yandex Cloud Serverless Container.
     */
    logOptions?: pulumi.Input<inputs.ServerlessContainerLogOptions>;
    /**
     * Memory in megabytes (**aligned to 128 MB**).
     */
    memory: pulumi.Input<number>;
    /**
     * Options set the access mode to revision's metadata endpoints.
     */
    metadataOptions?: pulumi.Input<inputs.ServerlessContainerMetadataOptions>;
    /**
     * Mounts for Yandex Cloud Serverless Container.
     */
    mounts?: pulumi.Input<pulumi.Input<inputs.ServerlessContainerMount>[]>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Provision policy. If specified the revision will have prepared instances.
     */
    provisionPolicy?: pulumi.Input<inputs.ServerlessContainerProvisionPolicy>;
    /**
     * Runtime for Yandex Cloud Serverless Container.
     */
    runtime?: pulumi.Input<inputs.ServerlessContainerRuntime>;
    /**
     * Secrets for Yandex Cloud Serverless Container.
     */
    secrets?: pulumi.Input<pulumi.Input<inputs.ServerlessContainerSecret>[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * (**DEPRECATED**, use `mounts.object_storage` instead) Storage mounts for Yandex Cloud Serverless Container.
     *
     * @deprecated The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
     */
    storageMounts?: pulumi.Input<pulumi.Input<inputs.ServerlessContainerStorageMount>[]>;
}
