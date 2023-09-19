package fruitninja

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getK8SService(namespace string) []string {
	svcs := []string{}

	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	fmt.Println("===============================")
	// 	fmt.Println("error")
	// 	// panic(err.Error())
	// }

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, v := range services.Items {
		svcs = append(svcs, v.Name)
	}
	return svcs
}
