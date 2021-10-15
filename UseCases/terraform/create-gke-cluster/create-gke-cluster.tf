
variable "projectmap" {}
resource "google_service_account" "default" {
  account_id   = "gke-service-account-id"
  display_name = "GKE Service Account"
  project      = var.projectmap.project
}

resource "google_container_cluster" "primary" {
  name     = "kambui-gke-cluster"
  location = "us-east1-c"
  project  = var.projectmap.project

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "kambui-node-pool"
  location   = "us-east1-c"
  project    = google_container_cluster.primary.project
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "e2-medium"

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.default.email
    oauth_scopes    = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
