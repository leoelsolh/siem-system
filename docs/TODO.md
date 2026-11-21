
## **Collect**

- [x] Read the log file into the program
- [x] Split into lines, no empty line at the end
- [x] `collectLogs` returns `[]string`
- [x] Line count matches `wc -l` (+1 if no trailing newline)

## **Parse**

- [x] Map out the fields in each line of `sample.log`
- [x] `Log` struct to hold them
- [x] Parse one line into a `Log`
- [x] Malformed lines get logged and skipped, not a panic
- [x] Every line in the sample becomes one struct with the right values

## **Store**

- [x] Keep parsed events in a slice
- [x] `filterEvents` returns only the matching events (sshd + "Failed password")
- [x] Count matches `grep -c` (5)
- [ ] Maybe later: move events into SQLite

## **Detect**

- [ ] Write the brute-force rule in plain English first: what's a hit, how many, what time window
- [ ] Figure out how to count failed logins per source IP
- [ ] Implement the rule
- [ ] Fires for `203.0.113.45`, doesn't fire for `198.51.100.7`

## **Alert**

- [ ] Print one alert line when the rule fires: rule name, IP, count
- [ ] One alert per attacking IP, not one per failed line

## **Done**

- [ ] All five stages pass on `sample.log`
- [ ] Second sample file with a different mix of events, make sure nothing breaks
