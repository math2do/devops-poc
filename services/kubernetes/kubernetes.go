package kubernetes

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// K8sClient ...
type K8sClient struct {
	ClientSet *kubernetes.Clientset
}

// Client ...
var Client *K8sClient

// NewK8sClient ...
func NewK8sClient(log *log.Logger) *K8sClient {
	if Client == nil {
		err := InitKubernetesClient()
		if err != nil {
			log.Fatalf("unable to init k8s client: %v", err)
		}
	}

	return Client
}

// InitKubernetesClient ...
func InitKubernetesClient() error {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Load config from kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	Client = &K8sClient{
		ClientSet: clientset,
	}

	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
