# go run main.go live/test/

globals {
  a = "b"
}

# replace terradude.*
# replace glocal.*
# replace dependency.*

# ignore local.*
# ignore var.*
# ignore aws_whatever.*
# ignore *.*.*....

terraform {
  module "a" {
    a = global.a
    b = upper(global.a)

    # handled - leave as is
    c = local.a

    # unhandled Unknown variable; There is no variable named \"local\". Did you mean \"global\"?"
    d = tolist([local.a, global.a])

    # expected output:
    # d = tolist([local.a, "b"])
  }
}
