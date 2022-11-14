import random
import time
import asyncio

async def CalculateValue(values):
    value = random.randint(0, 10)
    print("Calculated Random Value: {}".format(value))
    values.append(value)

async def main():
    random.seed(time.time())
    print("Go Channel Tutorial")

    values = []

    asyncio.create_task(CalculateValue(values))

    value = values[0]
    print(value)

asyncio.run(main())