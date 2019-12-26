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
    "os"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

type LauncherConfig struct {
	Image   string   `yaml:"image,omitempty"`
	Command []string `yaml:"command,omitempty"`
	Args    []string `yaml:"args,omitempty"`
}

func Launch(args []string) {

	// Verify that at least one arg has been supplied.
	if len(args) < 1 {
        log.Fatal("Must supply path to Launcher YAML as first argument.")
	}

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

	// Read target Launcher YAML and store into struct.
	launcherConfig := LauncherConfig{}

    // Get absolute path of the first argument.
	abs, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatal(err)
	}

    // Open the file as read-only and close after reading contents into memory.
    file, err := os.OpenFile(abs, os.O_RDWR, 0444)
    if err != nil {
        log.Fatal(err)
    }
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
    file.Close()

    // Unmarshal the raw data (YAML format) into the struct.
	err = yaml.Unmarshal(data, &launcherConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Construct the TaskRun object for creation.
	taskRun := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1alpha1",
			"kind":       "TaskRun",
			"metadata": map[string]interface{}{
				"name": "launched-taskrun",
			},
			"spec": map[string]interface{}{
				"taskSpec": map[string]interface{}{
					"steps": []map[string]interface{}{{
						"name":            "taskcontainer",
						"image":           launcherConfig.Image,
						"command":         launcherConfig.Command,
						"args":            launcherConfig.Args,
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
