from fabric.api import *

env.hosts = ['47.107.230.235']
env.user = 'root'


# from fabric import task
@task
def tar_task():
    # 打包
    local("make build_process_linux")
    with lcd("bin"):
        local('tar -czvf process-server.tar.gz process-server')


@task
def put_task():
    with settings(warn_only=True):
        put("bin/process-server.tar.gz", "/root/program/process/process-server.tar.gz")
    with cd("/root/program/process"):
        run('tar -xzvf process-server.tar.gz')
        run("rm -rf process-server.tar.gz")
    with cd("/root"):
        run('supervisorctl update')
        run('supervisorctl -c supervisord.conf reload')
        run("echo success")


@task
def deploy():
    tar_task()
    put_task()
