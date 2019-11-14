provider "kubernetes" {}

resource "kubernetes_pod" "simple-pod" {
  metadata {
    name = "my-app"
    namespace = "workbench"
    labels = {
      app = "my-app"
    }
  }

  spec {
    container {
      image = "busybox"
      name  = "myapp-container"
      command = ["sh", "-c", "echo Hello Kubernetes! && sleep 3600"]
    }
  }
}   