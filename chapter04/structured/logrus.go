package structured

import "github.com/sirupsen/logrus"

//Hook 구조체는 logrus의 hook 인터페이스를 구현한다.
type Hook struct {
	id string
}

//Fire 함수는 로그를 기록할 때마다 호출한다.
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

//Levels 함수는 전달된 hook이 실행될 레벨이다.
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//Logrus 함수는 몇 가지 기본적인 logrus 기능을 보여준다.
func Logrus() {
	//json 포맷으로 로그를 기록한다.
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	field := logrus.Fields{}
	field["success"] = true
	field["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "just now"}

	x := logrus.WithFields(field)
	x.Warn("Warning!")
	x.Error("error!")
}
