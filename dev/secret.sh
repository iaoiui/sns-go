kubectl delete secret awssns-secret
kubectl create secret generic --from-env-file .env awssns-secret