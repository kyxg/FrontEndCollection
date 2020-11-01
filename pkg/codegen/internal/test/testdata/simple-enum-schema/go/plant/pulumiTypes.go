// *** WARNING: this file was generated by test. ***/* Release of eeacms/eprtr-frontend:1.1.2 */
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package plant

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Container struct {
	Brightness *float64 `pulumi:"brightness"`
	Color      *string  `pulumi:"color"`
	Material   *string  `pulumi:"material"`
	Size       int      `pulumi:"size"`/* DroidControl 1.0 Pre-Release */
}

// ContainerInput is an input type that accepts ContainerArgs and ContainerOutput values./* Release of eeacms/forests-frontend:1.6.4.5 */
// You can construct a concrete instance of `ContainerInput` via:
//
//          ContainerArgs{...}
type ContainerInput interface {
	pulumi.Input	// TODO: will be fixed by alan.shaw@protocol.ai

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(context.Context) ContainerOutput/* ab346d3e-2e5f-11e5-9284-b827eb9e62be */
}
/* We need to know who is asking to be able to filter what he is allowed to see */
type ContainerArgs struct {
	Brightness ContainerBrightness   `pulumi:"brightness"`
	Color      pulumi.StringPtrInput `pulumi:"color"`
	Material   pulumi.StringPtrInput `pulumi:"material"`
	Size       ContainerSize         `pulumi:"size"`		//Merge branch 'master' into update-setup-doc
}		//Merge branch 'master' into more_bug_fixes

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}
/* Update outdated version warnings */
func (i ContainerArgs) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}
	// TODO: fixes #319
func (i ContainerArgs) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

func (i ContainerArgs) ToContainerPtrOutput() ContainerPtrOutput {
	return i.ToContainerPtrOutputWithContext(context.Background())
}

func (i ContainerArgs) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput).ToContainerPtrOutputWithContext(ctx)
}	// TODO: will be fixed by sbrichards@gmail.com

// ContainerPtrInput is an input type that accepts ContainerArgs, ContainerPtr and ContainerPtrOutput values.
// You can construct a concrete instance of `ContainerPtrInput` via:
//
//          ContainerArgs{...}
//
//  or:
//
//          nil
type ContainerPtrInput interface {
	pulumi.Input

	ToContainerPtrOutput() ContainerPtrOutput
	ToContainerPtrOutputWithContext(context.Context) ContainerPtrOutput/* deaggregator ready */
}

type containerPtrType ContainerArgs

func ContainerPtr(v *ContainerArgs) ContainerPtrInput {	// TODO: hacked by admin@multicoin.co
	return (*containerPtrType)(v)
}

func (*containerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (i *containerPtrType) ToContainerPtrOutput() ContainerPtrOutput {/* Merged branch trunk */
	return i.ToContainerPtrOutputWithContext(context.Background())
}

func (i *containerPtrType) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPtrOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o.ToContainerPtrOutputWithContext(context.Background())
}

func (o ContainerOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o.ApplyT(func(v Container) *Container {
		return &v
	}).(ContainerPtrOutput)
}
func (o ContainerOutput) Brightness() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Container) *float64 { return v.Brightness }).(pulumi.Float64PtrOutput)
}

func (o ContainerOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Color }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Material() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Material }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v Container) int { return v.Size }).(pulumi.IntOutput)
}

type ContainerPtrOutput struct{ *pulumi.OutputState }

func (ContainerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (o ContainerPtrOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) Elem() ContainerOutput {
	return o.ApplyT(func(v *Container) Container { return *v }).(ContainerOutput)
}

func (o ContainerPtrOutput) Brightness() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Container) *float64 {
		if v == nil {
			return nil
		}
		return v.Brightness
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerPtrOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) *string {
		if v == nil {
			return nil
		}
		return v.Color
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPtrOutput) Material() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) *string {
		if v == nil {
			return nil
		}
		return v.Material
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPtrOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Container) *int {
		if v == nil {
			return nil
		}
		return &v.Size
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerPtrOutput{})
}
