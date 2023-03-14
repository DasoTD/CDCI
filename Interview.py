# from collections import Counter
def one(self):
    pass

def two(schedule: list[str], time):
    from datetime import datetime
    s1 = '12:30'
    s2 = '14:00' # for example
    format = '%H:%M'
    timesDiff = []
    times = datetime.strptime(s2, format) - datetime.strptime(s1, format)
    print(times)
    scheduleTime = datetime.strptime(schedule[0], format) - datetime.strptime(time, format)
    print(scheduleTime)
    scheduledTime = datetime.strptime(schedule[0], format)
    currentTime = datetime.strptime(time, format)
    if scheduledTime > currentTime:
        print(" e pass")
    else:
        print("e no pass")
    for i in range(len(schedule)):
        data = datetime.strptime(schedule[i], format) - datetime.strptime(time, format)
        timesDiff.append(data)
        # print(data, currentTime)
    print(timesDiff)
    # seconds = seconds % (24 * 3600)
    # if schedule[0] == time:
    #     print(-1)
    # passschedule[0] == time:
    #     print(-1)
    # pass

def three(n,m, figures):
    pass

def four(numbers: list[int], x):
    number = [str(x) for x in numbers]
    print(number)
    count = 0
    x = str(x)
    count =0
    for i in range(len(numbers)):
        for j in range( len(numbers)):
            if i!=j and str(numbers[i]) + str(numbers[j]) == x:
                count +=1
    return count
    # for i in range(len(number)):
    #     for j in range(i+1, len(number)):
    #         if i!=j and number[i] + number[j] == x:
    #             count +=1
    #         if number[j]+ number[i] == x:
    #             count += 1
    # return count

def minutes_since_last_bus(schedule, time):
    # Convert current time to minutes
    current_time = int(time.split(':')[0]) * 60 + int(time.split(':')[1])

    print(current_time, "current time")

    # Convert arrival times in schedule to minutes and sort in increasing order
    arrival_times = sorted([int(t.split(':')[0]) * 60 + int(t.split(':')[1]) for t in schedule])

    print(arrival_times, "arrival")
    # Iterate over sorted arrival times in reverse order
    for t in reversed(arrival_times):
        print(t, "tee")
        if t <= current_time:
            # Last bus has already left
            return current_time - t
    # First bus for the day has yet to leave
    return -1
    

data = [1,212, 12, 12]
x = 1212
print(four(data, x))

# two(["12:00", "14:00", "19:55"], "14:30")
schedule = ["10:00", "11:30", "12:45", "14:00", "16:30", "19:00"]
time = "15:05"
print(minutes_since_last_bus(schedule, time))  # Output: 65

time = "20:00"
print(minutes_since_last_bus(schedule, time))  # Output: -1
