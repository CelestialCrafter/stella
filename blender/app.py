import json
import math
import os
import sys

import bpy

stella_prefix = "[stella]"
base_dir = os.path.dirname(os.path.realpath(__file__))
args = sys.argv[sys.argv.index("--") + 1:]
if len(args) < 1:
	raise Exception("no planet supplied")

def normalize_color(color):
	return (color[0] / 255, color[1] / 255, color[2] / 255, 1.0)

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

def apply_rings(amount, colors, rotations, size):
	for i in range(amount):
		set_selection(f'Ring{i}')

		ring = get_selection()
		ring.hide_set(False)
		apply_color(colors[i])
		apply_size(size)
		rotation = rotations[i]
		ring.rotation_mode = 'XYZ'
		ring.rotation_euler = (math.radians(rotation[0]), math.radians(rotation[1]), math.radians(rotation[2]))

def apply_surface(num):
	surfaces_dir = os.path.join(base_dir, 'surfaces/')
	available = os.listdir(surfaces_dir)
	surface = available[num % len(available) - 1]

	planet = bpy.data.objects['Planet']
	image = bpy.data.images.load(os.path.join(surfaces_dir, surface))
	planet.active_material.node_tree.nodes['Image Texture'].image = image

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

def apply_neutron_rod():
	bpy.data.objects['NeutronRod'].hide_set(False)

def apply_blackhole_colors(ring_style, colors):
	# sphere
	set_selection('BlackHole')
	get_selection().hide_set(False)
	apply_color(colors[0])

	# ring
	normalized = [normalize_color(color) for color in colors]
	size = (32, 32)
	image = bpy.data.images.new("BlackHoleRingImage", size[1], size[0])
	pixels = [None] * size[0] * size[1]

	for x in range(size[0]):
		for y in range(size[1]):
			vertical = ring_style == "vertical"
			selected_color = normalized[0]

			if vertical and y > size[1] / 2:
				selected_color = normalized[1]
			elif not vertical and x > size[0] / 2:
				selected_color = normalized[1]

			pixels[(y * size[0]) + x] = selected_color

	pixels = [chan for px in pixels for chan in px]
	image.pixels = pixels

	blackhole_ring = bpy.data.objects['BlackHoleRing']
	blackhole_ring.hide_set(False)
	blackhole_ring.active_material.node_tree.nodes['Image Texture'].image = image

def generate_planet(planet):
	set_selection("Planet")
	values = planet["values"]
	features = planet["features"]

	match features["type"]:
		case "normal":
			apply_size(values["normal_size"])
			apply_surface(values["normal_surface"])
			if features["normal_rings"]:
				apply_rings(values["normal_ring_amount"], values["normal_ring_colors"], values["normal_ring_rotation"],
				            values["normal_ring_size"])
		case "star":
			apply_size(values["star_size"])
			apply_emission_strength(values["star_brightness"])
			if features["star_neutron"]:
				apply_emission_color(values["star_neutron_color"])
				apply_neutron_rod()
		case "blackhole":
			apply_color(values["normal_color"])
			apply_color(values["blackhole_colors"][2])
			apply_size(values["blackhole_size"])
			apply_blackhole_colors(features["blackhole_style"], values["blackhole_colors"])

	bpy.ops.object.select_all(action='SELECT')
	bpy.ops.object.join()
	selected = get_selection()
	selected.name = 'Planet'
	selected.data.name = 'Planet'

	path = os.path.join(base_dir, "models/", planet["hash"] + ".glb")
	bpy.ops.export_scene.gltf(filepath=path, use_selection=True, export_apply=True)

bpy.ops.wm.open_mainfile(filepath=os.path.join(base_dir, "base.blend"))
generate_planet(json.loads(args[0]))
