# Class Scheduler
<br/><a name="home"></a>
> AI solutions by ChatGPT.

  * [Problem](#problem)
  * [Solutions](#solutions)
    - [Java](#java)
    - [JavaScript](javascript)
    - [Python](#python)

___
<br/><a name="problem"></a>
## Problem

Design a scheduling program for a college which has a few (less than 5) campuses, where each location has a limited numbers of classrooms (usually less than 30), plus any virtual classroom as facility resource set. Each physical classroom has attributes like size (allowing maximum numbers of students to fit), and availability of hours (morning 9 AM to 12 PM, afternoon 1 PM to 5 PM, and evening 6 PM to 10 PM).

For a specific quarter (12 weeks), there are a number of classes, which has attributes of length (assuming 1 or up to 2 hours) , repeats (less than 12 for each quarter, assuming only one in a week thru the whole quarter), enrollment (numbers of enrolled students assuming no more than the largest size of classes), preferred time (e.g. morning, afternoon and evening), and preferred list of campus locations and classrooms (including virtual classroom).

The virtual classroom (with almost unlimited availability) options must be set with each specific class, that can be as one of the following: not-allowed (for some class requiring e.g. hands-on experience), virtual-only, preferred choice, or an optional choice (meaning only to use if no physical classroom available).

The school has a number of instructors, and each of them has 3 preference lists, one for time availability of days (with specific morning, afternoon, and/or evening) across 12 weeks,  one for preferred classes (which allowed to teach), and one for preferred campus locations and class rooms (including preferred option to use virtual class). The virtual options for an instructor can be opt-out (no virtual at all), preferred (okay to use if physical available), or always (only use virtual).

Please write a program in, e.g. Python, to meet above requirements and balance the resources of instructors, campus locations, classrooms (e.g. only use virtual classroom if no physical classroom available), and time (e.g. only use evening if no available resources in morning or afternoon). Note to break down complex logic as well as to avoid from lengthy if-else and big functions.

[Back](#problem) | [Home](#home)


<br/><a name="solutions"></a>
## Solutions

<br/><a name="java"></a>
### Java

Java implementation:

* Defines clear **models** (`Classroom`, `Instructor`, `Course`, `ScheduleEntry`)
* Uses **enums** for time slots and virtual‐class rules
* Splits logic into small **finder** methods (`findClassroomCandidates`, `findSlot`, `findInstructor`)
* Keeps the **core scheduler** (`scheduleCourse`, `run`) concise and free of big `if` chains

Plug in actual data in `main(...)`, run `scheduler.run()`, then call `scheduler.printSchedule()`.

```java
import java.util.*;
import java.util.stream.Collectors;

public class CollegeScheduler {

    // --- Enums ---
    public enum TimeSlot {
        MORNING, AFTERNOON, EVENING
    }

    public enum VirtualOption {
        NOT_ALLOWED, VIRTUAL_ONLY, PREFERRED, OPTIONAL
    }

    public enum InstructorVirtualPref {
        OPT_OUT, PREFERRED, ALWAYS
    }

    // --- Data models ---
    public static class Classroom {
        public final String id;
        public final String campus;
        public final int size;
        public final Set<TimeSlot> availability;
        public final boolean isVirtual;

        public Classroom(String id, String campus, int size,
                         Set<TimeSlot> availability, boolean isVirtual) {
            this.id = id;
            this.campus = campus;
            this.size = size;
            this.availability = availability;
            this.isVirtual = isVirtual;
        }
    }

    public static class Instructor {
        public final String id;
        public final Map<Integer, Set<TimeSlot>> weeklyAvailability;
        public final Set<String> allowedClasses;
        public final Set<String> preferredCampuses;
        public final InstructorVirtualPref virtualPref;

        public Instructor(String id,
                          Map<Integer, Set<TimeSlot>> weeklyAvailability,
                          Set<String> allowedClasses,
                          Set<String> preferredCampuses,
                          InstructorVirtualPref virtualPref) {
            this.id = id;
            this.weeklyAvailability = weeklyAvailability;
            this.allowedClasses = allowedClasses;
            this.preferredCampuses = preferredCampuses;
            this.virtualPref = virtualPref;
        }
    }

    public static class Course {
        public final String id;
        public final int lengthHours;
        public final int repeats;
        public final int enrollment;
        public final Set<TimeSlot> preferredTimes;
        public final Set<String> preferredCampuses;
        public final VirtualOption virtualOption;

        public Course(String id, int lengthHours, int repeats, int enrollment,
                      Set<TimeSlot> preferredTimes,
                      Set<String> preferredCampuses,
                      VirtualOption virtualOption) {
            this.id = id;
            this.lengthHours = lengthHours;
            this.repeats = repeats;
            this.enrollment = enrollment;
            this.preferredTimes = preferredTimes;
            this.preferredCampuses = preferredCampuses;
            this.virtualOption = virtualOption;
        }
    }

    public static class ScheduleEntry {
        public final String courseId;
        public final String instructorId;
        public final String classroomId;
        public final TimeSlot timeSlot;
        public final List<Integer> weeks;

        public ScheduleEntry(String courseId, String instructorId,
                             String classroomId, TimeSlot timeSlot,
                             List<Integer> weeks) {
            this.courseId = courseId;
            this.instructorId = instructorId;
            this.classroomId = classroomId;
            this.timeSlot = timeSlot;
            this.weeks = weeks;
        }

        @Override
        public String toString() {
            return String.format("Course %s with Instructor %s in Room %s at %s Weeks %s",
                    courseId, instructorId, classroomId, timeSlot, weeks);
        }
    }

    // --- Scheduler ---
    public static class Scheduler {
        private final List<Classroom> classrooms;
        private final List<Instructor> instructors;
        private final List<Course> courses;
        private final List<ScheduleEntry> schedule = new ArrayList<>();

        public Scheduler(List<Classroom> classrooms,
                         List<Instructor> instructors,
                         List<Course> courses) {
            this.classrooms = classrooms;
            this.instructors = instructors;
            this.courses = courses;
        }

        public void run() {
            for (Course c : courses) {
                if (!scheduleCourse(c)) {
                    System.out.println("Failed to schedule: " + c.id);
                }
            }
        }

        private boolean scheduleCourse(Course course) {
            List<Classroom> candidates = findClassroomCandidates(course);
            for (Classroom room : candidates) {
                Optional<Slot> slot = findSlot(room, course);
                if (slot.isPresent()) {
                    Optional<Instructor> inst = findInstructor(course, slot.get());
                    if (inst.isPresent()) {
                        addEntry(course, room, inst.get(), slot.get());
                        return true;
                    }
                }
            }
            return false;
        }

        private List<Classroom> findClassroomCandidates(Course course) {
            return classrooms.stream()
                .filter(r -> {
                    boolean okSize = r.isVirtual
                        ? course.virtualOption != VirtualOption.NOT_ALLOWED
                        : (r.size >= course.enrollment && course.preferredCampuses.contains(r.campus));
                    return okSize;
                })
                .sorted(Comparator.comparing(r -> r.isVirtual))
                .collect(Collectors.toList());
        }

        private Optional<Slot> findSlot(Classroom room, Course course) {
            for (TimeSlot ts : course.preferredTimes) {
                if (!room.availability.contains(ts)) continue;
                List<Integer> weeks = IntStream.rangeClosed(1, 12)
                    .limit(course.repeats)
                    .boxed()
                    .collect(Collectors.toList());
                return Optional.of(new Slot(ts, weeks));
            }
            return Optional.empty();
        }

        private Optional<Instructor> findInstructor(Course course, Slot slot) {
            return instructors.stream()
                .filter(i -> i.allowedClasses.contains(course.id))
                .filter(i -> i.preferredCampuses.stream().anyMatch(course.preferredCampuses::contains))
                .filter(i -> isVirtualCompatible(i, course))
                .filter(i -> slot.weeks.stream()
                    .allMatch(w -> i.weeklyAvailability.getOrDefault(w, Collections.emptySet()).contains(slot.time)))
                .findFirst();
        }

        private boolean isVirtualCompatible(Instructor inst, Course course) {
            if (inst.virtualPref == InstructorVirtualPref.OPT_OUT
                    && (course.virtualOption == VirtualOption.PREFERRED
                        || course.virtualOption == VirtualOption.VIRTUAL_ONLY)) {
                return false;
            }
            return true;
        }

        private void addEntry(Course c, Classroom r, Instructor i, Slot s) {
            schedule.add(new ScheduleEntry(c.id, i.id, r.id, s.time, s.weeks));
        }

        public void printSchedule() {
            schedule.forEach(System.out::println);
        }

        private static class Slot {
            TimeSlot time;
            List<Integer> weeks;
            Slot(TimeSlot time, List<Integer> weeks) {
                this.time = time;
                this.weeks = weeks;
            }
        }
    }

    // --- Example main usage ---
    public static void main(String[] args) {
        // Build mock data here and invoke scheduler.run()
        // Then printSchedule()
    }
}

```

[Back](#java) | [Home](#home)


<br/><a name="javascript"></a>
### JavaScript

JavaScript (ES6) version:

* Defines **constants** for enums (TimeSlot, VirtualOption, InstructorVirtualPref)
* Uses **classes** to model `Classroom`, `Instructor`, `Course`, and `ScheduleEntry`
* Implements a **Scheduler** class with small helper methods:

  - `findClassroomCandidates`
  - `findSlot`
  - `findInstructor`
  - `isVirtualCompatible`

* Keeps the main flow in `run()` and `scheduleCourse()` minimal, avoiding large if/else blocks.

You can instantiate with real data in the commented “Example Usage” section, then call `scheduler.run()` and `scheduler.printSchedule()`. Let me know if you’d like sample data, testing harness, or export features!


```javascript
/*
  College Scheduling System in JavaScript
  - Models: Classroom, Instructor, Course, ScheduleEntry
  - Enums for TimeSlot, VirtualOption, InstructorVirtualPref
  - Scheduler class with small helper methods
*/

// --- Enums ---
const TimeSlot = Object.freeze({ MORNING: 'morning', AFTERNOON: 'afternoon', EVENING: 'evening' });
const VirtualOption = Object.freeze({ NOT_ALLOWED: 'not-allowed', VIRTUAL_ONLY: 'virtual-only', PREFERRED: 'preferred', OPTIONAL: 'optional' });
const InstructorVirtualPref = Object.freeze({ OPT_OUT: 'opt-out', PREFERRED: 'preferred', ALWAYS: 'always' });

// --- Models ---
class Classroom {
  constructor(id, campus, size, availability, isVirtual = false) {
    this.id = id;
    this.campus = campus;
    this.size = size;
    this.availability = new Set(availability);
    this.isVirtual = isVirtual;
  }
}

class Instructor {
  constructor(id, weeklyAvailability, allowedClasses, preferredCampuses, virtualPref) {
    this.id = id;
    // weeklyAvailability: { weekNumber: [TimeSlot,...], ... }
    this.weeklyAvailability = new Map(
      Object.entries(weeklyAvailability).map(([w, slots]) => [Number(w), new Set(slots)])
    );
    this.allowedClasses = new Set(allowedClasses);
    this.preferredCampuses = new Set(preferredCampuses);
    this.virtualPref = virtualPref;
  }
}

class Course {
  constructor(id, lengthHours, repeats, enrollment, preferredTimes, preferredCampuses, virtualOption) {
    this.id = id;
    this.lengthHours = lengthHours;
    this.repeats = repeats;
    this.enrollment = enrollment;
    this.preferredTimes = new Set(preferredTimes);
    this.preferredCampuses = new Set(preferredCampuses);
    this.virtualOption = virtualOption;
  }
}

class ScheduleEntry {
  constructor(courseId, instructorId, classroomId, timeSlot, weeks) {
    this.courseId = courseId;
    this.instructorId = instructorId;
    this.classroomId = classroomId;
    this.timeSlot = timeSlot;
    this.weeks = weeks;
  }
  toString() {
    return `Course ${this.courseId} with Instructor ${this.instructorId} in Room ${this.classroomId} at ${this.timeSlot} Weeks ${this.weeks}`;
  }
}

// --- Scheduler ---
class Scheduler {
  constructor(classrooms, instructors, courses) {
    this.classrooms = classrooms;
    this.instructors = instructors;
    this.courses = courses;
    this.schedule = [];
  }

  run() {
    this.courses.forEach(course => {
      if (!this.scheduleCourse(course)) {
        console.warn(`Could not schedule course ${course.id}`);
      }
    });
  }

  scheduleCourse(course) {
    const rooms = this.findClassroomCandidates(course);
    for (const room of rooms) {
      const slot = this.findSlot(room, course);
      if (!slot) continue;
      const inst = this.findInstructor(course, slot);
      if (!inst) continue;
      this.addEntry(course, room, inst, slot);
      return true;
    }
    return false;
  }

  findClassroomCandidates(course) {
    return this.classrooms
      .filter(room => {
        if (room.isVirtual) {
          return course.virtualOption !== VirtualOption.NOT_ALLOWED;
        } else {
          return room.size >= course.enrollment && course.preferredCampuses.has(room.campus);
        }
      })
      .sort((a, b) => a.isVirtual - b.isVirtual);
  }

  findSlot(room, course) {
    for (const ts of course.preferredTimes) {
      if (!room.availability.has(ts)) continue;
      // assume weeks 1-12 consecutive
      const weeks = Array.from({ length: course.repeats }, (_, i) => i + 1);
      return { timeSlot: ts, weeks };
    }
    return null;
  }

  findInstructor(course, slot) {
    return this.instructors.find(inst => {
      if (!inst.allowedClasses.has(course.id)) return false;
      if (![...inst.preferredCampuses].some(c => course.preferredCampuses.has(c))) return false;
      if (!this.isVirtualCompatible(inst, course)) return false;
      return slot.weeks.every(wk => inst.weeklyAvailability.get(wk)?.has(slot.timeSlot));
    });
  }

  isVirtualCompatible(inst, course) {
    if (inst.virtualPref === InstructorVirtualPref.OPT_OUT
        && [VirtualOption.PREFERRED, VirtualOption.VIRTUAL_ONLY].includes(course.virtualOption)) {
      return false;
    }
    return true;
  }

  addEntry(course, room, inst, slot) {
    this.schedule.push(
      new ScheduleEntry(course.id, inst.id, room.id, slot.timeSlot, slot.weeks)
    );
  }

  printSchedule() {
    this.schedule.forEach(entry => console.log(entry.toString()));
  }
}

// --- Example Usage ---
// const classrooms = [ new Classroom('R1', 'North', 30, [TimeSlot.MORNING, TimeSlot.AFTERNOON]), ... ];
// const instructors = [ new Instructor('I1', {1:['morning'],2:['morning', 'afternoon']}, ['C1'], ['North'], InstructorVirtualPref.PREFERRED), ... ];
// const courses = [ new Course('C1', 2, 12, 25, [TimeSlot.MORNING], ['North'], VirtualOption.OPTIONAL), ... ];
// const scheduler = new Scheduler(classrooms, instructors, courses);
// scheduler.run();
// scheduler.printSchedule();

```

[Back](#javascript) | [Home](#home)


<br/><a name="python"></a>
### Python

Here’s a modular scheduling system written in Python. It models classrooms, instructors, and classes with clear rules on physical/virtual resource use, instructor availability, and preferences.

You can fill in sample data for testing. Let me know if you'd like help generating mock data, adding constraints like priority weighting, or supporting conflict detection or GUI export (e.g., to CSV or HTML calendar).

```python
from dataclasses import dataclass, field
from typing import List, Dict, Optional, Tuple
from enum import Enum
import itertools

class TimeSlot(str, Enum):
    MORNING = 'morning'
    AFTERNOON = 'afternoon'
    EVENING = 'evening'

class VirtualOption(str, Enum):
    NOT_ALLOWED = 'not-allowed'
    VIRTUAL_ONLY = 'virtual-only'
    PREFERRED = 'preferred'
    OPTIONAL = 'optional'

class InstructorVirtualPref(str, Enum):
    OPT_OUT = 'opt-out'
    PREFERRED = 'preferred'
    ALWAYS = 'always'

@dataclass
class Classroom:
    id: str
    campus: str
    size: int
    availability: Dict[TimeSlot, bool]
    is_virtual: bool = False

@dataclass
class Instructor:
    id: str
    time_preferences: Dict[int, List[TimeSlot]]  # week -> times
    preferred_classes: List[str]
    preferred_locations: List[str]
    virtual_pref: InstructorVirtualPref

@dataclass
class Class:
    id: str
    length: int  # in hours
    repeats: int  # how many weeks it repeats
    enrollment: int
    preferred_time: List[TimeSlot]
    preferred_locations: List[str]
    virtual_option: VirtualOption

@dataclass
class ScheduleEntry:
    class_id: str
    instructor_id: str
    classroom_id: str
    campus: str
    timeslot: TimeSlot
    weeks: List[int]

class Scheduler:
    def __init__(self, classrooms: List[Classroom], instructors: List[Instructor], classes: List[Class]):
        self.classrooms = classrooms
        self.instructors = instructors
        self.classes = classes
        self.schedule: List[ScheduleEntry] = []

    def find_available_slots(self, classroom: Classroom, time_pref: List[TimeSlot], repeats: int) -> Optional[Tuple[TimeSlot, List[int]]]:
        for time in time_pref:
            if not classroom.availability.get(time, False):
                continue
            # Find weeks classroom is free (simplified logic for now)
            available_weeks = list(range(1, 13))  # All 12 weeks
            if len(available_weeks) >= repeats:
                return time, available_weeks[:repeats]
        return None

    def find_suitable_instructor(self, cls: Class, timeslot: TimeSlot, weeks: List[int]) -> Optional[Instructor]:
        for instructor in self.instructors:
            if cls.id not in instructor.preferred_classes:
                continue
            if not any(loc in instructor.preferred_locations for loc in cls.preferred_locations):
                continue
            if instructor.virtual_pref == InstructorVirtualPref.OPT_OUT and cls.virtual_option in [VirtualOption.PREFERRED, VirtualOption.VIRTUAL_ONLY]:
                continue
            if all(timeslot in instructor.time_preferences.get(week, []) for week in weeks):
                return instructor
        return None

    def schedule_class(self, cls: Class):
        candidates = sorted(
            [cr for cr in self.classrooms if (cr.is_virtual and cls.virtual_option != VirtualOption.NOT_ALLOWED) or
             (not cr.is_virtual and cr.size >= cls.enrollment and cr.campus in cls.preferred_locations)],
            key=lambda x: (x.is_virtual, x.size)
        )

        for cr in candidates:
            result = self.find_available_slots(cr, cls.preferred_time, cls.repeats)
            if result:
                timeslot, weeks = result
                instructor = self.find_suitable_instructor(cls, timeslot, weeks)
                if instructor:
                    self.schedule.append(ScheduleEntry(
                        class_id=cls.id,
                        instructor_id=instructor.id,
                        classroom_id=cr.id,
                        campus=cr.campus,
                        timeslot=timeslot,
                        weeks=weeks
                    ))
                    return True
        return False

    def run(self):
        for cls in self.classes:
            success = self.schedule_class(cls)
            if not success:
                print(f"Could not schedule class {cls.id}")

    def print_schedule(self):
        for entry in self.schedule:
            print(f"Class {entry.class_id} scheduled with Instructor {entry.instructor_id} in Room {entry.classroom_id} at {entry.timeslot} for weeks {entry.weeks}")

# Example usage (mock data can be filled in for test/demo purposes)
# scheduler = Scheduler(classrooms, instructors, classes)
# scheduler.run()
# scheduler.print_schedule()

```

[Back](#python) | [Home](#home)
