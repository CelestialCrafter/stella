import { GLTFLoader } from 'three/examples/jsm/Addons.js';
import { initScene as baseInitScene } from '../scene';
import * as THREE from 'three';
import { orbit, planets, selectedPlanet } from '../stores';

export const initScene = canvas => {
	const { controls, scene, camera, renderer, animate, intersectedObject, resize } =
		baseInitScene(canvas);
	camera.position.set(25, 25, 50);

	const light = new THREE.AmbientLight(0x404040, 80);
	scene.add(light);

	controls.autoRotate = true;
	controls.maxDistance = controls.getDistance();
	controls.minDistance = controls.getDistance() / 10;
	controls.enablePan = false;
	controls.update();

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
		})();
	});

	let selected;
	let doResize = false;
	const unsubscribeSelected = selectedPlanet.subscribe(newSelected => {
		if (!newSelected) return;
		selected = newSelected;
		doResize = true;
	});

	let distanceMultiplier, speed;
	const unsubscribeOrbit = orbit.subscribe(newOrbit => ([distanceMultiplier, speed] = newOrbit));

	const calculateCurrentOrbit = (d, s, t) => {
		const r = (((s * t) / d) % 360) * (Math.PI / 180);
		return [d * Math.cos(r) || 0, d * Math.sin(r) || 0];
	};

	renderer.setAnimationLoop(() => {
		const currentPlanets = Object.values(current);

		if (currentPlanets.length < 1) return animate();
		for (const [i, [hash, planet]] of Object.entries(current).entries()) {
			const [x, z] = calculateCurrentOrbit(i * distanceMultiplier, speed, Date.now());
			planet.position.setX(x);
			planet.position.setZ(z);
			if (selected == hash) {
				controls.target.setX(x);
				controls.target.setZ(z);
			}
		}

		if (doResize) {
			resize();
			doResize = false;
		}

		animate();
	});

	return [
		() => {
			unsubscribePlanets();
			unsubscribeSelected();
			unsubscribeOrbit();
			renderer.dispose();
		},
		() => intersectedObject(object => object.userData.name === 'Planet')
	];
};
