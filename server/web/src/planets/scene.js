import { GLTFLoader } from 'three/examples/jsm/Addons.js';
import { initScene as baseInitScene } from '../scene';
import * as THREE from 'three';
import { planets, selectedPlanet } from '../stores';

export const initScene = canvas => {
	const { controls, scene, camera, renderer, animate, intersectedObject } = baseInitScene(canvas);
	camera.position.set(25, 25, 50);

	const light = new THREE.AmbientLight(0x404040, 80);
	scene.add(light);

	controls.autoRotate = true;
	controls.maxDistance = controls.getDistance();
	controls.minDistance = controls.getDistance() / 10;
	controls.enablePan = false;
	controls.update();

	renderer.setAnimationLoop(animate);

	const loader = new GLTFLoader();

	// @FIX dont do weird stuff when multiple planet updates occur
	let current = {};
	const unsubscribePlanets = planets.subscribe(planets => {
		const currentHashes = Object.keys(current);

		const hashes = Object.keys(planets);
		const newHashes = hashes.filter(hash => !currentHashes.includes(hash));
		const oldHashes = currentHashes.filter(hash => !hashes.includes(hash));

		scene.remove(...oldHashes.map(hash => current[hash]));
		oldHashes.forEach(hash => delete current[hash]);

		// @NOTE please do not move this async inside of the for loop for concurrent planet loads. the browser will probably crash...
		(async () => {
			for (const hash of newHashes) {
				const planet = (await loader.loadAsync(`/models/${hash}.glb`)).scene;

				current[hash] = planet;
				planet.children[0].name = hash;
				scene.add(planet);
			}

			const currentPlanets = Object.values(current);
			for (const [i, planet] of currentPlanets.entries()) {
				planet.position.setX(30 * i);
			}

			const pos = currentPlanets[Math.round(Object.keys(current).length / 2)].position.clone();
			controls.target.set(pos.x, pos.y, pos.z);
		})();
	});

	const unsubscribeSelected = selectedPlanet.subscribe(selected => {
		if (!selected) return;
		const pos = current[selected].position.clone();
		controls.target.set(pos.x, pos.y, pos.z);
	});

	return [
		() => {
			unsubscribePlanets();
			unsubscribeSelected();
			renderer.dispose();
		},
		() => intersectedObject(object => object.userData.name === 'Planet')
	];
};
