/*
TEKTON LAUNCHER
created by: Nick Gerace

MIT License, Copyright (c) Nick Gerace
See 'LICENSE' file for more information

Please find license and further
information via the link below.
https://github.com/nickgerace/tekton-launcher
*/

package util

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

func Launch() {

	// Get kube config from default location or from flag.
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Build config from the selected path.
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create dynamic config.
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Define the TaskRun CRD for lookup.
	taskRunResource := schema.GroupVersionResource{Group: "tekton.dev", Version: "v1alpha1", Resource: "taskruns"}

	// Construct the TaskRun object for creation.
	taskRun := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1alpha1",
			"kind":       "TaskRun",
			"metadata": map[string]interface{}{
				"name": "example-taskrun",
			},
			"spec": map[string]interface{}{
				"taskSpec": map[string]interface{}{
					"steps": []map[string]interface{}{{
						"name":  "debian",
						"image": "debian:10-slim",
						"command": []string{
							"echo",
						},
						"args": []string{
							"Hello world!",
						},
						"imagePullPolicy": "Always",
					}},
				},
			},
		},
	}

	// Create the Tekton TaskRun in the cluster.
	fmt.Println("Creating TaskRun...")
	namespace := "default"
	result, err := client.Resource(taskRunResource).Namespace(namespace).Create(taskRun, metav1.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created TaskRun %q.\n", result.GetName())
}
