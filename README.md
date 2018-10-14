# drone-git-version

DroneCi plugin to read version from a file, write to the output file and push the new file to the current remote branch.

## Procedure

- `git config --global user.email "drone@wesense.cloud"`
- `git config --global user.name "DroneCi"`
- `VERSION=$(cat .image_tags)`
- `echo $VERSION > VERSION`
- `git commit VERSION -m "[CI SKIP] ci upgrade to version $VERSION"`
- `git push --set-upstream origin HEAD:$DRONE_COMMIT_BRANCH`

## Documentation

## Example

```YAML
publish:
    image: schrodit/drone-git-version
    git_name: testuser
    git_email: test@example.com
    input_file: .image-tags
    output_file: VERSION
```