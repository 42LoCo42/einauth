{
  outputs = { flake-utils, nixpkgs, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; }; in rec {
        packages.default = pkgs.buildGoModule {
          pname = "einauth";
          version = "1";
          src = ./.;
          vendorHash = "sha256-YQ6RiGBWD03rRu3MWm2HSxmYzfmOYiYMV4sPM3Y0H3o=";

          CGO_ENABLED = "0";
          ldflags = [ "-s" "-w" ];
        };

        devShells.default = pkgs.mkShell {
          inputsFrom = [ packages.default ];
          packages = with pkgs; [
            air
            caddy
            sqlite-interactive
          ];
        };
      });
}
