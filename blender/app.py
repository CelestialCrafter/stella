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
    material = bpy.data.materials.new(name="Planet Color")
    material.use_nodes = True

    nodes = material.node_tree.nodes
    links = material.node_tree.links

    # clear defaults
    nodes.clear()

    # create nodes
    node_rgb = nodes.new(type='ShaderNodeRGB')
    node_rgb.outputs[0].default_value = (color[0] / 255, color[1] / 255,
                                         color[2] / 255, 1)

    node_principled_bsdf = nodes.new(type='ShaderNodeBsdfPrincipled')
    node_output = nodes.new(type='ShaderNodeOutputMaterial')

    # link
    links.new(node_rgb.outputs['Color'],
              node_principled_bsdf.inputs['Base Color'])
    links.new(node_principled_bsdf.outputs['BSDF'],
              node_output.inputs['Surface'])

    # assign
    get_selection().data.materials.append(material)


def export(name):
    bpy.ops.export_scene.gltf(filepath=name, use_selection=True)


planet = json.loads(args[0])

sphere = new_base_sphere()
set_selection(sphere)

apply_subdiv(4)
apply_size(planet["values"]["size"])
apply_color(planet["values"]["color"])
export(planet["filepath"])

print(f"{stella_prefix}{planet}")
