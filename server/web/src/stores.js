import { readable, writable } from 'svelte/store';

export const selectedPlanet = writable(null);
export const planets = writable({});
