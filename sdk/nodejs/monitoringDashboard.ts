// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class MonitoringDashboard extends pulumi.CustomResource {
    /**
     * Get an existing MonitoringDashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitoringDashboardState, opts?: pulumi.CustomResourceOptions): MonitoringDashboard {
        return new MonitoringDashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/monitoringDashboard:MonitoringDashboard';

    /**
     * Returns true if the given object is an instance of MonitoringDashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitoringDashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitoringDashboard.__pulumiType;
    }

    /**
     * Dashboard ID.
     */
    public /*out*/ readonly dashboardId!: pulumi.Output<string>;
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
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Dashboard parametrization.
     */
    public readonly parametrizations!: pulumi.Output<outputs.MonitoringDashboardParametrization[]>;
    /**
     * Dashboard title.
     */
    public readonly title!: pulumi.Output<string | undefined>;
    /**
     * Widgets.
     */
    public readonly widgets!: pulumi.Output<outputs.MonitoringDashboardWidget[] | undefined>;

    /**
     * Create a MonitoringDashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MonitoringDashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitoringDashboardArgs | MonitoringDashboardState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitoringDashboardState | undefined;
            resourceInputs["dashboardId"] = state ? state.dashboardId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parametrizations"] = state ? state.parametrizations : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["widgets"] = state ? state.widgets : undefined;
        } else {
            const args = argsOrState as MonitoringDashboardArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parametrizations"] = args ? args.parametrizations : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["widgets"] = args ? args.widgets : undefined;
            resourceInputs["dashboardId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MonitoringDashboard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MonitoringDashboard resources.
 */
export interface MonitoringDashboardState {
    /**
     * Dashboard ID.
     */
    dashboardId?: pulumi.Input<string>;
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
     * Dashboard parametrization.
     */
    parametrizations?: pulumi.Input<pulumi.Input<inputs.MonitoringDashboardParametrization>[]>;
    /**
     * Dashboard title.
     */
    title?: pulumi.Input<string>;
    /**
     * Widgets.
     */
    widgets?: pulumi.Input<pulumi.Input<inputs.MonitoringDashboardWidget>[]>;
}

/**
 * The set of arguments for constructing a MonitoringDashboard resource.
 */
export interface MonitoringDashboardArgs {
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
     * Dashboard parametrization.
     */
    parametrizations?: pulumi.Input<pulumi.Input<inputs.MonitoringDashboardParametrization>[]>;
    /**
     * Dashboard title.
     */
    title?: pulumi.Input<string>;
    /**
     * Widgets.
     */
    widgets?: pulumi.Input<pulumi.Input<inputs.MonitoringDashboardWidget>[]>;
}
