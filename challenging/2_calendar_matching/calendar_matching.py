from typing import List


class StringMeeting:
    def __init__(self, start: str, end: str):
        self.start = start
        self.end = end


class Meeting:
    def __init__(self, start: int, end: int):
        self.start = start
        self.end = end


def calendar_matching(
    calendar1: List[StringMeeting],
    daily_bounds1: StringMeeting,
    calendar2: List[StringMeeting],
    daily_bounds2: StringMeeting,
    meeting_duration: int,
) -> List[StringMeeting]:
    updated_calendar1 = update_calendar(calendar1, daily_bounds1)
    updated_calendar2 = update_calendar(calendar2, daily_bounds2)
    merged_calendar = merge_calendars(updated_calendar1, updated_calendar2)
    flattened_calendar = flatten_calendar(merged_calendar)
    return get_matching_availabilities(flattened_calendar, meeting_duration)


def update_calendar(calendar: List[StringMeeting], daily_bounds: StringMeeting) ->