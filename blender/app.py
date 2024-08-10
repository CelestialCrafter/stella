import json
import math
import os
import sys

import bpy

stella_prefix = "[stella]"

args = sys.argv[sys.argv.index("--") + 1:]
if len(args) < 2:
    raise Exception("no base file or planet supplied")


def normalize_color(color):
    return (color[0] / 255, color[1] / 255, color[2] / 255, 1)


def set_selection(name):
    bpy.context.view_layer.objects.active = bpy.data.objects[name]


def get_selection():
    return bpy.context.view_layer.objects.active


def apply_size(size):
    get_selection().scale = (size, size, size)


def apply_color(color):
    material = get_selection().active_material
    nodes = material.node_tree.nodes

    bsdf = nodes['Principled BSDF']
    bsdf.inputs['Base Color'].default_value = normalize_color(color)


def apply_rings(amount, colors, rotations, size, planet_size):
    for i in range(amount):
        set_selection(f'Ring{i}')
        ring = get_selection()
        ring.hide_set(False)
        apply_color(colors[i])
        apply_size(planet_size + size)
        rotation = rotations[i]
        ring.rotation_mode = 'XYZ'
        ring.rotation_euler = (math.radians(rotation[0]),
                               math.radians(rotation[1]),
                               math.radians(rotation[2]))

    bpy.ops.object.select_all(action='DESELECT')
    for i in range(amount):
        bpy.data.objects[f'Ring{i}'].select_set(True)
    bpy.data.objects['Planet'].select_set(True)
    bpy.ops.object.join()
    get_selection().data.name = 'Planet'


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
    set_selection("Planet")
    values = planet["values"]
    features = planet["features"]
    match features["type"]:
        case "normal":
            apply_size(values["normal_size"])
            apply_color(values["normal_color"])
            if features["normal_rings"]:
                apply_rings(values["normal_ring_amount"],
                            values["normal_ring_colors"],
                            values["normal_ring_rotation"],
                            values["normal_ring_size"], values["normal_size"])
        case "star":
            apply_size(values["star_size"])
            apply_emission_strength(values["star_brightness"])
            if features["star_neutron"]:
                apply_emission_color(values["star_neutron_color"])
    export(os.path.join(planet["directory"], planet["hash"] + ".glb"))


bpy.ops.wm.open_mainfile(filepath=args[0])
generate_planet(json.loads(args[1]))
