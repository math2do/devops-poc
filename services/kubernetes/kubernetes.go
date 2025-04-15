package kubernetes

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// K8sClient ...
type K8sClient struct {
	ClientSet *kubernetes.Clientset
	log       *log.Logger
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

	Client.log = log
	return Client
}

// InitKubernetesClient ...
func InitKubernetesClient() error {
	var kubeconfig *string
	var config *rest.Config
	var err error

	env := os.Getenv("ENV")

	if env == "PROD" {
		// Set up in-cluster config
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("error building config: %v", err)
			return err
		}

	} else {
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// Load config from kubeconfig file
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			log.Fatalf("error building config: %v", err)
			return err
		}
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("error building clientset: %v", err)
		return nil
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

// LaunchK8sJob ...
func (c *K8sClient) LaunchK8sJob(jobName string, image string, cmd []string, args []string, env []v1.EnvVar) error {
	jobs := c.ClientSet.BatchV1().Jobs("default")
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    jobName,
							Image:   image,
							Command: cmd,
							Args:    args,
							Env:     env,
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		c.log.Printf("Failed to create K8s job.")
		return err
	}

	//print job details
	c.log.Println("Created K8s job successfully")
	return nil
}
