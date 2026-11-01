custom_build(
    'users-service',
    'docker build -t $EXPECTED_REF services/users && minikube image load $EXPECTED_REF',
    deps=['services/users'],
    disable_push=True,
    skips_local_docker=True,
)
k8s_yaml('deploy/k8s/users.yaml')
k8s_resource('users', port_forwards='8081:8080')
local_resource(
    'web-install',
    cmd='npm install',
    dir='web',
    deps=['web/package.json'],
    allow_parallel=True,
)
local_resource(
    'web',
    serve_cmd='npm run dev',
    serve_dir='web',
    resource_deps=['web-install'],
    links=['http://localhost:5173'],
)
