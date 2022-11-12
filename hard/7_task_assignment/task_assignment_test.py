import unittest
from task_assignment import taskAssignment

class TestProgram(unittest.TestCase):
  def test_case_1(self):
    k = 3
    tasks = [1, 3, 5, 3, 1, 4]
    expected = [
      [4, 2],
      [0, 5],
      [3, 1],
    ]
    actual = taskAssignment(k, tasks)
    self.assertEqual(actual, expected)
