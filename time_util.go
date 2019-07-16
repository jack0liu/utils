package utils

import "time"

const (
	TIME_LAYOUT string = "2006-01-02 15:04:05.999999999"
)

//获取系统当前UTC时间，返回字符串形式，精确到秒
func GetCurrentTime() string {
	t := time.Now().UTC()
	return t.Format(TIME_LAYOUT)
}

func TransStr2Time(s string) (t time.Time, error error) {
	t, err := time.Parse(TIME_LAYOUT, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
