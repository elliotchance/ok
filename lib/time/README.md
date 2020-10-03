# time

- [April number](#constants)
- [August number](#constants)
- [December number](#constants)
- [February number](#constants)
- [Hour number](#constants)
- [January number](#constants)
- [July number](#constants)
- [June number](#constants)
- [March number](#constants)
- [May number](#constants)
- [Microsecond number](#constants)
- [Millisecond number](#constants)
- [Minute number](#constants)
- [Nanosecond number](#constants)
- [November number](#constants)
- [October number](#constants)
- [Second number](#constants)
- [September number](#constants)

- [func Add(t Time, duration Duration) Time](#Add)
- [func After(a Time, b Time) bool](#After)
- [func Before(a Time, b Time) bool](#Before)
- [func Duration(seconds number) Duration](#Duration)
- [func Equal(a Time, b Time) bool](#Equal)
- [func FromUnix(seconds number) Time](#FromUnix)
- [func Now() Time](#Now)
- [func Sleep(duration Duration)](#Sleep)
- [func Sub(a Time, b Time) Duration](#Sub)
- [func Time(Year number, Month number, Day number, Hour number, Minute number, Second number) Time](#Time)
- [func Unix(t Time) number](#Unix)

## Constants

```
April = 4
```

```
August = 8
```

```
December = 12
```

```
February = 2
```

```
Hour = 3600
```

```
January = 1
```

```
July = 7
```

```
June = 6
```

```
March = 3
```

```
May = 5
```

```
Microsecond = 0.000001
```

```
Millisecond = 0.001
```

```
Minute = 60
```

```
Nanosecond = 0.000000001
```

```
November = 11
```

```
October = 10
```

```
Second = 1
```

```
September = 9
```

## Add

```
func Add(t Time, duration Duration) Time
```

Add returns a new time after applying a duration. You may use a negative
duration to subtract.

## After

```
func After(a Time, b Time) bool
```

After returns true if `a` is after `b`.

## Before

```
func Before(a Time, b Time) bool
```

Before returns true if `a` is before `b`.

## Duration

```
func Duration(seconds number) Duration
```

A Duration represents a length of time.

## Equal

```
func Equal(a Time, b Time) bool
```

Equal returns true only if both instances represent the exact same time
(including fractional seconds).

## FromUnix

```
func FromUnix(seconds number) Time
```

FromUnix performs the opposite operation as Unix. It receives the number of
seconds elapsed since January 1, 1970 UTC and returns the Time. The input
seconds may be fractional.

## Now

```
func Now() Time
```

Now returns the current time.

## Sleep

```
func Sleep(duration Duration)
```

Sleep will pause the execution for a specific duration of time.

## Sub

```
func Sub(a Time, b Time) Duration
```

Sub returns the duration between two time. The result will be negative if `b`
is before `a`, positive if `b` is after `a` and `0` if the two times are
equal.

## Time

```
func Time(Year number, Month number, Day number, Hour number, Minute number, Second number) Time
```

A Time represents a single point in time. The Second may be fractional.

## Unix

```
func Unix(t Time) number
```

Unix returns the unix timestamp, the number of seconds elapsed since January
1, 1970 UTC. The value returned may be fractional.

