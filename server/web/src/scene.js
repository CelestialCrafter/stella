import * as THREE from 'three';
import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';
import { OrbitControls } from 'three/addons/controls/OrbitControls.js';
const FOV = 70;

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(FOV, window.innerWidth / window.innerHeight, 0.1, 10000);
const pointer = new THREE.Vector2();
const raycaster = new THREE.Raycaster();
const loader = new GLTFLoader();
let renderer;
let controls;

const light = new THREE.AmbientLight(0x404040, 50);
scene.add(light);

export const selectedPlanet = () => {
	const intersects = raycaster.intersectObjects(Object.values(scene.children));

	if (intersects.length > 0)
		for (const intersect of intersects) {
			if (intersect.object.userData.name != 'Sphere') continue;
			return intersect.object;
		}

	return null;
};

const animate = () => {
	requestAnimationFrame(animate);

	raycaster.setFromCamera(pointer, camera);
	controls.update();
	renderer.render(scene, camera);
};

export const initScene = canvas => {
	renderer = new THREE.WebGLRenderer({
		antialias: true,
		canvas
	});

	controls = new OrbitControls(camera, renderer.domElement);

	camera.position.set(50, 0, 50);
	controls.target = new THREE.Vector3(0, 0, 0);
	controls.update();

	resize();
	animate();
};

const load = url => new Promise((res, rej) => loader.load(url, res, () => {}, rej));
export const addPlanets = async hashes => {
	const planetObjects = (
		await Promise.all(hashes.map(hash => `/models/${hash}.glb`).map(load))
	).map(gltf => gltf.scene);

	const totalBounding = new THREE.Box3();
	for (const [i, planet] of planetObjects.entries()) {
		scene.add(planet);
		planet.children[0].name = hashes[i];
		const spaceVector = new THREE.Vector3(50, 0, 0);
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

const resize = () => {
	renderer.setSize(window.innerWidth, window.innerHeight);
	camera.aspect = window.innerWidth / window.innerHeight;
	camera.updateProjectionMatrix();
};

window.addEventListener('pointermove', pointerMove);
window.addEventListener('resize', resize);
