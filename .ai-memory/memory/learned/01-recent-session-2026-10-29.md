# Recent session — 2026-10-29

This note records what was created and shipped on 2026-10-29. Paths below are relative git paths.

## go-email-reader

- Specifications for warm-up sequences 08 through 13 are in `02-spec/21-app/08-ui-and-warm-up-algorithm-improvements/` through `02-spec/21-app/13-ui-and-warm-up-algorithm-improvements/`.
- Plans are in `.ai-memory/plans/pending/08-ui-and-warm-up-algorithm-improvements.md` through `.ai-memory/plans/pending/13-ui-and-warm-up-algorithm-improvements.md`.
- Sequence 08 is the connector and gate map. Sequence 09 is brakes. Sequence 10 is the send lease and cool-down. Sequence 11 is the shadow compare. Sequence 12 is the pace envelope. Sequence 13 is the shared quiet window.
- The live tick now consults `internal/warmup/gatefeed.go`. A peer can tighten the remaining-send cap and cannot raise it. An absent or stale peer does not stop local warm-up. Local quiet hours hold. A quiet skip does not send. A dry run does not spend the cap. The gate order is health, cadence, connector, theme.
- The reader UI gate map reads `GET /api/warmup/gates`.
- Commit `f3864ee` is on `main` and is pushed. The remote uses SSH.
- Sequences 08 through 13 stay pending as catalogs. Only the gate engine is implemented.
- `user.email` is `local@gitmap`. The name is `gitmap automation`.

## gitmap-v28

- Cursor-like IDE setup landed in `cli/cmdide/ide_setup.go`, `cli/cmdide/setup/`, and `cli/cmdui/ui_server_ide_setup.go`.
- The spec is `02-spec/21-app/352-cursor-like-ide-setup-from-gitmap/`.
- Pushed commits include `3229df9b4`, `864012968`, and `04690b952`.
- `cli/cov` is a coverage profile and was not committed.

## repo-cache

- Commit `f474298` stores heatmap databases, including `.heatmap/06-cursor/go-email-reader/heatmap.db`.
- SQLite `-shm` and `-wal` sidecars were not committed.

## Push denied

These local commits are still unpushed because GitHub denied `alimtvnetwork`:

- `alim.karim.profile` (19 commits)
- `kubernetes-training-v1` (19 commits)
- `common-linux-installer` (1 commit)

## Rules learned

- Do not put a personal access token in a git remote URL. Use SSH.
- Do not commit coverage profiles or SQLite sidecar files.

## Last 30 commits in this repo

```text
0aea5cd docs(memory) - record the 2026-10-29 warm-up and git session.
2c08ded feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts (29-Oct-26 06:32 pm)
2794ac4 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts (29-Oct-26 04:23 pm)
3143895 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts (29-Oct-26 03:17 pm)
1d51540 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts (29-Oct-26 02:54 pm)
4fc4837 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts (28-Oct-26 10:18 am)
897f366 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
ece0028 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
dc99f0b feat(sync): sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager
a2dadcf feat(sync): sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager
f6a6233 feat(sync): sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager
5abd473 feat(sync): sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager
d174866 feat(sync): sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager
c485a57 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
aeb63a2 feat(sync): sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager
32c5b90 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
fd58950 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
4636fd3 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
1cc87ed feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
6d3aaf3 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
8306b6d feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
e44872a feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
e042a77 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
b5376c0 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
aa43312 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
e4403e8 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
2dd75a0 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
ec32872 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
6158940 feat(sync) - sync canonical v6 execute suite, gitmap heatmap telemetry, skills, and prompts
1761239 chore(sync): mirror canonical prompts, skills, specs
```
