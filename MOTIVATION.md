# Motivation

## Why not use terragrunt
terragrunt is great. It solves a lot of issues everybody is facing with native terraform.
That said, we faced some issues with terragrunt where fixing would require major breaking changes.
So integrating our ideas into terragrunt did not seem the right solution for us.

### Issues we faced with terragrunt as of today
- its output is not helping in identifying errors - what went wrong?
- it does not have a fully DRY (don't repeat yourself) way for defining backends and providers
- it does not support inherited configuration files - only a leaf and root configuration
  - in addition including root configuration is not dry
- it does not support the use of terraform modules from the registry in an easy way
- its code base is not easy to modify

### features we additionally want - based on our daily life experience
- enable people with restricted access to just run plan for faster development without need of CI
  - no need to be able to write to remote statefiles
    - read statefiles (`s3:GetObject`)
  - no need to lock the remote state (terraform option `-lock=false`)
  - no need to be able to change resources
    - refresh resources (run with `arn:aws:iam::aws:policy/ReadOnlyAccess`)
- run the commands from anywhere - no need to change into a specific directory for executing commands
- create a directory with a complete setup you can `cd` into and run terrform directly
  - not forcing injection of any values via environment variables to enable easy debugging
