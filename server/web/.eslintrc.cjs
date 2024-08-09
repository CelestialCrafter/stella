const baseConfig = {
	root: true,
	globals: {
		BigInt: true
	},
	env: {
		node: true,
		browser: true
	},
	extends: ['airbnb-base', 'prettier', 'plugin:svelte/recommended', 'plugin:svelte/prettier'],
	plugins: [],
	parserOptions: {
		ecmaVersion: 'latest',
		sourceType: 'module'
	},
	rules: {
		'no-restricted-globals': 'off',
		'no-underscore-dangle': 'off',
		'no-empty': [
			'error',
			{
				allowEmptyCatch: true
			}
		],
		curly: ['error', 'multi'],
		'max-len': [
			'error',
			{
				code: 120
			}
		],
		'no-unused-expressions': [
			'error',
			{
				allowTernary: true,
				allowShortCircuit: true
			}
		],
		'no-console': 'off',
		'no-plusplus': 'off',
		'consistent-return': 'off',
		'no-tabs': 'off',
		'no-return-assign': ['error', 'except-parens']
	}
};

const importOff = Object.keys(require('eslint-plugin-import').rules).reduce((acc, rule) => {
	acc[`import/${rule}`] = 'off';
	return acc;
}, {});

module.exports = {
	...baseConfig,
	rules: {
		...importOff,
		...baseConfig.rules
	}
};

