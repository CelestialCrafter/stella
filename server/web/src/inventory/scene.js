import { GLTFLoader } from 'three/examples/jsm/Addons.js';
import { initScene as baseInitScene } from '../scene';
import * as THREE from 'three';
import { planets } from '../stores';

export const initScene = canvas => {
	const { controls, scene, camera, renderer, animate, intersectedObject } = baseInitScene(canvas);
	camera.position.set(25, 25, 50);

	const light = new THREE.AmbientLight(0x404040, 70);
	scene.add(light);

	controls.autoRotate = true;
	controls.maxDistance = controls.getDistance();
	controls.minDistance = controls.getDistance() / 10;
	controls.enablePan = false;
	controls.update();

	renderer.setAnimationLoop(animate);

	const loader = new GLTFLoader();

	// @FIX dont do weird stuff when multiple planet updates occur
	const unsubscribe = planets.subscribe(planets => {
		const hashes = Object.keys(planets);
		(async () => {
			const planetObjects = (
				await Promise.all(hashes.map(hash => loader.loadAsync(`/models/${hash}.glb`)))
			).map(gltf => gltf.scene);

			for (const [i, planet] of planetObjects.entries()) {
				scene.add(planet);
				planet.children[0].name = hashes[i];
				const spaceVector = new THREE.Vector3(20, 0, 0);
				planet.position.add(spaceVector.multiplyScalar(i));
			}
		})();
	});

	return () =>
		intersectedObject(object => object.type === 'Group' && object.userData.name === 'Planet');
};
