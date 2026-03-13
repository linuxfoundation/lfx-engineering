# AI Prompt Template for Engineering Tasks

A guide for writing clear, well-scoped problem statements when using AI coding tools.
Following this structure helps the AI produce accurate, safe, and reviewable changes on
the first pass — reducing rework and preventing unintended side effects.

---

## Why This Matters

AI tools work best when they have a clear picture of three things: what the end result
looks like, what context surrounds the problem, and what boundaries they should not cross.
Vague prompts lead to vague output. Over-broad prompts lead to sprawling, hard-to-review
changes. This template gives you a repeatable structure to avoid both failure modes.

---

## The Template

Copy the block below and fill in each section. Remove any sections that genuinely don't
apply, but err on the side of including them — the more precise you are, the better the
result.

```markdown
## Task Title
<!-- A short, descriptive name for the feature or fix (used for commit messages, PR titles, etc.) -->

### End State
<!-- Describe the finished result as if you're looking at it after it's done.
     Be specific about what exists, where it appears, and what it does. -->

### Context
<!-- Provide the background the AI needs to orient itself in the codebase
     and understand the "why" behind the task. -->

- Where in the application does this change live?
- What existing modules, services, or components are relevant?
- Are there any related patterns, conventions, or prior art in the codebase the AI should follow?
- Link to any relevant tickets, specs, or design files if available.
- Reference any local configuration the AI should use (e.g. `.claude/`, `.editorconfig`, linting rules).

### Constraints
<!-- Draw the boundaries. Tell the AI what it should NOT do, what it must preserve,
     and what standards it must follow. -->

- What existing code or behavior must not be modified?
- What coding conventions, file structure patterns, or style rules apply?
- Are there performance, security, or accessibility requirements?
- What dependencies or libraries are approved (or off-limits)?
- What testing expectations apply (unit tests, integration tests, type coverage)?

### Acceptance Criteria
<!-- Define how you'll know the task is done correctly. Be concrete and verifiable. -->

- What should a reviewer see when they pull the branch and run the app?
- What edge cases should be handled?
- What automated checks should pass (tests, linting, type checks, build)?
- What output or confirmation should the AI provide when it's finished?
```

---

## Filled-In Examples

### Example 1: Adding a Feature

```markdown
## Task Title
Add CSV Export to the Groups Module

### End State
A new Export button is added to the Groups module, positioned in the top-right area
above the table, that exports the current group list as a well-formatted CSV file.

### Context
- The Groups module is accessible from the left navigation in the app.
- The export button should appear in the top-right corner of the page, above the
  groups table — consistent with other action buttons in the UI.
- The exported CSV should include all visible/relevant group fields (e.g. Group Name,
  Type, Status, Members Count, Created Date, etc.).
- Use all configuration from the `.claude` local directory.

### Constraints
- Don't modify existing group list logic or table components — only add the export
  feature on top.
- Follow the existing file-per-feature/module convention used elsewhere in the codebase.
- The export button should match the existing UI component library and style conventions
  already in use.
- CSV output must be well-formatted: proper headers, values escaped for commas/quotes,
  one row per group.
- Export should reflect the current filtered/searched state of the table (not always
  the full unfiltered dataset).

### Acceptance Criteria
- The export button renders correctly above the groups table.
- Clicking it downloads a `.csv` file.
- The CSV is well-formatted with correct headers and data.
- List what fields are included in the export and the component/file(s) modified
  or created.
```

### Example 2: Fixing a Bug

```markdown
## Task Title
Fix Pagination Reset on Filter Change in Members List

### End State
When a user changes a filter on the Members list while on page 3 (or any page beyond
the first), the table resets to page 1 and displays the correct filtered results.
Currently it stays on the current page number, which often results in an empty table
or stale data.

### Context
- The Members list is in `src/modules/members/MembersList.tsx`.
- Pagination state is managed via a `usePagination` hook in `src/hooks/usePagination.ts`.
- Filters are applied through query parameters and a `useFilters` hook.
- This bug was reported in JIRA-4521.

### Constraints
- Do not refactor the pagination or filter hooks — only fix the interaction between them.
- Do not change the API contract or query parameter structure.
- Ensure the fix works for all filter types (text search, dropdowns, date ranges).
- Add or update unit tests covering the filter-change-resets-page scenario.

### Acceptance Criteria
- Changing any filter resets pagination to page 1.
- The table displays the correct filtered results on page 1 after a filter change.
- Manually navigating to page 2+ still works normally when filters are unchanged.
- Existing pagination and filter tests still pass.
- New test(s) cover the specific regression.
```

### Example 3: Refactoring

```markdown
## Task Title
Extract Shared Date Formatting Utilities

### End State
A new `src/utils/dateFormat.ts` utility module exists with reusable date formatting
functions. All modules that previously had inline date formatting logic now import
from this shared utility. No user-facing behavior changes.

### Context
- At least four modules currently duplicate date formatting logic: Events, Reports,
  Members, and Audit Log.
- The project uses `date-fns` for date manipulation — continue using it.
- The coding standard requires utility functions to be pure, well-typed, and
  individually exported.

### Constraints
- This is a refactor only — no user-facing behavior should change.
- Do not introduce new dependencies.
- Every function must have JSDoc comments and TypeScript return types.
- Do not change date formats that are already displayed to users — match existing
  output exactly.
- Run the full test suite and confirm zero regressions.

### Acceptance Criteria
- `src/utils/dateFormat.ts` exists with exported, typed, documented functions.
- All previously duplicated date formatting code now uses the shared utility.
- The full test suite passes with no regressions.
- Provide a list of files modified and functions extracted.
```

---

## Writing Tips

**Be specific about the end state.** "Add a button that exports CSV" is adequate.
"Improve the export experience" is not. The AI needs a concrete target.

**Name the files and modules.** If you know where the change should happen, say so.
If you don't, say that too — the AI can search, but telling it where to look saves time
and reduces the chance it modifies the wrong thing.

**Constraints are the most important section.** AI tools are eager to help and will
happily refactor your entire codebase if you don't tell them to stay in their lane.
Be explicit about what should not change.

**Include the "why" when it affects the "how."** If a constraint exists because of a
performance concern, a security requirement, or a past incident, mention it. Context
helps the AI make better judgment calls in gray areas.

**Reference existing patterns.** "Follow the same approach used in the Events module"
is more effective than describing a pattern from scratch. The AI can read the codebase
— point it to good examples.

**Define done.** If you don't specify acceptance criteria, the AI will decide for itself
when it's finished. That's usually too early.

---

## Common Mistakes to Avoid

**Too vague:** "Fix the login bug." Which login bug? What's the expected behavior?
What should the AI leave alone?

**Too broad:** "Refactor the entire auth module to use the new API." Break large tasks
into smaller, reviewable units. Each prompt should map to roughly one PR's worth of work.

**Missing constraints:** "Add search to the dashboard." Without constraints, the AI might
rewrite your API layer, add a new dependency, or restructure your component hierarchy to
accommodate the feature.

**No acceptance criteria:** "Make it work." The AI will make _something_ work. Whether
it's what you wanted is another question.

**Assuming context:** "Use the standard approach." The AI doesn't know what your team
considers standard unless you tell it or point it to an example.

---

## Quick Checklist

Before submitting your prompt, verify:

- [ ] End state describes a specific, observable result
- [ ] Context identifies the relevant files, modules, or areas of the codebase
- [ ] Constraints define what must not change and what standards apply
- [ ] Acceptance criteria are concrete and verifiable
- [ ] The scope maps to a single, reviewable unit of work
- [ ] Any referenced conventions or patterns include a pointer to an example in the codebase
