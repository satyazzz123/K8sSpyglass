# K8sSpyglass

K8sSpyglass is an advanced Kubernetes cluster analysis tool designed to provide in-depth insights into your cluster's health and performance. Whether you're managing a small deployment or a complex microservices architecture, K8sSpyglass empowers you to quickly understand and optimize your Kubernetes environment.

## Installation

To start using K8sSpyglass, please do follow the below installation methods:

1. **Build from Source:**
   ```bash
   go build -o spyk8
    ```

 ## Features
- Comprehensive Cluster Overview:
Gain a holistic understanding of your Kubernetes cluster's status, identifying critical information and potential issues.

- Namespace-specific Analysis:
Tailor your analysis by specifying target namespaces, allowing for a more detailed and focused view of your cluster health.

- Flexible Logging Options:
Customize logging preferences to capture essential details during the tool's execution.      

 ## Run the Tool
 ```
./spyk8 check --kubeconfig=path/to/your/kubeconfig --logger --namespaces=namespace1,namespace2
 ```