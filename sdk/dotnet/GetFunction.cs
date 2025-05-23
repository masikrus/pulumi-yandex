// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetFunction
    {
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("yandex:index/getFunction:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("yandex:index/getFunction:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());

        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("yandex:index/getFunction:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : global::Pulumi.InvokeArgs
    {
        [Input("concurrency")]
        public int? Concurrency { get; set; }

        [Input("connectivity")]
        public Inputs.GetFunctionConnectivityArgs? Connectivity { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("functionId")]
        public string? FunctionId { get; set; }

        [Input("metadataOptions")]
        public Inputs.GetFunctionMetadataOptionsArgs? MetadataOptions { get; set; }

        [Input("mounts")]
        private List<Inputs.GetFunctionMountArgs>? _mounts;
        public List<Inputs.GetFunctionMountArgs> Mounts
        {
            get => _mounts ?? (_mounts = new List<Inputs.GetFunctionMountArgs>());
            set => _mounts = value;
        }

        [Input("name")]
        public string? Name { get; set; }

        [Input("secrets")]
        private List<Inputs.GetFunctionSecretArgs>? _secrets;
        public List<Inputs.GetFunctionSecretArgs> Secrets
        {
            get => _secrets ?? (_secrets = new List<Inputs.GetFunctionSecretArgs>());
            set => _secrets = value;
        }

        [Input("storageMounts")]
        private List<Inputs.GetFunctionStorageMountArgs>? _storageMounts;
        [Obsolete(@"to manage storage_mountss, please switch to using a separate resource type mounts")]
        public List<Inputs.GetFunctionStorageMountArgs> StorageMounts
        {
            get => _storageMounts ?? (_storageMounts = new List<Inputs.GetFunctionStorageMountArgs>());
            set => _storageMounts = value;
        }

        public GetFunctionArgs()
        {
        }
        public static new GetFunctionArgs Empty => new GetFunctionArgs();
    }

    public sealed class GetFunctionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("concurrency")]
        public Input<int>? Concurrency { get; set; }

        [Input("connectivity")]
        public Input<Inputs.GetFunctionConnectivityInputArgs>? Connectivity { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        [Input("metadataOptions")]
        public Input<Inputs.GetFunctionMetadataOptionsInputArgs>? MetadataOptions { get; set; }

        [Input("mounts")]
        private InputList<Inputs.GetFunctionMountInputArgs>? _mounts;
        public InputList<Inputs.GetFunctionMountInputArgs> Mounts
        {
            get => _mounts ?? (_mounts = new InputList<Inputs.GetFunctionMountInputArgs>());
            set => _mounts = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secrets")]
        private InputList<Inputs.GetFunctionSecretInputArgs>? _secrets;
        public InputList<Inputs.GetFunctionSecretInputArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.GetFunctionSecretInputArgs>());
            set => _secrets = value;
        }

        [Input("storageMounts")]
        private InputList<Inputs.GetFunctionStorageMountInputArgs>? _storageMounts;
        [Obsolete(@"to manage storage_mountss, please switch to using a separate resource type mounts")]
        public InputList<Inputs.GetFunctionStorageMountInputArgs> StorageMounts
        {
            get => _storageMounts ?? (_storageMounts = new InputList<Inputs.GetFunctionStorageMountInputArgs>());
            set => _storageMounts = value;
        }

        public GetFunctionInvokeArgs()
        {
        }
        public static new GetFunctionInvokeArgs Empty => new GetFunctionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        public readonly ImmutableArray<Outputs.GetFunctionAsyncInvocationResult> AsyncInvocations;
        public readonly int Concurrency;
        public readonly Outputs.GetFunctionConnectivityResult? Connectivity;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string Entrypoint;
        public readonly ImmutableDictionary<string, string> Environment;
        public readonly string ExecutionTimeout;
        public readonly string? FolderId;
        public readonly string? FunctionId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int ImageSize;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetFunctionLogOptionResult> LogOptions;
        public readonly int Memory;
        public readonly Outputs.GetFunctionMetadataOptionsResult MetadataOptions;
        public readonly ImmutableArray<Outputs.GetFunctionMountResult> Mounts;
        public readonly string? Name;
        public readonly string Runtime;
        public readonly ImmutableArray<Outputs.GetFunctionSecretResult> Secrets;
        public readonly string ServiceAccountId;
        public readonly ImmutableArray<Outputs.GetFunctionStorageMountResult> StorageMounts;
        public readonly ImmutableArray<string> Tags;
        public readonly int TmpfsSize;
        public readonly string Version;

        [OutputConstructor]
        private GetFunctionResult(
            ImmutableArray<Outputs.GetFunctionAsyncInvocationResult> asyncInvocations,

            int concurrency,

            Outputs.GetFunctionConnectivityResult? connectivity,

            string createdAt,

            string description,

            string entrypoint,

            ImmutableDictionary<string, string> environment,

            string executionTimeout,

            string? folderId,

            string? functionId,

            string id,

            int imageSize,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetFunctionLogOptionResult> logOptions,

            int memory,

            Outputs.GetFunctionMetadataOptionsResult metadataOptions,

            ImmutableArray<Outputs.GetFunctionMountResult> mounts,

            string? name,

            string runtime,

            ImmutableArray<Outputs.GetFunctionSecretResult> secrets,

            string serviceAccountId,

            ImmutableArray<Outputs.GetFunctionStorageMountResult> storageMounts,

            ImmutableArray<string> tags,

            int tmpfsSize,

            string version)
        {
            AsyncInvocations = asyncInvocations;
            Concurrency = concurrency;
            Connectivity = connectivity;
            CreatedAt = createdAt;
            Description = description;
            Entrypoint = entrypoint;
            Environment = environment;
            ExecutionTimeout = executionTimeout;
            FolderId = folderId;
            FunctionId = functionId;
            Id = id;
            ImageSize = imageSize;
            Labels = labels;
            LogOptions = logOptions;
            Memory = memory;
            MetadataOptions = metadataOptions;
            Mounts = mounts;
            Name = name;
            Runtime = runtime;
            Secrets = secrets;
            ServiceAccountId = serviceAccountId;
            StorageMounts = storageMounts;
            Tags = tags;
            TmpfsSize = tmpfsSize;
            Version = version;
        }
    }
}
