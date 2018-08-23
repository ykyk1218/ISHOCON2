import subprocess

from flask import Flask, request
from flask_cors import CORS

app = Flask(__name__)
CORS(app)


@app.route('/run', methods=['POST'])
def post_run():
    params = request.json
    ip = params['ip']
    workload = params['workload']
    username = params['name']
    options = ''
    if ip:
        options += ' --ip ' + ip
    if workload:
        options += ' --workload ' + str(workload)
    if username:
        options += ' --name ' + username
    command = "/root/benchmark {} &".format(options)
    print(command)
    subprocess.call(command, shell=True)
    print("POST /run")
    return "DONE"


@app.route('/')
def get_index():
    print("GET /")
    return "Hello"


if __name__ == "__main__":
    app.run()
