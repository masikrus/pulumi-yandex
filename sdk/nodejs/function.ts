// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Config for asynchronous invocations of Yandex Cloud Function.
     */
    public readonly asyncInvocation!: pulumi.Output<outputs.FunctionAsyncInvocation | undefined>;
    /**
     * The maximum number of requests processed by a function instance at the same time.
     */
    public readonly concurrency!: pulumi.Output<number>;
    /**
     * Function version connectivity. If specified the version will be attached to specified network.
     */
    public readonly connectivity!: pulumi.Output<outputs.FunctionConnectivity | undefined>;
    /**
     * Version deployment content for Yandex Cloud Function code. Can be only one `package` or `content` section. Either
     * `package` or `content` section must be specified.
     */
    public readonly content!: pulumi.Output<outputs.FunctionContent | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Entrypoint for Yandex Cloud Function.
     */
    public readonly entrypoint!: pulumi.Output<string>;
    /**
     * A set of key/value environment variables for Yandex Cloud Function. Each key must begin with a letter (A-Z, a-z).
     */
    public readonly environment!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Execution timeout in seconds for Yandex Cloud Function.
     */
    public readonly executionTimeout!: pulumi.Output<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Image size for Yandex Cloud Function.
     */
    public /*out*/ readonly imageSize!: pulumi.Output<number>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Options for logging from Yandex Cloud Function.
     */
    public readonly logOptions!: pulumi.Output<outputs.FunctionLogOptions | undefined>;
    /**
     * Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Function.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * Options set the access mode to function's metadata endpoints.
     */
    public readonly metadataOptions!: pulumi.Output<outputs.FunctionMetadataOptions>;
    /**
     * Mounts for Yandex Cloud Function.
     */
    public readonly mounts!: pulumi.Output<outputs.FunctionMount[]>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Version deployment package for Yandex Cloud Function code. Can be only one `package` or `content` section. Either
     * `package` or `content` section must be specified.
     */
    public readonly package!: pulumi.Output<outputs.FunctionPackage | undefined>;
    /**
     * Runtime for Yandex Cloud Function.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * Secrets for Yandex Cloud Function.
     */
    public readonly secrets!: pulumi.Output<outputs.FunctionSecret[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    public readonly serviceAccountId!: pulumi.Output<string | undefined>;
    /**
     * (**DEPRECATED**, use `mounts > objectStorage` instead). Storage mounts for Yandex Cloud Function.
     *
     * @deprecated The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
     */
    public readonly storageMounts!: pulumi.Output<outputs.FunctionStorageMount[]>;
    /**
     * Tags for Yandex Cloud Function. Tag `$latest` isn't returned.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * Tmpfs size for Yandex Cloud Function.
     */
    public readonly tmpfsSize!: pulumi.Output<number>;
    /**
     * User-defined string for current function version. User must change this string any times when function changed. Function
     * will be updated when hash is changed.
     */
    public readonly userHash!: pulumi.Output<string>;
    /**
     * Version of Yandex Cloud Function.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["asyncInvocation"] = state ? state.asyncInvocation : undefined;
            resourceInputs["concurrency"] = state ? state.concurrency : undefined;
            resourceInputs["connectivity"] = state ? state.connectivity : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["entrypoint"] = state ? state.entrypoint : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["executionTimeout"] = state ? state.executionTimeout : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["imageSize"] = state ? state.imageSize : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["logOptions"] = state ? state.logOptions : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["metadataOptions"] = state ? state.metadataOptions : undefined;
            resourceInputs["mounts"] = state ? state.mounts : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["package"] = state ? state.package : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["secrets"] = state ? state.secrets : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
            resourceInputs["storageMounts"] = state ? state.storageMounts : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tmpfsSize"] = state ? state.tmpfsSize : undefined;
            resourceInputs["userHash"] = state ? state.userHash : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.entrypoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entrypoint'");
            }
            if ((!args || args.memory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            if ((!args || args.userHash === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userHash'");
            }
            resourceInputs["asyncInvocation"] = args ? args.asyncInvocation : undefined;
            resourceInputs["concurrency"] = args ? args.concurrency : undefined;
            resourceInputs["connectivity"] = args ? args.connectivity : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["entrypoint"] = args ? args.entrypoint : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["executionTimeout"] = args ? args.executionTimeout : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["logOptions"] = args ? args.logOptions : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["metadataOptions"] = args ? args.metadataOptions : undefined;
            resourceInputs["mounts"] = args ? args.mounts : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["package"] = args ? args.package : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["secrets"] = args ? args.secrets : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["storageMounts"] = args ? args.storageMounts : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tmpfsSize"] = args ? args.tmpfsSize : undefined;
            resourceInputs["userHash"] = args ? args.userHash : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["imageSize"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * Config for asynchronous invocations of Yandex Cloud Function.
     */
    asyncInvocation?: pulumi.Input<inputs.FunctionAsyncInvocation>;
    /**
     * The maximum number of requests processed by a function instance at the same time.
     */
    concurrency?: pulumi.Input<number>;
    /**
     * Function version connectivity. If specified the version will be attached to specified network.
     */
    connectivity?: pulumi.Input<inputs.FunctionConnectivity>;
    /**
     * Version deployment content for Yandex Cloud Function code. Can be only one `package` or `content` section. Either
     * `package` or `content` section must be specified.
     */
    content?: pulumi.Input<inputs.FunctionContent>;
    createdAt?: pulumi.Input<string>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Entrypoint for Yandex Cloud Function.
     */
    entrypoint?: pulumi.Input<string>;
    /**
     * A set of key/value environment variables for Yandex Cloud Function. Each key must begin with a letter (A-Z, a-z).
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Execution timeout in seconds for Yandex Cloud Function.
     */
    executionTimeout?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Image size for Yandex Cloud Function.
     */
    imageSize?: pulumi.Input<number>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Options for logging from Yandex Cloud Function.
     */
    logOptions?: pulumi.Input<inputs.FunctionLogOptions>;
    /**
     * Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Function.
     */
    memory?: pulumi.Input<number>;
    /**
     * Options set the access mode to function's metadata endpoints.
     */
    metadataOptions?: pulumi.Input<inputs.FunctionMetadataOptions>;
    /**
     * Mounts for Yandex Cloud Function.
     */
    mounts?: pulumi.Input<pulumi.Input<inputs.FunctionMount>[]>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Version deployment package for Yandex Cloud Function code. Can be only one `package` or `content` section. Either
     * `package` or `content` section must be specified.
     */
    package?: pulumi.Input<inputs.FunctionPackage>;
    /**
     * Runtime for Yandex Cloud Function.
     */
    runtime?: pulumi.Input<string>;
    /**
     * Secrets for Yandex Cloud Function.
     */
    secrets?: pulumi.Input<pulumi.Input<inputs.FunctionSecret>[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * (**DEPRECATED**, use `mounts > objectStorage` instead). Storage mounts for Yandex Cloud Function.
     *
     * @deprecated The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
     */
    storageMounts?: pulumi.Input<pulumi.Input<inputs.FunctionStorageMount>[]>;
    /**
     * Tags for Yandex Cloud Function. Tag `$latest` isn't returned.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tmpfs size for Yandex Cloud Function.
     */
    tmpfsSize?: pulumi.Input<number>;
    /**
     * User-defined string for current function version. User must change this string any times when function changed. Function
     * will be updated when hash is changed.
     */
    userHash?: pulumi.Input<string>;
    /**
     * Version of Yandex Cloud Function.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Config for asynchronous invocations of Yandex Cloud Function.
     */
    asyncInvocation?: pulumi.Input<inputs.FunctionAsyncInvocation>;
    /**
     * The maximum number of requests processed by a function instance at the same time.
     */
    concurrency?: pulumi.Input<number>;
    /**
     * Function version connectivity. If specified the version will be attached to specified network.
     */
    connectivity?: pulumi.Input<inputs.FunctionConnectivity>;
    /**
     * Version deployment content for Yandex Cloud Function code. Can be only one `package` or `content` section. Either
     * `package` or `content` section must be specified.
     */
    content?: pulumi.Input<inputs.FunctionContent>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Entrypoint for Yandex Cloud Function.
     */
    entrypoint: pulumi.Input<string>;
    /**
     * A set of key/value environment variables for Yandex Cloud Function. Each key must begin with a letter (A-Z, a-z).
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Execution timeout in seconds for Yandex Cloud Function.
     */
    executionTimeout?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Options for logging from Yandex Cloud Function.
     */
    logOptions?: pulumi.Input<inputs.FunctionLogOptions>;
    /**
     * Memory in megabytes (**aligned to 128MB**) for Yandex Cloud Function.
     */
    memory: pulumi.Input<number>;
    /**
     * Options set the access mode to function's metadata endpoints.
     */
    metadataOptions?: pulumi.Input<inputs.FunctionMetadataOptions>;
    /**
     * Mounts for Yandex Cloud Function.
     */
    mounts?: pulumi.Input<pulumi.Input<inputs.FunctionMount>[]>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Version deployment package for Yandex Cloud Function code. Can be only one `package` or `content` section. Either
     * `package` or `content` section must be specified.
     */
    package?: pulumi.Input<inputs.FunctionPackage>;
    /**
     * Runtime for Yandex Cloud Function.
     */
    runtime: pulumi.Input<string>;
    /**
     * Secrets for Yandex Cloud Function.
     */
    secrets?: pulumi.Input<pulumi.Input<inputs.FunctionSecret>[]>;
    /**
     * [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
     */
    serviceAccountId?: pulumi.Input<string>;
    /**
     * (**DEPRECATED**, use `mounts > objectStorage` instead). Storage mounts for Yandex Cloud Function.
     *
     * @deprecated The 'storage_mounts' field has been deprecated. Please use 'mounts' instead.
     */
    storageMounts?: pulumi.Input<pulumi.Input<inputs.FunctionStorageMount>[]>;
    /**
     * Tags for Yandex Cloud Function. Tag `$latest` isn't returned.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tmpfs size for Yandex Cloud Function.
     */
    tmpfsSize?: pulumi.Input<number>;
    /**
     * User-defined string for current function version. User must change this string any times when function changed. Function
     * will be updated when hash is changed.
     */
    userHash: pulumi.Input<string>;
}
