package fruitninja

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type kubernetesMinion struct {
	client *kubernetes.Clientset
}

func setupKubernetesClient(mode string) (*kubernetes.Clientset, error) {
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
			zap.S().Fatalf("Error building kubeconfig: %s\n", err.Error())
		}

		client, err = kubernetes.NewForConfig(configFromFlags)
		if err != nil {
			zap.S().Fatalf("Error building Kubernetes client: %s\n", err.Error())
		}
	}
	if mode == "incluster" {
		config, err := rest.InClusterConfig()
		if err != nil {
			zap.S().Fatalf("Error building in cluster config: %s\n", err.Error())
		}
		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			zap.S().Fatalf("Error building Kubernetes client: %s\n", err.Error())
		}
		return client, nil
	}
	return client, nil
}

func newKubernetesMinion(settings *FruitNinjaSettings) (*kubernetesMinion, error) {
	cs, err := setupKubernetesClient(settings.Mode)
	if err != nil {
		return nil, err
	}
	return &kubernetesMinion{
		client: cs,
	}, nil
}

// func getK8SService(namespace string) []string {
func getK8SService() []string {
	svcs := []string{}

	// k8sConfig, err := rest.InClusterConfig()
	// if err != nil {
	// zap.S().Errorf("Get InClusterConfig Failed: %+v\n", err)
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
	zap.S().Infof("current namespace: %s\n", currentNS)

	services, err := clientset.CoreV1().Services(currentNS).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Error(err.Error())
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
