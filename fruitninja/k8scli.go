package fruitninja

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

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
