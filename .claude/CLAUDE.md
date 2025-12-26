# Principles:
DRY, YAGNI, KISS, SOLID.
Prioritize: Clarity>Cleverness, Consistency>Optimization.
Actions: !Hardcoding, Stateless, Confirm changes, Fail loudly, Sanitize inputs, Reversible changes, Test behavior, Explicit dependencies.

Fix root causes, not symptoms
Use String-based parsing rather than regex parsing

NEVER edit more than one module at a time. If you need to, stop and return to me first, letting me know what you plan on doing next.

I am the project manager. Implement only the features I specified and nothing more. That would be scope creep and NOT allowed.

do not over-engineer.

do not over-optimize.

do not over-document. be concise.

do not over-comment. be concise.

maintain code consistency. code should be consistent in style, naming, and structure. in par with other code in the project.

never use **. prefer using ##
never use long dashes.

follow titlecase for pr title.

stop and return to me to discuss if anything requires a significant refactor.

## Git commits:
- commit after each file change, not batch commits
- use conventional commit pattern: feat:, fix:, chore:, refactor:, docs:, test:
- no commit descriptions, no AI credits
- keep commit messages concise
