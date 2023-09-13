package fruitninja

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var config *rest.Config = &rest.Config{
	Host:        "https://10.0.0.112:42883",
	BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkljRW5yVXlJRG9jSWt4STJucXJZZlB0cEt2bWJDbmo2NXdYZXo1ZXNCdnMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Imdsb2JhbC1yZWFkZXItdG9rZW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZ2xvYmFsLXJlYWRlciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6Ijg5Y2I5MTRhLWM3MTEtNDhlNS1hMjgzLWIyYzc5MjE0NjFjZSIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0Omdsb2JhbC1yZWFkZXIifQ.XO41-knsGg0Ami3ZR-kBALTnKrwYb-MaHyEHjS7AbhQgNTIUdLvzP9nT4-AqSiM4nVJOarDgPyufFadDEjdk1P6A0Mcbu2NVcI80ItDdjm9SDsbw8TzGCW5Jm0rDmAlLrwTv6FnpMuFNd25Cf_YWwbjYVXM7Zmhbhi_dpbFJ0sG_qS8-r-45-vhqNOrWr1DGCsKdHBG2ZSulrc9I84gLT43buRuPiB-59_MkJUVqWTNB0Fnxn2cictETx30bM_pqpCcil4QCjK8MFYIskjmcJ6nZXZGV-YN5XGyA6jDPzHA7hCrL0rSIBn9NLH2e1EsiZEFF6-gjQKEXaVPedAXnUQ",
	TLSClientConfig: rest.TLSClientConfig{
		Insecure: true,
	},
}

func getK8SService(namespace string) []string {
	svcs := []string{}
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
