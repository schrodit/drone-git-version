# drone-git-version

DroneCi plugin to read version from a file, write to the output file and push the new file to the current remote branch.

## Procedure

- `git config --global user.email "test@example.com"`
- `git config --global user.name "testuser"`
- `VERSION=$(cat .image_tags)`
- `echo $VERSION > VERSION`
- `git commit VERSION -m "[CI SKIP] ci upgrade to version $VERSION"`
- `git push --set-upstream origin HEAD:$DRONE_COMMIT_BRANCH`

## Documentation

The drone-git-version plugin, reads comma separated versions from a input and writes the latest of these versions to an output file.
The output file is the pushed back to the cuirrent repository.

| Parameter name | Description                                                                      | Optional |
| -------------- | :------------------------------------------------------------------------------- | :------: |
| git_name       | git config.name                                                                  |          |
| git_email      | git config.email                                                                 |          |
| input_file     | Path to the versions input file. Versions file musst be comma separated versions |          |
| output_file    | Path to the version output file that is pushed back to the repo.                 |          |

## Example

```YAML
version:
    image: schrodit/drone-git-version
    git_name: testuser
    git_email: test@example.com
    input_file: .image_tags
    output_file: VERSION
```
