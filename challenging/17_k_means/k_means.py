import random


class Centroid:
    def __init__(self, location):
        self.location = location
        self.closest_users = set()


def get_k_means(user_feature_map, num_features_per_user, k):
    # Don't change the following two lines of code.
    random.seed(42)
    # Gets the inital users, to be used as centroids.
    inital_centroid_users = random.sample(sorted(list(user_feature_map.keys())), k)

    centroids = [Centroid(user_feature_map[inital_centroid_user]) for inital_centroid_user in inital_centroid_users]
    for _ in range(10):
        for uid, features in user_feature_map.items():
            closest_centroid_distance = float("inf")
            closest_centroid = None
            for centroid in centroids:
                features_to_centroid_distance = get_manhattan_distance(features, centroid.location)
                if features_to_centroid_distance < closest_centroid_distance:
                    closest_centroid_distance = features_to_centroid_distance
                    closest_centroid = centroid
            closest_centroid.closest_users.add(uid)

        for centroid in centroids:
            centroid.location = get_centroid_average(centroid, user_feature_map, num_features_per_user)
            centroid.closest_users.clear()
    return [centroid.location for centroid in centroids]


def get_manhattan_distance(features, other_features):
    absolute_differences = []
    for i in range(len(features)):
        absolute_differences.append(abs(features[i] - other_features[i]))
    return sum(absolute_differences)


def get_centroid_average(centroid, user_feature_map, num_features_per_user):
    centroid_average = [0] * num_features_per_user
    for i in range(num_features_per_user):
        for user in centroid.closest_users:
            centroid_average[i] = centroid_average[i] + user_feature_map[user][i]
    return [centroid_dimension / len(centroid.closest_users) for centroid_dimension in centroid_average]


if __name__ == "__main__":
    user_feature_map = {
        "user1": [1, 1, 1],
        "user2": [2, 2, 2],
        "user3": [3, 3, 3],
        "user4": [4, 4, 4],
        "user5": [5, 5, 5],
        "user6": [6, 6, 6],
        "user7": [7, 7, 7],
        "user8": [8, 8, 8],
        "user9": [9, 9, 9],
        "user10": [10, 10, 10],
    }
    num_features_per_user = 3
    k = 2
    centroids = get_k_means(user_feature_map, num_features_per_user, k)
    for centroid in centroids:
        print(centroid)