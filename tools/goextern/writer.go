package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var ownershipKeyPattern = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

type ownedOutput struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}

type ownershipManifest struct {
	SchemaVersion int           `json:"schemaVersion"`
	RootKey       string        `json:"rootKey"`
	Root          ownershipRoot `json:"root"`
	Files         []ownedOutput `json:"files"`
}

func writeEmission(emission *Emission) error {
	if emission == nil {
		return errors.New("nil emission")
	}
	if strings.TrimSpace(emission.OutputDir) == "" {
		return errors.New("empty output directory")
	}
	if strings.TrimSpace(emission.RootKey) == "" {
		return graphError("missing_root_identity", "emission has no root ownership key")
	}
	if !ownershipKeyPattern.MatchString(emission.RootKey) {
		return graphError("invalid_root_identity", "root ownership key %q is not a safe file token", emission.RootKey)
	}

	planned, contents, err := plannedOutputs(emission.Files)
	if err != nil {
		return err
	}
	manifestDir := filepath.Join(emission.OutputDir, ".goextern", "roots")
	if err := validatePathTarget(emission.OutputDir, filepath.Join(".goextern", "roots")); err != nil {
		return err
	}
	manifests, err := loadOwnershipManifests(manifestDir)
	if err != nil {
		return err
	}
	current := manifests[emission.RootKey]
	if current == nil {
		current = &ownershipManifest{SchemaVersion: 1, RootKey: emission.RootKey, Root: emission.Root}
	}
	if err := validateOwnershipPlan(emission.OutputDir, emission.RootKey, planned, manifests, current); err != nil {
		return err
	}

	if err := os.MkdirAll(manifestDir, 0o755); err != nil {
		return fmt.Errorf("create goextern ownership directory %q: %w", manifestDir, err)
	}
	for _, output := range planned {
		target := filepath.Join(emission.OutputDir, filepath.FromSlash(output.Path))
		if err := validateOutputTarget(emission.OutputDir, output.Path); err != nil {
			return err
		}
		if err := writeFileAtomically(target, contents[output.Path], 0o644); err != nil {
			return err
		}
	}

	otherClaims := ownershipClaims(manifests, emission.RootKey)
	plannedPaths := make(map[string]bool, len(planned))
	for _, output := range planned {
		plannedPaths[output.Path] = true
	}
	for _, old := range current.Files {
		if plannedPaths[old.Path] || otherClaims[old.Path] {
			continue
		}
		target := filepath.Join(emission.OutputDir, filepath.FromSlash(old.Path))
		if err := validateOutputTarget(emission.OutputDir, old.Path); err != nil {
			return err
		}
		if err := os.Remove(target); err != nil && !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("remove stale owned output %q: %w", target, err)
		}
	}

	next := ownershipManifest{
		SchemaVersion: 1,
		RootKey:       emission.RootKey,
		Root:          emission.Root,
		Files:         planned,
	}
	payload, err := json.MarshalIndent(next, "", "\t")
	if err != nil {
		return fmt.Errorf("marshal goextern ownership manifest: %w", err)
	}
	payload = append(payload, '\n')
	manifestRelativePath := filepath.Join(".goextern", "roots", emission.RootKey+".json")
	if err := validatePathTarget(emission.OutputDir, manifestRelativePath); err != nil {
		return err
	}
	return writeFileAtomically(filepath.Join(emission.OutputDir, manifestRelativePath), payload, 0o644)
}

func plannedOutputs(files []EmittedFile) ([]ownedOutput, map[string][]byte, error) {
	byPath := make(map[string]ownedOutput, len(files))
	contents := make(map[string][]byte, len(files))
	caseFolded := make(map[string]string, len(files))
	for _, file := range files {
		path, err := safeRelativeOutputPath(file.Name)
		if err != nil {
			return nil, nil, err
		}
		folded := strings.ToLower(path)
		if other, exists := caseFolded[folded]; exists && other != path {
			return nil, nil, graphError("output_path_collision", "%q and %q differ only by case", other, path)
		}
		caseFolded[folded] = path
		payload := []byte(file.Contents)
		digest := sha256.Sum256(payload)
		if existing, exists := contents[path]; exists && string(existing) != string(payload) {
			return nil, nil, graphError("output_path_collision", "two generated files target %q with different contents", path)
		}
		contents[path] = payload
		byPath[path] = ownedOutput{Path: path, SHA256: hex.EncodeToString(digest[:])}
	}
	outputs := make([]ownedOutput, 0, len(byPath))
	for _, output := range byPath {
		outputs = append(outputs, output)
	}
	sort.Slice(outputs, func(i, j int) bool { return outputs[i].Path < outputs[j].Path })
	return outputs, contents, nil
}

func safeRelativeOutputPath(path string) (string, error) {
	if strings.TrimSpace(path) == "" || filepath.IsAbs(path) {
		return "", graphError("unsafe_output_path", "output path %q must be relative", path)
	}
	clean := filepath.Clean(path)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return "", graphError("unsafe_output_path", "output path %q escapes the output root", path)
	}
	if filepath.Ext(clean) != ".hx" {
		return "", graphError("unsafe_output_path", "output path %q is not a Haxe source file", path)
	}
	return filepath.ToSlash(clean), nil
}

func loadOwnershipManifests(dir string) (map[string]*ownershipManifest, error) {
	out := make(map[string]*ownershipManifest)
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return out, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read goextern ownership directory %q: %w", dir, err)
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		if entry.Type()&os.ModeSymlink != 0 {
			return nil, graphError("invalid_ownership_manifest", "%q is a symbolic link", path)
		}
		payload, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read goextern ownership manifest %q: %w", path, err)
		}
		var manifest ownershipManifest
		if err := json.Unmarshal(payload, &manifest); err != nil {
			return nil, graphError("invalid_ownership_manifest", "%q: %v", path, err)
		}
		if manifest.SchemaVersion != 1 || manifest.RootKey == "" {
			return nil, graphError("invalid_ownership_manifest", "%q has unsupported identity or schema", path)
		}
		if !ownershipKeyPattern.MatchString(manifest.RootKey) || entry.Name() != manifest.RootKey+".json" {
			return nil, graphError("invalid_ownership_manifest", "%q has an unsafe or mismatched root key", path)
		}
		if _, exists := out[manifest.RootKey]; exists {
			return nil, graphError("invalid_ownership_manifest", "%q duplicates root %q", path, manifest.RootKey)
		}
		seenPaths := make(map[string]bool, len(manifest.Files))
		for _, output := range manifest.Files {
			clean, err := safeRelativeOutputPath(output.Path)
			if err != nil || clean != output.Path || seenPaths[output.Path] {
				return nil, graphError("invalid_ownership_manifest", "%q contains invalid output path %q", path, output.Path)
			}
			digest, err := hex.DecodeString(output.SHA256)
			if err != nil || len(digest) != sha256.Size {
				return nil, graphError("invalid_ownership_manifest", "%q contains invalid hash for %q", path, output.Path)
			}
			seenPaths[output.Path] = true
		}
		out[manifest.RootKey] = &manifest
	}
	return out, nil
}

func validateOwnershipPlan(outputRoot string, rootKey string, planned []ownedOutput, manifests map[string]*ownershipManifest, current *ownershipManifest) error {
	allClaims := make(map[string]ownedOutput)
	allClaimPathsByFold := make(map[string]string)
	claims := make(map[string]ownedOutput)
	claimPathsByFold := make(map[string]string)
	manifestRoots := make([]string, 0, len(manifests))
	for manifestRoot := range manifests {
		manifestRoots = append(manifestRoots, manifestRoot)
	}
	sort.Strings(manifestRoots)
	for _, otherRoot := range manifestRoots {
		manifest := manifests[otherRoot]
		for _, output := range manifest.Files {
			folded := strings.ToLower(output.Path)
			if other, exists := allClaimPathsByFold[folded]; exists && other != output.Path {
				return graphError("owned_output_conflict", "%q and %q differ only by case across roots", other, output.Path)
			}
			if claim, exists := allClaims[output.Path]; exists && claim.SHA256 != output.SHA256 {
				return graphError("owned_output_conflict", "%q has different recorded contents across roots", output.Path)
			}
			allClaimPathsByFold[folded] = output.Path
			allClaims[output.Path] = output
			if otherRoot == rootKey {
				continue
			}
			claimPathsByFold[folded] = output.Path
			claims[output.Path] = output
		}
	}
	currentByPath := make(map[string]ownedOutput, len(current.Files))
	for _, output := range current.Files {
		currentByPath[output.Path] = output
	}
	for _, output := range planned {
		if other, exists := claimPathsByFold[strings.ToLower(output.Path)]; exists && other != output.Path {
			return graphError("owned_output_conflict", "%q conflicts with owned path %q by case", output.Path, other)
		}
		if claim, exists := claims[output.Path]; exists && claim.SHA256 != output.SHA256 {
			return graphError("owned_output_conflict", "%q is owned by another root with different contents", output.Path)
		}
		target := filepath.Join(outputRoot, filepath.FromSlash(output.Path))
		if err := validateOutputTarget(outputRoot, output.Path); err != nil {
			return err
		}
		payload, err := os.ReadFile(target)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return fmt.Errorf("read existing output %q: %w", target, err)
		}
		digest := sha256.Sum256(payload)
		actual := hex.EncodeToString(digest[:])
		old, ownedHere := currentByPath[output.Path]
		claim, ownedElsewhere := claims[output.Path]
		if !ownedHere && !ownedElsewhere {
			return graphError("unowned_output_conflict", "%q already exists without a goextern owner", output.Path)
		}
		if ownedElsewhere && actual != claim.SHA256 {
			return graphError("owned_output_modified", "%q differs from the contents recorded by another root", output.Path)
		}
		if ownedHere && actual != old.SHA256 && actual != output.SHA256 {
			return graphError("owned_output_modified", "%q differs from its recorded and planned contents", output.Path)
		}
	}
	for _, old := range current.Files {
		found := false
		for _, output := range planned {
			if output.Path == old.Path {
				found = true
				break
			}
		}
		if found {
			continue
		}
		target := filepath.Join(outputRoot, filepath.FromSlash(old.Path))
		if err := validateOutputTarget(outputRoot, old.Path); err != nil {
			return err
		}
		payload, err := os.ReadFile(target)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return fmt.Errorf("read stale owned output %q: %w", target, err)
		}
		digest := sha256.Sum256(payload)
		if hex.EncodeToString(digest[:]) != old.SHA256 {
			return graphError("owned_output_modified", "%q was modified and cannot be removed", old.Path)
		}
	}
	return nil
}

func ownershipClaims(manifests map[string]*ownershipManifest, excludedRoot string) map[string]bool {
	out := make(map[string]bool)
	for root, manifest := range manifests {
		if root == excludedRoot {
			continue
		}
		for _, file := range manifest.Files {
			out[file.Path] = true
		}
	}
	return out
}

func validateOutputTarget(outputRoot string, relativePath string) error {
	clean, err := safeRelativeOutputPath(relativePath)
	if err != nil {
		return err
	}
	return validatePathTarget(outputRoot, clean)
}

func validatePathTarget(outputRoot string, relativePath string) error {
	if strings.TrimSpace(relativePath) == "" || filepath.IsAbs(relativePath) {
		return graphError("unsafe_output_path", "path %q must be relative to the output root", relativePath)
	}
	clean := filepath.Clean(relativePath)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return graphError("unsafe_output_path", "path %q escapes the output root", relativePath)
	}
	rootAbs, err := filepath.Abs(outputRoot)
	if err != nil {
		return fmt.Errorf("resolve output root %q: %w", outputRoot, err)
	}
	targetAbs := filepath.Join(rootAbs, clean)
	relative, err := filepath.Rel(rootAbs, targetAbs)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return graphError("unsafe_output_path", "output path %q escapes the output root", relativePath)
	}

	cursor := rootAbs
	parts := strings.Split(filepath.FromSlash(clean), string(filepath.Separator))
	for _, part := range parts {
		cursor = filepath.Join(cursor, part)
		info, err := os.Lstat(cursor)
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("inspect output path %q: %w", cursor, err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return graphError("unsafe_output_path", "output path %q crosses a symbolic link", relativePath)
		}
	}
	return nil
}

func writeFileAtomically(target string, payload []byte, mode fs.FileMode) error {
	dir := filepath.Dir(target)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create output directory %q: %w", dir, err)
	}
	temporary, err := os.CreateTemp(dir, ".goextern-write-*")
	if err != nil {
		return fmt.Errorf("create temporary output for %q: %w", target, err)
	}
	temporaryPath := temporary.Name()
	cleanup := true
	defer func() {
		if cleanup {
			_ = os.Remove(temporaryPath)
		}
	}()
	if _, err := temporary.Write(payload); err != nil {
		_ = temporary.Close()
		return fmt.Errorf("write temporary output for %q: %w", target, err)
	}
	if err := temporary.Chmod(mode); err != nil {
		_ = temporary.Close()
		return fmt.Errorf("set temporary output mode for %q: %w", target, err)
	}
	if err := temporary.Close(); err != nil {
		return fmt.Errorf("close temporary output for %q: %w", target, err)
	}
	if err := os.Rename(temporaryPath, target); err != nil {
		return fmt.Errorf("replace output %q: %w", target, err)
	}
	cleanup = false
	return nil
}
