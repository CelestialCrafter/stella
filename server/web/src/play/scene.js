import { FontLoader, TextGeometry } from 'three/examples/jsm/Addons.js';
import { initScene as baseInitScene } from '../scene';
import * as THREE from 'three';

const colorMap = [
	0x0000ff, 0x00ff00, 0x00ffff, 0xff0000, 0xff00ff, 0xffff00, 0xffffff, 0x0000ee, 0x00ee00,
	0x00eeee, 0xee0000, 0xee00ee, 0xeeee00, 0xeeeeee, 0xffffff
];

export const initScene = async (canvas, state) => {
	const { controls, scene, camera, renderer, animate } = baseInitScene(canvas);

	const light = new THREE.AmbientLight(0x404040, 70);
	scene.add(light);

	controls.enabled = false;
	controls.update();

	renderer.setAnimationLoop(animate);
	// @TODO replace azuki, it was the only one in my downloads folder
	const spacing = 20;
	const size = 10;
	const font = await new FontLoader().loadAsync('public/azuki_font.json');

	const bounding = new THREE.Box3();
	for (const [r, row] of state.board.entries())
		for (const [c, cell] of row.entries()) {
			const geometry = new TextGeometry(cell.toString(), { font, size, depth: 1 });
			const material = new THREE.MeshBasicMaterial({
				color: colorMap[Math.round(Math.max(Math.log(cell), 0))]
			});
			const object = new THREE.Mesh(geometry, material);
			object.position.setX(r * spacing);
			object.position.setY(c * -spacing);

			object.userData.r = r;
			object.userData.c = c;
			object.name = 'cell';

			scene.add(object);
			bounding.expandByObject(object);
		}

	const geometry = new TextGeometry('you lose!', { font, size, depth: 1 });
	const material = new THREE.MeshBasicMaterial({
		color: 0xffffff
	});
	const text = new THREE.Mesh(geometry, material);
	text.name = 'lose';
	text.position.setY(state.board[0].length * -spacing);
	text.visible = state.finished;
	scene.add(text);

	bounding.getCenter(controls.target);
	bounding.getCenter(camera.position);
	camera.position.setZ(100);
	controls.update();

	return [renderer.dispose, scene, font];
};

export const updateScene = (scene, font, state) => {
	const cells = scene.getObjectsByProperty('name', 'cell');
	const lose = scene.getObjectByProperty('name', 'lose');

	for (const object of cells) {
		const { r, c } = object.userData;
		const cell = state.board[r][c];

		object.material.color.setHex(colorMap[Math.round(Math.max(Math.log(cell), 0))]);
		object.geometry = new TextGeometry(cell.toString(), { font, size: 10, depth: 1 });
	}

	lose.visible = state.finished;
};
