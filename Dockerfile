FROM golang:1.22.3
COPY main.go .
RUN go mod init kubeacademy/k8sapp
RUN go build -o /server
CMD ["/server"]