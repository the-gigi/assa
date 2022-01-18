package kubectl_in_cluster

const (
	deploymentTemplate = `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: {{ .Name }}
  name: kubectl
  namespace: {{ .Namespace }}
spec:
  selector:
    matchLabels:
      app: {{ .Name }}
  template:
    metadata:
      labels:
        app: {{ .Name }}
    spec:
      serviceAccountName: system:serviceaccounts:{{ .Namespace }}:{{ .ServiceAccount }}
      containers:
      - image: g1g1/kubectl:{{ .Version }}
        name: kubectl
`
)
