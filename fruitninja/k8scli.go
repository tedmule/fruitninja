package fruitninja

import (
	"context"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var config *rest.Config = &rest.Config{
	// Host:        "https://10.0.0.112:42883",
	Host:        os.Getenv("K8S_API"),
	BearerToken: "ZXlKaGJHY2lPaUpTVXpJMU5pSXNJbXRwWkNJNklsOHdjbkZHUVZjeVFVczBiWGRHWmtodlQyWmhOMDF2VG5STWF6Qm5NSHBEYUhsSGFISkJjRXRxYlhNaWZRLmV5SnBjM01pT2lKcmRXSmxjbTVsZEdWekwzTmxjblpwWTJWaFkyTnZkVzUwSWl3aWEzVmlaWEp1WlhSbGN5NXBieTl6WlhKMmFXTmxZV05qYjNWdWRDOXVZVzFsYzNCaFkyVWlPaUprWlhZaUxDSnJkV0psY201bGRHVnpMbWx2TDNObGNuWnBZMlZoWTJOdmRXNTBMM05sWTNKbGRDNXVZVzFsSWpvaVpHVjJMV052Ym5SaGFXNWxjaTF6WldOeVpYUWlMQ0pyZFdKbGNtNWxkR1Z6TG1sdkwzTmxjblpwWTJWaFkyTnZkVzUwTDNObGNuWnBZMlV0WVdOamIzVnVkQzV1WVcxbElqb2laR1YyTFdOdmJuUmhhVzVsY2kxellTSXNJbXQxWW1WeWJtVjBaWE11YVc4dmMyVnlkbWxqWldGalkyOTFiblF2YzJWeWRtbGpaUzFoWTJOdmRXNTBMblZwWkNJNklqUXlabUV4WmpCbUxUVmhaV1V0TkdZeE15MDVORGN4TFdVMVlUTmxOVE16TkRSbU1pSXNJbk4xWWlJNkluTjVjM1JsYlRwelpYSjJhV05sWVdOamIzVnVkRHBrWlhZNlpHVjJMV052Ym5SaGFXNWxjaTF6WVNKOS5DdjVXRlFKZXg4bk13YnM3ZW5sY1phRHNrNWJJZS1kSG1iYnN4U0pWbkVpeUEzUEF3RUlfWlc5Unk4N3JGQkxSYUpBdnZhNzVIanJlbDczNEZCWmhPdUJ2ZE1sYVYxU1Z4UUc2TGxfVERZXzY1cTBEOU5TNTc4cVQ5Y2JZcXBIbFZZdGh6WjZpSzJDRXNFNi05dDJ6empuZ1lHSk5MZUlWWGZUcjN0WEJGSGMzMmJ4NDhyQS1SUUJGOW93VWZjMHJmb1NHTE93d3Z5UV9faEl4aFVESlFTYjYzeFFhelF5b3g4enBBUGpCWFFYdVRCRF9pZ1hsYWs5MzJoTjJNWnNESzlXZENnVS1KdmY4TjR4RWZTcmNDRkNsd084Tjh5UTdmdnB1eXVoSDdJSTZaN1RCbmx5WW9adWI3T09qcTUxS3Q3bmtTSHh2Y2lyVTJRTTd6RWlfbXc=",
	TLSClientConfig: rest.TLSClientConfig{
		Insecure: true,
	},
}

func getK8SService(namespace string) []string {
	svcs := []string{}

	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
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
