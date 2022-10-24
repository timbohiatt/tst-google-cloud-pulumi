// lint-staged.config.js
export default {
    '**/pulumi-native/ts/**/*.{js,ts}?(x)': () => 'tsc -p tsconfig.json --noEmit',
}