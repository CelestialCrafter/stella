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

  camera.position.set(0, 0, 100);

  controls.update();
  resize();
  animate();
};

export const addPlanets = planets => {
  let i = 0;
  planets.forEach(url => loader.load(url, gltf => {
    const planet = gltf.scene;
    scene.add(planet);
    planet.position.x = i * 50;
    console.log(planet);
    i++;
  }));
};

window.addEventListener('resize', resize);
