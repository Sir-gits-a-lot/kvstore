# kvstore
kvstore application implemented in Golang with all the Dockerfile and Kubernetes Manifests, can be tested in Kind Cluster

# Commands to reproduce in local kind based environment

1. Create a kind cluster using config file 
kind create cluster --config kind-cluster.yaml

2. Deploy Ingress to the kind cluster
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

3. Wait for the pod to get ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

4. Build docker image with tag ttl.sh/kvstore:24h (Useful for pushing to ephemeral registry)
cd kvstore
docker build -t ttl.sh/kvstore:24h . --no-cache

5. Apply Kubernetes manifest with the built image
cd ../kube_manifests
kubectl apply -f kube_manifests/


# To implement Observability for this service:- 

I will start by intrumenting the code using opentelemetry packages for instrumenting the code to generate telemetry data. Some of the libraries include these packages:-

"go.opentelemetry.io/otel/sdk/log"
"go.opentelemetry.io/otel/sdk/metric"
"go.opentelemetry.io/otel/sdk/trace"

Then, I will deploy the Opentelemetry collector into my kubernetes cluster and have different observability backends like Prometheus for storing Metrics, FluentBit as a DaemonSet for Logging, Distributed Tracing agent using Jaeger and ElasticSearch recent Continous Profiling Agent.

I will also make use of Kubernetes-monitoring dashboards available at Grafana.com and include different Data Sources like Prometheus to visualize the metrics and draws useful conclusions and even actions based on certain Alarms with the help of Alert-Manager for Closed-loop Service Assurance.

# Other tests to be performed before releasing the service to production:-

I will run some Chaos Engineering experiments on the Kubernetes deployments using tools like LitmusChaos, Gremlin and ChaosMonkey. 
I will simulte disaster scenario which can give me an idea on how to solve for challenges of quick replication and disaster recovery. 