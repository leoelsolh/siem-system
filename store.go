package main

func filterEvents(l []Log) []Log {
	var res []Log
	for _, e := range l {
		found := flagEntry(e)
		if found {
			res = append(res, e)
		}
	}
	return res
}
