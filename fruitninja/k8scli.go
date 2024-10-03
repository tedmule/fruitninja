package fruitninja

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type minionk8s struct {
	client *kubernetes.Clientset
}

func initMinionK8S(mode string) (*kubernetes.Clientset, error) {
	client := &kubernetes.Clientset{}

	if mode == "outofcluster" {
		home, err := os.UserHomeDir()
		if err != nil {
			home = "~"
		}

		// Use ~/.kube/config
		kubeconfig := filepath.Join(home, ".kube", "config")
		configFromFlags, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %s\n", err.Error())
		}

		client, err = kubernetes.NewForConfig(configFromFlags)
		if err != nil {
			log.Fatalf("Error building Kubernetes client: %s\n", err.Error())
		}
	}
	if mode == "incluster" {
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Error building in cluster config: %s\n", err.Error())
		}
		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("Error building Kubernetes client: %s\n", err.Error())
		}
		return client, nil
	}
	return client, nil
}

// func getK8SService(namespace string) []string {
func getK8SService() []string {
	svcs := []string{}

	// k8sConfig, err := rest.InClusterConfig()
	// if err != nil {
	// log.Errorf("Get InClusterConfig Failed: %+v\n", err)
	k8sConfig = &rest.Config{
		Host:        fruitNinjaSettings.K8SAPI,
		BearerToken: fruitNinjaSettings.K8SToken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	// }

	// create the clientset
	clientset, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		panic(err.Error())
	}

	// ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), "", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	currentNS := getNamespace()
	log.Infof("current namespace: %s\n", currentNS)

	services, err := clientset.CoreV1().Services(currentNS).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error(err.Error())
		return svcs
	}

	for _, v := range services.Items {
		svcs = append(svcs, v.Name)
	}
	return svcs
}

func queryService(name string, namespace string) {
	fmt.Println(name)
	fmt.Println(namespace)
}
