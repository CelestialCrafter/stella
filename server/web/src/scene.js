import * as THREE from 'three';
import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';
import { OrbitControls } from 'three/addons/controls/OrbitControls.js';
const FOV = 70;

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(FOV, window.innerWidth / window.innerHeight, 0.1, 20000);
const pointer = new THREE.Vector2();
const raycaster = new THREE.Raycaster();
const loader = new GLTFLoader();
let renderer;
let controls;

const light = new THREE.AmbientLight(0x404040, 70);
scene.add(light);

export const selectedPlanet = () => {
	const intersects = raycaster.intersectObjects(Object.values(scene.children));
	const searchObject = object => {
		if (object.type === 'Group' && object.userData.name === 'Planet') return object;
		return searchObject(object.parent);
	};

	for (const intersect of intersects) {
		const found = searchObject(intersect.object);
		if (found != null) return found;
	}

	return null;
};

const animate = () => {
	raycaster.setFromCamera(pointer, camera);
	controls.update();
	renderer.render(scene, camera);
};

const resize = () => {
	if (!renderer) return;
	renderer.setSize(window.innerWidth, window.innerHeight);
	camera.aspect = window.innerWidth / window.innerHeight;
	camera.updateProjectionMatrix();
};

export const initScene = canvas => {
	renderer = new THREE.WebGLRenderer({
		antialias: true,
		canvas
	});
	renderer.toneMapping = THREE.ACESFilmicToneMapping;
	renderer.toneMappingExposure = 0.5;
	camera.position.set(100, 100, 400);

	controls = new OrbitControls(camera, renderer.domElement);
	controls.autoRotate = true;
	controls.maxDistance = controls.getDistance();
	controls.minDistance = controls.getDistance() / 10;
	controls.enablePan = false;
	controls.update();

	const texture = new THREE.TextureLoader().load('public/skybox.jpg', () => {
		texture.mapping = THREE.EquirectangularReflectionMapping;
		texture.colorSpace = THREE.SRGBColorSpace;
		scene.background = texture;
	});

	resize();
	renderer.setAnimationLoop(animate);
};

const load = url =>
	new Promise((res, rej) => {
		loader.load(url, res, () => {}, rej);
	});
export const addPlanets = async hashes => {
	const planetObjects = (
		await Promise.all(hashes.map(hash => `/models/${hash}.glb`).map(load))
	).map(gltf => gltf.scene);

	const totalBounding = new THREE.Box3();
	for (const [i, planet] of planetObjects.entries()) {
		scene.add(planet);
		planet.children[0].name = hashes[i];
		const spaceVector = new THREE.Vector3(20, 0, 0);
		planet.position.add(spaceVector.multiplyScalar(i));
		totalBounding.expandByPoint(spaceVector);
	}

	const center = new THREE.Vector3(0, 0, 0);
	totalBounding.getCenter(center);
	controls.target = center;
};

const pointerMove = event => {
	pointer.x = (event.clientX / window.innerWidth) * 2 - 1;
	pointer.y = -(event.clientY / window.innerHeight) * 2 + 1;
};

window.addEventListener('pointermove', pointerMove);
window.addEventListener('resize', resize);
