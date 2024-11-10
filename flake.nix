{
  outputs = { flake-utils, nixpkgs, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; }; in rec {
        packages.default = pkgs.buildGoModule {
          pname = "einauth";
          version = "1";
          src = ./.;
          vendorHash = "sha256-4kIaKf1Noq5SwBVk8BGxaJcOSHtmPcmF+G7VG1hzOgw=";

          CGO_ENABLED = "0";
          ldflags = [ "-s" "-w" ];
        };

        devShells.default = pkgs.mkShell {
          inputsFrom = [ packages.default ];
          packages = with pkgs; [
            sqlite-interactive
          ];
        };
      });
}
