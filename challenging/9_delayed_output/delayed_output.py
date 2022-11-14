import time
import random

def log(s, delay, randomized):
  for c in s:
    print(c, end='', flush=True)
    time.sleep((random.random() if randomized else 1) * delay)
  print()

log('Hello, world!', 200, True)