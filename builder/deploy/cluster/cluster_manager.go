package cluster

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"path/filepath"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	knative "knative.dev/serving/pkg/client/clientset/versioned"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/config"
)

// Cluster holds basic information about a Kubernetes cluster
type Cluster struct {
	ID            string                // Unique identifier for the cluster
	ConfigPath    string                // Path to the kubeconfig file
	Client        *kubernetes.Clientset // Kubernetes client
	KnativeClient *knative.Clientset    // Knative client
}

// ClusterPool is a resource pool of cluster information
type ClusterPool struct {
	Clusters     []Cluster
	clusterStore *database.ClusterInfoStore
}

// NodeResourceInfo struct includes details about the node's resources and region
type NodeResourceInfo struct {
	NodeName  string  `json:"node_name"`
	GPUModel  string  `json:"gpu_model"`
	Region    string  `json:"region"`
	TotalCPU  float64 `json:"total_cpu"`
	UsedCPU   float64 `json:"used_cpu"`
	TotalGPU  int64   `json:"total_gpu"`
	UsedGPU   int64   `json:"used_gpu"`
	GPUVendor string  `json:"gpu_vendor"`
}

// NewClusterPool initializes and returns a ClusterPool by reading kubeconfig files from $HOME/.kube directory
func NewClusterPool() (*ClusterPool, error) {
	pool := &ClusterPool{}
	pool.clusterStore = database.NewClusterInfoStore()

	home := homedir.HomeDir()
	kubeconfigFolderPath := filepath.Join(home, ".kube")
	kubeconfigFiles, err := filepath.Glob(filepath.Join(kubeconfigFolderPath, "config*"))
	if err != nil {
		return nil, err
	}

	if len(kubeconfigFiles) == 0 {
		slog.Error("No kubeconfig files", slog.Any("path", kubeconfigFolderPath))
	}

	for _, kubeconfig := range kubeconfigFiles {
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			return nil, err
		}
		knativeClient, err := knative.NewForConfig(config)
		if err != nil {
			slog.Error("falied to create knative client", "error", err)
			return nil, fmt.Errorf("falied to create knative client,%w", err)
		}
		id := filepath.Base(kubeconfig)
		pool.Clusters = append(pool.Clusters, Cluster{
			ID:            id,
			ConfigPath:    kubeconfig,
			Client:        client,
			KnativeClient: knativeClient,
		})
		err = pool.clusterStore.Add(context.TODO(), id, "华中区")
		if err != nil {
			slog.Error("falied to add cluster info to db", "error", err)
		}
	}

	return pool, nil
}

// SelectCluster selects the most appropriate cluster to deploy the service to
func (p *ClusterPool) GetCluster() (*Cluster, error) {
	if len(p.Clusters) == 0 {
		return nil, fmt.Errorf("no available clusters")
	}
	// Randomly choose a cluster to deploy the service to
	// to do: The cluster should be selected based on criteria such as availability, performance, load, etc.
	randomIndex := rand.Intn(len(p.Clusters))

	// Select a cluster using the random index
	selectedCluster := p.Clusters[randomIndex]
	return &selectedCluster, nil
}

// GetClusterByID retrieves a cluster from the pool given its unique ID
func (p *ClusterPool) GetClusterByID(ctx context.Context, id string) (*Cluster, error) {
	cfId := "config"
	if len(id) != 0 {
		cInfo, _ := p.clusterStore.ByClusterID(ctx, id)
		cfId = cInfo.ClusterConfig
	}
	for _, Cluster := range p.Clusters {
		if Cluster.ID == cfId {
			return &Cluster, nil
		}
	}
	return nil, fmt.Errorf("cluster with the given ID does not exist")
}

// getNodeResources retrieves all node cpu and gpu info
func GetNodeResources(clientset *kubernetes.Clientset, config *config.Config) (map[string]NodeResourceInfo, error) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	nodeResourcesMap := make(map[string]NodeResourceInfo)

	for _, node := range nodes.Items {
		totalCPU := node.Status.Capacity.Cpu().MilliValue()
		allocatableCPU := node.Status.Allocatable.Cpu().MilliValue()
		totalGPU, found := node.Status.Capacity["nvidia.com/gpu"]
		if !found {
			totalGPU = resource.Quantity{}
		}

		region := node.Labels[config.Space.NodeRegion]
		gpuModelVendor := strings.Split(node.Labels[config.Space.GPUModelLablel], "-")
		gpuModel := ""
		if len(gpuModelVendor) > 1 {
			gpuModel = gpuModelVendor[1]
		}
		nodeResourcesMap[node.Name] = NodeResourceInfo{
			NodeName:  node.Name,
			Region:    region,
			TotalCPU:  millicoresToCores(totalCPU),
			UsedCPU:   millicoresToCores(allocatableCPU),
			GPUModel:  gpuModel,
			GPUVendor: gpuModelVendor[0],
			TotalGPU:  parseQuantityToInt64(totalGPU),
		}
	}

	for _, pod := range pods.Items {
		if pod.Spec.NodeName == "" || pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
			continue
		}

		nodeResource := nodeResourcesMap[pod.Spec.NodeName]
		for _, container := range pod.Spec.Containers {
			if requestedGPU, hasGPU := container.Resources.Requests["nvidia.com/gpu"]; hasGPU {
				nodeResource.UsedGPU += parseQuantityToInt64(requestedGPU)
			}
		}

		nodeResourcesMap[pod.Spec.NodeName] = nodeResource
	}

	return nodeResourcesMap, nil
}

func millicoresToCores(millicores int64) float64 {
	cores := float64(millicores) / 1000.0
	return math.Round(cores*10) / 10
}

func parseQuantityToInt64(q resource.Quantity) int64 {
	if q.IsZero() {
		return 0
	}
	value, _ := q.AsInt64()
	return value
}
