package main

import (
	"example.com/cdk8s/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	k8s.NewKubeConfigMap(chart, jsii.String("name"), &k8s.KubeConfigMapProps{
		Data: &map[string]*string{
			"HOSTNAME": jsii.String("0.0.0.0"),
			"PORT":     jsii.String("8085")},
		Immutable: new(bool),
		Metadata:  &k8s.ObjectMeta{Name: jsii.String("app"), Namespace: jsii.String("default")},
	})

	k8s.NewKubePod(chart, jsii.String("go-sample-app"), &k8s.KubePodProps{
		Metadata: &k8s.ObjectMeta{Name: jsii.String("go-sample-app"), Namespace: jsii.String("default")},
		Spec: &k8s.PodSpec{
			Containers: &[]*k8s.Container{
				{
					Name:  jsii.String("go-sample-app"),
					Image: jsii.String("ghcr.io/gsantoro/go-sample-app:wolfi"),
					EnvFrom: &[]*k8s.EnvFromSource{
						{
							ConfigMapRef: &k8s.ConfigMapEnvSource{Name: jsii.String("app")},
						},
					},
					Resources: &k8s.ResourceRequirements{
						Limits: &map[string]k8s.Quantity{
							"cpu":    k8s.Quantity_FromString(jsii.String("100m")),
							"memory": k8s.Quantity_FromString(jsii.String("100Mi")),
						},
						Requests: &map[string]k8s.Quantity{
							"cpu":    k8s.Quantity_FromString(jsii.String("100m")),
							"memory": k8s.Quantity_FromString(jsii.String("100Mi")),
						},
					},
					Ports: &[]*k8s.ContainerPort{
						{
							ContainerPort: jsii.Number(8085),
						},
					},
				},
			},
		},
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "cdk8s", nil)
	app.Synth()
}
