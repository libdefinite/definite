/** @type {import('vite').UserConfig} */
export default {
    build: {
        lib: {
            entry: 'internal/ctl/console/js/main.js',
            formats: ['iife'],
            name: 'app',
            fileName: () => 'output.js',
        },
        outDir: 'internal/ctl/console/static',
        emptyOutDir: false,
        rollupOptions: {
            onwarn(warning, warn) {
                // htmx uses eval internally — known and acceptable
                if (warning.code === 'EVAL' && warning.id?.includes('htmx.org')) return
                warn(warning)
            },
        },
    },
}
