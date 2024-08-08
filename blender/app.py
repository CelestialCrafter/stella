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


def apply_size(obj, size):
    set_selection(obj)
    bpy.ops.transform.resize(value=(size, size, size))


def export(obj, name):
    set_selection(obj)
    bpy.ops.export_scene.gltf(filepath=name, use_selection=True)


planet = json.loads(args[0])
sphere = new_base_sphere()

apply_size(sphere, planet["values"]["size"])
export(sphere, planet["filepath"])

print(f"{stella_prefix}{planet}")
