import random
import time

class Game:
	def __init__(self):
		self.health = 100
		self.current_location = "Bridge"

	def play(self):
		print("Welcome to the Starship Enterprise\n\n")
		while True:
			print(location_map[self.current_location].description)
			self.process_events(location_map[self.current_location].events)
			if self.health <= 0:
				print("You are dead, game over!!!")
				return
			print("Health: " + str(self.health))
			print("You can go to these places:")
			for index, loc in enumerate(location_map[self.current_location].transitions):
				print("\t" + str(index + 1) + " - " + loc)
			i = 0
			while i < 1 or i > len(location_map[self.current_location].transitions):
				i = int(input("Where do you want to go (0 - to quit), [" + str(1) + "..." + str(len(location_map[self.current_location].transitions)) + "]: "))
			new_loc = i - 1
			self.current_location = location_map[self.current_location].transitions[new_loc]

	def process_events(self, events):
		for evt_name in events:
			self.health += events_map[evt_name].process_event()

class Event:
	def __init__(self, event_type, chance, description, health, evt):
		self.type = event_type
		self.chance = chance
		self.description = description
		self.health = health
		self.evt = evt

	def process_event(self):
		if self.chance >= random.randrange(100):
			hp = self.health
			if self.type == "Combat":
				print("Combat Event")
			print("\t" + self.description)
			if self.evt != "":
				hp = hp + events_map[self.evt].process_event()
			return hp
		return 0

class Location:
	def __init__(self, description, transitions, events):
		self.description = description
		self.transitions = transitions
		self.events = events

events_map = {
	"alienAttack": Event("Combat", 20, "An alien beams in front of you and shoots you with a ray gun.", -50, "doctorTreatment"),
	"doctorTreatment": Event("Story", 10, "The doctor rushes in and inject you with a health boost.", 30, ""),
	"android": Event("Story", 50, "Data is in the turbo lift and says hi to you", 0, ""),
	"relaxing": Event("Story", 100, "In the lounge you are so relaxed that your health improves.", 10, ""),
}

location_map = {
	"Bridge": Location("You are on the bridge of a spaceship sitting in the Captain's chair.", ["Ready Room", "Turbo Lift"], ["alienAttack"]),
	"Ready Room": Location("The Captain's ready room.", ["Bridge"], []),
	"Turbo Lift": Location("A Turbo Lift that takes you anywhere in the ship.", ["Bridge", "Lounge", "Engineering"], ["android"]),
	"Engineering": Location("You are in engineering where you see the star drive", ["Turbo Lift"], ["alienAttack"]),
    "Lounge": Location("You are in the lounge where you can relax.", ["Turbo Lift"], ["relaxing"]),
}

if __name__ == "__main__":
    game = Game()
    game.play()
 

