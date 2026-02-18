# Profile Admission Criteria

A profile should exist only when it has a real, testable contract.

## Admission rules

All rules below must be true:

1. Behavior difference exists in compiler/runtime code.
2. Behavior difference is covered by automated tests.
3. Behavior difference is described in public docs with clear target audience.
4. Profile has explicit “choose this when…” guidance.
5. Compatibility/migration behavior is defined for deprecations/removals.

## Non-admission indicators

Do not add (or keep) a profile when:

- It is naming-only and behavior-equivalent to another profile.
- It has no dedicated tests.
- It increases cognitive load without clear product benefit.
- It is documented as different but implemented identically.

## Maintenance rules

- Keep profile set minimal by default.
- Remove or merge redundant profiles quickly.
- Experimental profiles must be labeled experimental and include strict boundary policy if they expose low-level interop.
