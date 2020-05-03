from fabric.api import *

env.hosts = ['47.107.230.235']
env.user = 'root'


# from fabric import task


@task
def tar_task():
  # 打包
  local('make all')


@task
def upload():
  with settings(warn_only=True):
    put(".env", "/root/process/.env")


@task
def put_task():
  # 创建远程服务器文件夹
  with cd("/root/process"):
    run('docker-compose up -d')
    run("echo success")
    run("exit")


@task
def deploy():
  tar_task()
  upload()
  put_task()
