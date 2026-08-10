# GitHub Merge Queues

## Overview

A merge queue serializes merges into a protected branch (usually `main`) and
re-tests each pull request against the branch *plus every pull request already
ahead of it in the queue*, instead of just against the branch as it was when
the pull request was last updated. Without a merge queue, two pull requests can
each pass CI individually, against a slightly stale `main`, and still break
`main` once both are merged because they conflict with each other in a way
neither author's CI run could see. The merge queue closes that gap: what CI
validates is what actually lands.

See GitHub's docs on
[merging a pull request with a merge queue](https://docs.github.com/en/pull-requests/how-tos/merge-and-close-pull-requests/merging-a-pull-request-with-a-merge-queue)
for the full reference.

## How It Works

1. Someone with merge permission (often the author) requests queuing —
  see [Adding a Pull Request to the Queue](#adding-a-pull-request-to-the-queue)
  below. This can happen before approvals or required checks finish.
2. GitHub holds the request until the pull request's own required approvals
  and checks pass, then adds it to the queue.
3. GitHub creates a temporary branch (visible as
  `gh-readonly-queue/<base>/pr-<number>-<sha>`) containing the base branch plus
  every pull request already queued ahead of this one, plus this pull
  request's changes.
4. Required checks run against that temporary branch, not the pull request's
  own branch.
5. If checks pass, the pull request merges into the base branch and leaves the
  queue.

Entries near the front of the queue are frequently tested speculatively and in
parallel rather than strictly one-at-a-time, so a short queue often drains
faster than "N entries times one check run."

## Adding a Pull Request to the Queue

Anyone with merge permission can request queuing, and can do so even before
the pull request's own approvals or required checks have finished:

- From the pull request page, use the merge button (**Merge when ready** on
  repositories with a merge queue enabled), or
- From the CLI: `gh pr merge --auto`.

GitHub holds the request and adds the pull request to the queue automatically
once branch protection is satisfied — required approvals in place and required
checks passing on the pull request's own branch. Once actually queued, the
pull request header shows an amber **Queued** badge next to the title,
replacing the usual merge button:

![Queued PR](./screenshots/github-merge-queue-pr-queued-badge.png)

## Removing a Pull Request from the Queue

A pull request leaves the queue if any of the following happen:

- You remove it manually, either from the pull request page or from the `...`
  overflow menu on its entry on the Merge queue page.
- You push a new commit to the pull request branch — queued entries are
  removed automatically on new pushes, since the entry no longer reflects what
  would be merged.
- A repository admin uses **Clear queue** (top-right of the Merge queue page)
  to drop every entry at once. This is a broad, admin-level action — use it
  deliberately, not to unblock a single stuck pull request.

## What Happens When a Pull Request in the Middle Fails

This is the behavior a merge queue exists to handle:

- The failing pull request is removed from the queue and its author is
  notified.
- Pull requests **ahead** of the failing one are unaffected — they were
  already tested without the failing change and keep merging normally.
- Pull requests **behind** the failing one are rebuilt and retested without the
  failing change, since the queue removed it from the base they are tested
  against. In most cases they merge without the author needing to do anything.

This is the default behavior. If a repository disables the **Only merge
non-failing pull requests** queue setting, a pull request with a failed check
can still be merged as part of a later batch if that batch's overall required
checks pass — so treat the description above as the common case, not a
guarantee, on repositories where that setting has been changed.

Practical consequence: if your queued pull request shows a failed check, the
failure may belong to a pull request that was ahead of you in the queue, not to
your own commit. Open the entry's **Queue entry status details** to see exactly
which check failed and against which combination of pull requests, before
assuming your own change is at fault. A dequeued pull request must have its
issue fixed and be re-added to the queue — it does not retry automatically.

## Viewing the Current Queue

Two places show what is queued:

- **Merge queue page** — `https://github.com/<org>/<repo>/queue/<branch>`.
  Lists every entry with its position, who enqueued it, and how long ago.
  Clicking an entry opens **Queue entry status details**, showing approval
  status and the state of each required check (for example `MegaLinter` still
  running, `DCO` already passed):

  ![Queue entry status details](./screenshots/github-merge-queue-entry-status-details.png)

- **Branches page** — `https://github.com/<org>/<repo>/branches`. The
  **Merge queue** column shows a live count of queued entries per branch, the
  fastest way to check whether `main` has anything queued without opening the
  queue page:

  ![Merge queue column on Branches page](./screenshots/github-merge-queue-branches-column.png)

## Tips

- Don't force-push a branch while its pull request is queued — the entry is
  removed automatically, and you will need to re-add it.
- Required checks must run against the temporary `gh-readonly-queue/*` branch
  or a queued entry stays blocked until the status-check timeout, after which
  GitHub removes the pull request for reporting no successful result. For
  GitHub Actions, add the `merge_group` trigger alongside `pull_request`. For
  third-party CI, configure it to run on pushes to `gh-readonly-queue/*`
  branches instead — `merge_group` is an Actions-only event and won't fire for
  other providers.
- Queue position is not a guarantee of merge order: if an entry ahead of yours
  fails and is removed, entries behind it are retested and may merge sooner
  than their original position suggested.
