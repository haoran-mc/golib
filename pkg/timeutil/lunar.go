package timeutil

import (
	"time"
)

var (
	chineseNumber         []string = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二"}
	gan                   []string = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	zhi                   []string = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	chineseTen            []string = []string{"初", "十", "廿", "卅"}
	animals               []string = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	year, month, day      int
	leap                  bool
	timeFormat_yyyy_MM_dd string = "2006-01-02 15:04:05"
	timeFormat_yyyyMMdd   string = "20060102"
	lunarInfo             []int  = []int{0x04bd8, 0x04ae0, 0x0a570,
		0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2,
		0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0,
		0x0ada2, 0x095b0, 0x14977, 0x04970, 0x0a4b0, 0x0b4b5, 0x06a50,
		0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, 0x06566,
		0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0,
		0x1c8d7, 0x0c950, 0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4,
		0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, 0x06ca0, 0x0b550,
		0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950,
		0x06aa0, 0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260,
		0x0f263, 0x0d950, 0x05b57, 0x056a0, 0x096d0, 0x04dd5, 0x04ad0,
		0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6,
		0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40,
		0x0af46, 0x0ab60, 0x09570, 0x04af5, 0x04970, 0x064b0, 0x074a3,
		0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0, 0x0c960,
		0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0,
		0x092d0, 0x0cab5, 0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9,
		0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, 0x07954, 0x06aa0,
		0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65,
		0x0d530, 0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0,
		0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, 0x0b5a0, 0x056d0, 0x055b2,
		0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0}
)

// 输入日期格式为 yyyyMMdd、20150101
func Lunar(date string) string {
	var monCyl, leapMonth = 0, 0
	t1, _ := time.Parse(timeFormat_yyyy_MM_dd, "1900-01-31 00:00:00")
	t2, err := time.Parse(timeFormat_yyyyMMdd, date)
	if err != nil {
		return "the date format is wrong"
	}
	offset := int((t2.UnixNano() - t1.UnixNano()) / 1000000 / 86400000)
	monCyl = 14
	var iYear, daysOfYear = 0, 0

	for iYear = 1900; iYear < 2050 && offset > 0; iYear++ {
		daysOfYear = yearDays(iYear)
		offset -= daysOfYear
		monCyl += 12
	}

	if offset < 0 {
		offset += daysOfYear
		iYear--
		monCyl -= 12
	}
	year = iYear
	leapMonth = leapMonthMethod(iYear)
	leap = false

	var iMonth, daysOfMonth = 0, 0

	for iMonth = 1; iMonth < 13 && offset > 0; iMonth++ {
		if leapMonth > 0 && iMonth == (leapMonth+1) && !leap {
			iMonth--
			leap = true
			daysOfMonth = leapDays(year)
		} else {
			daysOfMonth = monthDays(year, iMonth)
		}
		offset -= daysOfMonth
		if leap && iMonth == (leapMonth+1) {
			leap = false
		}
		if !leap {
			monCyl++
		}
	}

	if offset == 0 && leapMonth > 0 && iMonth == leapMonth+1 {
		if leap {
			leap = false
		} else {
			leap = true
			iMonth--
			monCyl--
		}
	}

	if offset < 0 {
		offset += daysOfMonth
		iMonth--
		monCyl--
	}
	month = iMonth
	day = offset + 1

	doubleMonth := ""
	if leap {
		doubleMonth = "闰"
	}
	return cyclical() + animalsYear() + "年" + doubleMonth + chineseNumber[month-1] + "月" + getChinaDayString(day)
}

func animalsYear() string {
	return animals[(year-4)%12]
}

func yearDays(y int) int {
	var i, sum = 348, 348
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (lunarInfo[y-1900] & i) != 0 {
			sum++
		}
	}
	return (sum + leapDays(y))
}

func getChinaDayString(day int) string {
	n := day
	if n%10 == 0 {
		n = 9
	} else {
		n = day%10 - 1
	}
	if day > 30 {
		return ""
	} else if day == 10 {
		return "初十"
	} else {
		return chineseTen[day/10] + chineseNumber[n]
	}
}

func leapMonthMethod(y int) int {
	return (int)(lunarInfo[y-1900] & 0xf)
}

func monthDays(y, m int) int {
	if (lunarInfo[y-1900] & (0x10000 >> uint(m))) == 0 {
		return 29
	}
	return 30
}

func leapDays(y int) int {
	if leapMonthMethod(y) != 0 {
		if (lunarInfo[y-1900] & 0x10000) != 0 {
			return 30
		}
		return 29
	}
	return 0
}

func cyclicalm(num int) string {
	return gan[num%10] + zhi[num%12]
}

func cyclical() string {
	num := year - 1900 + 36
	return (cyclicalm(num))
}
