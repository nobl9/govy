# Govy Skill Evals

[evals.json](evals.json) defines eight tasks with observable acceptance criteria.
The cases cover field boundaries, custom metadata, nested paths, optional data,
pointer presence, conditions, transformed values, validation plans,
hidden nested values, callback compatibility, and strict format selection.

## Run a comparison

1. Snapshot the old and revised skill before either run.
2. Give independent agents the same prompt and pinned govy checkout.
   Assign one skill snapshot per agent.
   Keep expectations, contract tests, and other outputs out of each agent's context.
3. Save each run's Go files and explanation in its own `outputs/` directory.
   Ask the agent to record the skill references it actually reads.
   Keep generated modules and logs outside `outputs/`.
4. Run the independent checks with [run-case.bash](run-case.bash).
   Use an installed Go toolchain that supports the checkout's `go.mod`.
   The checkout's dependencies must already be cached.
5. Grade every expectation from the outputs and test results.
   Record evidence in `grading.json` with `text`, `passed`, and `evidence` fields.
6. Compare correctness and reference selection before accepting a smaller skill.

Set `eval_outputs` to one run's output directory.
Set `govy_checkout` to the pinned source checkout.
From the govy repository root, run case 3:

```sh
bash .agents/skills/govy/evals/run-case.bash 3 "$eval_outputs" "$govy_checkout"
```

The runner copies the generated Go files into a temporary module.
It adds the selected [contract tests](testdata/) and runs both test sets.
It emits Go test JSON and returns a failing status for build or test failures.
It disables workspace inheritance and dependency downloads.
Output filenames beginning with `govy_eval_` are reserved for contract checks.

## Grade behavior and context use

A passing contract suite is necessary, but some expectations need review.
Check explanations, use of existing helpers, rule reuse, and plan descriptions.
An answer that names the expected APIs can still implement the wrong behavior.

Record the entrypoint size and references read for each task.
A narrow question must reach its relevant reference without loading every guide.
Treat file size as a context estimate, not a measured model token count.
One run per case checks coverage, not statistical improvement or output variance.

Before trusting a new assertion, run it against an intentionally incorrect output.
For example, omit required project names or replace collection limits with 999.
The checks must reject those outputs.
Keep such mutations in temporary copies.

## Inspect results

Use `skill-creator` to aggregate grades and create a static review page:

```sh
python <skill-creator-skill-dir>/eval-viewer/generate_review.py \
  <workspace>/iteration-N \
  --skill-name govy \
  --benchmark <workspace>/iteration-N/benchmark.json \
  --static <workspace>/iteration-N/review.html
```

Keep outputs and reports outside the skill package.
If token counts or timings are unavailable, mark them unmeasured.
Do not report placeholder zeroes as measurements.
