import { initScene as baseInitScene } from '../scene';
import * as THREE from 'three';
import { onDestroy } from 'svelte';

export const initScene = canvas => {
	const { controls, scene, camera, renderer, animate, resize, pointerMove } = baseInitScene(canvas);
	camera.position.set(0, 50, 0);

	const light = new THREE.AmbientLight(0x404040, 70);
	scene.add(light);

	controls.enabled = false;
	controls.update();

	renderer.setAnimationLoop(animate);

	window.onresize = resize;
	window.onpointermove = pointerMove;

	onDestroy(() => {
		document.removeEventListener('resize', resize);
		document.removeEventListener('pointermove', pointerMove);
	});
};
