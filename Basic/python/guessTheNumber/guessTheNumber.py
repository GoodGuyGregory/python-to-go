# This is a Guess that number game.
import random

secretNumber = random.randint(1, 20)
print('I am thinking of a number between 1 and 20')

# Ask the player to guess 6 times.
for guessesTaken in range(1, 7):
    print('Take a Guess at what it might be. ')
    guess = int(input())

    if guess < secretNumber:
        print('Your guess is too low.')
    elif guess > secretNumber:
        print('Your guess is too high.')
    else:
        # The person guessed it
        break

if guess == secretNumber:
    print('Good Job. You guessed the number I was thinking in ' +
          str(guessesTaken) + ' guesses!')
else:
    print('Nice try. But the number I was thinking of was ' + str(secretNumber))
