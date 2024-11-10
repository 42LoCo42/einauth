{
  outputs = { flake-utils, nixpkgs, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; }; in rec {
        packages.default = pkgs.buildGoModule {
          pname = "einauth";
          version = "1";
          src = ./.;
          vendorHash = "sha256-+LbuIRISEiwFlf4ewzRGz6T18oSdGYEsG/9G5NiT/b8=";

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
