package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalendarMatching(t *testing.T) {
	calendar1 := []StringMeeting{{"9:00", "10:30"}, {"12:00", "13:00"}, {"16:00", "18:00"}}
	dailyBounds1 := StringMeeting{"9:00", "20:00"}
	calendar2 := []StringMeeting{{"10:00", "11:30"}, {"12:30", "14:30"}, {"14:30", "15:00"}, {"16:00", "17:00"}}
	dailyBounds2 := StringMeeting{"10:00", "18:30"}
	meetingDuration := 30
	expected := []StringMeeting{{"11:30", "12:00"}, {"15:00", "16:00"}, {"18:00", "18:30"}}
	result := CalendarMatching(calendar1, dailyBounds1, calendar2, dailyBounds2, meetingDuration)
	require.Equal(t, result, expected)
}
