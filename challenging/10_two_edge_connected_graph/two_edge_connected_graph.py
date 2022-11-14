def two_edge_connected_graph(edges):
    if len(edges) == 0:
        return True

    arrival_times = [-1 for _ in range(len(edges))]
    start_vertex = 0

    if get_minimum_arrival_time_of_ancestors(start_vertex, -1, 0, arrival_times, edges) == -1:
        return False

    return are_all_vertices_visited(arrival_times)


def are_all_vertices_visited(arrival_times):
    for time in arrival_times:
        if time == -1:
            return False

    return True


def get_minimum_arrival_time_of_ancestors(current_vertex, parent, current_time, arrival_times, edges):
    arrival_times[current_vertex] = current_time

    minimum_arrival_time = current_time

    for destination in edges[current_vertex]:
        if arrival_times[destination] == -1:
            minimum_arrival_time = min(
                minimum_arrival_time,
                get_minimum_arrival_time_of_ancestors(destination, current_vertex, current_time + 1, arrival_times, edges),
            )
        elif destination != parent:
            minimum_arrival_time = min(minimum_arrival_time, arrival_times[destination])

    # A bridge was detected, which means the graph isn't two-edge-connected.
    if minimum_arrival_time == current_time and parent != -1:
        return -1

    return minimum_arrival_time


input1 = [
    [1, 2, 5],
    [0, 2],
    [0, 1, 3],
    [2, 4, 5],
    [3, 5],
    [0, 3, 4],
]
input2 = [
    [1, 2, 5],
    [0, 2],
    [0, 1, 3],
    [2, 4, 5],
    [3, 5],
    [0, 3, 4, 6],
    [5],
]
print(two_edge_connected_graph(input1))
print(two_edge_connected_graph(input2))