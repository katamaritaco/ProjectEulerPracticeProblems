/*Problem 19 - Counting Sundays

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

import "fmt"

const DAYS_IN_WEEK = 7;
const DAYS_IN_YEAR = 365;

const DAYS_IN_JAN_MAR_MAY_JUL_AUG_OCT_DEC = 31;
const DAYS_IN_APR_JUN_SEPT_NOV = 30;
const DAYS_IN_FEB = 28;

func main() {

    sundaysOnFirstOfTheMonth := 0;

    //jan 1 1900 was a monday. find what jan 1 1901 was.
    //1900 was not a leap year - A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

    //We'll treat 0 - Sunday, 1 - Monday, ..., 6 - Saturday;
  	day := ( 1 + DAYS_IN_YEAR ) % DAYS_IN_WEEK;//1, because started on a monday

    //go from jan 1 1901 to dec 31 2000 - 100 years
    for years := 1901; years <= 2000; years++ {

    	for months := 1; months <= 12; months++ {

    		if day == 0 {
    			sundaysOnFirstOfTheMonth++;
    		}

    		if months == 1 || months == 3 || months == 5 || months == 7 || months == 8 || months == 10 || months == 12 {
    			day = ( day + DAYS_IN_JAN_MAR_MAY_JUL_AUG_OCT_DEC ) % DAYS_IN_WEEK;
    		} else if months == 4 || months == 6 || months == 9 || months == 11 {
    			day = ( day + DAYS_IN_APR_JUN_SEPT_NOV ) % DAYS_IN_WEEK;
    		} else { //feb...
    			day = ( day + DAYS_IN_FEB ) % DAYS_IN_WEEK;
    			if years % 4 == 0 {
    				day = ( day + 1 ) % DAYS_IN_WEEK;
    			}
    		}

    	}

    }

    fmt.Printf("Sundays on first of the month: %v\n", sundaysOnFirstOfTheMonth );

}