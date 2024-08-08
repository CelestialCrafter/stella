import json
import sys

stella_prefix = "[stella]"

args = sys.argv[sys.argv.index("--") + 1:]
if len(args) < 1:
    raise Exception("no planet argument")

planet = json.loads(args[0])
print(f"{stella_prefix}{planet}")
