// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package shared

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./api/shared.Condition":                               schema__api_shared_Condition(ref),
		"./api/shared.NodeNetworkConfigurationEnactmentStatus": schema__api_shared_NodeNetworkConfigurationEnactmentStatus(ref),
		"./api/shared.NodeNetworkConfigurationPolicySpec":      schema__api_shared_NodeNetworkConfigurationPolicySpec(ref),
		"./api/shared.NodeNetworkConfigurationPolicyStatus":    schema__api_shared_NodeNetworkConfigurationPolicyStatus(ref),
		"./api/shared.NodeNetworkStateStatus":                  schema__api_shared_NodeNetworkStateStatus(ref),
		"./api/shared.State":                                   schema__api_shared_State(ref),
	}
}

func schema__api_shared_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"lastHearbeatTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema__api_shared_NodeNetworkConfigurationEnactmentStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeNetworkConfigurationEnactmentStatus defines the observed state of NodeNetworkConfigurationEnactment",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"desiredState": {
						SchemaProps: spec.SchemaProps{
							Description: "The desired state rendered for the enactment's node using the policy desiredState as template",
							Ref:         ref("./api/shared.State"),
						},
					},
					"policyGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "The generation from policy needed to check if an enactment condition status belongs to the same policy version",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./api/shared.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./api/shared.Condition", "./api/shared.State"},
	}
}

func schema__api_shared_NodeNetworkConfigurationPolicySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeNetworkConfigurationPolicySpec defines the desired state of NodeNetworkConfigurationPolicy",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the policy to be applied to the node. Selector which must match a node's labels for the policy to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"desiredState": {
						SchemaProps: spec.SchemaProps{
							Description: "The desired configuration of the policy",
							Ref:         ref("./api/shared.State"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./api/shared.State"},
	}
}

func schema__api_shared_NodeNetworkConfigurationPolicyStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeNetworkConfigurationPolicyStatus defines the observed state of NodeNetworkConfigurationPolicy",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./api/shared.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./api/shared.Condition"},
	}
}

func schema__api_shared_NodeNetworkStateStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeNetworkStateStatus is the status of the NodeNetworkState of a specific node",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"currentState": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./api/shared.State"),
						},
					},
					"lastSuccessfulUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./api/shared.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./api/shared.Condition", "./api/shared.State", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema__api_shared_State(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "State contains the namestatectl yaml [1] as string instead of golang struct so we don't need to be in sync with the schema.\n\n[1] https://github.com/nmstate/nmstate/blob/master/libnmstate/schemas/operational-state.yaml",
				Type:        []string{"object"},
			},
		},
	}
}