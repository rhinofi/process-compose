{ lib, flake-parts-lib, ... }:
let
  inherit (flake-parts-lib)
    mkPerSystemOption;
  inherit (lib)
    mdDoc
    mkOption
    types;
in
{
  options.perSystem = mkPerSystemOption ({ config, pkgs, lib, ... }:
    {
      options.process-compose = mkOption {
        description = mdDoc ''
          process-compose-flake: creates [process-compose](https://github.com/F1bonacc1/process-compose)
          executables from process-compose configurations written as Nix attribute sets.
        '';
        type = types.attrsOf (types.submoduleWith {
          specialArgs = { inherit pkgs; };
          modules = [
            ./process-compose
          ];
        });
      };

      config = {
        packages = lib.concatMapAttrs
          (name: cfg: {
            "${name}" = cfg.outputs.package;
            "${name}-up" = cfg.outputs.package-up;
          })
          config.process-compose;
        checks =
          let
            checks' = lib.mapAttrs
              (name: cfg: cfg.outputs.check)
              config.process-compose;
          in
          lib.filterAttrs (_: v: v != null) checks';
      };
    });
}

