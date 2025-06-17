# Class Scheduler

> AI solutions by Deepseek.

  * Problem
  * Deep Thoughts
    - The main components
    - Approach
    - Steps
  * Implementation
  * Solution Code (python)
  * Explanation
  * Java Implementation
  * Summary

___
## Problem

Design a scheduling program for a university which has a few (less than 5) campuses, where each location has a limited numbers of classrooms plus any virtual classroom (each classroom has attributes like `size` allowing maximum numbers of students to fit) and availability of hours (morning 9 AM to 12 PM, afternoon 1 PM to 5 PM, and evening 6 PM to 10 PM).

For a specific quarter (12 weeks), there are a number of classes, which has attributes of `length` (assuming 1 or up to 2 hours) , `repeats` (less than 12 for each quarter, assuming only one in a week), `enrollment` (numbers of students), and preferred list of campus locations and classrooms (including virtual classroom, and could be "any").

The school has a number of instructors, and each of them has 3 preference lists, one for time availability of days across 12 weeks, one for preferred campus locations and class rooms (including virtual class option), and one for preferred classes (which allowed to teach).

Please write a program in, e.g. Python, to meet above requirements and balance the resources of instructors, campus locations, classrooms (e.g. only use virtual classroom if no physical classroom available), and time (e.g. only use evening if no available resources in morning or afternoon).
We are going to design a scheduling program for a university with multiple campuses and virtual classrooms.


___
## Deep Thought

### The main components:

* Campuses: each has classrooms (physical and virtual) with attributes: `size`, `availability` hours (morning, afternoon, evening)

* Classrooms: each has a campus, a room_id, size, and type (physical or virtual). Also, we assume each classroom has availability for each time slot (we'll model time slots as `(day, time_block)` over 12 weeks).

* Classes: each class has:
  - `id`
  - `length` (1 or 2 hours)
  - `repeats` (number of sessions per week, but note: total weeks=`12`, so total sessions = `repeats * 12`)
  - `enrollment` (number of students)
  - `preferred_campuses`: list of campus locations (or "any")
  - `preferred_classrooms`: list of classroom types or specific rooms? We assume it's a list of preferences (like "physical", "virtual", or "any") or specific room ids?

> Let's assume for simplicity: it can be a list of campus names and a separate list of room types? But the problem says "**preferred list of campus locations and classrooms (including virtual)**".

> We'll interpret: the class has a list of preferred campuses (by name) and a list of preferred classroom types (e.g., `['physical', 'virtual']` or just `['any']`).

* Instructors: each has:
  - `id`
  - `time_availability`: a mapping of day (e.g., Monday to Sunday) to time blocks (morning, afternoon, evening) that they are available?

    > But note: the problem says "**time availability of days across 12 weeks**". So we assume the availability is per day of week and time block, and it's the same for all 12 weeks?

  - `preferred_campuses`: list of campuses they prefer (including virtual? note: virtual doesn't require physical presence, so maybe they can be anywhere)
  - `preferred_classrooms`: list of classroom types? (physical, virtual) or specific rooms? We'll assume types.
  - `preferred_classes`: list of class ids they are allowed to teach.

* Time:

  We assume the week days are Monday to Sunday? But typically university classes are Monday to Friday? Let's assume Monday to Friday for now.

* Time blocks:
  - morning: 9 AM - 12 PM -> we can break into 1-hour slots: `[9-10, 10-11, 11-12]` (but note: class length can be 1 or 2 hours)
  - afternoon: 1 PM - 5 PM -> `[1-2, 2-3, 3-4, 4-5]`
  - evening: 6 PM - 10 PM -> `[6-7, 7-8, 8-9, 9-10]`

> However, a class of 2 hours would need two consecutive slots? Actually, we can model the day in 1-hour slots and then if a class is 2 hours, it requires two consecutive slots.

> But note: the problem says the class has a length of 1 or up to 2 hours. So we assume either 1 hour or 2 hours.

> We'll model time as:

  - Days: `0=Monday, 1=Tuesday, ... 6=Sunday` (but we might not use Sunday) -> let's use Monday to Friday (0-4) and Saturday (5) if needed?

    > The problem doesn't specify, so we'll assume Monday to Friday only.

  - Slots per day:

    - morning: 3 slots (9, 10, 11) -> each slot represents the start hour, and the class will run for 1 or 2 hours starting at that slot.
    - afternoon: 4 slots (13, 14, 15, 16) [using 24-hour time? but we can represent as 1 PM = 13, ...]
    - evening: 4 slots (18, 19, 20, 21)

> But note: a 2-hour class starting at 9 would run from 9-11 (covering 9 and 10 o'clock hours). However, our slots are at 9, 10, 11.

> If we start at 9 for 2 hours, we need slots 9 and 10. Starting at 10 would run 10-12, which is also okay? But 12 is still in the morning block?

> Actually, the morning block is until 12 PM, so 10-12 is okay.

> Similarly, in the afternoon: 1 PM (13) for 2 hours -> 13-15 (covers 13 and 14). Starting at 14 (2 PM) would run 14-16 (covers 14 and 15) -> but 16 is 4 PM, which is in the afternoon?

> The afternoon block is until 5 PM, so that's okay.

> However, we must note that a 2-hour class starting at 16 (4 PM) would run until 6 PM, which is evening? But the evening block starts at 6 PM?

> So we cannot start a 2-hour class at 4 PM because it would run until 6 PM and the next block starts at 6 PM? Actually, the problem says the afternoon is until 5 PM, so 4 PM to 6 PM would overlap the evening?

> We need to adjust.

> Let me redefine the time blocks:

> The problem states:

> morning: 9 AM to 12 PM -> 3 hours (so 3 slots of 1 hour, but note: 9-10, 10-11, 11-12 -> so 3 slots, but a 2-hour class starting at 10 would go to 12, which is okay? but 11-12 is the last slot, so starting at 11 for 1 hour is okay, but for 2 hours? 11-1? but that overlaps lunch? and the next block starts at 1 PM.

> Actually, the problem says the afternoon is 1 PM to 5 PM. So there is a gap from 12 PM to 1 PM.

> Therefore, we can model:

  - Morning: 9:00, 10:00, 11:00 -> each 1-hour slot.
  - Afternoon: 13:00, 14:00, 15:00, 16:00 -> note: 16:00 is 4 PM, and a 1-hour class at 16:00 ends at 17:00 (5 PM). A 2-hour class at 15:00 (3 PM) would run until 17:00 (5 PM) -> okay.

> But a 2-hour class at 16:00 (4 PM) would run until 18:00 (6 PM) -> which is the start of evening. But the evening block is 6 PM to 10 PM. So we can have a class from 4 PM to 6 PM?

> However, the problem says the afternoon block is until 5 PM. So we must not schedule beyond 5 PM in the afternoon?

> Therefore, we have to restrict:

  > In the afternoon, the last slot for a 2-hour class must be 15:00 (so that it ends at 17:00). Similarly, the last slot for a 1-hour class is 16:00 (ending at 17:00).

  > Evening: 18:00, 19:00, 20:00, 21:00 -> a 2-hour class at 21:00 would run until 23:00? but the block is until 10 PM -> so we cannot start a 2-hour class at 21:00?

> Therefore, the last slot for a 2-hour class in the evening is 20:00 (ending at 22:00? but 10 PM is 22:00? Actually, 6 PM to 10 PM is 4 hours -> 18:00 to 22:00?

> But the problem says "evening 6 PM to 10 PM", so 4 hours. Therefore, we can have:

  - 1-hour: 18, 19, 20, 21 -> ending at 19,20,21,22? but 22 is 10 PM -> so 21:00 is acceptable for 1 hour?

> But 21:00 to 22:00 is within 10 PM?

> However, a 2-hour class at 20:00 would run until 22:00 -> which is acceptable?

> Therefore, we can define the time slots with constraints:

  - Morning (9:00, 10:00, 11:00):
    - 1-hour: any slot
    - 2-hour: only 9:00 and 10:00 (because 11:00 would run until 13:00, which is after the morning block and into the afternoon block? but we have a gap?

  > However, the problem says the morning block ends at 12 PM, so 11:00-12:00 is the last hour. A 2-hour class at 10:00 would be 10:00-12:00 -> acceptable.

  - Afternoon (13:00, 14:00, 15:00, 16:00):
    - 1-hour: all four
    - 2-hour: only 13:00, 14:00, 15:00 (because 16:00 would run until 18:00, which is in the evening block? and the problem says we should avoid evening if possible)

  - Evening (18:00, 19:00, 20:00, 21:00):
    - 1-hour: all four
    - 2-hour: 18:00, 19:00, 20:00 (because 21:00 would run until 23:00, which is beyond 10 PM? but 21:00 to 23:00 is 2 hours?

  > Actually, 21:00 to 22:00 is one hour, so 2-hour at 21:00 would be 21:00-23:00 -> not acceptable. Therefore, we only allow 18:00, 19:00, 20:00 for 2-hour.

> Alternatively, we can model the day as a sequence of 1-hour slots and then check the entire block?

> However, to simplify, we can break the day into 1-hour slots and then for a class of length L, we require L consecutive slots in the same block?

> But note: the problem says the availability is by blocks (morning, afternoon, evening). So we have to assign a class to a block? and then within the block, consecutive slots?

> How about we model:

  ```
  TimeBlock = namedtuple('TimeBlock', ['block_name', 'start_hour', 'end_hour'])
  ```

  - morning: 9-12 -> 3 hours (so 3 slots)
  - afternoon: 13-17 -> 4 hours (4 slots)
  - evening: 18-22 -> 4 hours (4 slots)

> Then we break each block into 1-hour slots. But note: the afternoon block is 4 hours? Actually, 1 PM to 5 PM is 4 hours? 1,2,3,4,5 -> 5 hours?

> But 1 PM to 2 PM is first hour, 2-3, 3-4, 4-5 -> 4 hours?

> So we have:

  - morning: 9,10,11 (3 slots) -> 9-10, 10-11, 11-12
  - afternoon: 13,14,15,16 (4 slots) -> 13-14,14-15,15-16,16-17
  - evening: 18,19,20,21 (4 slots) -> 18-19,19-20,20-21,21-22

> Now, for a class of 2 hours, it will require two consecutive slots in the same block?

> But note: a class starting at 11 in the morning would run until 13? which is the next block? That is not allowed. So we cannot start a 2-hour class at 11 in the morning?

> Similarly, in the afternoon: 16:00 to 18:00 would go into the evening? So we cannot start a 2-hour class at 16:00?

> In the evening: 21:00 to 23:00 -> but the block ends at 22:00? So we cannot start a 2-hour class at 21:00?

> Therefore, the available starting slots for 2-hour classes are:

  - morning: 9,10 -> because 11 would run until 13 (which is afternoon) -> not allowed.
  - afternoon: 13,14,15 -> because 16 would run until 18 (which is evening) -> not allowed.
  - evening: 18,19,20 -> because 21 would run until 23 -> not allowed.

> So we can precompute for each block the available starting slots for 1-hour and 2-hour classes.

### Approach:

We have:

  - 12 weeks (each week has 5 days: Monday to Friday) -> total 60 days? but we don't need to schedule every day? Actually, the class repeats per week: so we schedule one occurrence per week?

We can model the schedule as:

We have a set of classes and we need to assign:

  - An instructor
  - A classroom (on a campus) that is available and has size >= enrollment
  - A set of time slots (each session of the class) for each week: so we need to pick a day of the week and a starting slot (within a block) and then the consecutive slots (if 2 hours) for each week for 12 weeks.

However, the same class must meet at the same time and same classroom each week?

So for a class, we assign:

  - One day of the week (e.g., Monday)
  - One starting time (e.g., 9:00 AM) and then the length determines the end time.
  - One classroom

And then for 12 weeks, every Monday at 9:00 AM in that classroom.

But note: the instructor must be available at that time every week?

So we can model the time as: (day_of_week, start_slot) and then the session will occur on that day of the week for each of the 12 weeks.

### Steps:

We need to assign each class:

  - instructor
  - classroom
  - (day_of_week, start_time)

**Constraints**:

(1) The classroom must be in one of the preferred campuses (or "any" means any campus) and of preferred type (or "any").

And the classroom size must be at least the enrollment.

(2) The instructor must:

- be available at the chosen (day_of_week, time_block) -> note: the time_block is morning, afternoon, evening?

But we have the actual hour? How do we map?

Actually, the instructor's availability is given for time availability of days? We assume we have a set of (day, time_block) that the instructor is available?

And the time_block is the block (morning, afternoon, evening) that they are available?

So if the class is scheduled in the morning on Monday, then we require the instructor to be available on Monday morning.

(3) The classroom must be available at that (day_of_week, start_time) for the entire duration? and for all 12 weeks?

Since we are assigning the same time and same classroom for every week, we only need to check the weekly pattern.

(4) The class length must be compatible with the starting slot (as discussed above: for 2-hour, we must not start at the last slot of a block).

(5) The same instructor cannot teach two classes at the same time (same day and same time slot) in the same week?

But note: the same instructor might teach multiple classes, but they must not overlap in time.

(6) The same classroom cannot be used by two classes at the same time.

(7) We should balance resources:
    - Prefer physical classrooms over virtual.
    - Prefer morning and afternoon over evening.

How to model instructor availability?

We'll assume each instructor has a set of available (day, block) pairs. For example:

  `instructor.availability = { (Monday, 'morning'), (Monday, 'afternoon'), ... }`

But note: the class is scheduled at a specific start time (hour) and that falls in a block?

We can map the start time to block:

  ```
  if start_time in [9,10,11] -> block = 'morning'
  if start_time in [13,14,15,16] -> block = 'afternoon'
  if start_time in [18,19,20,21] -> block = 'evening'
  ```

And then we require that the instructor is available on that day and block.

Also, the classroom has availability? The problem says "availability of hours" per classroom?

We can assume that a classroom is available for a time block (morning, afternoon, evening) on a given day?

But note: the problem says "**availability of hours (morning, ...)**" -> meaning each classroom is available in those blocks?

So we can model classroom availability as a set of `(day, block)` as well?

However, the problem does not specify that a classroom is not available on certain days?

We assume that every classroom is available for every day and every block? unless specified otherwise?

But the problem says "each location has a limited numbers of classrooms plus any virtual classroom ... and availability of hours".

So we interpret: each classroom is available only during the blocks that its campus offers?

For example, a campus might say: we have classrooms available in morning and afternoon?

Actually, the problem says the campus has availability of hours (morning, afternoon, evening). So we can assume that every classroom in that campus has the same availability?

Therefore, we can model the campus availability: set of blocks (e.g., `['morning','afternoon','evening']`) and then the classroom inherits that?

But note: virtual classroom? It might be available 24/7? The problem says "any virtual classroom", and we are to use virtual only if no physical is available?

So we can assume that virtual classrooms are available for all blocks?

Therefore, for a physical classroom, we check the campus's available blocks?

We'll model:

  * `Campus`:

    - `name`: `str`
    - `available_blocks`: list of blocks (e.g., `['morning','afternoon','evening']`)
    - `classrooms`: list of `Classroom` objects

  * `Classroom`:

    - `id`: `str`
    - `campus`: Campus or None for virtual?
    - `size`: `int`
    - `type`: `'physical'` or `'virtual'`

Then, the availability of a classroom on a given (day, block) is:
  > - If the classroom is virtual: available for any day and any block?

  > - If the classroom is physical: available only if the campus has the block? and we assume every day?

However, what if the campus is closed on a particular day? The problem does not specify. We assume the campus is open every day?

So we can skip the day?

Therefore, for a classroom, we can define:

  * `available_blocks`:
    - if physical: same as the campus.available_blocks
    - if virtual: `['morning','afternoon','evening']`
      > but note: the virtual classroom might be available 24/7, but we are only scheduling in the three blocks?

So we assume the virtual classroom is available in all three blocks.

Then, for a class to be scheduled at (day, start_time) we must:

- Map the start_time to a block (as above)

- Check that the classroom is available for that block (on that day? we ignore day because we assume every day the same) -> so just the block.

But note: what if a campus is closed on a particular day? We don't model that. We assume the campus is open every day?

Therefore, we only check the block.

**Summary of constraints** for a class assignment (`instructor`, `classroom`, `day`, `start_time`):

Let **`C`** be a class.

1. **Classroom** must be:
    a. In the preferred campuses of `C` (if `C` has 'any', then any campus is okay) OR if the classroom is virtual, then we don't care about campus?

  > Actually, the class has a preferred list of campuses and classrooms (which may include virtual).

  > So we require:
  >  - The classroom must be either virtual OR physical in one of the preferred campuses (if the class has a non-any list) OR if the class has 'any', then any physical campus is okay?

  > How we model:

  > The class has:
    - `preferred_campuses`: list of campus names or the string "any"
    - `preferred_classroom_types`: list of strings, e.g., [`'physical','virtual']` or `['physical']` or [`'virtual']` or `['any']`?

  > Actually, the problem says: "preferred list of campus locations and classrooms (including virtual classroom, and could be 'any')"

  > We'll interpret:
  >   - The class has a list of campuses it prefers (could be a list of campus names, and if one of them is "any", then we treat as any campus) OR if the list is empty?

  > But the problem says "could be any", so we can assume that if the list contains "any", then any campus is acceptable?

  >   - Similarly, for classroom types: the class has a list of preferred classroom types (like ['physical','virtual'] or ['physical'] or ['virtual']).

  > If the list contains "any", then both are acceptable.

  > Therefore, we can have:

    > * classroom_campus_ok =
        if classroom is virtual -> then we don't require the campus? because virtual doesn't belong to a physical campus?

    > but note: the problem doesn't say that virtual is tied to a campus?

  > Let's assume virtual classroom is not tied to a campus. So for a virtual classroom, we skip the campus check?

  > However, the class may have a preference for virtual? so if the class has a preferred campus list that does not include virtual?

  > Actually, the class's preferred campuses are for physical locations? and then they separately can say they want virtual?

  > We'll assume the class has:
  >   * `preferred_campuses`: for physical campuses. And then separately, they can say they prefer virtual? via the preferred_classroom_types.

  > Therefore, we check:
  >   - If the classroom is physical: then its campus must be in the class's preferred_campuses (if the class's preferred_campuses is not "any") OR if the class has "any", then it's okay.
  >   - If the classroom is virtual: then we require that the class's preferred_classroom_types includes 'virtual' or 'any'.

  > Also, for classroom type:
  >   - The classroom's type must be in the class's preferred_classroom_types, or if the class has 'any' in the preferred_classroom_types, then it's okay?

  > But note: the class's preferred_classroom_types might be a list without 'any'.

  > So we can do:
    ```
    type_ok = (classroom.type in class.preferred_classroom_types) or ('any' in class.preferred_classroom_types)
    ```

  > And campus_ok =
    ```
    if classroom.type == 'physical':
      (classroom.campus.name in class.preferred_campuses) or ('any' in class.preferred_campuses)
    else:  # virtual
      True # because we already checked the type? and the class must have accepted virtual? but wait: the class might not have virtual in preferred_classroom_types?
    ```

  > Actually, we are checking `type_ok` above, so if the class doesn't want virtual, then we wouldn't assign a virtual room.

2. **Size**: `classroom.size >= class.enrollment`

3. **Instructor**:
  - The instructor must have the class in their preferred_classes (list of class ids) OR if they have 'any', then any class?
  > But the problem says "preferred classes (which allowed to teach)" -> so we assume the list of class ids they are allowed to teach?
  - The instructor must be available on the chosen (day, block) that we derive from the start_time.

4. **Classroom availability**:
  - The classroom must be available in the block (as determined by the start_time) -> we have the block from the start_time, then check the classroom.available_blocks includes that block.

5. **Time constraints** for the class length and block:
  - We must choose a start_time that is compatible with the class length and the block.

  > Specifically, for a 2-hour class, the start_time must be one that allows two consecutive slots in the same block without crossing the block boundary.

  > We precomputed:
  >   - morning: 2-hour starts: [9,10]
  >   - afternoon: [13,14,15]
  >   - evening: [18,19,20]

  - We can predefine a list of valid start times for each block and class length.

6. **Conflict checks**:

  a. For the instructor:

  > We are assigning the class to a specific (day, start_time) for every week. Then we must ensure that for that instructor, at the same (day, start_time) in any week, they are not teaching another class?

  > But note: we are scheduling the same class for 12 weeks at the same time. And we are assigning multiple classes.

  > So we need to check for the instructor: no two classes are assigned to the same (day, start_time).

  b. For the classroom:

  > Similarly, no two classes can be in the same classroom at the same (day, start_time) in the same week?

  > However, note: two classes at the same (day, start_time) but in different weeks? -> no, because we are scheduling every week?

  > So if two classes are scheduled at the same (day, start_time) and in the same classroom, then every week that classroom is occupied at that time by both classes? -> conflict.

Therefore, we model:

We have a schedule that is defined over:
  - 5 days (Monday to Friday)
  - The time slots: `[9,10,11,13,14,15,16,18,19,20,21]`
  > note: we skip the ones that are not allowed for 2-hour? but we are going to check per class length?

Actually, we can represent all the possible start times? and then for each class, we only consider the valid ones for its length and block?

We will create a data structure to track:

  - `instructor_schedule[instructor_id]` = set of (day, start_time) that the instructor is occupied.
  - `classroom_schedule[classroom_id]` = set of (day, start_time) that the classroom is occupied.

But note: if two classes are scheduled at the same (day, start_time) but one is 1-hour and the other is 2-hour, then they might overlap?

For example, class A at 9:00 (1 hour) ends at 10:00. Class B at 10:00 (1 hour) is okay?

But if class A is 2 hours (9-11), then class B at 10:00 (1 hour) would run from 10-11 -> which overlaps with class A?

Therefore, we cannot just check the start time. We need to check the entire time interval.

How about we model the time as 1-hour slots?

We have the day divided into 1-hour slots. For a class, we require:
  - For a 1-hour class at `start_time` s: it occupies the slot `s`.
  - For a 2-hour class at `start_time` s: it occupies s and `s+1`?

> But note: the `start_time` is an integer?

We can assign numbers to the slots:

Let the slots be:
    ```
    9 -> 0
    10 -> 1
    11 -> 2
    13 -> 3
    14 -> 4
    15 -> 5
    16 -> 6
    18 -> 7
    19 -> 8
    20 -> 9
    21 -> 10
    ```

Then the duration:
  - 1-hour: occupies one slot: `[s]`
  - 2-hour: occupies `[s, s+1]`

> But note: the gap between 12 and 13: so 11 (which is 2) and 13 (which is 3) are not consecutive?

We need to represent the day as a timeline?

Alternatively, we can avoid the gap by having a timeline that is not contiguous?

We can model the day as having 11 possible start times? and then for each class, we know the set of time slots (by index) that it occupies?

> But the indices are not contiguous?

Instead, we can represent the day as:

We have 11 time slots? but they are not consecutive in the day?

We can break the day into three segments?

How about we model the schedule per day and per block? and within a block, we have contiguous slots?

Then, for a class in a block, we know the starting slot index within the block?

For example, in the morning block, we have 3 slots: [0,1,2] corresponding to 9,10,11. Then
  - A 1-hour class at slot0 (9) occupies [0].
  - A 2-hour class at slot0 (9) occupies [0,1] -> which is valid?
  - A 2-hour class at slot1 (10) occupies [1,2] -> valid?
  - A 2-hour class at slot2 (11) would require [2,3] -> but 3 is beyond the block? -> invalid.

Similarly, in the afternoon: 4 slots [0,1,2,3] for 13,14,15,16.
  - 2-hour: at slot0 -> [0,1] -> valid; slot1 -> [1,2] -> valid; slot2 -> [2,3] -> valid; slot3 -> [3,4] -> invalid (because beyond the block).

> But note: the block is defined by the campus and we assume the classroom is available for the whole block?

Therefore, we can model:

We break the day into blocks. Then within a block, we have a fixed number of slots.

Then, for a class scheduled at (day, block, start_slot_index) with length L (in hours) -> it will occupy L consecutive slots in that block, starting at start_slot_index.

**Constraints***:

  ```
  start_slot_index + L <= total_slots_in_block
  ```

Then the global time identifier can be (day, block, slot_index) but note: we don't need the absolute time?

However, when we check for instructor and classroom, we need to know if there is an overlap?

But if two classes are in different blocks, they don't overlap?

But what if they are in the same block? Then we can have:

We can represent the occupation of an instructor (or classroom) as a set of (day, block, slot) for every slot they are occupied.

  - For a 1-hour class at (day, block, start_slot): occupies (day, block, start_slot)
  - For a 2-hour class: occupies (day, block, start_slot) and (day, block, start_slot+1)

Then, we can check for an instructor:

  - If we are trying to assign a class at (day, block, start_slot) for L hours, then we require that for each slot in [start_slot, start_slot+L-1],
    > the instructor is not already occupied at (day, block, slot) by another class.

Similarly for the classroom.

> But note: the same instructor might have two classes in the same block but at different slots? as long as the slots don't overlap?

For example, a class at slot0 (1-hour) and then a class at slot1 (1-hour) in the same block?

They are consecutive but do not overlap?

However, if the first class is 2-hour at slot0, then it occupies slot0 and slot1. Then a class at slot1 (even if 1-hour) would be at the same time? -> conflict.

Therefore, we must check all the slots the class will occupy.

So we change our representation:
  - We'll let the time assignment for a class be: (day, block, start_slot_index)
  - And then the class will occupy the set: {(day, block, start_slot_index + i) for i in range(length)}

Constraints for the block and length:

  * We must have start_slot_index + length <= total_slots_in_block[block]

    ```
    total_slots_in_block = {
      'morning': 3,
      'afternoon': 4,
      'evening': 4
    }
    ```

And the valid start_slot_index:
  * for a class of length `L`: `start_slot_index` can be from `0` to `total_slots_in_block[block] - L`.

This automatically rules out the cases that would cross the block boundary?

Now, the instructor availability: we are given (day, block) available.

We don't have slot-level availability?

> But note: the instructor might be available for the whole block? So if we assign a class that uses part of the block, that's fine?

However, we are only requiring the instructor to be available for the block? and then we can schedule multiple classes in the same block? as long as they don't overlap?

So we don't need slot-level availability from the instructor?

Therefore, we can assign a class to an instructor in a block if the instructor is available for that (day, block) and the specific slots we choose are free (i.e., not occupied by other classes of the same instructor).

Similarly, the classroom: we only require that the classroom is available for the block? and then we check the specific slots are free.

**Steps for the algorithm**:

We have a list of classes to schedule.

We want to assign to each class: instructor, classroom, and (day, block, start_slot_index)

We have to balance resources:
  - Prefer physical classrooms over virtual?
  - Prefer morning and afternoon over evening?

We can use a heuristic: assign the classes in order of difficulty? (e.g., by enrollment descending? because large classes are harder to schedule)

We'll try:
  - Sort classes by enrollment descending?
  - Then for each class:

We'll try to assign in the following order:
  - First, try physical classrooms
  - Then virtual
  - For `time`: try morning, then afternoon, then evening.
  - For `days`: we can try in order Monday to Friday?
  - For `start_slot_index`: we can try in increasing order?

But we have to choose an instructor as well.

We'll iterate:
  - For each candidate instructor (who is allowed to teach this class) and who has at least one available (day, block) that we haven't tried?

  - We can try to assign the class to the instructor on a day and block that the instructor is available?

  - Then for each candidate classroom (that meets the class's preferences and size) and that is available in the block?

  - Then for each day (Monday to Friday) and each `start_slot_index` in the block that fits the class length?

  - And then check if the instructor is free at the required (day, block, slots) and the classroom is free at the same?

  - We do this in the order of time preference: morning first, then afternoon, then evening.

  - And classroom: physical first, then virtual.

  - And within the same time block, we try the days in order?

However, this might be too slow? But the problem says the university has a few campuses (less than 5) and the number of classes is not specified?

We assume the number of classes is moderate (say 100?).

But the triple loop: instructor (say up to 100), classroom (say 50), day (5), block (3), `start_slot (max 4) -> 100*50*5*3*4 = 300000` per class? and 100 classes -> `30e6` -> acceptable in Python?

But worst-case it might be acceptable?

Alternatively, we can precompute for each instructor and classroom the occupied slots?

We'll represent:

  - instructor_occupation[instructor_id] = a set of (day, block, slot) that the instructor is teaching.
  - classroom_occupation[classroom_id] = a set of (day, block, slot)

Then for a candidate assignment for class C with instructor I, classroom R, at (day, block, start_slot) and length L:

  - required_slots = {(day, block, start_slot + i) for i in range(L)}

  - if required_slots is disjoint from instructor_occupation[I] and disjoint from classroom_occupation[R]:
    then assign:
    ```
    instructor_occupation[I] = instructor_occupation[I] | required_slots

    classroom_occupation[R] = classroom_occupation[R] | required_slots
    ```
    and record the assignment for class C.

    If we find one, we break and move to next class.

    If not, we try the next candidate.

We also have to consider that an instructor might be available for multiple `(day,block)` and we are iterating over days and blocks?

How we iterate:
  - For class in classes_sorted:
    * for block_order in [morning, afternoon, evening]:
    * for day in [Monday, Tuesday, Wednesday, Thursday, Friday]:
    * for instructor in candidate_instructors (that are available on (day, block_order) and that have the class in their preferred_classes) and that haven't been tried for every slot?
    * for classroom in candidate_classrooms (that are of type physical first, then virtual) and that meet the class preferences and size?
    * for start_slot in range(0, total_slots[block_order] - class.length + 1):

    > check conflict for instructor and classroom for the required_slots = (day, block_order, start_slot) ... with length class.length

> If we don't find in physical, then we try virtual?

> But note: we are iterating by block_order: so we do all mornings first, then afternoons, then evenings? and within morning, we try all days?

We break as soon as we find one.

However, we might want to balance the instructor workload?

But the problem doesn't specify. We just want to schedule all classes.

If we cannot schedule a class, we return an error.

### Implementation:

We'll have the **defination**:

  ```
  days = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'] # (0 to 4)
  blocks = {
    'morning': {'total_slots': 3, 'order': 0},
    'afternoon': {'total_slots': 4, 'order': 1},
    'evening': {'total_slots': 4, 'order': 2}
  }
  ```

We'll order blocks by their preference: first morning (order 0), then afternoon (1), then evening (2).

#### **Classes**:

  ```
  class Class:
    id: str
    length: int  # 1 or 2
    repeats: int  # but note: we are scheduling one session per week? so repeats per week? but the problem says repeats (less than 12 for each quarter) and one in a week?
  ```

  > Actually, the problem says "repeats (less than 12 for each quarter, assuming only one in a week)" -> so repeats = number of sessions per week?

  > But wait: it says "only one in a week", so repeats=1?

  > Let me read: "repeats (less than 12 for each quarter, assuming only one in a week)" -> so it means the class meets once per week? and the quarter has 12 weeks?

  > So repeats is the number of times per week? but it says only one in a week?

  > I think the problem means: each class meets once per week? so repeats=1?

  > But then why say "less than 12 for each quarter"?

  > Actually, the total number of sessions per quarter is 12?

  > So I think the class has an attribute: frequency per week? but the problem says "repeats (less than 12 for each quarter, assuming only one in a week)" ->
    > so it means that the class meets once per week? and the total sessions in the quarter is 12?

  > Therefore, we can ignore the "repeats" and assume every class meets once per week?

  > But what if a class meets twice per week?

  > The problem says "assuming only one in a week", so we assume one session per week?

  > So we schedule one session per week for 12 weeks?

  > Therefore, we assign one (day, block, start_slot) per class?

  > (continued `class Class`)

  ```
    enrollment: int
    preferred_campuses: # list of str (campus names) OR if the list contains "any", then any physical campus is acceptable?
    preferred_classroom_types: # list of str, e.g., ['physical','virtual'] or ['physical'] or ['virtual'] or ['any']?
  ```

#### **Instructors**:

  ```
  class Instructor:
    id: str
    availability: set of (day, block) # days are in ['Monday',...], blocks in ['morning','afternoon','evening']
    preferred_campuses: list of str?  # but note: we are using this?
    preferred_classroom_types: list of str?
    preferred_classes: list of class ids? # or if it contains 'any', then any class?
  ```

  > However, note: the problem says the instructor has 3 preference lists. We are using the time availability and the preferred classes.

  > The preferred campuses and classrooms for the instructor: we haven't used?

  > The problem says: "one for time availability ..., one for preferred campus locations and class rooms, and one for preferred classes"

  > We can use the instructor's preferred campuses and classrooms as a filter?

  > For example, when we assign a classroom to a class, we must also satisfy the instructor's preference?

  > So we add:

    > For the classroom:
    > * The instructor must also be okay with the campus and classroom type?
    > * Specifically, if the classroom is physical, then the campus must be in the instructor's preferred_campuses (or if the instructor has 'any', then okay)
    > * and the classroom type must be in the instructor's preferred_classroom_types (or 'any').
    > For virtual: the classroom type must be in the instructor's preferred_classroom_types.

  > So we update the classroom candidate filter:

  > For a candidate classroom R and candidate instructor I:

  > - if R is physical:
    (R.campus.name in I.preferred_campuses or 'any' in I.preferred_campuses) and
    (R.type in I.preferred_classroom_types or 'any' in I.preferred_classroom_types)

  > - else: # virtual
    (R.type in I.preferred_classroom_types or 'any' in I.preferred_classroom_types)

#### **Campuses and Classrooms**:

  ```
  class Campus:
    name: str
    available_blocks: list of str   # e.g., ['morning','afternoon','evening'] or a subset

  class Classroom:
    id: str
    campus: Campus or None # None for virtual?
    size: int
    type: str   # 'physical' or 'virtual'
  ```

  > We can also compute available_blocks:

  > - if type=='virtual':
  >     available_blocks = ['morning','afternoon','evening']
  > - else:
        available_blocks = campus.available_blocks

#### **Steps**:

**Precomputation**:

  * We'll create a list of all classrooms.

  * We'll initialize:

    - instructor_occupation: dict of instructor_id -> set of (day, block, slot_index)
    - classroom_occupation: dict of classroom_id -> set of (day, block, slot_index)

  * We'll also create a list of assignments for classes.

  * Then sort classes by enrollment descending?

  * For each class in sorted_classes:

    - candidate_instructors = [instructor for instructor in instructors if
      (class.id in instructor.preferred_classes or 'any' in instructor.preferred_classes) ]

  * Then, we try:

    - for block in ['morning','afternoon','evening']: # in that order
    - for day in days:   # Monday to Friday

      > candidate instructors available on (day, block)

    - available_instructors = [instructor for instructor in candidate_instructors

    - if (day, block) in instructor.availability]
      > candidate classrooms:
      >   We want to try physical first, then virtual.
      candidate_classrooms_physical = [room for room in classrooms

    - if room.type == 'physical'
        and room.size >= class.enrollment
        and (block in room.available_blocks)
        and (class.preferred_campuses is 'any' or (room.campus.name in class.preferred_campuses) or 'any' in class.preferred_campuses)   # for the class campus preference
        and (room.type in class.preferred_classroom_types or 'any' in class.preferred_classroom_types)   # for the class room type
      ]

      candidate_classrooms_virtual = [room for room in classrooms

  - if room.type == 'virtual'
      and room.size >= class.enrollment
      and (block in room.available_blocks)
      and ('virtual' in class.preferred_classroom_types or 'any' in class.preferred_classroom_types)   # for the class room type
    ]

    candidate_classrooms = candidate_classrooms_physical + candidate_classrooms_virtual

  > Now, for each available_instructor and each candidate_classroom, and for each start_slot in the valid range for the block and class.length:

    found = False

    for instructor in available_instructors:
      for room in candidate_classrooms:

  > Also, check the instructor's preference for the room?

  > We require that the room meets the instructor's preference for campus and classroom type?

  - if room.type == 'physical':

  - if not ((room.campus.name in instructor.preferred_campuses or 'any' in instructor.preferred_campuses) and
      (room.type in instructor.preferred_classroom_types or 'any' in instructor.preferred_classroom_types)):

      continue

  - else:
      - if not (room.type in instructor.preferred_classroom_types or 'any' in instructor.preferred_classroom_types):

        continue

    > Check the room's available_blocks already includes the block? we did that above?

    > Now, for the block, the valid start_slot indices: from 0 to (total_slots[block] - class.length)

    > total_slots = blocks[block]['total_slots']

    > for start_slot in range(0, total_slots - class.length + 1):
    >     required_slots = set()

    > for i in range(class.length):
    >     slot_index = start_slot + i
          required_slots.add((day, block, slot_index))

  > Check if instructor is free at all required_slots?

  - if instructor_occupation[instructor.id] & required_slots:
      > conflict

      continue

  > Check if room is free at all required_slots?

  - if classroom_occupation[room.id] & required_slots:

      continue

  > If we get here, we can assign

  > Update the occupation sets

      instructor_occupation[instructor.id] |= required_slots

      classroom_occupation[room.id] |= required_slots

  > Record the assignment for the class: (instructor.id, room.id, day, block, start_slot)

    ```
    assignment = {
      'class_id': class.id,
      'instructor_id': instructor.id,
      'classroom_id': room.id,
      'day': day,
      'block': block,
      'start_slot': start_slot
    }
    assignments.append(assignment)
    found = True

    break   # break out of start_slot loop

    if found:
      break   # break out of room loop

    if found:
      break   # break out of instructor loop

    if found:
      break   # break out of day loop

    if found:
      break   # break out of block loop
    ```

  > If we didn't find an assignment for the class, we try the next block?

  > But we are breaking as soon as we find?

  > If we didn't break (i.e., not found), then we have an error: class not scheduled.

  > We'll then output the assignments.

  > Note: the assignment for the class is for one session per week? and we assume the same schedule for all 12 weeks?

  > But note: the problem says the quarter is 12 weeks, and the class meets once per week? so the assignment is the same for each week?

  > Therefore, we only assign one session per class?

  > But note: what if a class meets multiple times per week?

  > The problem says "assuming only one in a week", so we assume one session per week.

  > Let's code accordingly.

  > We'll create classes for Campus, Classroom, Class, Instructor.

  > Then run the scheduling algorithm.

  > Note: the virtual classroom: we assume there is one virtual classroom? or multiple?

  > The problem says "any virtual classroom", and that we have limited physical classrooms.

  > We can assume there is one virtual classroom that can host multiple classes at the same time?

  > But wait: if we have two classes at the same time in the same virtual classroom, that is possible?

  > The problem doesn't specify.

  > However, the problem says "classroom" and we are tracking occupation.

  > We assume that the virtual classroom is a single resource? meaning one meeting at a time?

  > But in reality, a virtual meeting platform might allow multiple meetings?

  > However, the problem says "classroom" and we are to balance resources.

  > We'll treat virtual classrooms as individual classrooms?

  > How many virtual classrooms do we have?

  > The problem says "any virtual classroom" and we should use virtual only when physical is not available.

  > So we can have multiple virtual classrooms?

  > We'll assume the university has an unlimited supply of virtual classrooms?

  > But wait, the problem says "limited numbers of classrooms plus any virtual classroom".

  > So we interpret: the virtual classroom is a separate type and we have one virtual classroom?

  > Or we can have multiple virtual classrooms?

  > The problem does not specify a limit.

  > We'll assume there is one virtual classroom?

  > But then if we have many classes, we might run out of virtual classroom slots?

  > Alternatively, we can assume virtual classrooms are unlimited?

  > But then why schedule?

  > The problem says "balance the resources" and "only use virtual classroom if no physical classroom available".

  > So we assume there is one virtual classroom?

  > However, one virtual classroom can host multiple classes at the same time?

  > The problem doesn't specify.

  > Let me think: in reality, a virtual meeting can only host one class at a time?

  > So we treat a virtual classroom as a single resource? meaning one class at a time?

  > Therefore, we'll create one virtual classroom?

  > But what if we have two classes at the same time?

  > We'll create multiple virtual classrooms?

  > The problem says "any virtual classroom", which implies that there might be more than one?

  > We'll let the input specify: the university has a list of classrooms, and some of them are virtual.

  > So we can have multiple virtual classrooms?

  > For example, the university might have 5 virtual classrooms?

  > How do we know the number?

  > We'll let the input have a list of classrooms.

  > Therefore, we'll create:
  > - physical classrooms: with a campus and size
  > - virtual classrooms: we can create multiple?

  > We'll assume the input provides the classrooms.

  > Let's design the input and then the classes.

  > Example input:

  ```
  campuses = [
    Campus(name='CampusA', available_blocks=['morning','afternoon']),
    Campus(name='CampusB', available_blocks=['afternoon','evening']),
  ]

  classrooms = [
    # physical
    Classroom(id='A101', campus=campuses[0], size=50, type='physical'),
    Classroom(id='A102', campus=campuses[0], size=30, type='physical'),
    Classroom(id='B201', campus=campuses[1], size=40, type='physical'),

    # virtual
    Classroom(id='Virtual1', campus=None, size=100, type='virtual'),
    Classroom(id='Virtual2', campus=None, size=100, type='virtual')   # we have two virtual classrooms?
  ]

  classes = [
    Class(id='C1', length=1, enrollment=20,
      preferred_campuses=['CampusA'],
      preferred_classroom_types=['physical']),

    Class(id='C2', length=2, enrollment=45,
      preferred_campuses=['any'],
      preferred_classroom_types=['physical','virtual']),
  ]

  instructors = [
    Instructor(id='I1',
      availability={('Monday','morning'), ('Monday','afternoon')},
      preferred_campuses=['CampusA'],
      preferred_classroom_types=['physical'],
      preferred_classes=['C1','C2']),

    Instructor(id='I2',
      availability={('Monday','afternoon'), ('Tuesday','evening')},
      preferred_campuses=['any'],
      preferred_classroom_types=['physical','virtual'],
      preferred_classes=['any'])
  ]
  ```

  > Now, we run the scheduler.

  > We'll sort classes by enrollment: [C2 (45), C1 (20)]

  > For C2:

  >   enrollment=45
  >   candidate_instructors:
    > I1: preferred_classes has C2 -> yes
    > I2: preferred_classes has 'any' -> yes

  > We try blocks: morning, afternoon, evening.

  > Morning:
  > days: Monday to Friday
  > Monday:

  > available_instructors:
    > I1: available on (Monday, morning) -> yes
    > I2: available on (Monday, morning) -> no -> so only I1

  > candidate_classrooms:

physical:

  > A101: size=50>=45 -> yes, and available in morning? CampusA has morning -> yes.

  > class.preferred_campuses: 'any' -> so okay?

  > class.preferred_classroom_types: includes physical -> okay.

  > A102: 30<45 -> no.

  > B201: campus? CampusB, and class.preferred_campuses: 'any' -> okay, but CampusB does not have morning? -> not available in morning -> skip.

  > virtual:

  > Virtual1: size=100>=45 -> yes, and block morning -> available? virtual available in morning -> yes.

  > class.preferred_classroom_types: includes virtual -> okay.

  > Virtual2: same.

  > Now, we try physical first: so [A101, Virtual1, Virtual2] -> but we split: physical=[A101], virtual=[Virtual1, Virtual2] -> candidate_classrooms = [A101, Virtual1, Virtual2]?

  > But we want physical first: so we try A101 first.

  > For instructor I1 and room A101:

  > Check instructor preference for room A101:

  > A101 is physical ->

  > campus: CampusA -> in I1.preferred_campuses? ['CampusA'] -> yes.

  > type: physical -> in I1.preferred_classroom_types? ['physical'] -> yes.

  > Then, for the block 'morning', total_slots=3, class.length=2 -> valid start_slot: [0,1] (because 0 to 3-2=1)

  > Try start_slot=0:

  > required_slots: (Monday, morning,0) and (Monday, morning,1)

  > Check I1: initially instructor_occupation['I1'] is empty -> free.

  > Check A101: initially empty -> free.

  > So assign:

  > instructor_occupation['I1'] = { (Monday, morning,0), (Monday, morning,1) }

  > classroom_occupation['A101'] = same.

  > Assignment for C2: I1, A101, Monday, morning, start_slot=0.

  > Break.

  > So C2 is scheduled.

  > For C1:

  > enrollment=20

  > candidate_instructors:

  > I1: preferred_classes has C1 -> yes

  > I2: any -> yes

  > We try morning first:

  > Monday:

  > available_instructors:

  > I1: available on (Monday, morning) -> yes

  > I2: not available -> skip.

  > candidate_classrooms:

  > physical:

  > A101: occupied at (Monday, morning,0) and (1) -> so if we try to schedule at slot2?

  > We can try:

  > class.length=1 -> valid start_slot: [0,1,2]?

  > But A101 is occupied at slot0 and 1? so at slot2: free?

  > So A101 is available at slot2?

  > A102: size=30>=20 -> available in morning? yes. and class.preferred_campuses: ['CampusA'] -> yes.

  > We try A102 first.

  > For I1 and A102:

  > I1 preference:

  > campus: A102 is CampusA -> in I1.preferred_campuses -> yes.

  > type: physical -> in I1.preferred_classroom_types -> yes.

  > Then, for start_slot in [0,1,2] (for 1-hour):

  > start_slot=0:

  > required_slots: (Monday, morning,0) ->

  > I1 is already occupied at (Monday, morning,0) -> conflict.

  > start_slot=1: conflict (with (Monday, morning,1))

  > start_slot=2:

  > required_slots: (Monday, morning,2) ->

  > I1: currently occupied at (0,1) -> so 2 is free?

  > A102: free?

  > So assign.

  > So C1: I1, A102, Monday, morning, start_slot=2.

  > Done.

  > But note: I1 is teaching two classes on Monday morning:

  > C2: 9:00-11:00 (if slot0=9, slot1=10) -> 9-11

  > C1: 11:00-12:00 (slot2=11) -> 11-12

  > This is acceptable?

  > We'll code accordingly.

  > However, note: the classroom A101 is occupied for C2 from 9-11, so A102 is free?

  > But what if we had tried to put C1 in A101 at slot2?

  > That would be possible too?

  > But we try physical classrooms in the order given? and we try A102 first?

  > We might want to try the same classroom for multiple classes?

  > But in the candidate_classrooms, we have both A101 and A102?

  > We try A102 first because we sorted physical by the list?

  > We did: candidate_classrooms_physical = [room for room in classrooms ...] -> and the order is the order in the list?

  > So we try A101 first? then A102?

  > But in the example, for C1 we tried A102 first?

  > Actually, for C1, we have candidate_classrooms_physical = [A101, A102] because both are physical and in CampusA and available in the morning?

  > Then we try instructor I1 and room A101:

  > for start_slot in [0,1,2]:

  > slot0: conflict (because A101 is occupied at slot0 by C2? and we just scheduled C2) -> so skip.

  > slot1: conflict (slot1 is occupied by C2) -> skip.

  > slot2:

  > I1: at slot2? I1 is free?

  > A101: at slot2? free?

  > So we can assign.

  > So we would assign C1 to A101 at slot2?

  > Then A102 remains free.

  > So the algorithm would do that.

  > Therefore, we don't need to split physical by anything? we just iterate in the given order?

  > But we did candidate_classrooms_physical and then virtual?

  > However, we did: candidate_classrooms_physical = [room ...] -> so the order is the order of the classrooms list?

  > We might want to sort the candidate classrooms by size? to use the smallest room that fits?

  > But the problem doesn't specify.

  > Alternatively, we can sort the candidate_classrooms_physical by size? so that we use the smallest room first?

  > Similarly, we sort candidate_classrooms_virtual by size?

  > But note: we want to avoid using large rooms for small classes?

  > We can do:

  > candidate_classrooms_physical = sorted([room for room in ...], key=lambda x: x.size)

  > candidate_classrooms_virtual = sorted([room for room in ...], key=lambda x: x.size)

  > Then we try the smallest room first?

  > This would help in packing more classes?

  > We'll do that.

  > Also, for the candidate_instructors? we might want to sort by the number of classes they are already teaching? to balance the load?

  > But the problem doesn't specify.

  > We can sort available_instructors by the current number of assigned classes? or by the current number of occupied slots?

  > We can do:

  > available_instructors_sorted = sorted(available_instructors, key=lambda x: len(instructor_occupation[x.id]))

  > Then we try the instructor with the least load first?

  > Similarly, we can sort the candidate_classrooms by the current number of occupied slots? to spread the load?

  > But we are already sorting by size?

  > Alternatively, we can sort by the current utilization?

  > However, we'll do:

  > Sort candidate_classrooms_physical by size (ascending) -> so we try the smallest room that fits first?

  > Then candidate_classrooms_virtual by size (ascending) -> but virtual classrooms are the same?

  > And sort available_instructors by the number of occupied slots?

  > So we update:

  > candidate_classrooms_physical = sorted([room for ...], key=lambda x: x.size)

  > candidate_classrooms_virtual = sorted([room for ...], key=lambda x: x.size)

  > available_instructors = sorted(available_instructors, key=lambda x: len(instructor_occupation[x.id]))

  > Then we iterate over instructors in that order?

  > This might help balance.

  > We'll code accordingly.

Let's code.

Note: This is a greedy algorithm and might not find a solution even if one exists?

  > We'll assume we can use backtracking? but the problem doesn't require optimality.

  > We'll do a greedy assignment.

  > We'll create the data structures and then the algorithm.

  > Due to the complexity, we'll assume the input is provided as above.

  > We'll write the code accordingly.

  > Let's code the classes and then the scheduler.

We'll assume:

  ```
  days = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday']

  blocks_info = {
    'morning': {'total_slots': 3, 'order':0},
    'afternoon': {'total_slots':4, 'order':1},
    'evening': {'total_slots':4, 'order':2}
  }
  blocks_order = ['morning','afternoon','evening']
  ```

We'll create:

  ```
  class Campus:
    def __init__(self, name, available_blocks):
      self.name = name
      self.available_blocks = available_blocks   # list of block names

  class Classroom:
    def __init__(self, id, campus, size, type):
      self.id = id
      self.campus = campus   # None for virtual
      self.size = size
      self.type = type   # 'physical' or 'virtual'

    def get_available_blocks(self):
      if self.type == 'virtual':
        return ['morning','afternoon','evening']
      else:
        return self.campus.available_blocks

  class Class:
    def __init__(self, id, length, enrollment, preferred_campuses, preferred_classroom_types):
      self.id = id
      self.length = length
      self.enrollment = enrollment
      self.preferred_campuses = preferred_campuses   # list of campus names, and if 'any' is present then any campus is okay?
      self.preferred_classroom_types = preferred_classroom_types   # list of strings: 'physical','virtual','any'

  class Instructor:
    def __init__(self, id, availability, preferred_campuses, preferred_classroom_types, preferred_classes):
      self.id = id
      self.availability = availability   # set of tuples (day, block)
      self.preferred_campuses = preferred_campuses
      self.preferred_classroom_types = preferred_classroom_types
      self.preferred_classes = preferred_classes   # list of class ids, or if 'any' is present then any class
  ```

> Then the scheduling function:

  ```
  def schedule(campuses, classrooms, classes, instructors):
    # ......
  ```

  > Finally, we run the schedule function and print the assignments.

  > Note: the start_slot is an index within the block?

  > To interpret the actual start time, we can map:

  > - For morning:
  >>  slot0 -> 9:00
  >>  slot1 -> 10:00
  >>  slot2 -> 11:00

  > - For afternoon:
  >>  slot0 -> 13:00
  >>  slot1 -> 14:00
  >>  slot2 -> 15:00
  >>  slot3 -> 16:00

  > - For evening:
  >> slot0 -> 18:00
  >> slot1 -> 19:00
  >> slot2 -> 20:00
  >> slot3 -> 21:00

But note: the class might be 2-hour?

  > We can output:
  > - For a class:
  >> day,
  >> block,
  >> start_time: the time corresponding to the start_slot,
  >> end_time: start_time + (class.length in hours)

  > However, the exact time string we can derive from the slot and block?

  > We can create a helper function.

  > This is a basic implementation.

  > Let's test with the example above.

  > Note: We did not include the example data in the code. We'll create a main function to run a test.

  > Due to the length, we'll provide the full code.

  > Note: This is a complex scheduler and may have bugs. We assume the input is as expected.

  > We'll write a test.

  > To solve this problem, we need to design a scheduling program for a university with multiple campuses and virtual classrooms. The program must schedule classes over a 12-week quarter, considering constraints such as classroom availability, instructor preferences, and time slots. The goal is to balance resources by prioritizing physical classrooms over virtual ones and daytime slots over evening slots.

#### Approach

**Problem Analysis**: The problem involves scheduling classes for a university with multiple campuses. Each campus has physical classrooms with specific capacities and availability in morning, afternoon, or evening blocks. Virtual classrooms are available anytime. Classes have attributes like length (1 or 2 hours), enrollment count, and preferred campuses and classroom types. Instructors have preferences for time slots, campuses, classroom types, and classes they can teach.

**Key Insights**:

- Class Scheduling: Each class must be scheduled once per week for 12 weeks at the same time and location.

- Resource Constraints: Physical classrooms should be prioritized over virtual ones. Daytime slots (morning and afternoon) should be used before evening slots.

- Instructor and Classroom Availability: Instructors and classrooms must be available for the scheduled time slots without conflicts.

**Algorithm Selection**:

- Greedy Algorithm: Classes are sorted by enrollment in descending order to prioritize larger classes. For each class, we iterate over time blocks (morning, afternoon, evening), days of the week, and available instructors and classrooms to find a feasible schedule.

- Conflict Checking: For each potential schedule, we check if the instructor and classroom are available for the required time slots without overlapping with other classes.

- Complexity Analysis:

- Time Complexity: The algorithm processes each class (C), each time block (B), each day (D), each instructor (I), and each classroom (R), and each start slot (S). The worst-case complexity is O(C * B * D * I * R * S), which is manageable for typical university sizes (e.g., hundreds of classes and instructors).

- Space Complexity: O(I + R) for storing instructor and classroom schedules.


___
## Solution Code (python)

```python
class Campus:
    def __init__(self, name, available_blocks):
        self.name = name
        self.available_blocks = available_blocks

class Classroom:
    def __init__(self, id, campus, size, type):
        self.id = id
        self.campus = campus
        self.size = size
        self.type = type  # 'physical' or 'virtual'

    def get_available_blocks(self):
        if self.type == 'virtual':
            return ['morning', 'afternoon', 'evening']
        else:
            return self.campus.available_blocks

class Class:
    def __init__(self, id, length, enrollment, preferred_campuses, preferred_classroom_types):
        self.id = id
        self.length = length
        self.enrollment = enrollment
        self.preferred_campuses = preferred_campuses
        self.preferred_classroom_types = preferred_classroom_types

class Instructor:
    def __init__(self, id, availability, preferred_campuses, preferred_classroom_types, preferred_classes):
        self.id = id
        self.availability = availability  # set of (day, block)
        self.preferred_campuses = preferred_campuses
        self.preferred_classroom_types = preferred_classroom_types
        self.preferred_classes = preferred_classes  # list of class ids, or if contains 'any', then any class

def schedule_classes(campuses, classrooms, classes, instructors):
    days_list = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday']
    blocks_info = {
        'morning': {'total_slots': 3, 'order': 0},
        'afternoon': {'total_slots': 4, 'order': 1},
        'evening': {'total_slots': 4, 'order': 2}
    }
    blocks_order = ['morning', 'afternoon', 'evening']

    instructor_occupation = {instructor.id: set() for instructor in instructors}
    classroom_occupation = {classroom.id: set() for classroom in classrooms}
    assignments = []

    sorted_classes = sorted(classes, key=lambda c: c.enrollment, reverse=True)

    for cls in sorted_classes:
        found = False
        candidate_instructors = []
        for instr in instructors:
            if 'any' in instr.preferred_classes or cls.id in instr.preferred_classes:
                candidate_instructors.append(instr)

        for block in blocks_order:
            if found:
                break
            for day in days_list:
                if found:
                    break
                avail_instructors = [instr for instr in candidate_instructors if (day, block) in instr.availability]
                avail_instructors_sorted = sorted(avail_instructors, key=lambda x: len(instructor_occupation[x.id]))

                candidate_physical = []
                candidate_virtual = []
                for room in classrooms:
                    if room.size < cls.enrollment:
                        continue
                    room_avail_blocks = room.get_available_blocks()
                    if block not in room_avail_blocks:
                        continue
                    if room.type == 'physical':
                        if 'any' in cls.preferred_campuses or room.campus.name in cls.preferred_campuses:
                            if 'any' in cls.preferred_classroom_types or 'physical' in cls.preferred_classroom_types:
                                candidate_physical.append(room)
                    else:
                        if 'any' in cls.preferred_classroom_types or 'virtual' in cls.preferred_classroom_types:
                            candidate_virtual.append(room)

                candidate_physical_sorted = sorted(candidate_physical, key=lambda x: x.size)
                candidate_virtual_sorted = sorted(candidate_virtual, key=lambda x: x.size)
                candidate_classrooms = candidate_physical_sorted + candidate_virtual_sorted

                for instructor in avail_instructors_sorted:
                    if found:
                        break
                    for room in candidate_classrooms:
                        if found:
                            break
                        if room.type == 'physical':
                            if ('any' not in instructor.preferred_campuses) and (room.campus.name not in instructor.preferred_campuses):
                                continue
                            if ('any' not in instructor.preferred_classroom_types) and ('physical' not in instructor.preferred_classroom_types):
                                continue
                        else:
                            if ('any' not in instructor.preferred_classroom_types) and ('virtual' not in instructor.preferred_classroom_types):
                                continue

                        total_slots = blocks_info[block]['total_slots']
                        max_start = total_slots - cls.length
                        for start_slot in range(0, max_start + 1):
                            required_slots = set()
                            for i in range(cls.length):
                                slot_index = start_slot + i
                                required_slots.add((day, block, slot_index))

                            if instructor_occupation[instructor.id] & required_slots:
                                continue
                            if classroom_occupation[room.id] & required_slots:
                                continue

                            instructor_occupation[instructor.id] |= required_slots
                            classroom_occupation[room.id] |= required_slots
                            assignments.append({
                                'class_id': cls.id,
                                'instructor_id': instructor.id,
                                'classroom_id': room.id,
                                'day': day,
                                'block': block,
                                'start_slot': start_slot
                            })
                            found = True
                            break
        if not found:
            print(f"Warning: Class {cls.id} could not be scheduled.")
    return assignments

def print_schedule(assignments):
    block_slot_times = {
        'morning': {0: '9:00 AM', 1: '10:00 AM', 2: '11:00 AM'},
        'afternoon': {0: '1:00 PM', 1: '2:00 PM', 2: '3:00 PM', 3: '4:00 PM'},
        'evening': {0: '6:00 PM', 1: '7:00 PM', 2: '8:00 PM', 3: '9:00 PM'}
    }

    for assignment in assignments:
        start_time = block_slot_times[assignment['block']][assignment['start_slot']]
        length = next(cls.length for cls in classes if cls.id == assignment['class_id'])
        if length == 1:
            end_time = block_slot_times[assignment['block']].get(assignment['start_slot'] + 1, "?")
        else:
            end_slot = assignment['start_slot'] + 1
            end_time = block_slot_times[assignment['block']].get(end_slot, "?")
        print(f"Class {assignment['class_id']}: "
              f"Day: {assignment['day']}, "
              f"Time: {start_time} to {end_time}, "
              f"Instructor: {assignment['instructor_id']}, "
              f"Classroom: {assignment['classroom_id']}")

# Example usage:
if __name__ == "__main__":
    campusA = Campus('CampusA', ['morning', 'afternoon'])
    campusB = Campus('CampusB', ['afternoon', 'evening'])

    classrooms = [
        Classroom('A101', campusA, 50, 'physical'),
        Classroom('A102', campusA, 30, 'physical'),
        Classroom('B201', campusB, 40, 'physical'),
        Classroom('Virtual1', None, 100, 'virtual'),
        Classroom('Virtual2', None, 100, 'virtual')
    ]

    classes = [
        Class('C1', 1, 20, ['CampusA'], ['physical']),
        Class('C2', 2, 45, ['any'], ['physical', 'virtual']),
        Class('C3', 1, 60, ['CampusB'], ['physical', 'virtual']),
        Class('C4', 2, 25, ['CampusA', 'CampusB'], ['physical'])
    ]

    instructor1 = Instructor('I1',
                            {('Monday', 'morning'), ('Monday', 'afternoon'), ('Tuesday', 'morning')},
                            ['CampusA'],
                            ['physical'],
                            ['C1', 'C2', 'C4'])
    instructor2 = Instructor('I2',
                            {('Monday', 'afternoon'), ('Tuesday', 'evening'), ('Wednesday', 'evening')},
                            ['any'],
                            ['physical', 'virtual'],
                            ['any'])
    instructor3 = Instructor('I3',
                            {('Wednesday', 'afternoon'), ('Thursday', 'evening')},
                            ['CampusB'],
                            ['physical'],
                            ['C3'])

    instructors = [instructor1, instructor2, instructor3]

    assignments = schedule_classes([campusA, campusB], classrooms, classes, instructors)
    print_schedule(assignments)
```


___
## Explanation

### Classes and Data Structures:

* Campus: Represents a university campus with available time blocks.

* Classroom: Represents a physical or virtual classroom with size and availability.

* Class: Represents a university class with attributes like length, enrollment, and preferences.

* Instructor: Represents an instructor with availability and preferences.

### Scheduling Algorithm:

* Initialization: Prepare data structures to track instructor and classroom schedules.

* Class Processing: Sort classes by enrollment (descending) to prioritize larger classes.

* Time Block Iteration: For each class, iterate over time blocks (morning, afternoon, evening) and days (Monday to Friday).

* Instructor and Classroom Selection: For each time slot, find available instructors and classrooms that meet preferences and capacity requirements.

* Conflict Check: Ensure no scheduling conflicts for instructors and classrooms.

* Assignment: Record valid schedules and update occupation sets.

* Output: Print the schedule with class details, including day, time, instructor, and classroom.

This approach efficiently schedules classes while respecting constraints and prioritizing resources, ensuring optimal use of physical classrooms and daytime slots. The solution handles typical university scheduling scenarios with multiple campuses and virtual options.


___
## Java Implementation

Below is the Java implementation of the university scheduling program that balances resources, prioritizes physical classrooms and daytime slots, and respects instructor preferences.

```java
import java.util.*;
import java.util.stream.Collectors;

// Record to represent a day-block combination (without slot index)
record DayBlock(String day, String block) {}

// Record to represent a time slot (day, block, slot index)
record TimeSlot(String day, String block, int slotIndex) {}

// Record to store final assignment details
record Assignment(String courseId, String instructorId, String classroomId, String day, String block, int startSlot) {}

class Campus {
    private String name;
    private List<String> availableBlocks;

    public Campus(String name, List<String> availableBlocks) {
        this.name = name;
        this.availableBlocks = availableBlocks;
    }

    public String getName() { return name; }
    public List<String> getAvailableBlocks() { return availableBlocks; }
}

class Classroom {
    private String id;
    private Campus campus;
    private int size;
    private String type; // "physical" or "virtual"

    public Classroom(String id, Campus campus, int size, String type) {
        this.id = id;
        this.campus = campus;
        this.size = size;
        this.type = type;
    }

    public String getId() { return id; }
    public Campus getCampus() { return campus; }
    public int getSize() { return size; }
    public String getType() { return type; }

    public List<String> getAvailableBlocks() {
        if ("virtual".equals(type)) {
            return Arrays.asList("morning", "afternoon", "evening");
        }
        return campus.getAvailableBlocks();
    }
}

class Course {
    private String id;
    private int length; // 1 or 2 hours
    private int enrollment;
    private List<String> preferredCampuses; // campus names or "any"
    private List<String> preferredClassroomTypes; // "physical", "virtual", or "any"

    public Course(String id, int length, int enrollment,
                  List<String> preferredCampuses, List<String> preferredClassroomTypes) {
        this.id = id;
        this.length = length;
        this.enrollment = enrollment;
        this.preferredCampuses = preferredCampuses;
        this.preferredClassroomTypes = preferredClassroomTypes;
    }

    public String getId() { return id; }
    public int getLength() { return length; }
    public int getEnrollment() { return enrollment; }
    public List<String> getPreferredCampuses() { return preferredCampuses; }
    public List<String> getPreferredClassroomTypes() { return preferredClassroomTypes; }
}

class Instructor {
    private String id;
    private Set<DayBlock> availability; // Available day-block combinations
    private List<String> preferredCampuses; // campus names or "any"
    private List<String> preferredClassroomTypes; // "physical", "virtual", or "any"
    private List<String> preferredClasses; // course IDs or "any"

    public Instructor(String id, Set<DayBlock> availability,
                      List<String> preferredCampuses,
                      List<String> preferredClassroomTypes,
                      List<String> preferredClasses) {
        this.id = id;
        this.availability = availability;
        this.preferredCampuses = preferredCampuses;
        this.preferredClassroomTypes = preferredClassroomTypes;
        this.preferredClasses = preferredClasses;
    }

    public String getId() { return id; }
    public Set<DayBlock> getAvailability() { return availability; }
    public List<String> getPreferredCampuses() { return preferredCampuses; }
    public List<String> getPreferredClassroomTypes() { return preferredClassroomTypes; }
    public List<String> getPreferredClasses() { return preferredClasses; }
}

class Scheduler {
    private static final List<String> DAYS = Arrays.asList(
        "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"
    );
    private static final Map<String, Integer> BLOCK_SLOTS = Map.of(
        "morning", 3,
        "afternoon", 4,
        "evening", 4
    );
    private static final List<String> BLOCK_ORDER = Arrays.asList(
        "morning", "afternoon", "evening"
    );

    public List<Assignment> schedule(List<Campus> campuses,
                                     List<Classroom> classrooms,
                                     List<Course> courses,
                                     List<Instructor> instructors) {

        // Sort courses by enrollment descending (prioritize larger classes)
        courses.sort(Comparator.comparingInt(Course::getEnrollment).reversed());

        // Track occupied time slots for instructors and classrooms
        Map<String, Set<TimeSlot>> instructorOccupancy = new HashMap<>();
        Map<String, Set<TimeSlot>> classroomOccupancy = new HashMap<>();
        instructors.forEach(instructor -> instructorOccupancy.put(instructor.getId(), new HashSet<>()));
        classrooms.forEach(classroom -> classroomOccupancy.put(classroom.getId(), new HashSet<>()));

        List<Assignment> assignments = new ArrayList<>();

        for (Course course : courses) {
            boolean found = false;
            List<Instructor> candidateInstructors = instructors.stream()
                .filter(instructor ->
                    instructor.getPreferredClasses().contains("any") ||
                    instructor.getPreferredClasses().contains(course.getId())
                )
                .collect(Collectors.toList());

            // Try blocks in priority order: morning -> afternoon -> evening
            for (String block : BLOCK_ORDER) {
                if (found) break;

                for (String day : DAYS) {
                    if (found) break;

                    // Filter instructors available for this day-block
                    DayBlock currentDayBlock = new DayBlock(day, block);
                    List<Instructor> availableInstructors = candidateInstructors.stream()
                        .filter(instructor -> instructor.getAvailability().contains(currentDayBlock))
                        .sorted(Comparator.comparingInt(
                            instructor -> instructorOccupancy.get(instructor.getId()).size())
                        )
                        .collect(Collectors.toList());

                    // Get suitable classrooms (physical first, then virtual)
                    List<Classroom> candidateClassrooms = classrooms.stream()
                        .filter(room -> room.getSize() >= course.getEnrollment())
                        .filter(room -> room.getAvailableBlocks().contains(block))
                        .filter(room ->
                            room.getType().equals("virtual") ||
                            course.getPreferredCampuses().contains("any") ||
                            course.getPreferredCampuses().contains(room.getCampus().getName())
                        )
                        .filter(room ->
                            course.getPreferredClassroomTypes().contains("any") ||
                            course.getPreferredClassroomTypes().contains(room.getType())
                        )
                        .collect(Collectors.toList());

                    // Prioritize physical classrooms
                    List<Classroom> physicalRooms = candidateClassrooms.stream()
                        .filter(room -> "physical".equals(room.getType()))
                        .sorted(Comparator.comparingInt(Classroom::getSize))
                        .collect(Collectors.toList());

                    List<Classroom> virtualRooms = candidateClassrooms.stream()
                        .filter(room -> "virtual".equals(room.getType()))
                        .sorted(Comparator.comparingInt(Classroom::getSize))
                        .collect(Collectors.toList());

                    List<Classroom> sortedClassrooms = new ArrayList<>();
                    sortedClassrooms.addAll(physicalRooms);
                    sortedClassrooms.addAll(virtualRooms);

                    // Try each available instructor and classroom combination
                    for (Instructor instructor : availableInstructors) {
                        if (found) break;

                        for (Classroom classroom : sortedClassrooms) {
                            if (found) break;

                            // Check instructor preferences for classroom
                            if ("physical".equals(classroom.getType())) {
                                if (!instructor.getPreferredCampuses().contains("any") &&
                                    !instructor.getPreferredCampuses().contains(classroom.getCampus().getName())) {
                                    continue;
                                }
                                if (!instructor.getPreferredClassroomTypes().contains("any") &&
                                    !instructor.getPreferredClassroomTypes().contains("physical")) {
                                    continue;
                                }
                            } else {
                                if (!instructor.getPreferredClassroomTypes().contains("any") &&
                                    !instructor.getPreferredClassroomTypes().contains("virtual")) {
                                    continue;
                                }
                            }

                            // Check available time slots
                            int maxStart = BLOCK_SLOTS.get(block) - course.getLength();
                            for (int startSlot = 0; startSlot <= maxStart; startSlot++) {
                                Set<TimeSlot> requiredSlots = new HashSet<>();
                                for (int i = 0; i < course.getLength(); i++) {
                                    requiredSlots.add(new TimeSlot(day, block, startSlot + i));
                                }

                                // Check instructor availability
                                Set<TimeSlot> instructorSlots = instructorOccupancy.get(instructor.getId());
                                if (Collections.disjoint(instructorSlots, requiredSlots)) {
                                    // Check classroom availability
                                    Set<TimeSlot> classroomSlots = classroomOccupancy.get(classroom.getId());
                                    if (Collections.disjoint(classroomSlots, requiredSlots)) {
                                        // Book slots
                                        instructorSlots.addAll(requiredSlots);
                                        classroomSlots.addAll(requiredSlots);

                                        // Create assignment
                                        assignments.add(new Assignment(
                                            course.getId(),
                                            instructor.getId(),
                                            classroom.getId(),
                                            day,
                                            block,
                                            startSlot
                                        ));
                                        found = true;
                                        break;
                                    }
                                }
                            }
                        }
                    }
                }
            }
            if (!found) {
                System.err.println("Warning: Could not schedule course " + course.getId());
            }
        }
        return assignments;
    }
}

public class UniversityScheduler {
    public static void main(String[] args) {
        // Create campuses
        Campus campusA = new Campus("CampusA", Arrays.asList("morning", "afternoon"));
        Campus campusB = new Campus("CampusB", Arrays.asList("afternoon", "evening"));

        // Create classrooms
        List<Classroom> classrooms = Arrays.asList(
            new Classroom("A101", campusA, 50, "physical"),
            new Classroom("A102", campusA, 30, "physical"),
            new Classroom("B201", campusB, 40, "physical"),
            new Classroom("Virtual1", null, 100, "virtual"),
            new Classroom("Virtual2", null, 100, "virtual")
        );

        // Create courses
        List<Course> courses = Arrays.asList(
            new Course("C1", 1, 20, Arrays.asList("CampusA"), Arrays.asList("physical")),
            new Course("C2", 2, 45, Arrays.asList("any"), Arrays.asList("physical", "virtual")),
            new Course("C3", 1, 60, Arrays.asList("CampusB"), Arrays.asList("physical", "virtual")),
            new Course("C4", 2, 25, Arrays.asList("CampusA", "CampusB"), Arrays.asList("physical"))
        );

        // Create instructors
        Instructor instructor1 = new Instructor(
            "I1",
            new HashSet<>(Arrays.asList(
                new DayBlock("Monday", "morning"),
                new DayBlock("Monday", "afternoon"),
                new DayBlock("Tuesday", "morning")
            )),
            Arrays.asList("CampusA"),
            Arrays.asList("physical"),
            Arrays.asList("C1", "C2", "C4")
        );

        Instructor instructor2 = new Instructor(
            "I2",
            new HashSet<>(Arrays.asList(
                new DayBlock("Monday", "afternoon"),
                new DayBlock("Tuesday", "evening"),
                new DayBlock("Wednesday", "evening")
            )),
            Arrays.asList("any"),
            Arrays.asList("physical", "virtual"),
            Arrays.asList("any")
        );

        Instructor instructor3 = new Instructor(
            "I3",
            new HashSet<>(Arrays.asList(
                new DayBlock("Wednesday", "afternoon"),
                new DayBlock("Thursday", "evening")
            )),
            Arrays.asList("CampusB"),
            Arrays.asList("physical"),
            Arrays.asList("C3")
        );

        List<Instructor> instructors = Arrays.asList(instructor1, instructor2, instructor3);

        // Run scheduler
        Scheduler scheduler = new Scheduler();
        List<Assignment> schedule = scheduler.schedule(
            Arrays.asList(campusA, campusB),
            classrooms,
            courses,
            instructors
        );

        // Print schedule
        Map<String, Map<Integer, String>> slotTimes = Map.of(
            "morning", Map.of(
                0, "9:00 AM", 1, "10:00 AM", 2, "11:00 AM"
            ),
            "afternoon", Map.of(
                0, "1:00 PM", 1, "2:00 PM", 2, "3:00 PM", 3, "4:00 PM"
            ),
            "evening", Map.of(
                0, "6:00 PM", 1, "7:00 PM", 2, "8:00 PM", 3, "9:00 PM"
            )
        );

        for (Assignment assignment : schedule) {
            String startTime = slotTimes.get(assignment.block()).get(assignment.startSlot());
            int endSlot = assignment.startSlot() + 1;
            String endTime = endSlot < BLOCK_SLOTS.get(assignment.block())
                ? slotTimes.get(assignment.block()).get(endSlot)
                : "End of block";

            System.out.printf(
                "Class %s: Day: %s, Time: %s to %s, Instructor: %s, Classroom: %s%n",
                assignment.courseId(),
                assignment.day(),
                startTime,
                endTime,
                assignment.instructorId(),
                assignment.classroomId()
            );
        }
    }
}
```


___
## Summary

### Key Features and Design Choices

* Data Structures:
  - Records: Used for immutable data carriers (DayBlock, TimeSlot, Assignment)
  - Classes: Core entities (Campus, Classroom, Course, Instructor)
  - Collections: Efficiently track occupied time slots using HashSet

* Scheduling Algorithm:
  - Priority Handling: Processes larger classes first (sorted by enrollment)
  - Resource Prioritization:
    * Physical classrooms before virtual ones
    * Daytime slots (morning/afternoon) before evening slots
  - Conflict Checking: Uses disjoint set operations to ensure no overlapping bookings

* Preference Handling:
  - Instructor preferences for courses, campuses, and classroom types
  - Course preferences for campuses and classroom types
  - Flexible "any" option for preferences

* Time Management:
  - Time blocks divided into 1-hour slots
  - Morning (3 slots), Afternoon (4 slots), Evening (4 slots)
  - Consecutive slot booking for multi-hour classes

* Output:
  - Human-readable schedule with start/end times
  - Warnings for unscheduled courses

### Execution Workflow

* Initialize Entities:
  - Create campuses with available time blocks
  - Create classrooms with capacities and types
  - Define courses with enrollment and preferences
  - Set up instructors with availability and preferences

* Run Scheduler:
  - Processes courses in descending order of enrollment
  - For each course, tries time blocks in priority order (morning → afternoon → evening)
  - Finds compatible instructor-classroom pairs
  - Books earliest available consecutive slots

* Output Schedule:
  - Lists each assignment with day, time slot, instructor, and classroom
  - Converts slot indices to human-readable times
  - Shows warnings for any unscheduled courses

This implementation efficiently balances resource utilization while respecting institutional constraints and preferences, providing a practical solution for university scheduling needs.
