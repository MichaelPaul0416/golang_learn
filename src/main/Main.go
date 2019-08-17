package main

import "../chapter9"

func main()  {
	chapter9.CustomExclusiveLock(10,4)

	chapter9.ConcurrencyWithLock(10,3)

	chapter9.ReadWriteLock(9,3)
}