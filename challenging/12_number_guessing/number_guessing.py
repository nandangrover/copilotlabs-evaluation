import random

# Random number from 1 - 10
numberToGuess = random.randint(1, 10)
# This variable is used to determine if the app should continue prompting the user for input
foundCorrectNumber = False

while not foundCorrectNumber:
  # Get user input
  guess = input('Guess a number from 1 to 10: ')
  # Convert the string input to a number
  guess = int(guess)

  # Compare the guess to the secret answer and let the user know.
  if guess == numberToGuess:
    print('Congrats, you got it!')
    foundCorrectNumber = True
  else:
    print('Sorry, guess again!')