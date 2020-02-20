import subprocess
import os
import stat

binStructure = {
    "bin":{
        "eventserver":{
                "bin":{},
                "config":{},
                "plugins":{}
        }
    }
}

mainSrc = "cmd/start-default.go"
main = "bin/eventserver/bin/default-server"
pluginSrcDir = "plugins"
pluginDir = "bin/eventserver/plugins"
startScript = "bin/eventserver/start"

def build():
    createDirStructure(binStructure)

    subprocess.run(f"go build -o {main} {mainSrc}", shell=True)

    for filename in os.listdir(pluginSrcDir):
        if filename.endswith(".go"):
            subprocess.run(f"go build -buildmode=plugin -o {pluginDir}/{filename[:-3]}.so {pluginSrcDir}/{filename}", shell=True)

    with open(startScript, "w+") as f:
        f.write("#!/bin/bash\n./bin/default-server")

    makeExecutable(main)
    makeExecutable(startScript)

def createDirStructure(dirStructure):
    def traverse(dirStructure):
            for x in dirStructure:
                yield x
                for y in traverse(dirStructure[x]):
                    yield os.path.join(x, y)

    for path in traverse(dirStructure):
        if not os.path.exists(path):
            print(f"Creating dir {path}")
            os.mkdir(path)

def makeExecutable(filePath):
    st = os.stat(filePath)
    os.chmod(filePath, st.st_mode | stat.S_IEXEC)

if __name__ == "__main__":
    build()