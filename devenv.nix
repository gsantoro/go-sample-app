{ pkgs, ... }:

{
  # https://devenv.sh/packages/
  packages = [ 
    pkgs.git 
    pkgs.govulncheck
    pkgs.air
    pkgs.gofumpt
    pkgs.golangci-lint    
    pkgs.changie
    pkgs.go-swag
    pkgs.pre-commit
    pkgs.gocyclo
    pkgs.dive
    pkgs.docker-slim
    pkgs.curl
    pkgs.skaffold
    pkgs.trivy
    pkgs.kubernetes-helm
    pkgs.go-task
    pkgs.k9s
    # pkgs.kube3d # Note: Currently at 5.4.4 while github at 5.4.9
  ];

  devcontainer.enable = true;
  devcontainer.settings.image = "ghcr.io/gsantoro/devenv:latest";

  languages.go.enable = true;

  # https://devenv.sh/processes/
  # Run them with `devenv up` from root
  # Note: disable those port forward command for services that are not running
  # processes.grafana-pf.exec = "task grafana:pf";
  # processes.prometheus-pf.exec = "task prometheus:pf";
  processes.traefik-pf.exec = "task k3d:traefik:pf";

  # See full reference at https://devenv.sh/reference/options/
}
