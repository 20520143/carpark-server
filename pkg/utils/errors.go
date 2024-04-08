package utils

import "net/http"

var messageError map[int]string

func LoadMessageError() {
	messageError = make(map[int]string)
	messageError[http.StatusOK] = "Thành công"
	messageError[http.StatusForbidden] = "Thao tác của bạn đã bị từ chối"
	messageError[http.StatusInternalServerError] = "Lỗi hệ thống"
	messageError[http.StatusBadRequest] = "Không thể thực hiện thao tác"
	messageError[http.StatusUnauthorized] = "Quyền truy cập bị từ chối"
	messageError[http.StatusNotFound] = "Không tìm thấy bản ghi"
	messageError[http.StatusCreated] = "Đã tạo thành công"
	messageError[http.StatusGatewayTimeout] = "Yêu cầu vượt quá thời gian cho phép"
	messageError[http.StatusConflict] = "Dữ liệu đầu vào của bạn đã xung đột với một dữ liệu khác"
	messageError[http.StatusTooManyRequests] = "Thao tác quá nhanh"
	//messageError[http.StatusOK] = "Successfully"
	//messageError[http.StatusForbidden] = "Something when wrong, Your request has been rejected"
	//messageError[http.StatusInternalServerError] = "Internal server error"
	//messageError[http.StatusBadRequest] = "Something when wrong with your request"
	//messageError[http.StatusUnauthorized] = "Unauthorized, Permission denied"
	//messageError[http.StatusNotFound] = "Record not found, Please check your input"
	//messageError[http.StatusCreated] = "Created successfully"
	//messageError[http.StatusGatewayTimeout] = "Gateway time out"
	//messageError[http.StatusConflict] = "Your input has been conflict with another data"
	//messageError[http.StatusTooManyRequests] = "Too many request"
}

func MessageError() map[int]string {
	return messageError
}
