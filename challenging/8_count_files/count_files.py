import os
import re
from pathlib import Path

def getFiles(dir):
    files = []
    for r, d, f in os.walk(dir):
        if not re.match(r'.*node_modules.*', r):
            files += [os.path.join(r, file) for file in f]
    return files

files = getFiles('../../../copilotlabs-evaluation')
print(len(files))