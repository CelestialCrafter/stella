import * as THREE from 'three';
import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';
import { OrbitControls } from 'three/addons/controls/OrbitControls.js';
const FOV = 70;

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(FOV, window.innerWidth / window.innerHeight, 0.1, 1000);
const loader = new GLTFLoader();
let renderer;
let controls;

const light = new THREE.AmbientLight(0x404040, 50);
scene.add(light);

const animate = () => {
  requestAnimationFrame(animate);
  controls.update();
  renderer.render(scene, camera);
};

const resize = () => {
  renderer.setSize(window.innerWidth, window.innerHeight)
  camera.aspect = window.innerWidth / window.innerHeight;
  camera.updateProjectionMatrix();
};

export const initScene = canvas => {
  renderer = new THREE.WebGLRenderer({
    antialias: true,
    canvas,
  });

  controls = new OrbitControls(camera, renderer.domElement);

  camera.position.set(50, 0, 50);
  controls.target = new THREE.Vector3(0, 0, 0);
  controls.update();

  resize();
  animate();
};

const load = url => new Promise((res, rej) => loader.load(url, res, () => {}, rej));
export const addPlanets = async urls => {
  const planets = (await Promise.all(urls.map(load))).map(gltf => gltf.scene);

  let totalBounding = new THREE.Box3();
  for (const [i, planet] of planets.entries()) {
    scene.add(planet);
    const spaceVector = new THREE.Vector3(50, 0, 0)
    planet.position.add(spaceVector.multiplyScalar(i));
    totalBounding.expandByPoint(spaceVector);
  }

  let center = new THREE.Vector3(0, 0, 0);
  totalBounding.getCenter(center);
  controls.target = center;
};

window.addEventListener('resize', resize);
