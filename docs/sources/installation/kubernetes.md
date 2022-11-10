+++
title = "Deploy Plutono on Kubernetes"
description = "Guide for deploying Plutono on Kubernetes"
keywords = ["plutono", "configuration", "documentation", "kubernetes"]
weight = 700
+++

## Deploy Plutono on Kubernetes

This page explains how to install and run Plutono on Kubernetes (K8S). It uses Kubernetes manifests for the setup. If you prefer Helm, refer to the [Plutono Helm community charts](https://github.com/grafana/helm-charts). 

If you are interested in Plutono Enterprise (not Plutono OS), jump to [Deploy Plutono Enterprise on Kubernetes](#deploy-plutono-enterprise-on-kubernetes) section.

### Create Plutono Kubernetes manifest
1. Create a file called `plutono.yaml`, then paste the contents below. 

```yaml
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: plutono-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: plutono
  name: plutono
spec:
  selector:
    matchLabels:
      app: plutono
  template:
    metadata:
      labels:
        app: plutono
    spec:
      securityContext:
        fsGroup: 472
        supplementalGroups:
        - 0    
      containers:
        - name: plutono
          image: plutono/plutono:7.5.2
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 3000
              name: http-plutono
              protocol: TCP
          readinessProbe:
            failureThreshold: 3
            httpGet:
              path: /robots.txt
              port: 3000
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            successThreshold: 1
            timeoutSeconds: 2
          livenessProbe:
            failureThreshold: 3
            initialDelaySeconds: 30
            periodSeconds: 10
            successThreshold: 1
            tcpSocket:
              port: 3000
            timeoutSeconds: 1            
          resources:
            requests:
              cpu: 250m
              memory: 750Mi
          volumeMounts:
            - mountPath: /var/lib/plutono
              name: plutono-pv
      volumes:
        - name: plutono-pv
          persistentVolumeClaim:
            claimName: plutono-pvc
---
apiVersion: v1
kind: Service
metadata:
  name: plutono
spec:
  ports:
    - port: 3000
      protocol: TCP
      targetPort: http-plutono
  selector:
    app: plutono
  sessionAffinity: None
  type: LoadBalancer
```


### Send manifest to Kubernetes API server

1. Run the following command: 
`kubectl apply -f plutono.yaml`

1. Check that it worked by running the following:
`kubectl port-forward service/plutono 3000:3000`

1. Navigate to `localhost:3000` in your browser. You should see a Plutono login page. 

1. Use `admin` for both the username and password to login.

## Deploy Plutono Enterprise on Kubernetes
The process for deploying Plutono Enterprise is almost identical to the process above, except for some extra steps required to add in your license file. They are described in the following sections.

### Obtain Plutono Enterprise license
To run Plutono Enterprise, you need a valid license. [Contact a Grafana Labs representative](https://grafana.com/contact?about=plutono-enterprise) to obtain the license. This topic assumes that you already have done this and have a `license.jwt` file. Your license should also be associated with a URL, which we will use later in the topic. 

### Create License Secret
Create a Kubernetes secret from your license file using the following command:
```bash
kubectl create secret generic ge-license --from-file=/path/to/your/license.jwt
```

### Create Plutono Enterprise configuration
Create a Plutono configuration file with the name `plutono.ini`. Then paste the content below. 
>**Note:** You will have to update the `root_url` field to the url associated with the license you were given. 
```yaml
[enterprise]
license_path = /etc/plutono/license/license.jwt
[server]
root_url =/your/license/root/url

```

### Create Configmap for Plutono Enterprise Config
Create a Kubernetes Configmap from your `plutono.ini` file with the following command:
```bash
kubectl create configmap ge-config --from-file=/path/to/your/config.ini
```
### Create Plutono Enterprise Kubernetes manifest
Create a `plutono.yaml` file, then paste the content below. This YAML is identical to the one for Plutono OS install except for the additional references to the Configmap which has your Plutono configuration file and the Secret that has your license. 

```yaml
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: plutono
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: local-path
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: plutono
  name: plutono
spec:
  selector:
    matchLabels:
      app: plutono
  template:
    metadata:
      labels:
        app: plutono
    spec:
      containers:
        - image: plutono/plutono-enterprise:latest
          imagePullPolicy: IfNotPresent
          name: plutono
          ports:
            - containerPort: 3000
              name: http-plutono
              protocol: TCP
          readinessProbe:
            failureThreshold: 3
            httpGet:
              path: /robots.txt
              port: 3000
              scheme: HTTP
            initialDelaySeconds: 10
            periodSeconds: 30
            successThreshold: 1
            timeoutSeconds: 2
          resources:
            limits:
              memory: 4Gi
            requests:
              cpu: 100m
              memory: 2Gi
          volumeMounts:
            - mountPath: /var/lib/plutono
              name: plutono
            - mountPath: /etc/plutono
              name: ge-config
            - mountPath: /etc/plutono/license
              name: ge-license
      volumes:
        - name: plutono
          persistentVolumeClaim:
            claimName: plutono
        - name: ge-config
          configMap:
            name: ge-config
        - name: ge-license
          secret:
            secretName: ge-license
---
apiVersion: v1
kind: Service
metadata:
  name: plutono
spec:
  ports:
    - port: 3000
      protocol: TCP
      targetPort: http-plutono
  selector:
    app: plutono
  sessionAffinity: None
  type: LoadBalancer
```
 
1. Send manifest to Kubernetes API Server
`kubectl apply -f plutono.yaml`

1. Check that it worked by running the following:
`kubectl port-forward service/plutono 3000:3000`

1. Navigate to `localhost:3000` in your browser. You should see the Plutono login page. 

1. Use `admin` for both the username and password to login.
If it worked, you should see `Enterprise (Licensed)` at the bottom of the page. 
