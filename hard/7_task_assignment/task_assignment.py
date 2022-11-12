def taskAssignment(k, tasks):
    pairedTasks = []
    taskDurationToIndices = getTaskDurationToIndices(tasks)
    sortedTasks = sorted(tasks)
    for idx in range(k):
        task1Duration = sortedTasks[idx]
        indicesWithTask1Duration = taskDurationToIndices[task1Duration]
        task1Index = indicesWithTask1Duration.pop()
        task2SortedIndex = len(tasks) - 1 - idx
        task2Duration = sortedTasks[task2SortedIndex]
        indicesWithTask2Duration = taskDurationToIndices[task2Duration]
        task2Index = indicesWithTask2Duration.pop()
        pairedTasks.append([task1Index, task2Index])
    return pairedTasks
def getTaskDurationToIndices(tasks):
    taskDurationToIndices = {}
    for idx in range(len(tasks)):
        taskDuration = tasks[idx]
        if taskDuration in taskDurationToIndices:
            taskDurationToIndices[taskDuration].append(idx)
        else:
            taskDurationToIndices[taskDuration] = [idx]
    return taskDurationToIndices
tasks = [1, 3, 5, 3, 1, 4]
k = 3
print(taskAssignment(k, tasks))
