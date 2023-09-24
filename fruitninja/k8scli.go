package fruitninja

import (
	"context"
	"os"

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
		Host:        fruitNinjaConfig.K8SAPI,
		BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6InRBb1JyNzRaa3VYZmV6cmk4bHZybGJZcjVpOGN4cDhCSEtCdEJQMnp1RWMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXYiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zZWNyZXQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGV2LWNvbnRhaW5lci1zYSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjExMWY0OGRmLTFkYWEtNDljOS1hMzIzLTI0Nzc3ZWE0Y2U0ZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXY6ZGV2LWNvbnRhaW5lci1zYSJ9.o0ZKu_ziOO3-GJ_kzYDnNq3UslhjRkue0TJWFAC9wgAgndhQi37r6-HwtMx3syHnC8Q5sNdG_Df0vYAKSH5PjgA2RqbMIoOWUwRxEDIwNBHHZ9xJrOu4gCZoxWqHgBskmjsqE5zVw5D6ksltAEZKFke15t2NlYuiiaz1Mj9mcEdUk7ryo5Z18VGKe6lsdbqfu_6GkUvN5NvzvoZcSrnc6VTGxuBV_c1Mfhk0lJpIlzEZjjDCpi6w-V3aH1oIJE5xmBxSOo9i8GRCV1SmEMsOErF9Qsc2QRwIiuIe4R4ALS-xSxqbrDBEAnI95feZDlsJU8yrqMsm0zxpkpHWSHQ13Q",
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

	currentNS, err := getNamespace()
	if err != nil {
		log.Error("get current namespace error")
		os.Exit(1)
	}
	log.Infof("current namespace: %s\n", currentNS)

	services, err := clientset.CoreV1().Services(currentNS).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, v := range services.Items {
		svcs = append(svcs, v.Name)
	}
	return svcs
}
