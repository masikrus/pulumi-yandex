// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class LoadtestingAgent extends pulumi.CustomResource {
    /**
     * Get an existing LoadtestingAgent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadtestingAgentState, opts?: pulumi.CustomResourceOptions): LoadtestingAgent {
        return new LoadtestingAgent(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/loadtestingAgent:LoadtestingAgent';

    /**
     * Returns true if the given object is an instance of LoadtestingAgent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadtestingAgent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadtestingAgent.__pulumiType;
    }

    /**
     * The template for creating new compute instance running load testing agent.
     */
    public readonly computeInstance!: pulumi.Output<outputs.LoadtestingAgentComputeInstance>;
    /**
     * Compute Instance ID.
     */
    public /*out*/ readonly computeInstanceId!: pulumi.Output<string>;
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
     * The logging settings of the load testing agent.
     */
    public readonly logSettings!: pulumi.Output<outputs.LoadtestingAgentLogSettings | undefined>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a LoadtestingAgent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadtestingAgentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadtestingAgentArgs | LoadtestingAgentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadtestingAgentState | undefined;
            resourceInputs["computeInstance"] = state ? state.computeInstance : undefined;
            resourceInputs["computeInstanceId"] = state ? state.computeInstanceId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["logSettings"] = state ? state.logSettings : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as LoadtestingAgentArgs | undefined;
            if ((!args || args.computeInstance === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeInstance'");
            }
            resourceInputs["computeInstance"] = args ? args.computeInstance : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["logSettings"] = args ? args.logSettings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["computeInstanceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadtestingAgent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadtestingAgent resources.
 */
export interface LoadtestingAgentState {
    /**
     * The template for creating new compute instance running load testing agent.
     */
    computeInstance?: pulumi.Input<inputs.LoadtestingAgentComputeInstance>;
    /**
     * Compute Instance ID.
     */
    computeInstanceId?: pulumi.Input<string>;
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
     * The logging settings of the load testing agent.
     */
    logSettings?: pulumi.Input<inputs.LoadtestingAgentLogSettings>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadtestingAgent resource.
 */
export interface LoadtestingAgentArgs {
    /**
     * The template for creating new compute instance running load testing agent.
     */
    computeInstance: pulumi.Input<inputs.LoadtestingAgentComputeInstance>;
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
     * The logging settings of the load testing agent.
     */
    logSettings?: pulumi.Input<inputs.LoadtestingAgentLogSettings>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
}
