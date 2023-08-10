module.exports = {
    apps: [{
        name: 'real-time-server',
        script: './compiled/index.js',
        instances: 1,
        max_memory_restart: '500M'
    }]
}