import json
import os
import sys

import bpy

stella_prefix = "[stella]"

args = sys.argv[sys.argv.index("--") + 1:]
if len(args) < 2:
    raise Exception("no base file or planet supplied")


def normalize_color(color):
    return (color[0] / 255, color[1] / 255, color[2] / 255, 1)


def set_selection(obj):
    bpy.context.view_layer.objects.active = obj


def get_selection():
    return bpy.context.view_layer.objects.active


def apply_size(size):
    bpy.ops.transform.resize(value=(size, size, size))


def apply_color(color):
    material = get_selection().active_material
    nodes = material.node_tree.nodes

    bsdf = nodes['Principled BSDF']
    bsdf.inputs['Base Color'].default_value = normalize_color(color)


def apply_emission_strength(strength):
    material = get_selection().active_material
    nodes = material.node_tree.nodes

    bsdf = nodes['Principled BSDF']
    bsdf.inputs['Emission Strength'].default_value = strength


def apply_emission_color(color):
    material = get_selection().active_material
    nodes = material.node_tree.nodes

    bsdf = nodes['Principled BSDF']
    bsdf.inputs['Emission Color'].default_value = normalize_color(color)


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
            apply_emission_strength(values["star_brightness"])
            if features["star_neutron"]:
                apply_emission_color(values["star_neutron_color"])
    export(os.path.join(planet["directory"], planet["hash"] + ".glb"))


bpy.ops.wm.open_mainfile(filepath=args[0])
generate_planet(json.loads(args[1]))
