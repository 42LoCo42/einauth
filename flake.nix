{
  outputs = { flake-utils, nixpkgs, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; }; in rec {
        packages.default = pkgs.buildGoModule {
          pname = "einauth";
          version = "1";
          src = ./.;
          vendorHash = "sha256-P0dKaqav0Mvfq3xmpxhq1Z7FnGH+gtlAFWiUx+EZgII=";

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
