# Git Branching and Version Management Workflow

This document outlines our project's Git workflow, branch management strategy, and versioning approach.

## Branch Structure

Our repository maintains the following branch structure:

- **`main`**: Production-ready code, always stable
- **`develop`**: Integration branch for features before release
- **Temporary branches**:
  - **`feature/*`**: For new feature development
  - **`fix/*`**: For bug fixes
  - **`release/*`**: For release preparation

## Workflow Diagram

```txt
      Create release        Merge to main
          branch             & tag
develop ------+-----> release/1.0.0 --------+-----> main
   ^          |                  ^          |         |
   |          |                  |          |         |
   |          |             Final fixes     |         |
   |          |                             |         |
   |          |                             |         |
   |          v                             |         |
   |     QA Testing                         |         |
   |                                        |         |
   |                                        |         |
feature/xyz --------------------------+     |         |
   ^                                        |         |
   |                                        |         |
   +--- Created from develop                |         |
                                            |         |
                                            |         |
                                            |         v
                                       fix/critical ---+
                                            ^
                                            |
                                            +--- Created from main
```

## Detailed Workflow

### Feature Development

1. **Create feature branch**

   - Branch from: `develop`
   - Naming convention: `feature/short-description`

   ```bash
   git checkout develop
   git pull
   git checkout -b feature/new-command
   ```

2. **Work on feature**

   - Commit regularly with descriptive messages
   - Keep feature focused on a single functionality

3. **Create Pull Request to merge into `develop`**
   - PR should be created when feature is complete and tested
   - PR should target `develop` branch
   - After code review approval, merge feature into `develop`

### Bug Fixes

1. **Create fix branch**

   - For urgent production bugs: Branch from `main`
   - For non-urgent bugs: Branch from `develop`
   - Naming convention: `fix/issue-description`

   ```bash
   # For critical fixes
   git checkout main
   git pull
   git checkout -b fix/critical-issue

   # For non-critical fixes
   git checkout develop
   git pull
   git checkout -b fix/non-critical-issue
   ```

2. **Implement fix**

   - Keep changes minimal and focused on the bug
   - Include tests where appropriate

3. **Create Pull Request**
   - For critical fixes:
     - First PR to `main`
     - Second PR to `develop` (to ensure fix is incorporated in future releases)
   - For non-critical fixes:
     - Single PR to `develop`

### Release Process

1. **Create release branch**

   - Branch from: `develop`
   - Naming convention: `release/x.y.z` (following SemVer)

   ```bash
   git checkout develop
   git pull
   git checkout -b release/0.2.0
   ```

2. **Prepare release**

   - Update version numbers
   - Update changelog
   - Only bug fixes should be applied directly to this branch

3. **Testing and QA**

   - Perform thorough testing on the release branch
   - Any issues found should be fixed directly in the release branch

4. **Create Pull Requests**

   - First PR: From `release/x.y.z` to `main`
   - Second PR: From `release/x.y.z` to `develop`

5. **Tag the release**

   - After merging to `main`, create and push a tag

   ```bash
   git checkout main
   git pull
   git tag -a v0.2.0 -m "Release v0.2.0"
   git push origin v0.2.0
   ```

## Version Management

We follow Semantic Versioning (SemVer):

- **Major** (x.0.0): Breaking changes
- **Minor** (0.x.0): New features, backward compatible
- **Patch** (0.0.x): Bug fixes, spelling corrections, backward compatible

### When to increment version

- **Patch version**: For bug fixes, documentation updates, spelling corrections
- **Minor version**: For new features that don't break compatibility
- **Major version**: For changes that break backward compatibility

## Pull Request Guidelines

### When to Create Pull Requests

1. **Feature completion**: When a feature is completed and ready for review
2. **Bug fix**: When a bug fix is completed and tested
3. **Release preparation**: When a release branch is ready to be merged

### PR Workflow

1. **Push your branch to remote**:

   ```bash
   git push -u origin your-branch-name
   ```

2. **Create PR through GitHub/GitLab interface**:

   - Select source branch (your branch)
   - Select target branch (following workflow rules)
   - Add descriptive title and detailed description
   - Add reviewers
   - Link related issues

3. **Address review comments**:

   - Make requested changes
   - Push additional commits to the same branch
   - Request re-review if needed

4. **Merge**:
   - After approval, merge the PR
   - Delete the branch after merging (can be automated)

## Special Considerations

- **Hotfixes for production**: Create directly from `main`, merge to both `main` and `develop`
- **Long-running features**: Regularly merge `develop` into the feature branch to stay updated
- **Release conflicts**: Resolve in the release branch before merging
