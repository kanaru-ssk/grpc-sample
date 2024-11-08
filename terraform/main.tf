provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_cloud_run_v2_service" "default" {
  name                = "cloudrun-service"
  location            = var.region
  deletion_protection = false
  ingress             = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "us-docker.pkg.dev/cloudrun/container/hello"
    }
  }
}