package utils

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"time"
	"unicode"

	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type ConsumerRequest struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
	//Payload string `json:"payload"`
}

type UserHasBusiness struct {
	UserID     uuid.UUID `json:"user_id"`
	BusinessID uuid.UUID `json:"business_id"`
	Domain     string    `json:"domain"`
}

type Message struct {
	Type           string `json:"type"`
	OrderNumber    string `json:"order_number"`
	MessageContent string `json:"message_content"`
}

func CurrentUser(c *http.Request) (uuid.UUID, error) {
	userIdStr := c.Header.Get("x-user-id")
	if strings.Contains(userIdStr, "|") {
		userIdStr = strings.Split(userIdStr, "|")[0]
	}
	res, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, err
	}
	return res, nil
}

func String(in string) *string {
	return &in
}

func UUID(req *uuid.UUID) uuid.UUID {
	if req == nil {
		return uuid.Nil
	}
	return *req
}

func ConvertTimestampVN(dateTimeFrom *time.Time, dateTimeTo *time.Time) (string, string) {
	dateTimeFromStr := dateTimeFrom.Format("2006-01-02")
	dateTimeToStr := dateTimeTo.Format("2006-01-02")

	dateTimeFromStr = dateTimeFromStr + " 00:00:00+07"
	dateTimeToStr = dateTimeToStr + " 23:59:59+07"

	return dateTimeFromStr, dateTimeToStr
}

func TransformString(in string, uppercase bool) string {
	in = strings.TrimSpace(in)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, err := transform.String(t, in)
	if err != nil {
		logs.Error("Failed to transform %s ", in)
		return ""
	}
	result = strings.ReplaceAll(result, "Đ", "D")
	result = strings.ReplaceAll(result, "đ", "d")
	if uppercase {
		return strings.ToUpper(result)
	}
	return strings.ToLower(result)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func ConvertTimeIntToString(in int) string {
	if in < 10 {
		return "0" + strconv.Itoa(in)
	}
	return strconv.Itoa(in)
}

func ConvertTimeFormatForReport(in time.Time) string {
	return fmt.Sprintf("%s/%s/%s - %s:%s",
		ConvertTimeIntToString(in.Day()),
		ConvertTimeIntToString(int(in.Month())),
		ConvertTimeIntToString(in.Year()),
		ConvertTimeIntToString(in.Hour()),
		ConvertTimeIntToString(in.Minute()),
	)
}

func RemoveSpace(str string) string {
	re := regexp.MustCompile(`\s+`)
	out := re.ReplaceAllString(str, " ")
	out = strings.TrimSpace(out)
	return out
}

func ValidTimeRequest(inStartTime string, inEndTime string) (outStartTime string, outEndTime string, err error) {
	t1, err := time.Parse(time.RFC3339, inStartTime)
	if err != nil {
		return "", "", ginext.NewError(http.StatusBadRequest, "Lỗi sai định dạng ngày bắt đầu")
	}
	outStartTime = t1.UTC().Format(TIME_FORMAT_FOR_QUERRY)
	t2, err := time.Parse(time.RFC3339, inEndTime)
	if err != nil {
		return "", "", ginext.NewError(http.StatusBadRequest, "Lỗi sai định dạng ngày kết thúc")
	}
	outEndTime = t2.UTC().Format(TIME_FORMAT_FOR_QUERRY)
	return outStartTime, outEndTime, nil
}

func EndOfWeek(date time.Time) time.Time {
	return date.AddDate(0, 0, +6)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func DayTime(i time.Time) *time.Time {
	if i.IsZero() {
		return nil
	} else {
		return &i
	}
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func ResultFloat(input1 *float64, input2 *float64) *float64 {
	if input1 != nil && input2 != nil {
		result := *input1 + *input2
		return &result
	}
	if input1 == nil && input2 != nil {
		result := *input2
		return &result
	}
	if input1 != nil && input2 == nil {
		result := *input1
		return &result
	}
	return nil
}

func CurrentBusiness(c *http.Request) (uuid.UUID, error) {
	userIdStr := c.Header.Get("x-business-id")
	if strings.Contains(userIdStr, "|") {
		userIdStr = strings.Split(userIdStr, "|")[0]
	}
	res, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, err
	}
	return res, nil
}

func RandomFloat(min float64, max float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	res := min + rand.Float64()*(max-min)
	return math.Round(res*ratio) / ratio
}
