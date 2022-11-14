const getMovieList = require('./movie_list_fetch');

it('Test Case #1', function () {
  const expected = ["Processing our list of movies",
  "Castle in the Sky, 1986",
  "Grave of the Fireflies, 1988",
  "My Neighbor Totoro, 1988",
  "Kiki's Delivery Service, 1989",
  "Only Yesterday, 1991",
  "Porco Rosso, 1992",
  "Pom Poko, 1994",
  "Whisper of the Heart, 1995",
  "Princess Mononoke, 1997",
  "My Neighbors the Yamadas, 1999",
  "Spirited Away, 2001",
  "The Cat Returns, 2002",
  "Howl's Moving Castle, 2004",
  "Tales from Earthsea, 2006",
  "Ponyo, 2008",
  "Arrietty, 2010",
  "From Up on Poppy Hill, 2011",
  "The Wind Rises, 2013",
  "The Tale of the Princess Kaguya, 2013",
  "When Marnie Was There, 2014",
  "The Red Turtle, 2016",
  "Earwig and the Witch, 2021"];

  const logSpy = jest.spyOn(console._stdout, 'write')
  getMovieList(input);
  // wait
  setTimeout(() => {
    
  }, 100);
  const sanitizedOutput = logSpy.mock.calls.toString().replace(/,/g, '').trim();
  expect(sanitizedOutput).toBe(expected);
});
