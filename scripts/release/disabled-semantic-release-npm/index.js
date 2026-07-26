/**
 * What: Reject every lifecycle hook exposed by the unused npm publish plugin.
 *
 * Why: haxe.go uses semantic-release only to select a version and create the
 * exact tested tag. The upstream meta-package depends on its npm publisher even
 * when an explicit plugin list excludes it, which otherwise installs a complete
 * unused npm CLI and its bundled vulnerability surface.
 *
 * How: Satisfy semantic-release's package dependency with this local,
 * deterministic sentinel. If configuration drift ever activates an npm publish
 * hook, fail before it can mutate a registry or the tested checkout.
 */
async function disabledNpmPublishHook() {
  throw new Error(
    "@semantic-release/npm is disabled by haxe.go release policy; releases publish only reviewed GitHub and Haxelib assets"
  );
}

export {
  disabledNpmPublishHook as addChannel,
  disabledNpmPublishHook as prepare,
  disabledNpmPublishHook as publish,
  disabledNpmPublishHook as verifyConditions,
};
