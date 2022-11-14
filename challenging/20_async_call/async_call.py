async def resolveAfter2Seconds():
  return "resolved"

async def asyncCall():
  print("calling");
  result = await resolveAfter2Seconds()
  print(result)

asyncCall()