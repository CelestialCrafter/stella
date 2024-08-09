import json
import os
import sys

import bpy
from bpy.app.handlers import load_post, persistent

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


def get_selection():
    return bpy.context.active_object


def apply_size(size):
    bpy.ops.transform.resize(value=(size, size, size))


def apply_subdiv(level):
    obj = get_selection()
    modifier = obj.modifiers.new("Subdivision Surface", "SUBSURF")
    modifier.levels = level
    modifier.render_levels = level


def apply_color(color):
    material = get_selection().active_material
    nodes = material.node_tree.nodes

    bsdf = nodes['Principled BSDF']
    bsdf.inputs['Base Color'].default_value = (color[0] / 255, color[1] / 255,
                                               color[2] / 255, 1)


def apply_brightness(strength):
    material = get_selection().active_material
    nodes = material.node_tree.nodes

    bsdf = nodes['Principled BSDF']
    bsdf.inputs['Emission Strength'].default_value = strength


def export(name):
    bpy.ops.export_scene.gltf(filepath=name, use_selection=True)


def generate_planet(planet):
    values = planet["values"]
    features = planet["features"]
    match features["type"]:
        case "normal":
            apply_size(values["normal_size"])
            apply_color(values["normal_color"])
        case "star":
            apply_size(values["star_size"])
            apply_brightness(values["star_brightness"])
    export(os.path.join(planet["directory"], planet["hash"] + ".glb"))
    print(f"{stella_prefix}{planet}")


bpy.ops.wm.open_mainfile(filepath=args[0])
generate_planet(json.loads(args[1]))
