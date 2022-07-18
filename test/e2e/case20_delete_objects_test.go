// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"open-cluster-management.io/config-policy-controller/test/utils"
)

const (
	case20ConfigPolicyNameCreate   string = "policy-pod-create"
	case20ConfigPolicyNameEdit     string = "policy-pod-edit"
	case20ConfigPolicyNameExisting string = "policy-pod-already-created"
	case20ConfigPolicyNameInform   string = "policy-pod-inform"
	case20PolicyYamlCreate         string = "../resources/case20_delete_objects/case20_create_pod.yaml"
	case20PolicyYamlEdit           string = "../resources/case20_delete_objects/case20_edit_pod.yaml"
	case20PolicyYamlExisting       string = "../resources/case20_delete_objects/case20_enforce_noncreated_pod.yaml"
	case20PolicyYamlInform         string = "../resources/case20_delete_objects/case20_inform_pod.yaml"
)

var _ = Describe("Test status fields being set for object deletion", func() {
	Describe("Create a policy on managed cluster in ns:"+testNamespace, func() {
		It("should update status fields properly for created objects", func() {
			By("Creating " + case20ConfigPolicyNameCreate + " on managed")
			utils.Kubectl("apply", "-f", case20PolicyYamlCreate, "-n", testNamespace)
			plc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
				case20ConfigPolicyNameCreate, testNamespace, true, defaultTimeoutSeconds)
			Expect(plc).NotTo(BeNil())
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameCreate, testNamespace, true, defaultTimeoutSeconds)

				return utils.GetComplianceState(managedPlc)
			}, defaultTimeoutSeconds, 1).Should(Equal("Compliant"))
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameCreate, testNamespace, true, defaultTimeoutSeconds)
				relatedObj := managedPlc.Object["status"].(map[string]interface{})["relatedObjects"].([]interface{})[0]
				properties := relatedObj.(map[string]interface{})["properties"].(map[string]interface{})

				return properties["createdByPolicy"].(bool)
			}, defaultTimeoutSeconds, 1).Should(Equal(true))
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameCreate, testNamespace, true, defaultTimeoutSeconds)
				relatedObj := managedPlc.Object["status"].(map[string]interface{})["relatedObjects"].([]interface{})[0]
				properties := relatedObj.(map[string]interface{})["properties"].(map[string]interface{})

				return properties["uid"].(string)
			}, defaultTimeoutSeconds, 1).Should(Not(Equal("")))
		})
		It("should update status fields properly for non-created objects", func() {
			By("Creating " + case20ConfigPolicyNameExisting + " on managed")
			utils.Kubectl("apply", "-f", case20PolicyYamlExisting, "-n", testNamespace)
			plc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
				case20ConfigPolicyNameExisting, testNamespace, true, defaultTimeoutSeconds)
			Expect(plc).NotTo(BeNil())
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameExisting, testNamespace, true, defaultTimeoutSeconds)

				return utils.GetComplianceState(managedPlc)
			}, defaultTimeoutSeconds, 1).Should(Equal("Compliant"))
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameExisting, testNamespace, true, defaultTimeoutSeconds)
				relatedObj := managedPlc.Object["status"].(map[string]interface{})["relatedObjects"].([]interface{})[0]
				properties := relatedObj.(map[string]interface{})["properties"].(map[string]interface{})

				return properties["createdByPolicy"].(bool)
			}, defaultTimeoutSeconds, 1).Should(Equal(false))
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameExisting, testNamespace, true, defaultTimeoutSeconds)
				relatedObj := managedPlc.Object["status"].(map[string]interface{})["relatedObjects"].([]interface{})[0]
				properties := relatedObj.(map[string]interface{})["properties"].(map[string]interface{})

				return properties["uid"]
			}, defaultTimeoutSeconds, 1).Should(BeNil())
		})
		It("should update status fields properly for edited objects", func() {
			By("Creating " + case20ConfigPolicyNameEdit + " on managed")
			utils.Kubectl("apply", "-f", case20PolicyYamlEdit, "-n", testNamespace)
			plc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
				case20ConfigPolicyNameEdit, testNamespace, true, defaultTimeoutSeconds)
			Expect(plc).NotTo(BeNil())
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameEdit, testNamespace, true, defaultTimeoutSeconds)

				return utils.GetComplianceState(managedPlc)
			}, defaultTimeoutSeconds, 1).Should(Equal("Compliant"))
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameEdit, testNamespace, true, defaultTimeoutSeconds)
				relatedObj := managedPlc.Object["status"].(map[string]interface{})["relatedObjects"].([]interface{})[0]
				properties := relatedObj.(map[string]interface{})["properties"].(map[string]interface{})

				return properties["createdByPolicy"].(bool)
			}, defaultTimeoutSeconds, 1).Should(Equal(false))
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameEdit, testNamespace, true, defaultTimeoutSeconds)
				relatedObj := managedPlc.Object["status"].(map[string]interface{})["relatedObjects"].([]interface{})[0]
				properties := relatedObj.(map[string]interface{})["properties"].(map[string]interface{})

				return properties["uid"]
			}, defaultTimeoutSeconds, 1).Should(BeNil())
		})
		It("should not update status field for inform policies", func() {
			By("Creating " + case20ConfigPolicyNameInform + " on managed")
			utils.Kubectl("apply", "-f", case20PolicyYamlInform, "-n", testNamespace)
			plc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
				case20ConfigPolicyNameInform, testNamespace, true, defaultTimeoutSeconds)
			Expect(plc).NotTo(BeNil())
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameInform, testNamespace, true, defaultTimeoutSeconds)

				return utils.GetComplianceState(managedPlc)
			}, defaultTimeoutSeconds, 1).Should(Equal("Compliant"))
			Eventually(func() interface{} {
				managedPlc := utils.GetWithTimeout(clientManagedDynamic, gvrConfigPolicy,
					case20ConfigPolicyNameInform, testNamespace, true, defaultTimeoutSeconds)
				relatedObj := managedPlc.Object["status"].(map[string]interface{})["relatedObjects"].([]interface{})[0]
				properties := relatedObj.(map[string]interface{})["properties"]

				return properties
			}, defaultTimeoutSeconds, 1).Should(BeNil())
		})
		It("Cleans up", func() {
			utils.Kubectl("delete", "pod", "nginx-pod-e2e20", "-n", "default")
			policies := []string{
				case20ConfigPolicyNameCreate,
				case20ConfigPolicyNameExisting,
				case20ConfigPolicyNameEdit,
				case20ConfigPolicyNameInform,
			}

			deleteConfigPolicies(policies)
		})
	})
})