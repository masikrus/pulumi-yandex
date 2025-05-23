// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LoggingGroup extends pulumi.CustomResource {
    /**
     * Get an existing LoggingGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoggingGroupState, opts?: pulumi.CustomResourceOptions): LoggingGroup {
        return new LoggingGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/loggingGroup:LoggingGroup';

    /**
     * Returns true if the given object is an instance of LoggingGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoggingGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoggingGroup.__pulumiType;
    }

    /**
     * The `Cloud ID` which resource belongs to. If it is not provided, the default provider `cloud-id` is used.
     */
    public /*out*/ readonly cloudId!: pulumi.Output<string>;
    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Data Stream.
     */
    public readonly dataStream!: pulumi.Output<string | undefined>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Log entries retention period for the Yandex Cloud Logging group.
     */
    public readonly retentionPeriod!: pulumi.Output<string>;
    /**
     * The Yandex Cloud Logging group status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a LoggingGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LoggingGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoggingGroupArgs | LoggingGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoggingGroupState | undefined;
            resourceInputs["cloudId"] = state ? state.cloudId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dataStream"] = state ? state.dataStream : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as LoggingGroupArgs | undefined;
            resourceInputs["dataStream"] = args ? args.dataStream : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
            resourceInputs["cloudId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoggingGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoggingGroup resources.
 */
export interface LoggingGroupState {
    /**
     * The `Cloud ID` which resource belongs to. If it is not provided, the default provider `cloud-id` is used.
     */
    cloudId?: pulumi.Input<string>;
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Data Stream.
     */
    dataStream?: pulumi.Input<string>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Log entries retention period for the Yandex Cloud Logging group.
     */
    retentionPeriod?: pulumi.Input<string>;
    /**
     * The Yandex Cloud Logging group status.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoggingGroup resource.
 */
export interface LoggingGroupArgs {
    /**
     * Data Stream.
     */
    dataStream?: pulumi.Input<string>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Log entries retention period for the Yandex Cloud Logging group.
     */
    retentionPeriod?: pulumi.Input<string>;
}
