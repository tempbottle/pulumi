// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Check codegen of functions with default values.
func FuncWithDefaultValue(ctx *pulumi.Context, args *FuncWithDefaultValueArgs, opts ...pulumi.InvokeOption) (*FuncWithDefaultValueResult, error) {
	var rv FuncWithDefaultValueResult
	err := ctx.Invoke("mypkg::funcWithDefaultValue", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FuncWithDefaultValueArgs struct {
	A string  `pulumi:"a"`
	B *string `pulumi:"b"`
}

// Defaults sets the appropriate defaults for FuncWithDefaultValueArgs
func (val *FuncWithDefaultValueArgs) Defaults() *FuncWithDefaultValueArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.B == nil {
		b_ := "b-default"
		tmp.B = &b_
	}
	return &tmp
}

type FuncWithDefaultValueResult struct {
	R string `pulumi:"r"`
}

func FuncWithDefaultValueOutput(ctx *pulumi.Context, args FuncWithDefaultValueOutputArgs, opts ...pulumi.InvokeOption) FuncWithDefaultValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (FuncWithDefaultValueResult, error) {
			args := v.(FuncWithDefaultValueArgs)
			r, err := FuncWithDefaultValue(ctx, &args, opts...)
			var s FuncWithDefaultValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(FuncWithDefaultValueResultOutput)
}

type FuncWithDefaultValueOutputArgs struct {
	A pulumi.StringInput    `pulumi:"a"`
	B pulumi.StringPtrInput `pulumi:"b"`
}

func (FuncWithDefaultValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithDefaultValueArgs)(nil)).Elem()
}

type FuncWithDefaultValueResultOutput struct{ *pulumi.OutputState }

func (FuncWithDefaultValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithDefaultValueResult)(nil)).Elem()
}

func (o FuncWithDefaultValueResultOutput) ToFuncWithDefaultValueResultOutput() FuncWithDefaultValueResultOutput {
	return o
}

func (o FuncWithDefaultValueResultOutput) ToFuncWithDefaultValueResultOutputWithContext(ctx context.Context) FuncWithDefaultValueResultOutput {
	return o
}

func (o FuncWithDefaultValueResultOutput) R() pulumi.StringOutput {
	return o.ApplyT(func(v FuncWithDefaultValueResult) string { return v.R }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FuncWithDefaultValueResultOutput{})
}
