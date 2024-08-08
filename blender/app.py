import json
import sys

import bpy

stella_prefix = "[stella]"

args = sys.argv[sys.argv.index("--") + 1:]
if len(args) < 1:
    raise Exception("no planet argument")


def new_base_sphere():
    bpy.ops.mesh.primitive_uv_sphere_add(radius=1,
                                         enter_editmode=False,
                                         location=(0, 0, 0))
    return bpy.context.active_object


def set_selection(obj):
    bpy.context.view_layer.objects.active = obj


def apply_size(size):
    bpy.ops.transform.resize(value=(size, size, size))


def export(name):
    bpy.ops.export_scene.gltf(filepath=name, use_selection=True)


planet = json.loads(args[0])
sphere = new_base_sphere()

set_selection(sphere)
apply_size(planet["values"]["size"])
export(planet["filepath"])

print(f"{stella_prefix}{planet}")
