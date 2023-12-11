# Tag Total Timewarrior extension
This [Timewarrior extension](https://timewarrior.net/docs/extensions/) calculates the total time registered on each tag for the closed intervals received in the input. If an interval has more than one tag, the duration of the interval will be added to all the tags.
Total time is printed out in descending order, from highest to lowest.

## Installation
1. Download the latest executable for your operating system from the [releases page](https://github.com/crossbone-magister/<REPLACE_ME>/releases).
2. Add it to the Timewarrior extension folder as described in the [documentation](https://timewarrior.net/docs/api/).
3. Verify that the extension is active and installed by running `timew extensions`.

## Usage
In a terminal window, run `timew tag-total`. An example output could be:
```bash
tag1          - 50h  5m  0s
very-long-tag -  3h  0m 10s
tag2          -  0h 55m  0s

```
