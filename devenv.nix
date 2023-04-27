{ pkgs, ... }:

{
  # https://devenv.sh/basics/
  # env.GREET = "devenv";

  # https://devenv.sh/packages/
  packages = [ 
    pkgs.git 
    pkgs.govulncheck
    pkgs.air
    pkgs.gofumpt
    pkgs.golangci-lint    
    pkgs.changie
    pkgs.kind
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
  ];

  # devcontainer.enable = true;

  languages.go.enable = true;

  # https://devenv.sh/processes/
  processes.grafana-pf.exec = "task grafana:pf";
  processes.prometheus-pf.exec = "task prometheus:pf";

  # See full reference at https://devenv.sh/reference/options/
}
