import App from './App.svelte';
import 'ress';
import './styles/global.scss';

const app = new App({
	target: document.getElementById('app')
});

export default app;
