import requests

def getMovieList():
    r = requests.get('https://ghibliapi.herokuapp.com/films')

    if r.status_code != 200:
        print(f'Could not send request to API: {r.status_code}')
        exit()

    print('Processing our list of movies')
    movies = r.json()
    for movie in movies:
        print(f"{movie['title']}, {movie['release_date']}")

getMovieList()