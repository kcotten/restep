# rest endpoint
This repo is meant as an example of base components in go with tests and with deployment to Docker. A sample deployment to Kubernetes might look something like the following:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: restep-deployment
  labels:
    app: restep
spec:
  selector:
    matchLabels:
      app: restep
  template:
    metadata:
      labels:
        app: restep
    spec:
      containers:
      - name: restep
        image: restep:local
        imagePullPolicy: Never
        ports:
        - containerPort: 8000
```

You could apply the configuration with `k apply -f <filename>.yaml`.

----

> ℹ️ The following provide an overview of some of the things the currently restep includes as an example. Ideally, both will become inputs in a future pull request

## Prometheus Client
Example of exporting metrics to Prometheus that will be available at `localhost:8000/metrics` with no expectation that prometheus is running.

## Redis Client
The simple Redis rate limiter *does* expect Redis to be running and available on the default port.
