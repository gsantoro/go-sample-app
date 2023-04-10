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
    pkgs.trivy
    pkgs.curl
    pkgs.skaffold
  ];

  # https://devenv.sh/scripts/
  # scripts.hello.exec = "echo hello from $GREET";

  enterShell = ''
  '';

  # https://devenv.sh/languages/
  # languages.nix.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # https://devenv.sh/processes/
  # processes.ping.exec = "ping example.com";

  # See full reference at https://devenv.sh/reference/options/
}
