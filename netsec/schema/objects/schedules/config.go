package schedules

/*
Config ""

Parent chains:
* "Config"

Args:

Param ObjectId (string, read-only): The ObjectId param. Example: "abcd-1234".

Param Name (required, string): The Name param. String length must not exceed 31 characters.

Param ScheduleType (required, ScheduleTypeObject): The ScheduleType param. ScheduleTypeObject instance.
*/
type Config struct {
    ObjectId string `json:"id,omitempty"`
    Name string `json:"name"`
    ScheduleType ScheduleTypeObject `json:"schedule_type"`
}

/*
ScheduleTypeObject ""

Parent chains:
* "Config schedule_type"

Args:

Param NonRecurring (list of strings): The NonRecurring param.

Param Recurring (RecurringObject): The Recurring param. RecurringObject instance.
*/
type ScheduleTypeObject struct {
    NonRecurring []string `json:"non_recurring,omitempty"`
    Recurring *RecurringObject `json:"recurring,omitempty"`
}

/*
RecurringObject ""

Parent chains:
* "Config schedule_type recurring"

Args:

Param Daily (list of strings): The Daily param.

Param Weekly (WeeklyObject): The Weekly param. WeeklyObject instance.
*/
type RecurringObject struct {
    Daily []string `json:"daily,omitempty"`
    Weekly *WeeklyObject `json:"weekly,omitempty"`
}

/*
WeeklyObject ""

Parent chains:
* "Config schedule_type recurring weekly"

Args:

Param Friday (list of strings): The Friday param.

Param Monday (list of strings): The Monday param.

Param Saturday (list of strings): The Saturday param.

Param Sunday (list of strings): The Sunday param.

Param Thursday (list of strings): The Thursday param.

Param Tuesday (list of strings): The Tuesday param.

Param Wednesday (list of strings): The Wednesday param.
*/
type WeeklyObject struct {
    Friday []string `json:"friday,omitempty"`
    Monday []string `json:"monday,omitempty"`
    Saturday []string `json:"saturday,omitempty"`
    Sunday []string `json:"sunday,omitempty"`
    Thursday []string `json:"thursday,omitempty"`
    Tuesday []string `json:"tuesday,omitempty"`
    Wednesday []string `json:"wednesday,omitempty"`
}