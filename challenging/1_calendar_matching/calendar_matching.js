const calendarMatching = (
	calendar1,
	dailyBounds1,
	calendar2,
	dailyBounds2,
	meetingDuration,
  ) => {
	const updatedCalendar1 = updateCalendar(calendar1, dailyBounds1);
	const updatedCalendar2 = updateCalendar(calendar2, dailyBounds2);
	const mergedCalendar = mergeCalendars(updatedCalendar1, updatedCalendar2);
	const flattenedCalendar = flattenCalendar(mergedCalendar);
	return getMatchingAvailabilities(flattenedCalendar, meetingDuration);
  };
  
  const updateCalendar = (calendar, dailyBounds) => {
	const updatedCalendar = [
	  { start: 0, end: timeToMinutes(dailyBounds.start) },
	].concat(calendar);
	updatedCalendar.push({
	  start: timeToMinutes(dailyBounds.end),
	  end: 24 * 60,
	});
  
	const meetings = [];
	for (const i of updatedCalendar) {
	  meetings.push({ start: timeToMinutes(i.start), end: timeToMinutes