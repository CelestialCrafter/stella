import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';

export const initScene = canvas => {
	const renderer = new THREE.WebGLRenderer({
		antialias: true,
		canvas
	});
	const pointer = new THREE.Vector2();
	const raycaster = new THREE.Raycaster();
	const scene = new THREE.Scene();
	const camera = new THREE.PerspectiveCamera(70, 0, 0.1, 20000);
	const controls = new OrbitControls(camera, renderer.domElement);

	renderer.toneMapping = THREE.ACESFilmicToneMapping;
	renderer.toneMappingExposure = 0.5;

	const animate = () => {
		raycaster.setFromCamera(pointer, camera);
		controls.update();
		renderer.render(scene, camera);
	};

	const texture = new THREE.TextureLoader().load('public/skybox.jpg', () => {
		texture.mapping = THREE.EquirectangularReflectionMapping;
		texture.colorSpace = THREE.SRGBColorSpace;
		scene.background = texture;
	});

	const intersectedObject = predicate => {
		const intersects = raycaster.intersectObjects(Object.values(scene.children));
		const searchObject = object => {
			if (predicate(object)) return object;
			return searchObject(object.parent);
		};

		for (const intersect of intersects) {
			const found = searchObject(intersect.object);
			if (found != null) return found;
		}

		return null;
	};

	const resize = () => {
		if (!renderer) return;
		const { offsetWidth: width, offsetHeight: height } = renderer.domElement;
		renderer.setSize(width, height);
		camera.aspect = width / height;
		camera.updateProjectionMatrix();
	};

	const pointerMove = event => {
		const { clientHeight, clientWidth } = renderer.domElement;
		pointer.x = (event.offsetX / clientWidth) * 2 - 1;
		pointer.y = -(event.offsetY / clientHeight) * 2 + 1;
	};

	resize();
	window.addEventListener('resize', resize);
	renderer.domElement.addEventListener('pointermove', pointerMove);

	return { controls, scene, camera, animate, renderer, intersectedObject, resize, pointerMove };
};
